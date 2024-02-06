package messages

import (
	"io"
)

type UnSubAck struct {
	Header           FixedHeader
	PacketIdentifier uint16 //可变报头包含等待确认的UNSUBSCRIBE报文的报文标识符
}

func (c *UnSubAck) Encode(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (c *UnSubAck) Decode(r io.Reader, hdr FixedHeader, config DecoderConfig) error {
	//TODO implement me
	panic("implement me")
}

// GetUnSubAck 服务端发送UNSUBACK报文给客户端用于确认收到UNSUBSCRIBE报文。
func GetUnSubAck() *UnSubAck {
	return &UnSubAck{
		Header: FixedHeader{MessageType: MsgUnsubAck},
	}
}
