package main

import (
	"context"
	"log"

	pb "github.com/DHolmanCoding/GoPlayground/shippy-service-consignment/proto/consignment"
	"github.com/micro/go-micro/v2"
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
	consignments []*pb.Consignment
}

// AddConsignment a new consignment
//
// This is an example of a method that acts on the ConsignmentList struct
// func (<struct-instance-variable-name> <struct-name>) <MethodName(argument Type)> (<return-type(s)>)
func (repo *ConsignmentList) AddConsignment(consignment *pb.Consignment) (*pb.Consignment, error) {
	locallyUpdatedConsignments := append(repo.consignments, consignment)
	repo.consignments = locallyUpdatedConsignments
	return consignment, nil
}

// GetAll consignments
func (repo *ConsignmentList) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the consignmentService
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type consignmentService struct {
	repo repository
}

// CreateConsignment - we created just one method on our consignmentService,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Save our consignment
	consignment, err := s.repo.AddConsignment(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments -
func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	// Initialize an empty repository to store incoming consignments
	repo := &ConsignmentList{}

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)
	// Init will parse the command line flags.
	service.Init()

	// Register service
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
