package main

import (
	"context"
	"testing"

	keystoreclient "github.com/brotherlogic/keystore/client"
	rcc "github.com/brotherlogic/recordcollection/client"

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
