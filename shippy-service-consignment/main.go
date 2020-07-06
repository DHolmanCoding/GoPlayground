package main

import (
	"context"
	"log"
	"net"
	"sync"

	// Import the generated protobuf code
	pb "GoPlayground/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type repository interface {
	AddConsignment(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// ConsignmentList - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
//
// Below is an example of building your own custom type:
//  type <custom-type-name> <existing-type>
type ConsignmentList struct {
	rwMutex      sync.RWMutex // Mutex that allows multiple readers but allows locking for writers
	consignments []*pb.Consignment
}

// AddConsignment a new consignment
//
// This is an example of a method that acts on the ConsignmentList struct
// func (<struct-instance-variable-name> <struct-name>) <MethodName(argument Type)> (<return-type(s)>)
func (repo *ConsignmentList) AddConsignment(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.rwMutex.Lock()
	locallyUpdatedConsignments := append(repo.consignments, consignment)
	repo.consignments = locallyUpdatedConsignments
	repo.rwMutex.Unlock()
	return consignment, nil
}

// GetAll consignments
func (repo *ConsignmentList) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Save our consignment
	consignment, err := s.repo.AddConsignment(req)
	if err != nil {
		return nil, err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

// GetConsignments -
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	// Initialize an empty repository to store incoming consignments
	repo := &ConsignmentList{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceServer(server, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(server)

	log.Println("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}