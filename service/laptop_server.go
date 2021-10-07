package service

import (
	"context"
	"errors"
	"log"

	"grpc-proto/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore;
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}	
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
	) (*pb.CreateLaptopResponse, error) {
		laptop := req.GetLaptop()
		log.Printf("Received a create laptop request: %s", laptop.Id)

		if len(laptop.Id) > 0 {
			_, err := uuid.Parse(laptop.Id)
			if err != nil {
				log.Printf("Cannot create a laptop with an invalid ID: %s", laptop.Id)
				return nil, status.Errorf(codes.InvalidArgument, "Laptop id is not valid: %v", err)
			}
		} else {
			id, err := uuid.NewRandom()
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Cannot generation a new laptop ID: %v", err)
			}
			laptop.Id = id.String()
		}

		err := server.Store.Save(laptop)

		if err != nil {
			code := codes.Internal
			if errors.Is(err, ErrAlreadyExists) {
				code = codes.AlreadyExists
			}
			return nil, status.Errorf(code, "Cannot save laptop to store: %v", err)
		}
		log.Printf("Laptop saved: %s", laptop.Id)
		
		res := &pb.CreateLaptopResponse{
			Id: laptop.Id,
		}

		return res, nil
}