// Code generated by Kitex v0.6.0. DO NOT EDIT.

package serviceb

import (
	api "RPC_Server/kitex_gen/api"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MethodA(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
	MethodB(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
	MethodC(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kServiceBClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kServiceBClient struct {
	*kClient
}

func (p *kServiceBClient) MethodA(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MethodA(ctx, req)
}

func (p *kServiceBClient) MethodB(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MethodB(ctx, req)
}

func (p *kServiceBClient) MethodC(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MethodC(ctx, req)
}
