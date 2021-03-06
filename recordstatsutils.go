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

	rec, err := s.getRecord(ctx, id)
	if err != nil {
		return err
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
	s.Log(fmt.Sprintf("Folders = %v", folders))

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
