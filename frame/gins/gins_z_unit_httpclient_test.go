// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins_test

import (
	"fmt"
	"testing"
	
	"github.com/888go/goframe/frame/gins"
	"github.com/888go/goframe/test/gtest"
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
