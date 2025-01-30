package protocol

import (
	"encoding/binary"
	"errors"
)

const (
	HeaderSize      = 16
	MagicNumber     = 0x1314
	ProtocolVersion = 1
)

type MessageType uint8

const (
	MessageTypeRequest MessageType = iota
	MessageTypeResponse
)

type Header struct {
	Magic       uint16
	Version     uint8
	MessageType MessageType
	Compress    uint8
	Serialize   uint8
	RequestID   uint64
	BodyLen     uint32
}

func EncodeHeader(h *Header) []byte {
	buf := make([]byte, HeaderSize)
	binary.BigEndian.PutUint16(buf[0:2], h.Magic)
	buf[2] = h.Version
	buf[3] = byte(h.MessageType)
	buf[4] = h.Compress
	buf[5] = h.Serialize
	binary.BigEndian.PutUint64(buf[6:14], h.RequestID)
	binary.BigEndian.PutUint32(buf[14:18], h.BodyLen)
	return buf
}
func DecodeHeader(data []byte) (*Header, error) {
	if len(data) < HeaderSize {
		return nil, errors.New("invalid header length")
	}
	return &Header{
		Magic:       binary.BigEndian.Uint16(data[0:2]),
		Version:     data[2],
		MessageType: MessageType(data[3]),
		Compress:    data[4],
		Serialize:   data[5],
		RequestID:   binary.BigEndian.Uint64(data[6:14]),
		BodyLen:     binary.BigEndian.Uint32(data[14:18]),
	}, nil
}
