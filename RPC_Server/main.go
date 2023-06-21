package main

import (
	api "RPC_Server/kitex_gen/api/servicea"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/cloudwego/kitex/server"
)

func main() {

	err1 := registerOnNacos("ServiceA", 8080)

	if err1 != nil {
		log.Fatal("Failed to register server on Nacos:", err1)
	}

	numInstances := 3
	// count := 0

	// Create a wait group to wait for all instances to start
	var group sync.WaitGroup
	group.Add(numInstances)

	// Start multiple instances in Goroutines
	for i := 0; i < numInstances; i++ {
		go func(instanceID int) {
			defer group.Done()

			// fmt.Println(count)
			addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:808%d", instanceID))
			// count++

			// r, err := registry.NewDefaultNacosRegistry()
			// if err != nil {
			// 	panic(err)
			// }

			svr := api.NewServer(
				new(ServiceAImpl),
				server.WithServiceAddr(addr),
				// server.WithRegistry(r),
				// server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "ServiceA"}),
			)

			err1 := svr.Run()

			if err1 != nil {
				log.Println(err1.Error())
			}
		}(i)
	}

	// svr := api.NewServer(
	// 	new(ServiceAImpl))

	// err := svr.Run()
	group.Wait()

}
