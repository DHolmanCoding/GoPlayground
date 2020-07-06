if [ "$PWD" -ef "/home/douglas/go/src/GoPlayground" ];
  then
    protoc --proto_path=. --go_out=. --micro_out=. \
      shippy-service-consignment/proto/consignment/consignment.proto
fi
