package main

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func registerOnNacos(serviceName string, port int) error {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	serverConfig := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfig,
		},
	)

	if err != nil {
		return err
	}

	// register
	// success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
	// 	Ip:          "127.0.0.1",
	// 	Port:        uint64(port),
	// 	ServiceName: serviceName,
	// 	Weight:      10,
	// 	Enable:      true,
	// 	Healthy:     true,
	// })

	// if !success || err != nil {
	// 	return fmt.Errorf("RegisterServiceInstance failed!" + err.Error())
	// }

	// batch register
	instanceCount := 3

	// Create an instance list for batch registration
	instances := make([]vo.RegisterInstanceParam, instanceCount)
	for i := 0; i < instanceCount; i++ {
		instances[i] = vo.RegisterInstanceParam{
			Ip:          "127.0.0.1",
			Port:        uint64(port + i),
			ServiceName: serviceName,
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
		}
	}

	success, err := namingClient.BatchRegisterInstance(vo.BatchRegisterInstanceParam{
		ServiceName: serviceName,
		Instances:   instances,
	})

	if !success || err != nil {
		return fmt.Errorf("BatchRegisterServiceInstance failed!" + err.Error())
	}

	// check instance count
	// time.Sleep(5 * time.Second)

	// service, err := namingClient.GetService(vo.GetServiceParam{
	// 	ServiceName: serviceName,
	// 	Clusters:    []string{},
	// })
	// if err != nil {
	// 	return fmt.Errorf("Failed to get service: %s", err.Error())
	// }

	// // Print the instance count
	// fmt.Println("Instance count:", len(service.Hosts))

	// Deregister
	// for i := 0; i < 3; i++ {
	// 	namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
	// 		Ip:          "127.0.0.1",
	// 		Port:        uint64(port + i),
	// 		ServiceName: serviceName,
	// 	})
	// }

	return nil
}