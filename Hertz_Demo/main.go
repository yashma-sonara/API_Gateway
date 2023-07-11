// Package main provides an API Gateway server that handles
// incoming requests and performs operations such as request validation,
// protocol translation, service discovery, load balancing and request forwarding.
//
// The main package sets up a Hertz server and defines
// a handler function, `decode`, to handle incoming requests. The `decode`
// function performs the necessary operations to handle and forward the request and
// returns the response or an error if any operation fails.
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
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

var (
	ge  generic.Generic
	lb  loadbalance.Loadbalancer
	reg discovery.Resolver
)

// validateContentType checks if the content type of ctx is valid.
// It returns true if the content type is invalid, otherwise false.
func invalidContentType(ctx *app.RequestContext) bool {
	return string(ctx.ContentType()) != "application/json"
}

// readPath reads the ctx request path and returns an array
// of the path splitted into segments by "/"
func readPath(ctx *app.RequestContext) []string {
	path := ctx.Request.Path()
	pathStr := string(path)
	splitArr := strings.Split(pathStr, "/")
	return splitArr
}

// parseRequestBody parses the request body into a JSON map using Unmarshal.
// It returns the parsed JSON map and an error if parsing fails.
func parseRequestBody(body []byte) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(body, &jsonMap)
	return jsonMap, err
}

// getServiceName returns the service name requested from arr.
//
// Index out of range error if length of arr is lesser than 2.
func getServiceName(arr []string) string {
	return arr[1]
}

// getMethod returns the method name requested from  arr.
//
// Index out of range error if length of arr is lesser than 3.
func getMethod(arr []string) string {
	return arr[2]
}

// createNacosRegistry creates a new default Nacos service resolver.
// It returns the Nacos resolver and an error if creation fails.
func createNacosRegistry() (discovery.Resolver, error) {
	return resolver.NewDefaultNacosResolver()
}

// resolveService resolves serviceName and c using the Nacos resolver re.
// It returns the result of the service discovery process and an error if resolution fails.
func resolveService(c context.Context, re discovery.Resolver, serviceName string) (discovery.Result, error) {
	return re.Resolve(c, serviceName)
}

// checkInstances prints the number of instances found in result to log.
func checkInstances(result discovery.Result) {
	instances := result.Instances
	log.Println("Number of instances found:", len(instances))
	for index, instance := range instances {
		log.Println("Instance", index, "at address", instance.Address())
	}
}

// translateThrift creates a new JSON to Thrift Generic.
// The provided IDL file is at path ../RPC_Server/serviceA.thrift
// It returns the translated Generic object and an error if file parse or translation fails.
func translateThrift() (generic.Generic, error) {
	p, err := generic.NewThriftFileProvider("../RPC_Server/serviceA.thrift")
	if err != nil {
		return nil, err
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// makeGenericCall performs generic call on c and body,
// using the generic client cli to the specified method.
// It returns the response from the call and an error if the call fails.
func makeGenericCall(c context.Context, cli genericclient.Client, method string, body string) (interface{}, error) {
	return cli.GenericCall(c, method, body)
}

// initialise initialises the global variables before starting the server.
// It returns any error when calling other functions create new instances.
func initialise() error {
	var err error
	ge, err = translateThrift()
	if err != nil {
		return err
	}
	lb = loadbalance.NewWeightedRandomBalancer()
	reg, err = createNacosRegistry()
	if err != nil {
		return err
	}
	return nil
}

// decode handles the incoming request and performs the necessary operations.
// It validates the context ctx, parses the request body, discover the service,
// and makes a generic call with load balancer. Finally, it returns the response in JSON, or an error if any operation fails.
func decode(c context.Context, ctx *app.RequestContext) {
	if invalidContentType(ctx) {
		log.Println("Invalid Content-Type:", string(ctx.ContentType()))
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid Content-Type, expected application/json")
		return
	}

	splitArr := readPath(ctx)
	if len(splitArr) < 3 {
		log.Println("Invalid URL path:", ctx.Request.Path())
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid URL path")
		return
	}

	serviceName := getServiceName(splitArr)
	method := getMethod(splitArr)

	body, err := ctx.Body()
	if err != nil {
		log.Println("Error reading request body:", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Internal Server Error")
		return
	}

	_, err = parseRequestBody(body)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid JSON data")
		return
	}

	result, err := resolveService(c, reg, serviceName)
	if err != nil {
		log.Println("Error resolving service:", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Error resolving service")
		return
	}

	checkInstances(result)

	cli, err := genericclient.NewClient(serviceName, ge, client.WithResolver(reg), client.WithLoadBalancer(lb))
	if err != nil {
		log.Println("Error creating generic client:", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Error creating generic client")
		return
	}

	resp, err := makeGenericCall(c, cli, method, string(body))
	if err != nil {
		log.Println("Error making generic call:", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Error making generic call")
		return
	}

	fmt.Println("Response:", resp)
	var response map[string]interface{}
	json.Unmarshal([]byte(resp.(string)), &response)
	ctx.JSON(consts.StatusOK, response)
}

// main acts as the entry point of the server application. It sets up a server
// using the Hertz framework and registers the `decode` function as the
// handler for incoming requests. The server listens on 127.0.0.1:8888
// and handles requests for any registered routes.
func main() {
	err := initialise()
	if err != nil {
		panic(err)
	}

	hz := server.Default(
		server.WithHostPorts("127.0.0.1:8888"),
	)

	hz.Any("/", decode)
	hz.NoRoute(decode)
	hz.NoMethod(decode)

	hz.Spin()
}
