// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcmd_test

import (
	"context"
	"os"
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/test/gtest"
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
		var ctx = gctx.New()
		cmd, err := gcmd.NewFromObject(TestParamsCase{})
		t.AssertNil(err)

		os.Args = []string{"root", "-name=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}
