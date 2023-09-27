package auth

import (
	"API_Gateway/configs"
	"API_Gateway/internal/auth/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *configs.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect auth svc")
		log.Println(err)
	}

	return pb.NewAuthServiceClient(cc)
}
