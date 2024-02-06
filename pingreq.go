package messages

import "io"

type PingResp struct {
	Header FixedHeader
}

func (c *PingResp) Encode(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (c *PingResp) Decode(r io.Reader, hdr FixedHeader, config DecoderConfig) error {
	//TODO implement me
	panic("implement me")
}

func GetPingResp() *PingResp {
	return &PingResp{
		Header: FixedHeader{MessageType: MsgPingResp},
	}
}
