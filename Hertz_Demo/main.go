package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

func translate(ctx *app.RequestContext) generic.Generic {
	p, err5 := generic.NewThriftFileProvider("../RPC_Server/serviceA.thrift")
	if err5 != nil {
		log.Println("Error", err5.Error())
		ctx.SetStatusCode(http.StatusBadGateway)
		ctx.AbortWithStatus(502)
	}

	g, err6 := generic.JSONThriftGeneric(p)
	if err6 != nil {
		panic(err6)
	}
	return g
}

func decode(c context.Context, ctx *app.RequestContext) {
	if string(ctx.ContentType()) != "application/json" {
		log.Println("Invalid Content-Type:", ctx.ContentType())
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid Content-Type, expected application/json")
		return
	}

	path := ctx.Request.Path()
	pathStr := string(path)
	splitArr := strings.Split(pathStr, "/")
	if len(splitArr) < 3 {
		log.Println("Invalid URL path:", path)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadGateway, "Invalid URL path")
		return
	}

	body, err := ctx.Body()
	if err != nil {
		log.Println("Error reading request body:", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Internal Server Error")
		return
	}

	var jsonMap map[string]interface{}
	err2 := json.Unmarshal(body, &jsonMap)
	if err2 != nil {
		log.Println("Error parsing JSON:", err2)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid JSON data")
		return
	}

	serviceName := splitArr[1]
	method := splitArr[2]

	re, err3 := resolver.NewDefaultNacosResolver()
	if err3 != nil {
		log.Println("Error creating new Nacos Resolver:", err3.Error())
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Error creating new Nacos Resolver")
	}

	// Discover number of instances registered
	// result, err4 := re.Resolve(c, serviceName)
	// if err4 != nil {
	// 	log.Println("Error", err4.Error())
	// }

	// instances := result.Instances
	// fmt.Println("Number of instances found: ", len(instances))
	// for index, instance := range instances {
	// 	fmt.Println("Instance", index, "at address", instance.Address())
	// }

	// Create the generic client
	g := translate(ctx)

	cli, err7 := genericclient.NewClient(serviceName, g, client.WithResolver(re), client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
	if err7 != nil {
		panic(err7)
	}

	fmt.Println(string(body))

	// Make a request to the server
	resp, err8 := cli.GenericCall(c, method, string(body))
	if err8 != nil {
		ctx.JSON(consts.StatusBadRequest, err8.Error())
		return
	}

	fmt.Println("Response:", resp)
	ctx.JSON(consts.StatusOK, resp)
}

func main() {
	hz := server.Default(
		server.WithHostPorts("127.0.0.1:8888"),
	)

	hz.Any("/", decode)
	hz.NoRoute(decode)
	hz.NoMethod(decode)

	hz.Spin()
}
