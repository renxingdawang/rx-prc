package codec

// Codec 编解码器统一接口
type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}

var DefaultCodec Codec = &JSONCodec{}
