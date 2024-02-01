// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcode_test
import (
	"testing"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/test/gtest"
	)

func Test_Case(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gcode.CodeNil.String(), "-1")
		t.Assert(gcode.CodeInternalError.String(), "50:Internal Error")
	})
}

func Test_Nil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c := gcode.New(1, "custom error", "detailed description")
		t.Assert(c.Code(), 1)
		t.Assert(c.Message(), "custom error")
		t.Assert(c.Detail(), "detailed description")
	})
}

func Test_WithCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c := gcode.WithCode(gcode.CodeInternalError, "CodeInternalError")
		t.Assert(c.Code(), gcode.CodeInternalError.Code())
		t.Assert(c.Detail(), "CodeInternalError")
	})
}
