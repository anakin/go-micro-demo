package main

import (
	pb "github.com/anakin/gomicro/con-service/proto/consignment"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

const (
	serviceName = "shippy.service.consignment"
)

func main() {
	reg := consul.NewRegistry()
	repo := &ConRepository{}
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Registry(reg),
	)
	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
