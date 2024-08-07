// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package 配置类_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	gcfg "github.com/888go/goframe/os/gcfg"
	gtest "github.com/888go/goframe/test/gtest"
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

		c := gcfg.X创建并按适配器(adapter)
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
	gtest.C(t, func(t *gtest.T) {
		adapter, err := gcfg.NewAdapterContent()
		t.AssertNil(err)
		t.Assert(adapter.Available(ctx), false)

		content := `{"a": 1, "b": 2, "c": {"d": 3}}`
		err = adapter.SetContent(content)
		t.AssertNil(err)
		c := gcfg.X创建并按适配器(adapter)
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
