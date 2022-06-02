package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vitor-porfirio/fc3-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	//connection with grpc server
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	//close connection when stop using
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Vitor",
		Email: "vhgporfirio@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Vitor",
		Email: "vhgporfirio@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status:", stream.Status, "-", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "v1",
			Name:  "Vitor",
			Email: "vitor@123.com",
		},
		&pb.User{
			Id:    "v2",
			Name:  "Vitor 2",
			Email: "vitor2@123.com",
		},
		&pb.User{
			Id:    "v3",
			Name:  "Vitor 3",
			Email: "vitor3@123.com",
		},
		&pb.User{
			Id:    "v4",
			Name:  "Vitor 4",
			Email: "vitor4@123.com",
		},
		&pb.User{
			Id:    "v5",
			Name:  "Vitor 5",
			Email: "vitor5@123.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		fmt.Println("Sending: ", req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
