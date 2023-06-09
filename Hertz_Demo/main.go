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
)

func decode(c context.Context, ctx *app.RequestContext) {
	if string(ctx.ContentType()) != "application/json" {
		log.Println("Invalid Content-Type:", ctx.ContentType())
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteString("Invalid Content-Type, expected application/json")
		return
	}
	
	path := ctx.Request.Path()
	pathStr := string(path)
	splitArr := strings.Split(pathStr, "/")
	if len(splitArr) < 3 {
		log.Println("Invalid URL path:", path)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteString("Invalid URL path")
		return
	}
	
	body, err := ctx.Body()
	if err != nil {
		log.Println("Error reading request body:", err)
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.WriteString("Internal Server Error")
		return
	}
	
	var jsonMap map[string]interface{}
	err2 := json.Unmarshal(body, &jsonMap)
	if err2 != nil {
		log.Println("Error parsing JSON:", err2)
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteString("Invalid JSON data")
		return
	}

	serviceName := splitArr[1]
	method := splitArr[2]
	directMethod := ctx.Request.Method()
	
	fmt.Println("Service name: ", serviceName, "Method", method, "/", directMethod)
	fmt.Println(jsonMap)
}

func main() {
	hz := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
	)

	hz.Any("/", decode)
	hz.NoRoute(decode)
	hz.NoMethod(decode)

	hz.Spin()
}
