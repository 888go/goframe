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

func Test_LeEncodeAndLeDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k, v := range testData {
			ve := gbinary.LeEncode(v)
			ve1 := gbinary.LeEncodeByLength(len(ve), v)

						// 使用t.Logf格式化输出，内容为键(k)、值(v)和编码后的值(ve)，格式为："key:value, encoded:value\n". md5:7e3d283f1e0efde3
			switch v.(type) {
			case int:
				t.Assert(gbinary.LeDecodeToInt(ve), v)
				t.Assert(gbinary.LeDecodeToInt(ve1), v)
			case int8:
				t.Assert(gbinary.LeDecodeToInt8(ve), v)
				t.Assert(gbinary.LeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(gbinary.LeDecodeToInt16(ve), v)
				t.Assert(gbinary.LeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(gbinary.LeDecodeToInt32(ve), v)
				t.Assert(gbinary.LeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(gbinary.LeDecodeToInt64(ve), v)
				t.Assert(gbinary.LeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(gbinary.LeDecodeToUint(ve), v)
				t.Assert(gbinary.LeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(gbinary.LeDecodeToUint8(ve), v)
				t.Assert(gbinary.LeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(gbinary.LeDecodeToUint16(ve1), v)
				t.Assert(gbinary.LeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(gbinary.LeDecodeToUint32(ve1), v)
				t.Assert(gbinary.LeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(gbinary.LeDecodeToUint64(ve), v)
				t.Assert(gbinary.LeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(gbinary.LeDecodeToBool(ve), v)
				t.Assert(gbinary.LeDecodeToBool(ve1), v)
			case string:
				t.Assert(gbinary.LeDecodeToString(ve), v)
				t.Assert(gbinary.LeDecodeToString(ve1), v)
			case float32:
				t.Assert(gbinary.LeDecodeToFloat32(ve), v)
				t.Assert(gbinary.LeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(gbinary.LeDecodeToFloat64(ve), v)
				t.Assert(gbinary.LeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := gbinary.LeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_LeEncodeStruct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := gbinary.LeEncode(user)
		s := gbinary.LeDecodeToString(ve)
		t.Assert(s, s)
	})
}
