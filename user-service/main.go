package main

import (
	pb "github.com/anakin/gomicro/user-service/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {
	reg := consul.NewRegistry()
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("could not connect to DB:%v", err)
	}
	//db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}
	srv := micro.NewService(
		micro.Name("shippy.service.user"),
		micro.Registry(reg),
	)
	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
