// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package cmd类_test

import (
	"context"
	"os"
	"testing"

	"github.com/888go/goframe/frame/g"
	gcmd "github.com/888go/goframe/os/gcmd"
	gctx "github.com/888go/goframe/os/gctx"
	gtest "github.com/888go/goframe/test/gtest"
)

type TestParamsCase struct {
	g.Meta `name:"root" root:"root"`
}

type TestParamsCaseRootInput struct {
	g.Meta `name:"root"`
	Name   string
}

type TestParamsCaseRootOutput struct {
	Content string
}

func (c *TestParamsCase) Root(ctx context.Context, in TestParamsCaseRootInput) (out *TestParamsCaseRootOutput, err error) {
	out = &TestParamsCaseRootOutput{
		Content: in.Name,
	}
	return
}

func Test_Command_ParamsCase(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.X创建()
		cmd, err := gcmd.NewFromObject(TestParamsCase{})
		t.AssertNil(err)

		os.Args = []string{"root", "-name=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}
