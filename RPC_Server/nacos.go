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

	instanceCount := 3

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

	return nil
}