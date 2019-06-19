package main

import (
	"github.com/anakin/gomicro/middleware"
	pb "github.com/anakin/gomicro/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	reg := consul.NewRegistry()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	t, _, err := middleware.NewTracer("test.svc")
	if err != nil {
		log.Fatal("tracer error", err)
	}
	opentracing.InitGlobalTracer(t)
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
		micro.Registry(reg),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(opentracing.GlobalTracer())), // add tracing plugin in to middleware
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
