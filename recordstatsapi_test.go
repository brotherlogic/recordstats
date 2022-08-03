package main

import (
	"context"
	"testing"

	gdpb "github.com/brotherlogic/godiscogs"
	rcpb "github.com/brotherlogic/recordcollection/proto"
)

func TestUpdate(t *testing.T) {
	s := InitTest()
	s.rcclient.AddRecord(&rcpb.Record{Release: &gdpb.Release{InstanceId: 12}, Metadata: &rcpb.ReleaseMetadata{}})
	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: 12})
	if err != nil {
		t.Errorf("Could not process update: %v", err)
	}
}
