package main

import (
	"context"
	"testing"
)

func InitTest() *Server {
	s := Init()
	s.SkipLog = true
	return s
}

func TestBasic(t *testing.T) {
	s := InitTest()
	s.computeOldest(context.Background())
}
