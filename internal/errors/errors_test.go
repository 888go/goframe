// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package errors_test

import (
	"testing"
	
	"github.com/888go/goframe/internal/errors"
	"github.com/888go/goframe/test/gtest"
)

func Test_IsStackModeBrief(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(errors.IsStackModeBrief(), true)
	})
}
