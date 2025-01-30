package codec

import "encoding/json"

type JSONCodec struct{}

func (j *JSONCodec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (j *JSONCodec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (j *JSONCodec) Name() string {
	return "json"
}

// JSON 实例
var JSON = &JSONCodec{}
