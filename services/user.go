package services

import (
	"context"
	"fmt"
	"time"

	"github.com/vitor-porfirio/fc3-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	//Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 2)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User:   &pb.User{},
	})

	return nil
}
