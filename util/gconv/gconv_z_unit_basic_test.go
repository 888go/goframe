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

func Test_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		f32 := float32(123.456)
		i64 := int64(1552578474888)
		t.AssertEQ(转换类.X取整数(f32), int(123))
		t.AssertEQ(转换类.X取整数8位(f32), int8(123))
		t.AssertEQ(转换类.X取整数16位(f32), int16(123))
		t.AssertEQ(转换类.X取整数32位(f32), int32(123))
		t.AssertEQ(转换类.X取整数64位(f32), int64(123))
		t.AssertEQ(转换类.X取整数64位(f32), int64(123))
		t.AssertEQ(转换类.X取正整数(f32), uint(123))
		t.AssertEQ(转换类.X取正整数8位(f32), uint8(123))
		t.AssertEQ(转换类.X取正整数16位(f32), uint16(123))
		t.AssertEQ(转换类.X取正整数32位(f32), uint32(123))
		t.AssertEQ(转换类.X取正整数64位(f32), uint64(123))
		t.AssertEQ(转换类.X取小数32位(f32), float32(123.456))
		t.AssertEQ(转换类.X取小数64位(i64), float64(i64))
		t.AssertEQ(转换类.X取布尔(f32), true)
		t.AssertEQ(转换类.String(f32), "123.456")
		t.AssertEQ(转换类.String(i64), "1552578474888")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		s := "-0xFF"
		t.Assert(转换类.X取整数(s), int64(-0xFF))
	})
}

func Test_Duration(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		d := 转换类.X取时长("1s")
		t.Assert(d.String(), "1s")
		t.Assert(d.Nanoseconds(), 1000000000)
	})
}

func Test_ConvertWithRefer(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertEQ(转换类.X按参考值类型转换("1", 100), 1)
		t.AssertEQ(转换类.X按参考值类型转换("1.01", 1.111), 1.01)
		t.AssertEQ(转换类.X按参考值类型转换("1.01", "1.111"), "1.01")
		t.AssertEQ(转换类.X按参考值类型转换("1.01", false), true)
		t.AssertNE(转换类.X按参考值类型转换("1.01", false), false)
	})
}
