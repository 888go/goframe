// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 效验类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

func Test_CI(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := g.X效验类().Data("id").Rules("in:Id,Name").Run(ctx)
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := g.X效验类().Data("id").Rules("ci|in:Id,Name").Run(ctx)
		t.AssertNil(err)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err := g.X效验类().Ci().Rules("in:Id,Name").Data("id").Run(ctx)
		t.AssertNil(err)
	})
}
