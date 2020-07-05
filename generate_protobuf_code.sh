if [ "$PWD" -ef "/home/douglas/go/src/GoPlayground" ];
  then
    protoc -I . --go_out=plugins=grpc:. go-playground-service-consignment/proto/consignment/consignment.proto
fi
