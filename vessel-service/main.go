package main

import (
	pb "github.com/anakin/gomicro/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {
	reg := consul.NewRegistry()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
		micro.Registry(reg),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
