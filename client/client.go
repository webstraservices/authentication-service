package client

import (
	"fmt"

	"github.com/webstrasuite/webstra-auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(addr string) pb.AuthServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(conn)
}
