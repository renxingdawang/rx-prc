package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/renxingdawang/rx-rpc/internal/codec"
	"github.com/renxingdawang/rx-rpc/internal/compress"
)

// 解码器接口，定义了协议的解码方法
type Decoder interface {
	Decode(data []byte) (header *Header, decodedData interface{}, err error) // 解码
}

// ProtocolDecoder 实现了 Decoder 接口
type ProtocolDecoder struct{}

func (d *ProtocolDecoder) Decode(data []byte) (header *Header, decodedData interface{}, err error) {
	// 1. 解包协议头
	header = &Header{}
	reader := bytes.NewReader(data)

	err = binary.Read(reader, binary.BigEndian, &header.Length)
	if err != nil {
		return nil, nil, err
	}
	err = binary.Read(reader, binary.BigEndian, &header.MsgID)
	if err != nil {
		return nil, nil, err
	}

	// 2. 解压数据
	compressedData := data[8:] // 排除掉消息头部分
	var compressorInstance compress.Compress
	switch header.Compress {
	case "gzip":
		compressorInstance = compress.Gzip
	case "snappy":
		compressorInstance = compress.Snappy
	default:
		return nil, nil, errors.New("unsupported compressor")
	}

	decompressedData, err := compressorInstance.DeCompress(compressedData)
	if err != nil {
		return nil, nil, err
	}

	// 3. 解码数据
	var codecInstance codec.Codec
	switch header.Codec {
	case "json":
		codecInstance = codec.JSON
	case "protobuf":
		codecInstance = codec.Proto
	default:
		return nil, nil, errors.New("unsupported codec")
	}

	// 反序列化数据
	err = codecInstance.Unmarshal(decompressedData, &decodedData)
	if err != nil {
		return nil, nil, err
	}

	return header, decodedData, nil
}
