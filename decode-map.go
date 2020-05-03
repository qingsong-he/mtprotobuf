package mtprotobuf

type TL interface {
	Encode() []byte
	CRC32() uint32
}

var DefaultDecodeMap = make(map[uint32]func(m *DecodeBuf) TL)
