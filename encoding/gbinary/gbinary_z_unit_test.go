// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 字节集类_test

import (
	"math"
	"testing"

	gbinary "github.com/888go/goframe/encoding/gbinary"
	gtest "github.com/888go/goframe/test/gtest"
)

type User struct {
	Name string
	Age  int
	Url  string
}

var testData = map[string]interface{}{
	//"nil":         nil,
	"int":         int(123),
	"int8":        int8(-99),
	"int8.max":    math.MaxInt8,
	"int16":       int16(123),
	"int16.max":   math.MaxInt16,
	"int32":       int32(-199),
	"int32.max":   math.MaxInt32,
	"int64":       int64(123),
	"uint":        uint(123),
	"uint8":       uint8(123),
	"uint8.max":   math.MaxUint8,
	"uint16":      uint16(9999),
	"uint16.max":  math.MaxUint16,
	"uint32":      uint32(123),
	"uint64":      uint64(123),
	"bool.true":   true,
	"bool.false":  false,
	"string":      "hehe haha",
	"byte":        []byte("hehe haha"),
	"float32":     float32(123.456),
	"float32.max": math.MaxFloat32,
	"float64":     float64(123.456),
}

var testBitData = []int{0, 99, 122, 129, 222, 999, 22322}

func Test_EncodeAndDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k, v := range testData {
			ve := gbinary.Encode(v)
			ve1 := gbinary.EncodeByLength(len(ve), v)

						// 使用t.Logf格式化输出，内容为键(k)、值(v)和编码后的值(ve)，格式为："key:value, encoded:value\n". md5:7e3d283f1e0efde3
			switch v.(type) {
			case int:
				t.Assert(gbinary.DecodeToInt(ve), v)
				t.Assert(gbinary.DecodeToInt(ve1), v)
			case int8:
				t.Assert(gbinary.DecodeToInt8(ve), v)
				t.Assert(gbinary.DecodeToInt8(ve1), v)
			case int16:
				t.Assert(gbinary.DecodeToInt16(ve), v)
				t.Assert(gbinary.DecodeToInt16(ve1), v)
			case int32:
				t.Assert(gbinary.DecodeToInt32(ve), v)
				t.Assert(gbinary.DecodeToInt32(ve1), v)
			case int64:
				t.Assert(gbinary.DecodeToInt64(ve), v)
				t.Assert(gbinary.DecodeToInt64(ve1), v)
			case uint:
				t.Assert(gbinary.DecodeToUint(ve), v)
				t.Assert(gbinary.DecodeToUint(ve1), v)
			case uint8:
				t.Assert(gbinary.DecodeToUint8(ve), v)
				t.Assert(gbinary.DecodeToUint8(ve1), v)
			case uint16:
				t.Assert(gbinary.DecodeToUint16(ve1), v)
				t.Assert(gbinary.DecodeToUint16(ve), v)
			case uint32:
				t.Assert(gbinary.DecodeToUint32(ve1), v)
				t.Assert(gbinary.DecodeToUint32(ve), v)
			case uint64:
				t.Assert(gbinary.DecodeToUint64(ve), v)
				t.Assert(gbinary.DecodeToUint64(ve1), v)
			case bool:
				t.Assert(gbinary.DecodeToBool(ve), v)
				t.Assert(gbinary.DecodeToBool(ve1), v)
			case string:
				t.Assert(gbinary.DecodeToString(ve), v)
				t.Assert(gbinary.DecodeToString(ve1), v)
			case float32:
				t.Assert(gbinary.DecodeToFloat32(ve), v)
				t.Assert(gbinary.DecodeToFloat32(ve1), v)
			case float64:
				t.Assert(gbinary.DecodeToFloat64(ve), v)
				t.Assert(gbinary.DecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := gbinary.Decode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_EncodeStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := gbinary.Encode(user)
		s := gbinary.DecodeToString(ve)
		t.Assert(s, s)
	})
}

func Test_Bits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := range testBitData {
			bits := make([]gbinary.Bit, 0)
			res := gbinary.EncodeBits(bits, testBitData[i], 64)

			t.Assert(gbinary.DecodeBits(res), testBitData[i])
			t.Assert(gbinary.DecodeBitsToUint(res), uint(testBitData[i]))

			t.Assert(gbinary.DecodeBytesToBits(gbinary.EncodeBitsToBytes(res)), res)
		}
	})
}
