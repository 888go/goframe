// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gutil"
)

func Test_StructToSlice(t *testing.T) {
	type A struct {
		K1 int
		K2 string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		a := &A{
			K1: 1,
			K2: "v2",
		}
		s := 工具类.X结构体到数组(a)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], g.Slice别名{"K1", "K2", 1, "v2"})
		t.AssertIN(s[1], g.Slice别名{"K1", "K2", 1, "v2"})
		t.AssertIN(s[2], g.Slice别名{"K1", "K2", 1, "v2"})
		t.AssertIN(s[3], g.Slice别名{"K1", "K2", 1, "v2"})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 工具类.X结构体到数组(1)
		t.Assert(s, nil)
	})
}

func Test_FillStructWithDefault(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type A struct {
			V1 int     `d:"1.01"`
			V2 string  `d:"1.01"`
			V3 float32 `d:"1.01"`
		}
		a := A{}
		err := 工具类.FillStructWithDefault(&a)
		t.AssertNil(err)

		t.Assert(a.V1, `1`)
		t.Assert(a.V2, `1.01`)
		t.Assert(a.V3, `1.01`)
	})
}
