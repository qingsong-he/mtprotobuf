package example_tl

import (
	"github.com/qingsong-he/mtprotobuf"
	"testing"
)

func TestNewTL_something_layer128(t *testing.T) {
	obj := NewTL_something_layer128()
	obj.F1 = 1
	obj.F2 = 2
	obj.F3 = 3.4
	obj.F4 = "hello"
	obj.F5 = mtprotobuf.NewTL_boolTrue_layer0()
	obj.F6 = []int32{1, 2, 3}
	obj.F7 = []int64{4, 5, 6}
	obj.F8 = []float64{1.1, 2.2, 3.3}
	obj.F9 = []string{"hello1", "hello2", "hello3"}
	obj.F10 = []mtprotobuf.TL{mtprotobuf.NewTL_boolTrue_layer0(), mtprotobuf.NewTL_boolFalse_layer0()}

	obj.Flags |= (1 << 0)
	obj.F11 = 11

	obj.Flags |= (1 << 1)
	obj.F12 = 22

	obj.Flags |= (1 << 2)
	obj.F13 = 33.33

	obj.Flags |= (1 << 3)
	obj.F14 = "world"

	obj.F14_1 = true

	obj.Flags |= (1 << 4)
	obj.F15 = mtprotobuf.NewTL_boolTrue_layer0()

	obj.F15_1 = true

	obj.Flags |= (1 << 5)
	obj.F16 = []int32{11, 22, 33}

	obj.Flags |= (1 << 6)
	obj.F17 = []int64{44, 55, 66}

	obj.Flags |= (1 << 7)
	obj.F18 = []float64{11.11, 22.22, 33.33}

	obj.Flags |= (1 << 8)
	obj.F19 = []string{"hello11", "hello22", "hello33"}

	obj.Flags |= (1 << 9)
	obj.F20 = []mtprotobuf.TL{mtprotobuf.NewTL_boolFalse_layer0(), mtprotobuf.NewTL_boolTrue_layer0()}

	obj.F21 = true

	objBin := obj.Encode()
	t.Log("objBin size:", len(objBin))
	tlObj := mtprotobuf.NewDecodeBuf(objBin).Object()
	t.Logf("%#v", *tlObj.(*TL_something_layer128))

	objBinByGZip := mtprotobuf.GetBufWithGZip(objBin)
	t.Log("objBinByGZip size:", len(objBinByGZip))
	tlObjByGZip := mtprotobuf.NewDecodeBuf(objBinByGZip).Object()
	t.Logf("%#v", *tlObjByGZip.(*TL_something_layer128))
}

func TestNewTL_something1_layer128(t *testing.T) {
	obj := NewTL_something1_layer128()
	obj.F1 = 1
	obj.F2 = 2
	obj.F3 = 3.4
	obj.F4 = "hello"
	obj.F5 = mtprotobuf.NewTL_boolTrue_layer0()
	obj.F6 = []int32{1, 2, 3}
	obj.F7 = []int64{4, 5, 6}
	obj.F8 = []float64{1.1, 2.2, 3.3}
	obj.F9 = []string{"hello1", "hello2", "hello3"}
	obj.F10 = []mtprotobuf.TL{mtprotobuf.NewTL_boolTrue_layer0(), mtprotobuf.NewTL_boolFalse_layer0()}

	obj.Flags |= (1 << 0)
	obj.F11 = 11

	obj.Flags |= (1 << 1)
	obj.F12 = 22

	obj.Flags |= (1 << 2)
	obj.F13 = 33.33

	obj.Flags |= (1 << 3)
	obj.F14 = "world"

	obj.F14_1 = true

	obj.Flags |= (1 << 4)
	obj.F15 = mtprotobuf.NewTL_boolTrue_layer0()

	obj.F15_1 = true

	obj.Flags |= (1 << 5)
	obj.F16 = []int32{11, 22, 33}

	obj.Flags |= (1 << 6)
	obj.F17 = []int64{44, 55, 66}

	obj.Flags |= (1 << 7)
	obj.F18 = []float64{11.11, 22.22, 33.33}

	obj.Flags |= (1 << 8)
	obj.F19 = []string{"hello11", "hello22", "hello33"}

	obj.Flags |= (1 << 9)
	obj.F20 = []mtprotobuf.TL{mtprotobuf.NewTL_boolFalse_layer0(), mtprotobuf.NewTL_boolTrue_layer0()}

	obj.F21 = true

	objBin := obj.Encode()
	t.Log("objBin size:", len(objBin))
	tlObj := mtprotobuf.NewDecodeBuf(objBin).Object()
	t.Logf("%#v", *tlObj.(*TL_something1_layer128))

	objBinByGZip := mtprotobuf.GetBufWithGZip(objBin)
	t.Log("objBinByGZip size:", len(objBinByGZip))
	tlObjByGZip := mtprotobuf.NewDecodeBuf(objBinByGZip).Object()
	t.Logf("%#v", *tlObjByGZip.(*TL_something1_layer128))
}
