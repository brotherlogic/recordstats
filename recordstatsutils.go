package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
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
	if err == nil {
		s.Log(fmt.Sprintf("Folders = %v", folders))
	}

	return nil
}
