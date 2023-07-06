package service

import (
	"context"

	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	gc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(gc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{gc: gc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.gc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
