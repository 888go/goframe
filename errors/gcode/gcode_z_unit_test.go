// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcode_test//bm:错误码类_test

import (
	"testing"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/test/gtest"
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
