// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 字节集类_test

import (
	"testing"

	gbinary "github.com/888go/goframe/encoding/gbinary"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_BeEncodeAndBeDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k, v := range testData {
			ve := gbinary.BeEncode(v)
			ve1 := gbinary.BeEncodeByLength(len(ve), v)

						// 使用t.Logf格式化输出，内容为键(k)、值(v)和编码后的值(ve)，格式为："key:value, encoded:value\n". md5:7e3d283f1e0efde3
			switch v.(type) {
			case int:
				t.Assert(gbinary.BeDecodeToInt(ve), v)
				t.Assert(gbinary.BeDecodeToInt(ve1), v)
			case int8:
				t.Assert(gbinary.BeDecodeToInt8(ve), v)
				t.Assert(gbinary.BeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(gbinary.BeDecodeToInt16(ve), v)
				t.Assert(gbinary.BeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(gbinary.BeDecodeToInt32(ve), v)
				t.Assert(gbinary.BeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(gbinary.BeDecodeToInt64(ve), v)
				t.Assert(gbinary.BeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(gbinary.BeDecodeToUint(ve), v)
				t.Assert(gbinary.BeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(gbinary.BeDecodeToUint8(ve), v)
				t.Assert(gbinary.BeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(gbinary.BeDecodeToUint16(ve1), v)
				t.Assert(gbinary.BeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(gbinary.BeDecodeToUint32(ve1), v)
				t.Assert(gbinary.BeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(gbinary.BeDecodeToUint64(ve), v)
				t.Assert(gbinary.BeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(gbinary.BeDecodeToBool(ve), v)
				t.Assert(gbinary.BeDecodeToBool(ve1), v)
			case string:
				t.Assert(gbinary.BeDecodeToString(ve), v)
				t.Assert(gbinary.BeDecodeToString(ve1), v)
			case float32:
				t.Assert(gbinary.BeDecodeToFloat32(ve), v)
				t.Assert(gbinary.BeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(gbinary.BeDecodeToFloat64(ve), v)
				t.Assert(gbinary.BeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := gbinary.BeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_BeEncodeStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := gbinary.BeEncode(user)
		s := gbinary.BeDecodeToString(ve)
		t.Assert(string(s), s)
	})
}
