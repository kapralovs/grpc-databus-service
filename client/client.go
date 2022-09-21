package client

import (
	"context"

	"github.com/kapralovs/grpc-databus-service/databus"
)

type GRPCClient struct{}

func (c *GRPCClient) Send(ctx context.Context, req *databus.SendRequest) (*databus.SendResponse, error) {

	return resp, nil
}
