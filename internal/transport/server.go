package transport

import (
	"errors"
	"fmt"
	"github.com/renxingdawang/rx-rpc/internal/protocol"
	"github.com/renxingdawang/rx-rpc/internal/registry"
	"net"
	"sync"
)

type RPCServer struct {
	address  string
	listener net.Listener
	registry registry.Registry
	methods  map[string]func([]byte) ([]byte, error)
	mu       sync.RWMutex
}

// NewRPCServer 创建 RPC 服务器
func NewRPCServer(address string, reg registry.Registry) *RPCServer {
	return &RPCServer{
		address:  address,
		registry: reg,
		methods:  make(map[string]func([]byte) ([]byte, error)),
	}
}

// RegisterMethod 注册可调用的方法
func (s *RPCServer) RegisterMethod(name string, handler func([]byte) ([]byte, error)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.methods[name] = handler
}

// Start 启动服务器
func (s *RPCServer) Start() error {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.listener = ln

	// 注册服务
	err = s.registry.Register(&registry.ServiceInstance{Name: "rpc-service", Address: s.address})
	if err != nil {
		return err
	}
	fmt.Println("RPC Server started at", s.address)

	// 处理请求
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

// 处理客户端连接
func (s *RPCServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		req, err := protocol.DecodeRequest(conn)
		if err != nil {
			return
		}

		s.mu.RLock()
		handler, exists := s.methods[req.Method]
		s.mu.RUnlock()

		var response []byte
		if exists {
			response, err = handler(req.Payload)
		} else {
			err = errors.New("method not found")
		}

		protocol.EncodeResponse(conn, response, err)
	}
}

// Stop 关闭服务器
func (s *RPCServer) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
