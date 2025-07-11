package main

import (
	"fmt"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	rcc "github.com/brotherlogic/recordcollection/client"
	rcpb "github.com/brotherlogic/recordcollection/proto"
	ropb "github.com/brotherlogic/recordsorganiser/proto"
	pb "github.com/brotherlogic/recordstats/proto"
)

// Server main server type
type Server struct {
	*goserver.GoServer
	rcclient *rcc.RecordCollectionClient
	Testing  bool
}

func cleanAuditions(config *pb.Config) {
	mapper := make(map[int32]*pb.Auditioned)
	for _, aud := range config.Auditions {
		if val, ok := mapper[aud.GetInstanceId()]; ok {
			if aud.GetLastAudition() > val.GetLastAudition() {
				mapper[aud.GetInstanceId()] = aud
			}
		} else {
			mapper[aud.GetInstanceId()] = aud
		}
	}

	var auds []*pb.Auditioned
	for _, aud := range mapper {
		auds = append(auds, aud)
	}
	config.Auditions = auds
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	s.rcclient = &rcc.RecordCollectionClient{Gs: s.GoServer}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	rcpb.RegisterClientUpdateServiceServer(server, s)
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

	conn, err := s.FDialServer(ctx, "recordcollection")
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
		return []int32{12, 13, 14}, nil
	}
	conn, err := s.FDialServer(ctx, "recordsorganiser")
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
	res, err := s.rcclient.GetRecord(ctx, &rcpb.GetRecordRequest{InstanceId: iid})
	if err != nil {
		return nil, err
	}
	return res.GetRecord(), err
}

func main() {
	server := Init()
	server.PrepServer("recordstats")
	server.Register = server

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	/*ctx2, cancel2 := utils.ManualContext("recordbudget-trigger", time.Minute*5)
	server.update(ctx2, 1)
	err = server.cleanCategories(ctx2)
	if err != nil {
		log.Fatalf("Unable to clean categories: %v", err)
	}
	server.CtxLog(ctx2, fmt.Sprintf("Cleaned categories: %v", err))
	cancel2()
	*/
	fmt.Printf("%v", server.Serve())
}
