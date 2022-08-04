package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shelter-map-bot/proto"
)

type Client struct {
	api proto.ShelterApiClient
}

func CreateAndConnect(addr string) (*Client, error) {
	fmt.Println(addr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := &Client{
		api: proto.NewShelterApiClient(conn),
	}

	return c, nil
}
