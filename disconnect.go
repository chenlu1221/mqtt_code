package messages

import (
	"io"
)

type Disconnect struct {
	Header FixedHeader
}

func (d Disconnect) Encode(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (d Disconnect) Decode(r io.Reader, hdr FixedHeader, config DecoderConfig) error {
	//TODO implement me
	panic("implement me")
}

func GetDisconnect() *Disconnect {
	return &Disconnect{
		Header: FixedHeader{MessageType: MsgDisconnect},
	}
}
