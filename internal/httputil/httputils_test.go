// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package httputil_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/httputil"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func TestBuildParams(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"a": "1",
			"b": "2",
		}
		params := httputil.BuildParams(data)
		t.Assert(gstr.X是否包含(params, "a=1"), true)
		t.Assert(gstr.X是否包含(params, "b=2"), true)
	})
	gtest.C(t, func(t *gtest.T) {
		data := g.Map{
			"a": "1",
			"b": nil,
		}
		params := httputil.BuildParams(data)
		t.Assert(gstr.X是否包含(params, "a=1"), true)
		t.Assert(gstr.X是否包含(params, "b="), false)
		t.Assert(gstr.X是否包含(params, "b"), false)
	})
}
