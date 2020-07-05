if [ "$PWD" -ef "/home/douglas/go/src/GoPlayground" ];
  then
    protoc -I . --go_out=plugins=grpc:. GoProtoServiceConsignment/proto/consignment/consignment.proto
fi
