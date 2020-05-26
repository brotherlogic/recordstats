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

	if time.Now().Sub(time.Unix(oldestTime, 0)) > time.Hour*24*365*2 {
		s.RaiseIssue(ctx, "Haven't Listened", fmt.Sprintf("Haven't listened to %v in a while (since %v)", r.GetRelease().GetInstanceId(), time.Unix(r.GetMetadata().GetLastListenTime(), 0)), false)
	}

	return nil
}
