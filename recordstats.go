package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/brotherlogic/goserver"
	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	rcpb "github.com/brotherlogic/recordcollection/proto"
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

func (s *Server) getInstanceIds(ctx context.Context, folder int32) ([]int32, error) {
	if s.Testing {
		return []int32{12, 13}, nil
	}

	conn, err := s.NewBaseDial("recordcollection")
	if err != nil {
		return []int32{}, nil
	}
	defer conn.Close()

	client := rcpb.NewRecordCollectionServiceClient(conn)
	ids, err := client.QueryRecords(ctx, &rcpb.QueryRecordsRequest{Query: &rcpb.QueryRecordsRequest_FolderId{int32(folder)}})
	if err != nil {
		return []int32{}, nil
	}

	iids := []int32{}
	for _, id := range ids.GetInstanceIds() {
		iids = append(iids, id)
	}

	return iids, nil
}

func (s *Server) getPhysicalFolders(ctx context.Context) ([]int32, error) {
	if s.Testing {
		return []int32{12, 13}, nil
	}
	conn, err := s.NewBaseDial("recordsorganiser")
	if err != nil {
		return []int32{}, err
	}
	defer conn.Close()

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

func (s *Server) getRecord(ctx context.Context, iid int32) (*rcpb.Record, error) {
	if s.Testing {
		return &rcpb.Record{Metadata: &rcpb.ReleaseMetadata{}}, nil
	}
	conn, err := s.NewBaseDial("recordcollection")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := rcpb.NewRecordCollectionServiceClient(conn)
	res, err := client.GetRecord(ctx, &rcpb.GetRecordRequest{InstanceId: iid})
	if err != nil {
		return nil, err
	}

	return res.GetRecord(), nil
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

	server.RegisterRepeatingTask(server.computeOldest, "compute_oldest", time.Hour)

	fmt.Printf("%v", server.Serve())
}
