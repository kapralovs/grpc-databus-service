package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/kapralovs/grpc-databus-service/databus"
	"google.golang.org/grpc"
)

func main() {
	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	ipAddr := os.Args[1]
	conn, err := grpc.Dial(ipAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := databus.NewDatabusServiceClient(conn)
	res, err := c.Send(context.Background(), &databus.SendRequest{Prm1: float32(x), Prm2: float32(y)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetResult())
}
