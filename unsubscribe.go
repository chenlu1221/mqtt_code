package messages

import (
	"io"
)

type UnSubScribe struct {
	Header           FixedHeader
	PacketIdentifier uint16   //客户端标识符
	TopicList        []string //想要取消的主题列表
}

func (c *UnSubScribe) Encode(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (c *UnSubScribe) Decode(r io.Reader, hdr FixedHeader, config DecoderConfig) error {
	//TODO implement me
	panic("implement me")
}

// GetUnSubScribe 客户端发送UNSUBSCRIBE报文给服务端，用于取消订阅主题。
func GetUnSubScribe() *UnSubScribe {
	return &UnSubScribe{
		Header: FixedHeader{MessageType: MsgUnsubscribe},
	}
}
