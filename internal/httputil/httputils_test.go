// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package httputil_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/httputil"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func TestBuildParams(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"a": "1",
			"b": "2",
		}
		params := httputil.BuildParams(data)
		t.Assert(文本类.X是否包含(params, "a=1"), true)
		t.Assert(文本类.X是否包含(params, "b=2"), true)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		data := g.Map{
			"a": "1",
			"b": nil,
		}
		params := httputil.BuildParams(data)
		t.Assert(文本类.X是否包含(params, "a=1"), true)
		t.Assert(文本类.X是否包含(params, "b="), false)
		t.Assert(文本类.X是否包含(params, "b"), false)
	})
}
