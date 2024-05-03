package main

import (
	"context"
	"log"
	"net"

	pb "A4/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	users []*pb.User
}

func (s *server) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	s.users = append(s.users, user)
	return user, nil
}

func (s *server) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	for _, u := range s.users {
		if u.Id == userID.Id {
			return u, nil
		}
	}
	return nil, grpc.Errorf(codes.NotFound, "User not found")
}

func (s *server) ListUsers(ctx context.Context, empty *pb.Empty) (*pb.UserList, error) {
	return &pb.UserList{Users: s.users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("Server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// protoc --go_out=. --go-grpc_out=. user.proto
