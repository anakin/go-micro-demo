package main

import (
	"context"
	pb "github.com/anakin/gomicro/vessel-service/proto/vessel"
)

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}
