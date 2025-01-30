package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GzipCompressor struct{}

func (g *GzipCompressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	writer.Close()
	return buf.Bytes(), nil
}
func (g *GzipCompressor) DeCompress(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func (g *GzipCompressor) Name() string {
	return "gzip"
}

// Gzip 实例
var Gzip = &GzipCompressor{}
