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
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

var lb loadbalance.Loadbalancer
var reg discovery.Resolver
var idlFile []string = []string{"../RPC_Server/serviceA.thrift"}
var idlContent = make(map[string]string)
var serviceIdlMap = make(map[string]string)
var serviceClientMap = make(map[string]genericclient.Client)

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

// initIdl initialises the generic call features of the API Gateway based on each file in slice idlFile.
// It returns an error if any process in between fails.
func initIdl() error {
	for _, file := range idlFile {
		err := readIdl(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// readIdl open and read file, map the content to the file path,
// read file line by line to get the services it defines and map the file to services,
// prepare necessary variables for generic call, then map the generic clients to services.
// It returns an error if any process in between fails.
func readIdl(file string) error {
	readFile, err := os.Open(file)
	if err != nil {
		return err
	}

	err = mapContent(file)
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	services := getServices(fileScanner)

	var gen generic.Generic

	if len(services) > 0 {
		gen, err = genericUtil(file)
		if err != nil {
			return err
		}
	}

	for _, service := range services {
		err = mapping(service, file, gen)
		if err != nil {
			return err
		}
	}
	readFile.Close()
	return nil
}

// mapContent store the content of the file and its path in map idlContent.
// It returns an error if file reading fails.
func mapContent(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	idlContent[file] = string(content)
	return nil
}

// mapping habdles the mapping of services and idl and generic client
// It returns an error if fail to create generic client
func mapping(service, file string, gen generic.Generic) error {
	serviceIdlMap[service] = file
	cli, err := genericClient(service, gen)
	if err != nil {
		return err
	}
	serviceClientMap[service] = cli
	return nil
}

// getServices read each line in a file to identify the services being defined in the idl.
// The returned string slice consists of all the service names.
func getServices(fileScanner *bufio.Scanner) []string {
	var services []string
	var line string
	for fileScanner.Scan() {
		line = fileScanner.Text()
		if !strings.HasPrefix(line, "service") {
			continue
		}
		sa := strings.SplitN(line, "service ", 2)
		if len(sa) == 1 {
			continue
		}
		sa = strings.SplitN(sa[1], " {", 2)
		if len(sa) == 1 {
			continue
		}
		services = append(services, sa[0])
	}
	return services
}

// genericProvider creates a new thrift content provider which
// maintains a map between idl paths and contents.
// It returns the pointer to the provider and an error if fails.
func genericProvider(file string) (*generic.ThriftContentWithAbsIncludePathProvider, error) {
	p, err := generic.NewThriftContentWithAbsIncludePathProvider(file, idlContent)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// translateThrift creates a new JSON to Thrift Generic.
// It returns the translated Generic object and an error if translation fails.
func translateThrift(provider *generic.ThriftContentWithAbsIncludePathProvider) (generic.Generic, error) {
	g, err := generic.JSONThriftGeneric(provider)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// genericUtil creates descriptor provider and generic.Generic for the specified idl file.
// It returns generic.Generic and error if anything fails.
func genericUtil(file string) (generic.Generic, error) {
	p, err := genericProvider(file)
	if err != nil {
		return nil, err
	}
	gen, err := translateThrift(p)
	if err != nil {
		return nil, err
	}
	return gen, nil
}

// genericClient creates the genericClient for serviceName using the given generic ge.
// It returns the created generic client and an error if fails.
func genericClient(serviceName string, ge generic.Generic) (genericclient.Client, error) {
	cli, err := genericclient.NewClient(serviceName, ge, client.WithResolver(reg), client.WithLoadBalancer(lb))
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// makeGenericCall performs generic call on c and body,
// using the generic client cli to the specified method.
// It returns the response from the call and an error if the call fails.
func makeGenericCall(c context.Context, cli genericclient.Client, method string, body string) (interface{}, error) {
	return cli.GenericCall(c, method, body)
}

// updateIdl will update the idl mapping of serviceName to the given file.
// It returns an error if fails.
func updateIDL(serviceName, file string) error {
	err := mapContent(file)
	if err != nil {
		return err
	}

	gen, err := genericUtil(file)
	if err != nil {
		return err
	}

	mapping(serviceName, file, gen)
	return nil
}

// initialise initialises the global variables before starting the server.
// It returns any error when calling other functions create new instances.
func initialise() error {
	var err error

	lb = loadbalance.NewWeightedRandomBalancer()
	reg, err = createNacosRegistry()
	if err != nil {
		return err
	}

	err = initIdl()
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

	reqBody, err := parseRequestBody(body)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid JSON data")
		return
	}

	file, ok := reqBody["file"]

	if ok {
		fmt.Println("update idl")
		err = updateIDL(serviceName, file.(string))
	}

	if err != nil {
		log.Println("Error updating IDL", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.String(consts.StatusInternalServerError, "Internal server error, fail to update IDL")
		return
	}

	if ok {
		log.Println("Updated idl of ", serviceName, " to ", file)
		ctx.SetStatusCode(http.StatusAccepted)
		ctx.String(consts.StatusAccepted, "Updated IDL")
		return
	}

	_, ok = serviceClientMap[serviceName]
	if !ok {
		log.Println("Invalid service name:")
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.String(consts.StatusBadRequest, "Invalid service name, service undefined")
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

	resp, err := makeGenericCall(c, serviceClientMap[serviceName], method, string(body))
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
	hz := server.Default(
		server.WithHostPorts("127.0.0.1:8888"),
	)

	err := initialise()
	if err != nil {
		panic(err.Error())
	}

	hz.Any("/", decode)
	hz.NoRoute(decode)
	hz.NoMethod(decode)

	hz.Spin()
}
