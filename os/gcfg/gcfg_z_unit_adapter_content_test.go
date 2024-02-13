// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package 配置类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/test/gtest"
)

func TestAdapterContent_Available_Get_Data(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		adapter, err := 配置类.NewAdapterContent()
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `{"a": 1, "b": 2, "c": {"d": 3}}`
		adapter, err := 配置类.NewAdapterContent(content)
		t.AssertNil(err)

		c := 配置类.X创建并按适配器(adapter)
		t.Assert(c.X是否可用(ctx), true)
		t.Assert(c.X取值PANI(ctx, "a"), 1)
		t.Assert(c.X取值PANI(ctx, "b"), 2)
		t.Assert(c.X取值PANI(ctx, "c.d"), 3)
		t.Assert(c.X取值PANI(ctx, "d"), nil)
		t.Assert(c.X取MapPANI(ctx), g.Map{
			"a": 1,
			"b": 2,
			"c": g.Map{
				"d": 3,
			},
		})
	})
}

func TestAdapterContent_SetContent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		adapter, err := 配置类.NewAdapterContent()
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), false)

		content := `{"a": 1, "b": 2, "c": {"d": 3}}`
		err = adapter.SetContent(content)
		t.AssertNil(err)
		c := 配置类.X创建并按适配器(adapter)
		t.Assert(c.X是否可用(ctx), true)
		t.Assert(c.X取值PANI(ctx, "a"), 1)
		t.Assert(c.X取值PANI(ctx, "b"), 2)
		t.Assert(c.X取值PANI(ctx, "c.d"), 3)
		t.Assert(c.X取值PANI(ctx, "d"), nil)
		t.Assert(c.X取MapPANI(ctx), g.Map{
			"a": 1,
			"b": 2,
			"c": g.Map{
				"d": 3,
			},
		})
	})

}
