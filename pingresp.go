package messages

import "io"

type PingReq struct {
	Header FixedHeader
}

func (c *PingReq) Encode(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (c *PingReq) Decode(r io.Reader, hdr FixedHeader, config DecoderConfig) error {
	//TODO implement me
	panic("implement me")
}

func GetPingReq() *PingReq {
	return &PingReq{
		Header: FixedHeader{MessageType: MsgPingReq},
	}
}
