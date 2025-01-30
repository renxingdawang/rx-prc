package protocol

// 消息头结构，包含了协议的基本信息
type Header struct {
	Length   uint32 // 消息长度
	MsgID    uint32 // 消息ID
	Codec    string // 编码格式（json、msgpack、protobuf）
	Compress string // 压缩算法（gzip、snappy）
}
