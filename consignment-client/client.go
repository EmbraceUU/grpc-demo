package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/consignment-service/proto/consignment"
	"io/ioutil"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment-client/consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	errJ := json.Unmarshal(data, &consignment)
	if errJ != nil {
		return nil, err
	}
	return &consignment, nil
}

func main() {
	// connect grpc server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("connect err: %s", err.Error()))
	}

	// defer deal connect close
	defer func() {
		errC := conn.Close()
		if errC != nil {
			panic(errC)
		}
	}()

	// get client
	client := pb.NewShippingServiceClient(conn)
	consignment, err := parseFile(defaultFilename)
	if err != nil {
		panic(err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		panic(err)
	}

	println(r.Created)

	result, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		println(err)
	}
	println(fmt.Sprintf("%v", result))
}
