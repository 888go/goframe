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

func Test_BeEncodeAndBeDecode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for k, v := range testData {
			ve := gbinary.BeEncode(v)
			ve1 := gbinary.BeEncodeByLength(len(ve), v)

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
