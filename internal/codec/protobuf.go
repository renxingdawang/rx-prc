package codec

import "github.com/golang/protobuf/proto"

type ProtobufCodec struct{}

func (c *ProtobufCodec) Encode(msg interface{}) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}
func (c *ProtobufCodec) Decode(data []byte, msg interface{}) error {
	return proto.Unmarshal(data, msg.(proto.Message))
}

func (c *ProtobufCodec) Name() string {
	return "protobuf"
}
