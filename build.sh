protoc --proto_path ../../../ -I=./proto --go_out=plugins=grpc:./proto proto/recordstats.proto
mv proto/github.com/brotherlogic/recordstats/proto/* ./proto
