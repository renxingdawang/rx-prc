package server

import (
	"fmt"
	"github.com/renxingdawang/rx-rpc/internal/registry"
	"github.com/renxingdawang/rx-rpc/internal/transport"
	"log"
)

func main() {
	// 使用 Consul 注册中心
	reg, err := registry.NewConsulRegistry("127.0.0.1:8500")
	if err != nil {
		log.Fatalf("Failed to connect to registry: %v", err)
	}

	server := transport.NewRPCServer(":9000", reg)

	// 注册服务方法
	server.RegisterMethod("echo", func(payload []byte) ([]byte, error) {
		fmt.Println("Received:", string(payload))
		return []byte("Echo: " + string(payload)), nil
	})

	// 启动服务器
	err = server.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
