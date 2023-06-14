// Code generated by Kitex v0.5.2. DO NOT EDIT.

package servicea

import (
	api "RPC_Server/kitex_gen/api"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return serviceAServiceInfo
}

var serviceAServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "serviceA"
	handlerType := (*api.ServiceA)(nil)
	methods := map[string]kitex.MethodInfo{
		"methodA": kitex.NewMethodInfo(methodAHandler, newServiceAMethodAArgs, newServiceAMethodAResult, false),
		"methodB": kitex.NewMethodInfo(methodBHandler, newServiceAMethodBArgs, nil, true),
		"methodC": kitex.NewMethodInfo(methodCHandler, newServiceAMethodCArgs, newServiceAMethodCResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func methodAHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	err := handler.(api.ServiceA).MethodA(ctx)
	if err != nil {
		return err
	}

	return nil
}
func newServiceAMethodAArgs() interface{} {
	return api.NewServiceAMethodAArgs()
}

func newServiceAMethodAResult() interface{} {
	return api.NewServiceAMethodAResult()
}

func methodBHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	err := handler.(api.ServiceA).MethodB(ctx)
	if err != nil {
		return err
	}

	return nil
}
func newServiceAMethodBArgs() interface{} {
	return api.NewServiceAMethodBArgs()
}

func methodCHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	err := handler.(api.ServiceA).MethodC(ctx)
	if err != nil {
		return err
	}

	return nil
}
func newServiceAMethodCArgs() interface{} {
	return api.NewServiceAMethodCArgs()
}

func newServiceAMethodCResult() interface{} {
	return api.NewServiceAMethodCResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) MethodA(ctx context.Context) (err error) {
	var _args api.ServiceAMethodAArgs
	var _result api.ServiceAMethodAResult
	if err = p.c.Call(ctx, "methodA", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) MethodB(ctx context.Context) (err error) {
	var _args api.ServiceAMethodBArgs
	if err = p.c.Call(ctx, "methodB", &_args, nil); err != nil {
		return
	}
	return nil
}

func (p *kClient) MethodC(ctx context.Context) (err error) {
	var _args api.ServiceAMethodCArgs
	var _result api.ServiceAMethodCResult
	if err = p.c.Call(ctx, "methodC", &_args, &_result); err != nil {
		return
	}
	return nil
}
