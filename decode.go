package mtprotobuf

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
)

var decodeErrBySizeNotRight = errors.New("size not right")
var decodeErrByNotVector = errors.New("not vector type")

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b)}
}

func (m *DecodeBuf) Long() int64 {
	if m.off+8 > m.size {
		panic(decodeErrBySizeNotRight)
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.off+8 > m.size {
		panic(decodeErrBySizeNotRight)
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Int() int32 {

	if m.off+4 > m.size {
		panic(decodeErrBySizeNotRight)
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.off+4 > m.size {
		panic(decodeErrBySizeNotRight)
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.off+size > m.size {
		panic(decodeErrBySizeNotRight)
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x
}

func (m *DecodeBuf) StringBytes() []byte {
	var size, padding int

	if m.off+1 > m.size {
		panic(decodeErrBySizeNotRight)
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			panic(decodeErrBySizeNotRight)
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		panic(decodeErrBySizeNotRight)
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		panic(decodeErrBySizeNotRight)
	}
	m.off += padding

	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	x := string(b)
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *DecodeBuf) VectorInt() []int32 {
	constructor := m.UInt()
	if constructor != CRC32_TL_vector_layer0 {
		panic(decodeErrByNotVector)
	}
	size := m.Int()
	if size < 0 {
		panic(decodeErrBySizeNotRight)

	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorLong() []int64 {
	constructor := m.UInt()
	if constructor != CRC32_TL_vector_layer0 {
		panic(decodeErrByNotVector)
	}
	size := m.Int()
	if size < 0 {
		panic(decodeErrBySizeNotRight)
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorDouble() []float64 {
	constructor := m.UInt()
	if constructor != CRC32_TL_vector_layer0 {
		panic(decodeErrByNotVector)
	}
	size := m.Int()
	if size < 0 {
		panic(decodeErrBySizeNotRight)
	}
	x := make([]float64, size)
	i := int32(0)
	for i < size {
		y := m.Double()
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorString() []string {
	constructor := m.UInt()
	if constructor != CRC32_TL_vector_layer0 {
		panic(decodeErrByNotVector)

	}
	size := m.Int()
	if size < 0 {
		panic(decodeErrBySizeNotRight)

	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Bool() bool {
	constructor := m.UInt()
	switch constructor {
	case CRC32_TL_boolFalse_layer0:
		return false
	case CRC32_TL_boolTrue_layer0:
		return true
	}
	return false
}

func (m *DecodeBuf) Vector() []TL {
	constructor := m.UInt()
	if constructor != CRC32_TL_vector_layer0 {
		panic(decodeErrByNotVector)
	}
	size := m.Int()
	if size < 0 {
		panic(decodeErrBySizeNotRight)
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) Object() (r TL) {
	constructor := m.UInt()

	switch constructor {
	case CRC32_TL_gzipPacked_layer0:
		gz, err := gzip.NewReader(bytes.NewReader(m.StringBytes()))
		if err != nil {
			panic(err)
		}
		obj, err := ioutil.ReadAll(gz)
		if err != nil {
			panic(err)
		}

		err = gz.Close()
		if err != nil {
			panic(err)
		}

		d := NewDecodeBuf(obj)
		r = d.Object()

	default:
		r = m.objectGenerated(constructor)
	}
	return
}

func (m *DecodeBuf) objectGenerated(constructor uint32) (r TL) {
	f, ok := DefaultDecodeMap[constructor]
	if !ok {
		panic(fmt.Errorf("constructor 0x%x not find", constructor))
	} else {
		return f(m)
	}
}
