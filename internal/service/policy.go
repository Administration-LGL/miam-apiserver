package service

import (
	"context"

	pb "miam-apiserver/api/policy/v1"
	"miam-apiserver/internal/biz"
)

type PolicyService struct {
	pb.UnimplementedPolicyServer
	uc *biz.PolicyUsecase
}

func NewPolicyService(uc *biz.PolicyUsecase) *PolicyService {
	return &PolicyService{uc: uc}
}

func (s *PolicyService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}

func (s *PolicyService) Create(ctx context.Context, req *pb.CreatePolicyRequest) (*pb.CreatePolicyReply, error) {

	return &pb.CreatePolicyReply{}, nil
}
