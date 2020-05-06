package mtprotobuf

import "errors"

/*
// layer 0
vector#1cb5c415 {t:Type} # [ t ] = Vector t;
gzipPacked#0x814a3c0 = GZipPacked;
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;
*/

var (
	CRC32_TL_vector_layer0     uint32 = 0x1cb5c415
	CRC32_TL_gzipPacked_layer0 uint32 = 0x814a3c0
	CRC32_TL_boolFalse_layer0  uint32 = 0xbc799737
	CRC32_TL_boolTrue_layer0   uint32 = 0x997275b5
)

var TLErrByTLRepeatReg = errors.New("tl repeat")

func init() {

	if _, has := DefaultDecodeMap[CRC32_TL_boolFalse_layer0]; !has {
		DefaultDecodeMap[CRC32_TL_boolFalse_layer0] = func(m *DecodeBuf) TL {
			return NewTL_boolFalse_layer0().Decode(m)
		}
	} else {
		panic(TLErrByTLRepeatReg)
	}

	if _, has := DefaultDecodeMap[CRC32_TL_boolTrue_layer0]; !has {
		DefaultDecodeMap[CRC32_TL_boolTrue_layer0] = func(m *DecodeBuf) TL {
			return NewTL_boolTrue_layer0().Decode(m)
		}
	} else {
		panic(TLErrByTLRepeatReg)
	}
}

// begin of 'boolFalse#bc799737 = Bool;'
type TL_boolFalse_layer0 struct {
}

func NewTL_boolFalse_layer0() *TL_boolFalse_layer0 {
	return new(TL_boolFalse_layer0)
}

func (*TL_boolFalse_layer0) CRC32() uint32 {
	return CRC32_TL_boolFalse_layer0
}

func (*TL_boolFalse_layer0) GetLayer() int32 {
	return 0
}

func (*TL_boolFalse_layer0) Encode() []byte {
	x := NewEncodeBuf(4)
	x.UInt(CRC32_TL_boolFalse_layer0)
	return x.buf
}

func (tl *TL_boolFalse_layer0) Decode(m *DecodeBuf) TL {
	return tl
}

// end of 'boolFalse#bc799737 = Bool;'

// begin of 'boolTrue#997275b5 = Bool;'

type TL_boolTrue_layer0 struct {
}

func NewTL_boolTrue_layer0() *TL_boolTrue_layer0 {
	return new(TL_boolTrue_layer0)
}

func (*TL_boolTrue_layer0) CRC32() uint32 {
	return CRC32_TL_boolTrue_layer0
}

func (*TL_boolTrue_layer0) GetLayer() int32 {
	return 0
}

func (*TL_boolTrue_layer0) Encode() []byte {
	x := NewEncodeBuf(4)
	x.UInt(CRC32_TL_boolTrue_layer0)
	return x.buf
}

func (tl *TL_boolTrue_layer0) Decode(m *DecodeBuf) TL {
	return tl
}

// end of 'boolTrue#997275b5 = Bool;'
