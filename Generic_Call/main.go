package main
import (
        "context"
        "fmt"
        "log"
        "time"

        "github.com/cloudwego/kitex/client"
        "github.com/cloudwego/kitex/client/genericclient"
        "github.com/cloudwego/kitex/pkg/generic"
        "github.com/cloudwego/kitex/server/genericserver"
        "github.com/kitex-contrib/registry-nacos/resolver"
)

type GenericServiceImpl struct{}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
        // use jsoniter or other JSON parsing library to assert request
        m := request.(string)
        fmt.Printf("Recv: %v\n", m)
        return "{\"Msg\": \"world\"}", nil
}

func main() {
        // Create a context
        ctx := context.Background()

        // Parse IDL with Local Files
        p, err := generic.NewThriftFileProvider("./serviceA.thrift")
        if err != nil {
                panic(err)
        }

        g, err := generic.JSONThriftGeneric(p)
        if err != nil {
                panic(err)
        }

        re, err3 := resolver.NewDefaultNacosResolver()
        if err3 != nil {
                log.Println("Error creating new Nacos Resolver:", err3.Error())
        }

        // Create the client
        cli, err := genericclient.NewClient("ServiceA", g, client.WithResolver(re))
        if err != nil {
                panic(err)
        }

        // Create the server
        svr := genericserver.NewServer(new(GenericServiceImpl), g)

        // Start the server
        go func() {
                err := svr.Run()
                if err != nil {
                        panic(err)
                }
        }()

        // Wait for the server to start
        time.Sleep(time.Second)

        // Make a request to the server
        resp, err := cli.GenericCall(ctx, "methodB", "{\"userId\": \"123\", \"message\": \"hello\"}")
        if err != nil {
                panic(err)
        }

        fmt.Println("Response:", resp)
}
