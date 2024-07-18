// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins_test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/frame/gins"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Client(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			c  = gins.HttpClient()
			c1 = gins.HttpClient("c1")
			c2 = gins.HttpClient("c2")
		)
		c.SetAgent("test1")
		c.SetAgent("test2")
		t.AssertNE(fmt.Sprintf(`%p`, c), fmt.Sprintf(`%p`, c1))
		t.AssertNE(fmt.Sprintf(`%p`, c), fmt.Sprintf(`%p`, c2))
		t.AssertNE(fmt.Sprintf(`%p`, c1), fmt.Sprintf(`%p`, c2))
	})
}
