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

func Test_Copy(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(工具类.X深拷贝(0), 0)
		t.Assert(工具类.X深拷贝(1), 1)
		t.Assert(工具类.X深拷贝("a"), "a")
		t.Assert(工具类.X深拷贝(nil), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		src := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		dst := 工具类.X深拷贝(src)
		t.Assert(dst, src)

		dst.(g.Map)["k3"] = "v3"
		t.Assert(src, g.Map{
			"k1": "v1",
			"k2": "v2",
		})
		t.Assert(dst, g.Map{
			"k1": "v1",
			"k2": "v2",
			"k3": "v3",
		})
	})
}
