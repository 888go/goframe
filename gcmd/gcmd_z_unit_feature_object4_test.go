// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcmd_test

import (
	"context"
	"os"
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/888go/goframe/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

type TestNoNameTagCase struct {
	g.Meta `name:"root"`
}

type TestNoNameTagCaseRootInput struct {
	Name string
}

type TestNoNameTagCaseRootOutput struct {
	Content string
}

func (c *TestNoNameTagCase) TEST(ctx context.Context, in TestNoNameTagCaseRootInput) (out *TestNoNameTagCaseRootOutput, err error) {
	out = &TestNoNameTagCaseRootOutput{
		Content: in.Name,
	}
	return
}

func Test_Command_NoNameTagCase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()
		cmd, err := gcmd.NewFromObject(TestNoNameTagCase{})
		t.AssertNil(err)

		os.Args = []string{"root", "TEST", "-name=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}
