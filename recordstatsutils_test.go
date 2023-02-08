package main

import (
	"context"
	"testing"

	gdpb "github.com/brotherlogic/godiscogs/proto"
	keystoreclient "github.com/brotherlogic/keystore/client"
	rcc "github.com/brotherlogic/recordcollection/client"
	rcpb "github.com/brotherlogic/recordcollection/proto"

	pb "github.com/brotherlogic/recordstats/proto"
)

func InitTest() *Server {
	s := Init()
	s.rcclient = &rcc.RecordCollectionClient{Test: true}
	s.SkipLog = true
	s.Testing = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient(".test")
	s.GoServer.KSclient.Save(context.Background(), CONFIG, &pb.Config{})

	return s
}

func TestBasic(t *testing.T) {
	s := InitTest()
	s.computeOldest(context.Background())
}

func TestTrackWeight(t *testing.T) {
	s := InitTest()

	s.rcclient.AddRecord(
		&rcpb.Record{
			Release:  &gdpb.Release{InstanceId: 12345},
			Metadata: &rcpb.ReleaseMetadata{WeightInGrams: 300}})

	s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: 12345})

	data, _, err := s.KSclient.Read(context.Background(), CONFIG, &pb.Config{})
	if err != nil {
		t.Fatalf("Bad read: %v", err)
	}
	config := data.(*pb.Config)

	if val, ok := config.GetWeights()[12345]; !ok {
		t.Errorf("Could not locate weight: %v", config)
	} else {
		if val != 300 {
			t.Errorf("Bad weight: %v from %v", val, config)
		}
	}
}
