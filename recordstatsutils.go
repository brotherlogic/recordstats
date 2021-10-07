package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordstats/proto"
)

var (
	oldest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_oldest",
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
	unlistened = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_unlistened",
		Help: "The number of records processed",
	})
)

const (
	// CONFIG - where we store the stats config
	CONFIG = "/github.com/brotherlogic/recordstats/config"
)

func (s *Server) update(ctx context.Context, id int32) error {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if err != nil {
		return err
	}
	config := data.(*pb.Config)

	if config.GetFiled() == nil {
		config.Filed = make(map[int32]rcpb.ReleaseMetadata_FileSize)
	}
	if config.GetFiledTime() == nil {
		config.FiledTime = make(map[int32]int64)
	}
	if config.GetLbLastTime() == nil {
		s.Log("RESTTING LB")
		config.LbLastTime = make(map[int32]int64)
	}
	if config.GetLbLastTimeHigh() == nil {
		s.Log("RESETTING LB HIGH")
		config.LbLastTimeHigh = make(map[int32]int64)
	}
	if config.GetLastListen() == nil {
		s.Log("RESETTING LB HIGH")
		config.LastListen = make(map[int32]int64)
	}

	defer func() {
		tA := 0
		tAA := 0

		for _, auditioned := range config.GetAuditions() {
			if auditioned.GetValid() {
				tA++
				if auditioned.GetLastAudition() > 0 {
					tAA++
				}
			}
		}
		totalToAuditions.Set(float64(tA))
		totalAuditions.Set(float64(tAA))

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
		id := int32(-1)
		for iid, v := range config.GetLbLastTime() {
			if v < lax {
				lax = v
				id = iid
			}
		}
		oldestLBStaged.Set(float64(time.Since(time.Unix(lax, 0)).Seconds()))

		ll := time.Now().Unix()
		idll := int32(-1)
		unlisten := float64(0)
		for iid, v := range config.GetLastListen() {
			if v < ll && v > 0 {
				ll = v
				idll = iid
			}
			if v == 0 {
				unlisten++
			}
		}
		lastListen.Set(float64(time.Since(time.Unix(ll, 0)).Seconds()))
		unlistened.Set(unlisten)

		laxhs := time.Now().Unix()
		idhs := int32(-1)
		for iid, v := range config.GetLbLastTimeHigh() {
			if v < laxhs {
				laxhs = v
				idhs = iid
			}
		}
		oldestLBHigh.Set(float64(time.Since(time.Unix(laxhs, 0)).Seconds()))
		s.Log(fmt.Sprintf("THE OLDEST LB HIGH is %v (%v) but %v (%v) AND (%v), (%v)", idhs, laxhs, id, lax, ll, idll))
	}()

	rec, err := s.getRecord(ctx, id)
	if err != nil {
		return err
	}

	if rec.Metadata.GetCategory() != rcpb.ReleaseMetadata_SOLD_ARCHIVE &&
		rec.Metadata.GetCategory() != rcpb.ReleaseMetadata_PARENTS {
		config.LastListen[rec.GetRelease().GetInstanceId()] = rec.GetMetadata().GetLastListenTime()
	} else {
		delete(config.LastListen, rec.GetRelease().GetInstanceId())
	}

	if rec.Release.GetFolderId() == 673768 {
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
			aud.Valid = rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_IN_COLLECTION
			aud.LastAudition = rec.GetMetadata().GetLastAudition()
			found = true
			break
		}
	}

	if !found {
		config.Auditions = append(config.Auditions, &pb.Auditioned{
			Valid:        rec.GetMetadata().GetCategory() == rcpb.ReleaseMetadata_IN_COLLECTION,
			LastAudition: rec.GetMetadata().GetLastAudition(),
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

	return s.KSclient.Save(ctx, CONFIG, config)
}

func (s *Server) computeOldest(ctx context.Context) (err error) {
	folders, err := s.getPhysicalFolders(ctx)
	if err != nil {
		return err
	}
	oldestTime := time.Now().Unix()
	var r *rcpb.Record
	s.Log(fmt.Sprintf("Folders are %v", folders))

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

	s.Log(fmt.Sprintf("Found %v - %v", r.GetRelease().GetInstanceId(), r.GetRelease().GetTitle()))
	oldest.Set(float64(oldestTime))

	return nil
}
