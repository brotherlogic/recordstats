package main

import (
	"log"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pbrc "github.com/brotherlogic/recordcollection/proto"
)

func main() {
	ctx2, cancel2 := utils.ManualContext("recordcollectioncli", time.Hour)
	defer cancel2()

	conn, err := utils.LFDialServer(ctx2, "recordstats")
	if err != nil {
		log.Fatalf("Argh: %v", err)
	}
	client := pbrc.NewClientUpdateServiceClient(conn)

	conn2, err := utils.LFDialServer(ctx2, "recordcollection")
	if err != nil {
		log.Fatalf("Cannot reach rc: %v", err)
	}
	defer conn2.Close()

	registry := pbrc.NewRecordCollectionServiceClient(conn2)
	ids, err := registry.QueryRecords(ctx2, &pbrc.QueryRecordsRequest{Query: &pbrc.QueryRecordsRequest_UpdateTime{int64(0)}})
	if err != nil {
		log.Fatalf("Bad query: %v, %v", err, ids)
	}

	for i, id := range ids.GetInstanceIds() {
		log.Printf("PING %v -> %v", i, id)
		_, err = client.ClientUpdate(ctx2, &pbrc.ClientUpdateRequest{InstanceId: int32(id)})
		if err != nil {
			log.Fatalf("Error on GET: %v", err)
		}

	}
}
