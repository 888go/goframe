		// 版权归GoFrame作者(https:		//goframe.org)所有。保留所有权利。
		//
		// 本源代码形式受MIT许可证条款约束。
		// 如果未随本文件一同分发MIT许可证副本，
		// 您可以在https:		//github.com/gogf/gf处获取。
		// md5:a9832f33b234e3f3

package gutil_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gutil"
)

func Test_Copy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gutil.Copy(0), 0)
		t.Assert(gutil.Copy(1), 1)
		t.Assert(gutil.Copy("a"), "a")
		t.Assert(gutil.Copy(nil), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		src := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		dst := gutil.Copy(src)
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
