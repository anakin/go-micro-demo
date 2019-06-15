package main

import (
	"context"
	"encoding/json"
	pb "github.com/anakin/gomicro/con-service/proto/consignment"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var con *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(data, &con)
	return con, nil
}

func main() {
	reg := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("shippy.consignment.cli"),
		micro.Registry(reg),
	)
	service.Init()
	client := pb.NewShippingServiceClient("shippy.service.consignment", service.Client())
	file := defaultFilename

	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	con, err := parseFile(file)
	if err != nil {
		log.Fatalf("parse error:%v", err)
	}
	r, err := client.CreateConsignment(context.Background(), con)
	if err != nil {
		log.Fatalf("not greet:%v", err)
	}
	log.Printf("created:%t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
