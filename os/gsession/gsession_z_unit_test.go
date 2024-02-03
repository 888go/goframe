// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gsession

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
)

func Test_NewSessionId(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		id1 := NewSessionId()
		id2 := NewSessionId()
		t.AssertNE(id1, id2)
		t.Assert(len(id1), 32)
	})
}
