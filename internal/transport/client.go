package transport

import (
	"errors"
	"github.com/renxingdawang/rx-rpc/internal/protocol"
)

// RPCClient RPC 客户端
type RPCClient struct {
	pool *ConnPool
}

// NewRPCClient 创建 RPC 客户端
func NewRPCClient(address string, maxConns int) *RPCClient {
	return &RPCClient{pool: NewConnPool(address, maxConns)}
}

// Call 发送 RPC 请求
func (c *RPCClient) Call(method string, payload []byte) ([]byte, error) {
	conn, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	defer c.pool.Put(conn)
	err = protocol.EncodeRequest(conn, method, payload)
	if err != nil {
		return nil, err
	}

	resp, err := protocol.DecodeResponse(conn)
	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Result, nil
}

// Close 关闭连接池
func (c *RPCClient) Close() {
	c.pool.Close()
}
