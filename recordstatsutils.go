package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordstats/proto"
)

var (
	unlistenedCDs = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unlistened_cds",
		Help: "The oldest physical record",
	})
	unlistened12s = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unlistened_12s",
		Help: "The oldest physical record",
	})
	unlistened45s = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unlistened_45s",
		Help: "The oldest physical record",
	})
	unlistenedDigital = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unlistened_digital",
		Help: "The oldest physical record",
	})

	oldest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_oldest",
		Help: "The oldest physical record",
	})

	oldestIC = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_oldest_in_collection",
		Help: "The oldest physical record",
	})

	processed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_processed",
		Help: "The number of records processed",
	})

	totalSales = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_total_sales",
		Help: "The number of records processed",
	})

	totalCompleteSales = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_total_sales_complete",
		Help: "The number of records processed",
	})

	totalToAuditions = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_total_to_be_auditioned",
		Help: "The number of records processed",
	})
	totalAuditions = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_total_auditioned",
		Help: "The number of records processed",
	})
	weekAuditioned = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_auditioned_week",
		Help: "Number of auditions this week",
	})
	scoreAuditioned = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordstats_auditioned_scores",
		Help: "All the auditioned scores",
	}, []string{"score"})
	totalUnfilled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unfiled",
		Help: "The number of records processed",
	})
	totalFilled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_filed",
		Help: "The number of records processed",
	})
	rateFiled = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_filed_rate",
		Help: "The number of records processed",
	})
	oldestLBStaged = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_oldest_lb_staged",
		Help: "The number of records processed",
	})
	oldestLBHigh = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_oldest_lb_highschool",
		Help: "The number of records processed",
	})
	lastListen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_last_listen",
		Help: "The number of records processed",
	})
	listenRate = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_listen_rate",
	})
	listenTodayRate = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_listen_today_rate",
	})
	listenTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_days_to_listen",
	})
	cwidths = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_cwidths",
	})
	medianListen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_median_listen",
		Help: "The number of records processed",
	})
	unlistened = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unlistened",
		Help: "The number of records processed",
	})
	tValue = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordstats_total_value",
		Help: "The number of records processed",
	}, []string{"category", "filed"})

	aValue = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordstats_average_value",
		Help: "The number of records processed",
	}, []string{"category", "filed"})

	keeps = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordstats_keepers",
		Help: "The number of records kept",
	}, []string{"folder"})
	categories = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordstats_categories",
		Help: "The number of records kept",
	}, []string{"category", "filed"})
	withWeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordstats_weights",
		Help: "Records that have weights",
	}, []string{"category"})

	cdsToListen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_cdstogo",
	})
)

const (
	// CONFIG - where we store the stats config
	CONFIG = "/github.com/brotherlogic/recordstats/config"
)

func (s *Server) cleanCategories(ctx context.Context) error {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if err != nil {
		return err
	}
	config := data.(*pb.Config)

	for id := range config.GetCategories() {
		_, err := s.getRecord(ctx, id)
		if status.Code(err) == codes.OutOfRange {
			delete(config.Categories, id)
		}
	}

	return s.KSclient.Save(ctx, CONFIG, config)
}

func (s *Server) update(ctx context.Context, id int32) error {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if err != nil {
		return err
	}
	config := data.(*pb.Config)
	cleanAuditions(config)

	if config.GetFiled() == nil {
		config.Filed = make(map[int32]rcpb.ReleaseMetadata_FileSize)
	}
	if config.GetFiledTime() == nil {
		config.FiledTime = make(map[int32]int64)
	}
	if config.GetLbLastTime() == nil {
		s.CtxLog(ctx, "RESTTING LB")
		config.LbLastTime = make(map[int32]int64)
	}
	if config.GetLbLastTimeHigh() == nil {
		s.CtxLog(ctx, "RESETTING LB HIGH")
		config.LbLastTimeHigh = make(map[int32]int64)
	}
	if config.GetLastListen() == nil {
		s.CtxLog(ctx, "RESETTING LB HIGH")
		config.LastListen = make(map[int32]int64)
	}
	if config.GetKeeps() == nil {
		config.Keeps = make(map[int32]rcpb.ReleaseMetadata_KeepState)
	}
	if config.GetFolder() == nil {
		config.Folder = make(map[int32]int32)
	}
	if config.GetCategories() == nil {
		config.Categories = make(map[int32]rcpb.ReleaseMetadata_Category)
	}
	if config.GetWeights() == nil {
		config.Weights = make(map[int32]int32)
	}
	if config.GetSleeves() == nil {
		config.Sleeves = make(map[int32]rcpb.ReleaseMetadata_SleeveState)
	}
	if config.GetIsDirty() == nil {
		config.IsDirty = make(map[int32]bool)
	}
	if config.GetScore() == nil {
		config.Score = make(map[int32]int32)
	}

	defer func() {
		s.computeUnlistenedCDs(ctx, config)
		s.computeUnlistened45s(ctx, config)
		s.computeUnlistened12s(ctx, config)
		s.computeUnlistenedDigital(ctx, config)

		sleeveCount := float64(0)
		for id, state := range config.GetSleeves() {
			if state == rcpb.ReleaseMetadata_VINYL_STORAGE_NO_INNER {
				if config.GetCategories()[id] != rcpb.ReleaseMetadata_LISTED_TO_SELL && config.GetCategories()[id] != rcpb.ReleaseMetadata_SOLD_ARCHIVE {
					sleeveCount++
				}
			}
		}
		cwidths.Set(sleeveCount)
		weightCount := make(map[string]float64)
		for id, weight := range config.GetWeights() {
			if weight > 0 {
				weightCount[config.GetCategories()[id].String()]++
			}
		}
		for cat, count := range weightCount {
			withWeight.With(prometheus.Labels{"category": cat}).Set(count)
		}

		keepCount := make(map[int32]float64)
		for key, value := range config.GetKeeps() {
			if value == rcpb.ReleaseMetadata_KEEPER {
				keepCount[config.GetFolder()[key]]++
			}
		}
		for folder, count := range keepCount {
			keeps.With(prometheus.Labels{"folder": fmt.Sprintf("%v", folder)}).Set(count)
		}

		catCount := make(map[string]float64)
		for id, value := range config.GetCategories() {
			catCount[fmt.Sprintf("%v|%v", value, config.GetFiled()[id])]++
		}
		for category, count := range catCount {
			elems := strings.Split(category, "|")
			categories.With(prometheus.Labels{"category": elems[0], "filed": elems[1]}).Set(count)
		}

		for _, value := range config.GetValues() {
			tl := float32(0)
			for _, v := range value.GetValue() {
				tl += v
			}

			tValue.With(prometheus.Labels{"category": value.GetCategory(), "filed": value.GetFilling()}).Set(float64(tl))
			aValue.With(prometheus.Labels{"category": value.GetCategory(), "filed": value.GetFilling()}).Set(float64(tl / float32(len(value.GetValue()))))
		}

		tA := 0
		tAA := 0

		wCount := 0
		sa := make(map[int32]int32)
		for _, auditioned := range config.GetAuditions() {
			if auditioned.GetAudScore() > 0 {
				sa[auditioned.GetAudScore()]++
			}
			if auditioned.GetValid() {
				tA++
				if auditioned.GetLastAudition() > 0 {
					tAA++
				}
			}
			if time.Since(time.Unix(auditioned.GetLastAudition(), 0)) < time.Hour*24*7 {
				wCount++
			}
		}
		weekAuditioned.Set(float64(wCount))
		totalToAuditions.Set(float64(tA))
		totalAuditions.Set(float64(tAA))
		for key, val := range sa {
			scoreAuditioned.With(prometheus.Labels{"score": fmt.Sprintf("%v", key)}).Set(float64(val))
		}

		tF := float64(0)
		tUF := float64(0)
		for _, fu := range config.GetFiled() {
			if fu == rcpb.ReleaseMetadata_FILE_UNKNOWN {
				tUF++
			} else {
				tF++
			}
		}
		totalFilled.Set(tF)
		totalUnfilled.Set(tUF)

		rate := float64(0)
		for _, t := range config.GetFiledTime() {
			if time.Since(time.Unix(t, 0)) < time.Hour*24*7 {
				rate++
			}
		}
		rateFiled.Set(rate / 7)

		lax := time.Now().Unix()
		for _, v := range config.GetLbLastTime() {
			if v < lax {
				lax = v
			}
		}
		oldestLBStaged.Set(float64(time.Since(time.Unix(lax, 0)).Seconds()))

		ll := time.Now().Unix()
		unlisten := float64(0)
		listens := make([]int, 0)
		last14 := float64(0)
		today := float64(0)
		for iid, v := range config.GetLastListen() {
			if time.Since(time.Unix(v, 0)) < time.Hour*24*14 {
				last14++
			}
			if time.Since(time.Unix(v, 0)) < time.Hour*24 {
				s.CtxLog(ctx, fmt.Sprintf("FOUND TODAY: %v", iid))
				today++
			}
			if v < ll && v > 0 {
				ll = v
			}
			if v == 0 {
				unlisten++
			} else {
				listens = append(listens, int(v))
			}
		}
		sort.Ints(listens)
		if len(listens) > 0 {
			medianListen.Set(float64(listens[len(listens)/2]))
		}

		lastListen.Set(float64(ll))
		unlistened.Set(unlisten)

		listenTodayRate.Set(float64(today))
		listenRate.Set(float64(last14 / 14))
		listenTotal.Set(float64(len(config.GetLastListen())) / (last14 / 14))

		laxhs := time.Now().Unix()
		for _, v := range config.GetLbLastTimeHigh() {
			if v < laxhs {
				laxhs = v
			}
		}
		oldestLBHigh.Set(float64(time.Since(time.Unix(laxhs, 0)).Seconds()))

		oldestInCollection := int64(math.MaxInt64)
		boing := int32(0)
		for iid, v := range config.GetLastListen() {
			if config.GetCategories()[iid] == rcpb.ReleaseMetadata_IN_COLLECTION {
				if v < int64(oldestInCollection) {
					oldestInCollection = v
					boing = iid
				}
			}
		}
		oldestIC.Set(float64(oldestInCollection))

		s.CtxLog(ctx, fmt.Sprintf("THE OLDEST IC IS %v", boing))
	}()

	if id > 1 {
		rec, err := s.getRecord(ctx, id)
		if err != nil {
			if status.Convert(err).Code() == codes.OutOfRange {
				delete(config.LastListen, id)
				delete(config.LbLastTimeHigh, id)
				delete(config.LbLastTime, id)
				delete(config.Categories, id)
				return s.KSclient.Save(ctx, CONFIG, config)
			}
			return err
		}

		config.IsDirty[id] = rec.GetMetadata().GetDirty()
		config.Weights[id] = rec.GetMetadata().GetWeightInGrams()
		config.Folder[id] = rec.GetRelease().GetFolderId()
		config.Keeps[id] = rec.GetMetadata().GetKeep()
		config.Categories[id] = rec.GetMetadata().GetCategory()
		config.Sleeves[id] = rec.GetMetadata().GetSleeve()
		config.Score[id] = rec.GetRelease().GetRating()
		if rec.GetMetadata().GetSetRating() != 0 {
			if rec.GetMetadata().GetSetRating() < 0 {
				config.Score[id] = 0
			} else {
				config.Score[id] = rec.GetMetadata().GetSetRating()
			}
		}
		s.CtxLog(ctx, fmt.Sprintf("Score (%v) -> %v", id, config.Score[id]))

		vfound := false
		for _, value := range config.GetValues() {
			if rec.Metadata.GetCategory().String() == value.GetCategory() && rec.Metadata.GetFiledUnder().String() == value.GetFilling() {
				vfound = true
				if value.Value == nil {
					value.Value = make(map[int32]float32)
				}
				value.Value[rec.GetRelease().GetInstanceId()] = float32(rec.GetMetadata().GetCurrentSalePrice())
			} else {
				delete(value.GetValue(), rec.GetRelease().GetInstanceId())
			}
		}
		if !vfound {
			val := &pb.Values{
				Category: rec.GetMetadata().GetCategory().String(),
				Filling:  rec.GetMetadata().GetFiledUnder().String(),
				Value:    make(map[int32]float32),
			}
			val.Value[rec.Release.GetInstanceId()] = float32(rec.GetMetadata().GetCurrentSalePrice())
			config.Values = append(config.Values, val)
		}

		if rec.Metadata.GetCategory() != rcpb.ReleaseMetadata_SOLD_ARCHIVE &&
			rec.Metadata.GetCategory() != rcpb.ReleaseMetadata_PARENTS {
			config.LastListen[rec.GetRelease().GetInstanceId()] = rec.GetMetadata().GetLastListenTime()
		} else {
			delete(config.LastListen, rec.GetRelease().GetInstanceId())
		}

		if rec.Release.GetFolderId() == 673768 || rec.Release.GetFolderId() == 3578980 {
			if rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_STAGED {
				config.LbLastTime[rec.GetRelease().GetInstanceId()] = rec.GetMetadata().GetLastListenTime()
				delete(config.LbLastTimeHigh, rec.GetRelease().GetInstanceId())
			} else if rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_HIGH_SCHOOL {
				config.LbLastTimeHigh[rec.GetRelease().GetInstanceId()] = rec.GetMetadata().GetLastListenTime()
				delete(config.LbLastTime, rec.GetRelease().GetInstanceId())
			} else {
				delete(config.LbLastTimeHigh, rec.GetRelease().GetInstanceId())
				delete(config.LbLastTime, rec.GetRelease().GetInstanceId())
			}
		} else {
			delete(config.LbLastTime, rec.GetRelease().GetInstanceId())
			delete(config.LbLastTimeHigh, rec.GetRelease().GetInstanceId())
		}

		exist, ok := config.Filed[id]
		config.Filed[id] = rec.Metadata.GetFiledUnder()
		if rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_SOLD_ARCHIVE {
			delete(config.Filed, id)
			delete(config.FiledTime, id)
		} else {
			if (!ok || exist == rcpb.ReleaseMetadata_FILE_UNKNOWN) && rec.GetMetadata().GetFiledUnder() != rcpb.ReleaseMetadata_FILE_UNKNOWN {
				config.FiledTime[id] = time.Now().Unix()
			}
		}

		found := false
		for _, aud := range config.GetAuditions() {
			if aud.GetInstanceId() == id {
				aud.Valid = rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_IN_COLLECTION &&
					(rec.GetMetadata().GetBoxState() == rcpb.ReleaseMetadata_BOX_UNKNOWN ||
						rec.GetMetadata().GetBoxState() == rcpb.ReleaseMetadata_OUT_OF_BOX)
				aud.LastAudition = rec.GetMetadata().GetLastAudition()
				aud.AudScore = rec.GetMetadata().GetAuditionScore()
				found = true
				break
			}
		}

		if !found {
			config.Auditions = append(config.Auditions, &pb.Auditioned{
				Valid: rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_IN_COLLECTION &&
					(rec.GetMetadata().GetBoxState() == rcpb.ReleaseMetadata_BOX_UNKNOWN ||
						rec.GetMetadata().GetBoxState() == rcpb.ReleaseMetadata_OUT_OF_BOX),
				LastAudition: rec.GetMetadata().GetLastAudition(),
				AudScore:     rec.GetMetadata().GetAuditionScore(),
				InstanceId:   rec.GetRelease().GetInstanceId(),
			})
		}

		if rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_SOLD_ARCHIVE {
			found := false
			for _, entry := range config.GetCompleteSales() {
				if entry.GetInstanceId() == rec.GetRelease().GetInstanceId() {
					found = true
					entry.HasCost = rec.GetMetadata().GetSoldPrice() > 0
				}
			}

			if !found {
				config.CompleteSales = append(config.CompleteSales,
					&pb.CompleteSale{
						InstanceId: rec.GetRelease().GetInstanceId(),
						HasCost:    rec.GetMetadata().GetSoldPrice() > 0,
					})
			}
		} else {
			var centries []*pb.CompleteSale
			for _, entry := range config.GetCompleteSales() {
				if entry.GetInstanceId() != rec.GetRelease().GetInstanceId() {
					centries = append(centries, entry)
				}
			}
			config.CompleteSales = centries
		}

		completes := 0
		fullCompletes := 0
		for _, entry := range config.GetCompleteSales() {
			completes++
			if entry.HasCost {
				fullCompletes++
			}
		}
		totalSales.Set(float64(completes))
		totalCompleteSales.Set(float64(fullCompletes))

		if rec.GetMetadata().GetLastListenTime() < config.GetLastListenTime() && rec.GetMetadata().GetLastListenTime() > 0 {
			config.LastListenTime = rec.GetMetadata().GetLastListenTime()
			oldest.Set(float64(config.LastListenTime))
		}
	}

	return s.KSclient.Save(ctx, CONFIG, config)
}

func (s *Server) computeUnlistenedCDs(ctx context.Context, config *pb.Config) {
	count := 0
	for id, val := range config.GetCategories() {
		if val == rcpb.ReleaseMetadata_UNLISTENED {
			if config.GetFiled()[id] == rcpb.ReleaseMetadata_FILE_CD && config.GetScore()[id] == 0 && config.GetWeights()[id] == 0 && !config.GetIsDirty()[id] {
				s.CtxLog(ctx, fmt.Sprintf("FOUND_CD %v (%v)", id, config.GetScore()[id]))
				count++
			}
		}
	}
	unlistenedCDs.Set(float64(count))
}

func (s *Server) computeUnlistened12s(ctx context.Context, config *pb.Config) {
	count := 0
	for id, val := range config.GetCategories() {
		s.CtxLog(ctx, fmt.Sprintf("%v %v", id, config.GetFiled()[id]))
		if val == rcpb.ReleaseMetadata_UNLISTENED {
			if config.GetFiled()[id] == rcpb.ReleaseMetadata_FILE_12_INCH && config.GetScore()[id] == 0 && !config.GetIsDirty()[id] {
				s.CtxLog(ctx, fmt.Sprintf("FOUND_12 %v (%v)", id, config.GetScore()[id]))
				count++
			}
		}
	}
	unlistened12s.Set(float64(count))
}

func (s *Server) computeUnlistened45s(ctx context.Context, config *pb.Config) {
	count := 0
	for id, val := range config.GetCategories() {
		if val == rcpb.ReleaseMetadata_UNLISTENED || val == rcpb.ReleaseMetadata_ARRIVED {
			if config.GetFiled()[id] == rcpb.ReleaseMetadata_FILE_7_INCH && config.GetScore()[id] == 0 && !config.GetIsDirty()[id] {
				s.CtxLog(ctx, fmt.Sprintf("FOUND_45 %v (%v)", id, config.GetScore()[id]))
				count++
			}
		}

		if config.GetFiled()[id] == rcpb.ReleaseMetadata_FILE_7_INCH {
			s.CtxLog(ctx, fmt.Sprintf("FOUND_45 %v", id))
		}
	}
	unlistened45s.Set(float64(count))
}

func (s *Server) computeUnlistenedDigital(ctx context.Context, config *pb.Config) {
	count := 0
	for id, val := range config.GetCategories() {
		if val == rcpb.ReleaseMetadata_UNLISTENED {
			s.CtxLog(ctx, fmt.Sprintf("FOUND_DIG: %v -> %v", id, config.GetFiled()[id]))
			if config.GetFiled()[id] == rcpb.ReleaseMetadata_FILE_DIGITAL && config.GetScore()[id] == 0 && config.GetWeights()[id] == 0 && !config.GetIsDirty()[id] {
				count++
			}
		}
	}
	unlistenedDigital.Set(float64(count))
}

func (s *Server) computeOldest(ctx context.Context) (err error) {
	folders, err := s.getPhysicalFolders(ctx)
	if err != nil {
		return err
	}
	oldestTime := time.Now().Unix()
	var r *rcpb.Record
	s.CtxLog(ctx, fmt.Sprintf("Folders are %v", folders))

	for _, folder := range folders {
		ids, err := s.getInstanceIds(ctx, folder)
		if err != nil {
			return err
		}

		for _, id := range ids {
			rec, err := s.getRecord(ctx, id)
			if err != nil {
				return err
			}
			if rec.GetMetadata().GetLastListenTime() < oldestTime {
				oldestTime = rec.GetMetadata().GetLastListenTime()
				r = rec
			}
		}
	}

	s.CtxLog(ctx, fmt.Sprintf("Found %v - %v", r.GetRelease().GetInstanceId(), r.GetRelease().GetTitle()))
	oldest.Set(float64(oldestTime))

	return nil
}
