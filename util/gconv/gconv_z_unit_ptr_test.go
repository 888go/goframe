// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Ptr_Functions(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var v interface{} = 1
		t.AssertEQ(转换类.X取any指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v string = "1"
		t.AssertEQ(转换类.X取文本指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v bool = true
		t.AssertEQ(转换类.X取布尔指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v int = 1
		t.AssertEQ(转换类.X取整数指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v int8 = 1
		t.AssertEQ(转换类.X取整数8位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v int16 = 1
		t.AssertEQ(转换类.X取整数16位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v int32 = 1
		t.AssertEQ(转换类.X取整数32位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v int64 = 1
		t.AssertEQ(转换类.X取整数64位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v uint = 1
		t.AssertEQ(转换类.X取正整数指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v uint8 = 1
		t.AssertEQ(转换类.X取正整数8位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v uint16 = 1
		t.AssertEQ(转换类.X取正整数16位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v uint32 = 1
		t.AssertEQ(转换类.X取正整数32位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v uint64 = 1
		t.AssertEQ(转换类.X取正整数64位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v float32 = 1.01
		t.AssertEQ(转换类.X取小数32位指针(v), &v)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var v float64 = 1.01
		t.AssertEQ(转换类.X取小数64位指针(v), &v)
	})
}
