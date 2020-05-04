package mtprotobuf

import (
	"math/big"
	"reflect"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	enc := NewEncodeBuf(0xff)
	enc.Int(1)
	enc.UInt(2)
	enc.Long(3)
	enc.Double(1.23)
	enc.String("hello world")
	enc.BigInt(big.NewInt(0xff))
	enc.StringBytes([]byte("xyz"))
	enc.Bytes(NewTL_boolTrue_layer0().Encode())
	enc.VectorInt([]int32{1, 2, 3})
	enc.VectorLong([]int64{4, 5, 6})
	enc.VectorString([]string{"a", "b", "c"})
	enc.Vector([]TL{NewTL_boolTrue_layer0(), NewTL_boolFalse_layer0()})

	dec := NewDecodeBuf(enc.GetBuf())
	t.Log(dec.Int() == 1)
	t.Log(dec.UInt() == 2)
	t.Log(dec.Long() == 3)
	t.Log(dec.Double() == 1.23)
	t.Log(dec.String() == "hello world")
	t.Log(dec.BigInt().String() == big.NewInt(0xff).String())
	t.Log(reflect.DeepEqual(dec.StringBytes(), []byte("xyz")))
	t.Log(dec.Object().CRC32() == CRC32_TL_boolTrue_layer0)
	t.Log(reflect.DeepEqual(dec.VectorInt(), []int32{1, 2, 3}))
	t.Log(reflect.DeepEqual(dec.VectorLong(), []int64{4, 5, 6}))
	t.Log(reflect.DeepEqual(dec.VectorString(), []string{"a", "b", "c"}))
	vecByTL := dec.Vector()
	t.Log(len(vecByTL) == 2)
	t.Log(vecByTL[0].CRC32() == CRC32_TL_boolTrue_layer0)
	t.Log(vecByTL[1].CRC32() == CRC32_TL_boolFalse_layer0)
}
