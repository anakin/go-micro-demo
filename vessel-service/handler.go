package main

import (
	"context"
	pb "github.com/anakin/gomicro/vessel-service/proto/vessel"
	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
)

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	var sp opentracing.Span
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	// create new span and bind with context
	sp = opentracing.StartSpan("Find", opentracing.ChildOf(wireContext))
	// record request
	sp.SetTag("req", req)
	defer func() {
		// record response
		sp.SetTag("res", res)
		// before function return stop span, cuz span will counted how much time of this function spent
		sp.Finish()
	}()
	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}
