package example_tl

import (
	"github.com/qingsong-he/mtprotobuf"
)

var (
	CRC32_TL_something_layer128  uint32 = 0xfdaefd46
	CRC32_TL_something1_layer128 uint32 = 0x52d9dbe0
)

func init() {

	if _, has := mtprotobuf.DefaultDecodeMap[CRC32_TL_something_layer128]; !has {
		mtprotobuf.DefaultDecodeMap[CRC32_TL_something_layer128] = func(m *mtprotobuf.DecodeBuf) mtprotobuf.TL {
			return NewTL_something_layer128().Decode(m)
		}
	} else {
		panic(mtprotobuf.TLErrByTLRepeatReg)
	}

	if _, has := mtprotobuf.DefaultDecodeMap[CRC32_TL_something1_layer128]; !has {
		mtprotobuf.DefaultDecodeMap[CRC32_TL_something1_layer128] = func(m *mtprotobuf.DecodeBuf) mtprotobuf.TL {
			return NewTL_something1_layer128().Decode(m)
		}
	} else {
		panic(mtprotobuf.TLErrByTLRepeatReg)
	}

}

// begin of 'something#ff544e65 flags:# f1:int f2:long f3:double f4:string f5:Foobar f6:Vector<int> f7:Vector<long> f8:Vector<double> f9:Vector<string> f10:Vector<Foobar> f11:flags.0?int f12:flags.1?long f13:flags.2?double f14:flags.3?string f14_1:flags.11?true f15:flags.4?Foobar f15_1:flags.12?true f16:flags.5?Vector<int> f17:flags.6?Vector<long> f18:flags.7?Vector<double> f19:flags.8?Vector<string> f20:flags.9?Vector<Foobar> f21:flags.10?true = SomeThing;'
type TL_something_layer128 struct {
	Flags int32
	F1    int32
	F2    int64
	F3    float64
	F4    string
	F5    mtprotobuf.TL
	F6    []int32
	F7    []int64
	F8    []float64
	F9    []string
	F10   []mtprotobuf.TL
	F11   int32
	F12   int64
	F13   float64
	F14   string
	F14_1 bool
	F15   mtprotobuf.TL
	F15_1 bool
	F16   []int32
	F17   []int64
	F18   []float64
	F19   []string
	F20   []mtprotobuf.TL
	F21   bool
}

func NewTL_something_layer128() *TL_something_layer128 {
	return new(TL_something_layer128)
}

func (*TL_something_layer128) CRC32() uint32 {
	return CRC32_TL_something_layer128
}

func (*TL_something_layer128) GetLayer() int32 {
	return 128
}

func (t *TL_something_layer128) Encode() []byte {
	x := mtprotobuf.NewEncodeBuf(256)
	x.UInt(CRC32_TL_something_layer128)
	// f14_1:flags.11?true
	if t.F14_1 {
		t.Flags |= 2048
	} else {
		t.Flags &^= 2048
	}

	// f15_1:flags.12?true
	if t.F15_1 {
		t.Flags |= 4096
	} else {
		t.Flags &^= 4096
	}

	// f21:flags.10?true
	if t.F21 {
		t.Flags |= 1024
	} else {
		t.Flags &^= 1024
	}

	x.Int(t.Flags)
	// f1:int
	x.Int(t.F1)
	// f2:long
	x.Long(t.F2)
	// f3:double
	x.Double(t.F3)
	// f4:string
	x.String(t.F4)
	// f5:Foobar
	x.Bytes(t.F5.Encode())
	// f6:Vector<int>
	x.VectorInt(t.F6)
	// f7:Vector<long>
	x.VectorLong(t.F7)
	// f8:Vector<double>
	x.VectorDouble(t.F8)
	// f9:Vector<string>
	x.VectorString(t.F9)
	// f10:Vector<Foobar>
	x.Vector(t.F10)
	// f11:flags.0?int
	if (t.Flags & 1) != 0 {
		x.Int(t.F11)
	}

	// f12:flags.1?long
	if (t.Flags & 2) != 0 {
		x.Long(t.F12)
	}

	// f13:flags.2?double
	if (t.Flags & 4) != 0 {
		x.Double(t.F13)
	}

	// f14:flags.3?string
	if (t.Flags & 8) != 0 {
		x.String(t.F14)
	}

	// f15:flags.4?Foobar
	if (t.Flags & 16) != 0 {
		x.Bytes(t.F15.Encode())
	}

	// f16:flags.5?Vector<int>
	if (t.Flags & 32) != 0 {
		x.VectorInt(t.F16)
	}

	// f17:flags.6?Vector<long>
	if (t.Flags & 64) != 0 {
		x.VectorLong(t.F17)
	}

	// f18:flags.7?Vector<double>
	if (t.Flags & 128) != 0 {
		x.VectorDouble(t.F18)
	}

	// f19:flags.8?Vector<string>
	if (t.Flags & 256) != 0 {
		x.VectorString(t.F19)
	}

	// f20:flags.9?Vector<Foobar>
	if (t.Flags & 512) != 0 {
		x.Vector(t.F20)
	}

	return x.GetBuf()
}

func (t *TL_something_layer128) Decode(d *mtprotobuf.DecodeBuf) mtprotobuf.TL {
	t.Flags = d.Int()
	// f14_1:flags.11?true
	t.F14_1 = (t.Flags & 2048) != 0
	// f15_1:flags.12?true
	t.F15_1 = (t.Flags & 4096) != 0
	// f21:flags.10?true
	t.F21 = (t.Flags & 1024) != 0
	// f1:int
	t.F1 = d.Int()
	// f2:long
	t.F2 = d.Long()
	// f3:double
	t.F3 = d.Double()
	// f4:string
	t.F4 = d.String()
	// f5:Foobar
	t.F5 = d.Object()
	// f6:Vector<int>
	t.F6 = d.VectorInt()
	// f7:Vector<long>
	t.F7 = d.VectorLong()
	// f8:Vector<double>
	t.F8 = d.VectorDouble()
	// f9:Vector<string>
	t.F9 = d.VectorString()
	// f10:Vector<Foobar>
	t.F10 = d.Vector()
	// f11:flags.0?int
	if (t.Flags & 1) != 0 {
		t.F11 = d.Int()
	}

	// f12:flags.1?long
	if (t.Flags & 2) != 0 {
		t.F12 = d.Long()
	}

	// f13:flags.2?double
	if (t.Flags & 4) != 0 {
		t.F13 = d.Double()
	}

	// f14:flags.3?string
	if (t.Flags & 8) != 0 {
		t.F14 = d.String()
	}

	// f15:flags.4?Foobar
	if (t.Flags & 16) != 0 {
		t.F15 = d.Object()
	}

	// f16:flags.5?Vector<int>
	if (t.Flags & 32) != 0 {
		t.F16 = d.VectorInt()
	}

	// f17:flags.6?Vector<long>
	if (t.Flags & 64) != 0 {
		t.F17 = d.VectorLong()
	}

	// f18:flags.7?Vector<double>
	if (t.Flags & 128) != 0 {
		t.F18 = d.VectorDouble()
	}

	// f19:flags.8?Vector<string>
	if (t.Flags & 256) != 0 {
		t.F19 = d.VectorString()
	}

	// f20:flags.9?Vector<Foobar>
	if (t.Flags & 512) != 0 {
		t.F20 = d.Vector()
	}

	return t
}

// end of 'something#ff544e65 flags:# f1:int f2:long f3:double f4:string f5:Foobar f6:Vector<int> f7:Vector<long> f8:Vector<double> f9:Vector<string> f10:Vector<Foobar> f11:flags.0?int f12:flags.1?long f13:flags.2?double f14:flags.3?string f14_1:flags.11?true f15:flags.4?Foobar f15_1:flags.12?true f16:flags.5?Vector<int> f17:flags.6?Vector<long> f18:flags.7?Vector<double> f19:flags.8?Vector<string> f20:flags.9?Vector<Foobar> f21:flags.10?true = SomeThing;'

// begin of 'something1#ff544e65 flags:# f1:int f2:long f3:double f4:string f5:Foobar f6:Vector<int> f7:Vector<long> f8:Vector<double> f9:Vector<string> f10:Vector<Foobar> f11:flags.0?int f12:flags.1?long f13:flags.2?double f14:flags.3?string f14_1:flags.11?true f15:flags.4?Foobar f15_1:flags.12?true f16:flags.5?Vector<int> f17:flags.6?Vector<long> f18:flags.7?Vector<double> f19:flags.8?Vector<string> f20:flags.9?Vector<Foobar> f21:flags.10?true = SomeThing1;'
type TL_something1_layer128 struct {
	Flags int32
	F1    int32
	F2    int64
	F3    float64
	F4    string
	F5    mtprotobuf.TL
	F6    []int32
	F7    []int64
	F8    []float64
	F9    []string
	F10   []mtprotobuf.TL
	F11   int32
	F12   int64
	F13   float64
	F14   string
	F14_1 bool
	F15   mtprotobuf.TL
	F15_1 bool
	F16   []int32
	F17   []int64
	F18   []float64
	F19   []string
	F20   []mtprotobuf.TL
	F21   bool
}

func NewTL_something1_layer128() *TL_something1_layer128 {
	return new(TL_something1_layer128)
}

func (*TL_something1_layer128) CRC32() uint32 {
	return CRC32_TL_something1_layer128
}

func (*TL_something1_layer128) GetLayer() int32 {
	return 128
}

func (t *TL_something1_layer128) Encode() []byte {
	x := mtprotobuf.NewEncodeBuf(256)
	x.UInt(CRC32_TL_something1_layer128)
	// f14_1:flags.11?true
	if t.F14_1 {
		t.Flags |= 2048
	} else {
		t.Flags &^= 2048
	}

	// f15_1:flags.12?true
	if t.F15_1 {
		t.Flags |= 4096
	} else {
		t.Flags &^= 4096
	}

	// f21:flags.10?true
	if t.F21 {
		t.Flags |= 1024
	} else {
		t.Flags &^= 1024
	}

	x.Int(t.Flags)
	// f1:int
	x.Int(t.F1)
	// f2:long
	x.Long(t.F2)
	// f3:double
	x.Double(t.F3)
	// f4:string
	x.String(t.F4)
	// f5:Foobar
	x.Bytes(t.F5.Encode())
	// f6:Vector<int>
	x.VectorInt(t.F6)
	// f7:Vector<long>
	x.VectorLong(t.F7)
	// f8:Vector<double>
	x.VectorDouble(t.F8)
	// f9:Vector<string>
	x.VectorString(t.F9)
	// f10:Vector<Foobar>
	x.Vector(t.F10)
	// f11:flags.0?int
	if (t.Flags & 1) != 0 {
		x.Int(t.F11)
	}

	// f12:flags.1?long
	if (t.Flags & 2) != 0 {
		x.Long(t.F12)
	}

	// f13:flags.2?double
	if (t.Flags & 4) != 0 {
		x.Double(t.F13)
	}

	// f14:flags.3?string
	if (t.Flags & 8) != 0 {
		x.String(t.F14)
	}

	// f15:flags.4?Foobar
	if (t.Flags & 16) != 0 {
		x.Bytes(t.F15.Encode())
	}

	// f16:flags.5?Vector<int>
	if (t.Flags & 32) != 0 {
		x.VectorInt(t.F16)
	}

	// f17:flags.6?Vector<long>
	if (t.Flags & 64) != 0 {
		x.VectorLong(t.F17)
	}

	// f18:flags.7?Vector<double>
	if (t.Flags & 128) != 0 {
		x.VectorDouble(t.F18)
	}

	// f19:flags.8?Vector<string>
	if (t.Flags & 256) != 0 {
		x.VectorString(t.F19)
	}

	// f20:flags.9?Vector<Foobar>
	if (t.Flags & 512) != 0 {
		x.Vector(t.F20)
	}

	return x.GetBuf()
}

func (t *TL_something1_layer128) Decode(d *mtprotobuf.DecodeBuf) mtprotobuf.TL {
	t.Flags = d.Int()
	// f14_1:flags.11?true
	t.F14_1 = (t.Flags & 2048) != 0
	// f15_1:flags.12?true
	t.F15_1 = (t.Flags & 4096) != 0
	// f21:flags.10?true
	t.F21 = (t.Flags & 1024) != 0
	// f1:int
	t.F1 = d.Int()
	// f2:long
	t.F2 = d.Long()
	// f3:double
	t.F3 = d.Double()
	// f4:string
	t.F4 = d.String()
	// f5:Foobar
	t.F5 = d.Object()
	// f6:Vector<int>
	t.F6 = d.VectorInt()
	// f7:Vector<long>
	t.F7 = d.VectorLong()
	// f8:Vector<double>
	t.F8 = d.VectorDouble()
	// f9:Vector<string>
	t.F9 = d.VectorString()
	// f10:Vector<Foobar>
	t.F10 = d.Vector()
	// f11:flags.0?int
	if (t.Flags & 1) != 0 {
		t.F11 = d.Int()
	}

	// f12:flags.1?long
	if (t.Flags & 2) != 0 {
		t.F12 = d.Long()
	}

	// f13:flags.2?double
	if (t.Flags & 4) != 0 {
		t.F13 = d.Double()
	}

	// f14:flags.3?string
	if (t.Flags & 8) != 0 {
		t.F14 = d.String()
	}

	// f15:flags.4?Foobar
	if (t.Flags & 16) != 0 {
		t.F15 = d.Object()
	}

	// f16:flags.5?Vector<int>
	if (t.Flags & 32) != 0 {
		t.F16 = d.VectorInt()
	}

	// f17:flags.6?Vector<long>
	if (t.Flags & 64) != 0 {
		t.F17 = d.VectorLong()
	}

	// f18:flags.7?Vector<double>
	if (t.Flags & 128) != 0 {
		t.F18 = d.VectorDouble()
	}

	// f19:flags.8?Vector<string>
	if (t.Flags & 256) != 0 {
		t.F19 = d.VectorString()
	}

	// f20:flags.9?Vector<Foobar>
	if (t.Flags & 512) != 0 {
		t.F20 = d.Vector()
	}

	return t
}

// end of 'something1#ff544e65 flags:# f1:int f2:long f3:double f4:string f5:Foobar f6:Vector<int> f7:Vector<long> f8:Vector<double> f9:Vector<string> f10:Vector<Foobar> f11:flags.0?int f12:flags.1?long f13:flags.2?double f14:flags.3?string f14_1:flags.11?true f15:flags.4?Foobar f15_1:flags.12?true f16:flags.5?Vector<int> f17:flags.6?Vector<long> f18:flags.7?Vector<double> f19:flags.8?Vector<string> f20:flags.9?Vector<Foobar> f21:flags.10?true = SomeThing1;'
