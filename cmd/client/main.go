package client

import (
	"fmt"
	"github.com/renxingdawang/rx-rpc/internal/transport"
	"log"
)

func main() {
	client := transport.NewRPCClient(":9000", 100)

	resp, err := client.Call("echo", []byte("Hello RPC!"))
	if err != nil {
		log.Fatalf("RPC call failed: %v", err)
	}

	fmt.Println("Server response:", string(resp))
}
