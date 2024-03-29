// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.12.4
// source: policy/v1/policy.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPolicyCreate = "/policy.v1.Policy/Create"
const OperationPolicySayHello = "/policy.v1.Policy/SayHello"

type PolicyHTTPServer interface {
	Create(context.Context, *CreatePolicyRequest) (*CreatePolicyReply, error)
	// SayHello Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterPolicyHTTPServer(s *http.Server, srv PolicyHTTPServer) {
	r := s.Route("/")
	r.GET("/policy/{name}", _Policy_SayHello1_HTTP_Handler(srv))
	r.GET("/policy1/name", _Policy_Create0_HTTP_Handler(srv))
}

func _Policy_SayHello1_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPolicySayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Policy_Create0_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPolicyCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*CreatePolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePolicyReply)
		return ctx.Result(200, reply)
	}
}

type PolicyHTTPClient interface {
	Create(ctx context.Context, req *CreatePolicyRequest, opts ...http.CallOption) (rsp *CreatePolicyReply, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type PolicyHTTPClientImpl struct {
	cc *http.Client
}

func NewPolicyHTTPClient(client *http.Client) PolicyHTTPClient {
	return &PolicyHTTPClientImpl{client}
}

func (c *PolicyHTTPClientImpl) Create(ctx context.Context, in *CreatePolicyRequest, opts ...http.CallOption) (*CreatePolicyReply, error) {
	var out CreatePolicyReply
	pattern := "/policy1/name"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPolicyCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PolicyHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/policy/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPolicySayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
