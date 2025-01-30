package codec

import (
	"github.com/golang/protobuf/proto"
)

type ProtoCodec struct{}

func (p *ProtoCodec) Marshal(v any) ([]byte, error) {
	msg, _ := v.(proto.Message)
	return proto.Marshal(msg)
}

func (p *ProtoCodec) Unmarshal(data []byte, v any) error {
	msg, _ := v.(proto.Message)
	return proto.Unmarshal(data, msg)
}

func (p *ProtoCodec) Name() string {
	return "protobuf"
}

// Proto 实例
var Proto = &ProtoCodec{}
