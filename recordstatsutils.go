package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
)

var (
	oldest = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordstats_oldest",
		Help: "The oldest physical record",
	})
)

func (s *Server) computeOldest(ctx context.Context) {
	time.Sleep(time.Minute)
	oldest.Set(float64(time.Now().Unix()))
}
