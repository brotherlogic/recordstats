package main

import (
	"context"
	"testing"

	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/recordstats/proto"
)

func InitTest() *Server {
	s := Init()
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
