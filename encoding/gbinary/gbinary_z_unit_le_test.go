// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gbinary_test
import (
	"testing"
	
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/test/gtest"
	)

func Test_LeEncodeAndLeDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k, v := range testData {
			ve := gbinary.LeEncode(v)
			ve1 := gbinary.LeEncodeByLength(len(ve), v)

			// t.Logf("%s:%v, encoded:%v\n", k, v, ve)
// 使用t.Logf格式化输出信息，其中：
// %s 表示字符串类型，此处代表变量k的值
// %v 表示变量的值，此处代表变量v的值
// %v 同样表示变量的值，此处代表变量ve的值
// 整个输出语句表示：键（k）对应的值是（v），编码后为（ve）
// 翻译成中文注释：
// 使用t.Logf函数打印日志，格式为：键（k）的值为（v），编码后的结果为（ve）
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
