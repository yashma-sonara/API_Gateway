package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func decode(c context.Context, ctx *app.RequestContext) {
	if string(ctx.ContentType()) != "application/json" {
		log.Println("Invalid Content-Type:", ctx.ContentType(), "Expected Content-Type: application/json")
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteString("Invalid Content-Type")
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

	serviceName := splitArr[1]
	method := splitArr[2]
	directMethod := ctx.Request.Method()
	fmt.Println("Service name: ", serviceName, "Method", method, "/", directMethod)
}

func main() {
	hz := server.Default(
		server.WithHostPorts("127.0.0.1:8080"),
	)

	hz.GET("/", decode)

	hz.Spin()
}
