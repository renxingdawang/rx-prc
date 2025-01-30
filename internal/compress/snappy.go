package compress

import (
	"github.com/golang/snappy"
)

type SnappyCompressor struct{}

func (s *SnappyCompressor) Compress(data []byte) ([]byte, error) {
	return snappy.Encode(nil, data), nil
}

func (s *SnappyCompressor) DeCompress(data []byte) ([]byte, error) {
	return snappy.Decode(nil, data)
}

func (s *SnappyCompressor) Name() string {
	return "snappy"
}

// Snappy 实例
var Snappy = &SnappyCompressor{}
