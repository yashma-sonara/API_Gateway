package main

import (
	"RPC_Server/kitex_gen/api"
	"context"
	"fmt"
)

// ServiceAImpl implements the last service interface defined in the IDL.
type ServiceAImpl struct{}

// MethodA implements the ServiceAImpl interface.
func (s *ServiceAImpl) MethodA(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	msg := fmt.Sprint("User", req.UserId, " Connected to serviceA, methodA.\nMessage content:", req.Message)
	return &api.Response{Message: msg}, nil
}

// MethodB implements the ServiceAImpl interface.
func (s *ServiceAImpl) MethodB(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	msg := fmt.Sprint("User", req.UserId, " Connected to serviceA, methodB.\nMessage content:", req.Message)
	return &api.Response{Message: msg}, nil
}

// MethodC implements the ServiceAImpl interface.
func (s *ServiceAImpl) MethodC(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	msg := fmt.Sprint("User", req.UserId, " Connected to serviceA, methodB.\nMessage content:", req.Message)
	return &api.Response{Message: msg}, nil
}
