// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 错误码类_test

import (
	"testing"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/test/gtest"
)

func Test_Case(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(错误码类.CodeNil.String(), "-1")
		t.Assert(错误码类.CodeInternalError.String(), "50:Internal Error")
	})
}

func Test_Nil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c := 错误码类.New(1, "custom error", "detailed description")
		t.Assert(c.Code(), 1)
		t.Assert(c.Message(), "custom error")
		t.Assert(c.Detail(), "detailed description")
	})
}

func Test_WithCode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		c := 错误码类.WithCode(错误码类.CodeInternalError, "CodeInternalError")
		t.Assert(c.Code(), 错误码类.CodeInternalError.Code())
		t.Assert(c.Detail(), "CodeInternalError")
	})
}
