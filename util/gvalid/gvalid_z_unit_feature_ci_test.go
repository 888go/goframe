// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 效验类_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_CI(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := g.X效验类().Data("id").Rules("in:Id,Name").Run(ctx)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		err := g.X效验类().Data("id").Rules("ci|in:Id,Name").Run(ctx)
		t.AssertNil(err)
	})
	gtest.C(t, func(t *gtest.T) {
		err := g.X效验类().Ci().Rules("in:Id,Name").Data("id").Run(ctx)
		t.AssertNil(err)
	})
}
