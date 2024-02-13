// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 字节集类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gbinary"
	"github.com/888go/goframe/test/gtest"
)

func Test_BeEncodeAndBeDecode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		for k, v := range testData {
			ve := 字节集类.BeEncode(v)
			ve1 := 字节集类.BeEncodeByLength(len(ve), v)

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
				t.Assert(字节集类.BeDecodeToInt(ve), v)
				t.Assert(字节集类.BeDecodeToInt(ve1), v)
			case int8:
				t.Assert(字节集类.BeDecodeToInt8(ve), v)
				t.Assert(字节集类.BeDecodeToInt8(ve1), v)
			case int16:
				t.Assert(字节集类.BeDecodeToInt16(ve), v)
				t.Assert(字节集类.BeDecodeToInt16(ve1), v)
			case int32:
				t.Assert(字节集类.BeDecodeToInt32(ve), v)
				t.Assert(字节集类.BeDecodeToInt32(ve1), v)
			case int64:
				t.Assert(字节集类.BeDecodeToInt64(ve), v)
				t.Assert(字节集类.BeDecodeToInt64(ve1), v)
			case uint:
				t.Assert(字节集类.BeDecodeToUint(ve), v)
				t.Assert(字节集类.BeDecodeToUint(ve1), v)
			case uint8:
				t.Assert(字节集类.BeDecodeToUint8(ve), v)
				t.Assert(字节集类.BeDecodeToUint8(ve1), v)
			case uint16:
				t.Assert(字节集类.BeDecodeToUint16(ve1), v)
				t.Assert(字节集类.BeDecodeToUint16(ve), v)
			case uint32:
				t.Assert(字节集类.BeDecodeToUint32(ve1), v)
				t.Assert(字节集类.BeDecodeToUint32(ve), v)
			case uint64:
				t.Assert(字节集类.BeDecodeToUint64(ve), v)
				t.Assert(字节集类.BeDecodeToUint64(ve1), v)
			case bool:
				t.Assert(字节集类.BeDecodeToBool(ve), v)
				t.Assert(字节集类.BeDecodeToBool(ve1), v)
			case string:
				t.Assert(字节集类.BeDecodeToString(ve), v)
				t.Assert(字节集类.BeDecodeToString(ve1), v)
			case float32:
				t.Assert(字节集类.BeDecodeToFloat32(ve), v)
				t.Assert(字节集类.BeDecodeToFloat32(ve1), v)
			case float64:
				t.Assert(字节集类.BeDecodeToFloat64(ve), v)
				t.Assert(字节集类.BeDecodeToFloat64(ve1), v)
			default:
				if v == nil {
					continue
				}
				res := make([]byte, len(ve))
				err := 字节集类.BeDecode(ve, res)
				if err != nil {
					t.Errorf("test data: %s, %v, error:%v", k, v, err)
				}
				t.Assert(res, v)
			}
		}
	})
}

func Test_BeEncodeStruct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		user := User{"wenzi1", 999, "www.baidu.com"}
		ve := 字节集类.BeEncode(user)
		s := 字节集类.BeDecodeToString(ve)
		t.Assert(string(s), s)
	})
}
