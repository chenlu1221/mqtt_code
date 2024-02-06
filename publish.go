package messages

import "io"

type Publish struct {
	Header           FixedHeader
	TopicName        string
	PacketIdentifier uint16 //报文标识符
	Payload          PayloadIntFace
}

func (p Publish) Encode(w io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (p Publish) Decode(r io.Reader, hdr FixedHeader, config DecoderConfig) error {
	var err error
	defer func() {
		err = recoverError(err, recover())
	}()
	p.Header = hdr
	//剩余长度
	remainingLength := decodeLength(r)

	p.TopicName = getString(r, &remainingLength)
	if p.Header.QoS.HasId() {
		p.PacketIdentifier = getUint16(r, &remainingLength)
	}
	payloadReader := &io.LimitedReader{R: r, N: int64(remainingLength)}

	if p.Payload, err = config.MakePayload(&p, payloadReader, int(remainingLength)); err != nil {
		return err
	}

	return p.Payload.ReadPayload(payloadReader)
}

func GetPublish() *Publish {
	return &Publish{
		Header: FixedHeader{MessageType: MsgPublish},
	}
}
