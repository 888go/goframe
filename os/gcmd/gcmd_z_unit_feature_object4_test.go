		// 版权归GoFrame作者(https:		//goframe.org)所有。保留所有权利。
		//
		// 本源代码形式受MIT许可证条款约束。
		// 如果未随本文件一同分发MIT许可证副本，
		// 您可以在https:		//github.com/gogf/gf处获取。
		// md5:a9832f33b234e3f3

package gcmd_test

import (
	"context"
	"os"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
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
