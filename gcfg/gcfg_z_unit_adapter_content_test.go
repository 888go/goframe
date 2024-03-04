// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gcfg_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gcfg"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestAdapterContent_Available_Get_Data(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapter, err := gcfg.NewAdapterContent()
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), false)
	})
	gtest.C(t, func(t *gtest.T) {
		content := `{"a": 1, "b": 2, "c": {"d": 3}}`
		adapter, err := gcfg.NewAdapterContent(content)
		t.AssertNil(err)

		c := gcfg.NewWithAdapter(adapter)
		t.Assert(c.Available(ctx), true)
		t.Assert(c.MustGet(ctx, "a"), 1)
		t.Assert(c.MustGet(ctx, "b"), 2)
		t.Assert(c.MustGet(ctx, "c.d"), 3)
		t.Assert(c.MustGet(ctx, "d"), nil)
		t.Assert(c.MustData(ctx), g.Map{
			"a": 1,
			"b": 2,
			"c": g.Map{
				"d": 3,
			},
		})
	})
}

func TestAdapterContent_SetContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		adapter, err := gcfg.NewAdapterContent()
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), false)

		content := `{"a": 1, "b": 2, "c": {"d": 3}}`
		err = adapter.SetContent(content)
		t.AssertNil(err)
		c := gcfg.NewWithAdapter(adapter)
		t.Assert(c.Available(ctx), true)
		t.Assert(c.MustGet(ctx, "a"), 1)
		t.Assert(c.MustGet(ctx, "b"), 2)
		t.Assert(c.MustGet(ctx, "c.d"), 3)
		t.Assert(c.MustGet(ctx, "d"), nil)
		t.Assert(c.MustData(ctx), g.Map{
			"a": 1,
			"b": 2,
			"c": g.Map{
				"d": 3,
			},
		})
	})

}
