package main

import (
	"testing"
)

func InitTest() *Server {
	s := Init()
	s.SkipLog = true
	return s
}

func TestBasic(t *testing.T) {
	doNothing()
}
