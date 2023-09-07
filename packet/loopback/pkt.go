package loopback

import (
	"github.com/MaxAFriedrich/go.pkt/packet"
)

type Packet struct {
	Content     int32
	pkt_payload packet.Packet `cmp:"skip" string:"skip"`
}

func (p *Packet) Unpack(buf *packet.Buffer) error {
	var content int32
	buf.ReadL(&content)

	p.Content = content
	return nil
}

func (p *Packet) Answers(other packet.Packet) bool {
	return false
}

func (p *Packet) Equals(other packet.Packet) bool {
	return packet.Compare(p, other)
}

func (p *Packet) GetLength() uint16 {
	if p.pkt_payload != nil {
		return p.pkt_payload.GetLength() + 20
	}

	return 20
}

func (p *Packet) GetType() packet.Type {
	return packet.None
}

func (p *Packet) GuessPayloadType() packet.Type {
	switch p.Content {
	case 2:
		return packet.IPv4
	case 24, 28, 30:
		return packet.IPv6
	default:
		return packet.Raw
	}
}

func (p *Packet) InitChecksum(csum uint32) {
}

func (p *Packet) Pack(buf *packet.Buffer) error {
	return nil
}

func (p *Packet) Payload() packet.Packet {
	return p.pkt_payload
}

func (p *Packet) SetPayload(pl packet.Packet) error {
	p.pkt_payload = pl

	return nil
}

func (p *Packet) String() string {
	return packet.Stringify(p)
}
