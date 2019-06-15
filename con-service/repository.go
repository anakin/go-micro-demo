package main

import (
	pb "github.com/anakin/gomicro/con-service/proto/consignment"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type ConRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}
