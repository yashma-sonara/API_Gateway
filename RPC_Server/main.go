package main

import (
	api "RPC_Server/kitex_gen/api/servicea"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

func serverA(addr *net.TCPAddr, count int) server.Server {
	c := 2 + count - 1
	if c == 4 {
		c = 5
	}
	svr := api.NewServer(
		new(ServiceAImpl),
		server.WithServiceAddr(addr),
		server.WithTracer(prometheus.NewServerTracer(fmt.Sprintf(":909%d", 2+count), fmt.Sprintf("/kitexserverA%d", addr.Port))),
	)
	return svr
}

func serverB(addr *net.TCPAddr, count int) server.Server {
	svr := api.NewServer(
		new(ServiceBImpl),
		server.WithServiceAddr(addr),
	)
	return svr
}

func startServer(serviceName string, port int, f func(*net.TCPAddr, int) server.Server) {

	err1 := registerOnNacos(serviceName, port)

	if err1 != nil {
		log.Fatal("Failed to register server on Nacos:", err1)
	}

	numInstances := 3

	var group sync.WaitGroup
	group.Add(numInstances)

	for i := 0; i < numInstances; i++ {
		go func(instanceID int) {
			defer group.Done()

			addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", int(port+instanceID)))

			svr := f(addr, instanceID)

			err1 := svr.Run()

			if err1 != nil {
				log.Println(err1.Error())
			}
		}(i)
	}

	group.Wait()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer("ServiceA", 8080, serverA)
	go startServer("ServiceB", 8085, serverB)
	wg.Wait()
}
