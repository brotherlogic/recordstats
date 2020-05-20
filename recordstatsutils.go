package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"

	rcpb "github.com/brotherlogic/recordcollection/proto"
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
)

func (s *Server) computeOldest(ctx context.Context) (err error) {
	folders, err := s.getPhysicalFolders(ctx)
	oldestTime := time.Now().Unix()
	var r rcpb.Record
	if err == nil {
		s.Log(fmt.Sprintf("Folders = %v", folders))

		for _, folder := range folders {
			ids, err2 := s.getInstanceIds(ctx, folder)
			err = err2
			if err == nil {
				for _, id := range ids {
					rec, err3 := s.getRecord(ctx, id)
					err = err3
					if rec.GetMetadata().GetLastListenTime() < oldestTime {
						oldestTime = rec.GetMetadata().GetLastListenTime()
						r = rec
					}
				}
			}
		}
	}

	s.Log(fmt.Sprintf("Found %v - %v", r.GetRelease().GetInstanceId(), r.GetRelease().GetTitle()))
	oldest.Set(float64(oldestTime))

	return err
}
