package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/renxingdawang/rx-rpc/internal/codec"
	"github.com/renxingdawang/rx-rpc/internal/compress"
)

// 编码器接口，定义了协议的编码方法
type Encoder interface {
	Encode(header *Header, data interface{}) ([]byte, error) // 编码
}

// ProtocolEncoder 实现了 Encoder 接口
type ProtocolEncoder struct{}

func (e *ProtocolEncoder) Encode(header *Header, data interface{}) ([]byte, error) {
	// 1. 编码数据
	var codecInstance codec.Codec
	switch header.Codec {
	case "json":
		codecInstance = codec.JSON
	case "protobuf":
		codecInstance = codec.Proto
	default:
		return nil, errors.New("unsupported codec")
	}

	encodedData, err := codecInstance.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 2. 压缩数据
	var compressorInstance compress.Compress
	switch header.Compress {
	case "gzip":
		compressorInstance = compress.Gzip
	case "snappy":
		compressorInstance = compress.Snappy
	default:
		return nil, errors.New("unsupported compressor")
	}

	compressedData, err := compressorInstance.Compress(encodedData)
	if err != nil {
		return nil, err
	}

	// 3. 构建协议包（包含头和数据）
	var buffer bytes.Buffer
	err = binary.Write(&buffer, binary.BigEndian, header.Length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buffer, binary.BigEndian, header.MsgID)
	if err != nil {
		return nil, err
	}

	// 返回编码后的数据包
	return append(buffer.Bytes(), compressedData...), nil
}
