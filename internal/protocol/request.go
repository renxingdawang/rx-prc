package protocol

import (
	"encoding/json"
	"io"
	"net"
)

// RPCRequest 表示 RPC 请求格式
type RPCRequest struct {
	Method  string `json:"method"`
	Payload []byte `json:"payload"`
}

// EncodeRequest 序列化并发送请求
func EncodeRequest(conn net.Conn, method string, payload []byte) error {
	req := RPCRequest{
		Method:  method,
		Payload: payload,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	return err
}

// DecodeRequest 解析 RPC 请求
func DecodeRequest(conn net.Conn) (*RPCRequest, error) {
	data, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	var req RPCRequest
	err = json.Unmarshal(data, &req)
	return &req, err
}
