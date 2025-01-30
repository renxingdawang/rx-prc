package compress

type Compress interface {
	Compress(data []byte) ([]byte, error)
	DeCompress(data []byte) ([]byte, error)
	Name() string
}
