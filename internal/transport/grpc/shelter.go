package grpc

import (
	"context"
	"shelter-map-bot/proto"
)

func (c *Client) GetFirstNearbyShelter(latitude, longitude float32) (*proto.ShelterResponse, error) {
	res, err := c.api.GetFirstNearbyShelter(context.Background(), &proto.ShelterRequest{
		Latitude:  latitude,
		Longitude: longitude,
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
