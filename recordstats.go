package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	ropb "github.com/brotherlogic/recordsorganiser/proto"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	Testing bool
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {

}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{}
}

func (s *Server) getPhysicalFolders(ctx context.Context) ([]int32, error) {
	if s.Testing {
		return []int32{12, 13}, nil
	}
	conn, err := s.NewBaseDial("discovery:///recordsorganiser")
	defer conn.Close()
	if err != nil {
		return []int32{}, err
	}

	client := ropb.NewOrganiserServiceClient(conn)
	res, err := client.GetOrganisation(ctx, &ropb.GetOrganisationRequest{})
	if err != nil {
		return []int32{}, err
	}

	folders := []int32{}
	for _, location := range res.GetLocations() {
		if location.GetMediaType() == ropb.Location_PHYSICAL {
			folders = append(folders, location.GetFolderIds()...)
		}
	}

	return folders, nil
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.DialMaster)
	server.PrepServer()
	server.Register = server

	err := server.RegisterServerV2("recordstats", false, false)
	if err != nil {
		return
	}

	fmt.Printf("%v", server.Serve())
}
