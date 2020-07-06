module GoPlayground/shippy-cli-consignment

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/DHolmanCoding/GoPlayground/shippy-service-consignment v0.0.0-20200706025652-5595d8733894
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/micro/go-micro/v2 v2.9.1
)
