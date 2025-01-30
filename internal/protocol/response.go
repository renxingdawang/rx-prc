package protocol

import (
	"encoding/json"
	"net"
)

// RPCResponse RPC 响应格式
type RPCResponse struct {
	Result []byte `json:"result"`
	Error  string `json:"error"`
}

// EncodeResponse 发送 RPC 响应
func EncodeResponse(conn net.Conn, result []byte, err error) error {
	resp := RPCResponse{
		Result: result,
	}
	if err != nil {
		resp.Error = err.Error()
	}
	data, _ := json.Marshal(resp)
	_, err = conn.Write(data)
	return err
}

// DecodeResponse 解析 RPC 响应
func DecodeResponse(conn net.Conn) (*RPCResponse, error) {
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		return nil, err
	}
	var resp RPCResponse
	err = json.Unmarshal(data[:n], &resp)
	return &resp, err
}
