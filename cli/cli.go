package main

import (
	"log"
	"os"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pbrc "github.com/brotherlogic/recordcollection/proto"
)

func main() {
	ctx2, cancel2 := utils.ManualContext("recordcollectioncli-"+os.Args[1], time.Hour)
	defer cancel2()

	conn, err := utils.LFDialServer(ctx2, "recordstats")
	if err != nil {
		log.Fatalf("Argh: %v", err)
	}
	client := pbrc.NewClientUpdateServiceClient(conn)

	_, err = client.ClientUpdate(ctx2, &pbrc.ClientUpdateRequest{InstanceId: int32(19866960)})
	if err != nil {
		log.Fatalf("Bad: %v", err)
	} else {
		log.Fatalf("Good: %v", err)
	}

	conn2, err := utils.LFDialServer(ctx2, "recordcollection")
	if err != nil {
		log.Fatalf("Cannot reach rc: %v", err)
	}
	defer conn2.Close()

	registry := pbrc.NewRecordCollectionServiceClient(conn2)
	ids, err := registry.QueryRecords(ctx2, &pbrc.QueryRecordsRequest{Query: &pbrc.QueryRecordsRequest_All{true}})
	if err != nil {
		log.Fatalf("Bad query: %v", err)
	}

	for i, id := range ids.GetInstanceIds() {
		log.Printf("PING %v -> %v", i, id)
		_, err = client.ClientUpdate(ctx2, &pbrc.ClientUpdateRequest{InstanceId: int32(id)})
		if err != nil {
			log.Fatalf("Error on GET: %v", err)
		}

	}
}
