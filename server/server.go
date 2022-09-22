package server

import (
	"context"
	"errors"
	"os"

	"github.com/kapralovs/grpc-databus-service/databus"
)

type GRPCServer struct {
	databus.UnimplementedDatabusServiceServer
	Port string
}

func (c *GRPCServer) Send(ctx context.Context, req *databus.SendRequest) (*databus.SendResponse, error) {
	var result float32
	switch os.Args[2] {
	case "mul":
		result = req.GetPrm1() * req.GetPrm2()
	case "div":
		result = req.GetPrm1() / req.GetPrm2()
	case "add":
		result = req.GetPrm1() + req.GetPrm2()
	case "sub":
		result = req.GetPrm1() - req.GetPrm2()
	default:
		return nil, errors.New("no action argument")
	}
	return &databus.SendResponse{Result: result}, nil
}
