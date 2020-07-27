package main

import (
	"context"
	"testing"

	rcpb "github.com/brotherlogic/recordcollection/proto"
)

func TestUpdate(t *testing.T) {
	s := InitTest()
	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: 12})
	if err != nil {
		t.Errorf("Could not process update: %v", err)
	}
}
