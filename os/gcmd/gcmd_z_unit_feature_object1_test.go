// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

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

type TestCmdObject struct {
	g.Meta `name:"root" usage:"root env/test" brief:"root env command" dc:"description" ad:"ad"`
}

type TestCmdObjectEnvInput struct {
	g.Meta `name:"env" usage:"root env" brief:"root env command" dc:"root env command description" ad:"root env command ad"`
}

type TestCmdObjectEnvOutput struct{}

type TestCmdObjectTestInput struct {
	g.Meta  `name:"test" usage:"root test" brief:"root test command" dc:"root test command description" ad:"root test command ad"`
	Name    string `name:"yourname" v:"required" short:"n" orphan:"false" brief:"name for test command" d:"tom"`
	Version bool   `name:"version" short:"v" orphan:"true" brief:"show version"`
}

type TestCmdObjectTestOutput struct {
	Name    string
	Version bool
}

func (TestCmdObject) Env(ctx context.Context, in TestCmdObjectEnvInput) (out *TestCmdObjectEnvOutput, err error) {
	return
}

func (TestCmdObject) Test(ctx context.Context, in TestCmdObjectTestInput) (out *TestCmdObjectTestOutput, err error) {
	out = &TestCmdObjectTestOutput{
		Name:    in.Name,
		Version: in.Version,
	}
	return
}

func Test_Command_NewFromObject_Help(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx      = gctx.New()
			cmd, err = gcmd.NewFromObject(&TestCmdObject{})
		)
		t.AssertNil(err)
		t.Assert(cmd.Name, "root")

		os.Args = []string{"root"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, nil)
	})
}

func Test_Command_NewFromObject_Run(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx      = gctx.New()
			cmd, err = gcmd.NewFromObject(&TestCmdObject{})
		)
		t.AssertNil(err)
		t.Assert(cmd.Name, "root")

		os.Args = []string{"root", "test", "-n=john"}

		cmd.Run(ctx)
	})
}

func Test_Command_NewFromObject_RunWithValue(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx      = gctx.New()
			cmd, err = gcmd.NewFromObject(&TestCmdObject{})
		)
		t.AssertNil(err)
		t.Assert(cmd.Name, "root")

		// test short name
		os.Args = []string{"root", "test", "-n=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Name":"john","Version":false}`)

		// 测试名称 标签名称
		os.Args = []string{"root", "test", "-yourname=hailaz"}
		value1, err1 := cmd.RunWithValueError(ctx)
		t.AssertNil(err1)
		t.Assert(value1, `{"Name":"hailaz","Version":false}`)

		// 测试默认标签值
		os.Args = []string{"root", "test"}
		value2, err2 := cmd.RunWithValueError(ctx)
		t.AssertNil(err2)
		t.Assert(value2, `{"Name":"tom","Version":false}`)

		// 测试名称标签和孤儿标签为真
		os.Args = []string{"root", "test", "-v"}
		value3, err3 := cmd.RunWithValueError(ctx)
		t.AssertNil(err3)
		t.Assert(value3, `{"Name":"tom","Version":true}`)
	})
}

func Test_Command_AddObject(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx     = gctx.New()
			command = gcmd.Command{
				Name: "start",
			}
		)
		err := command.AddObject(&TestCmdObject{})
		t.AssertNil(err)

		os.Args = []string{"start", "root", "test", "-n=john"}
		value, err := command.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Name":"john","Version":false}`)
	})
}

type TestObjectForRootTag struct {
	g.Meta `name:"root" root:"root"`
}

type TestObjectForRootTagEnvInput struct {
	g.Meta `name:"env" usage:"root env" brief:"root env command" dc:"root env command description" ad:"root env command ad"`
}

type TestObjectForRootTagEnvOutput struct{}

type TestObjectForRootTagTestInput struct {
	g.Meta `name:"root"`
	Name   string `v:"required" short:"n" orphan:"false" brief:"name for test command"`
}

type TestObjectForRootTagTestOutput struct {
	Content string
}

func (TestObjectForRootTag) Env(ctx context.Context, in TestObjectForRootTagEnvInput) (out *TestObjectForRootTagEnvOutput, err error) {
	return
}

func (TestObjectForRootTag) Root(ctx context.Context, in TestObjectForRootTagTestInput) (out *TestObjectForRootTagTestOutput, err error) {
	out = &TestObjectForRootTagTestOutput{
		Content: in.Name,
	}
	return
}

func Test_Command_RootTag(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
		)
		cmd, err := gcmd.NewFromObject(TestObjectForRootTag{})
		t.AssertNil(err)

		os.Args = []string{"root", "-n=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
	// Pointer.
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
		)
		cmd, err := gcmd.NewFromObject(&TestObjectForRootTag{})
		t.AssertNil(err)

		os.Args = []string{"root", "-n=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}

type TestObjectForNeedArgs struct {
	g.Meta `name:"root" root:"root"`
}

type TestObjectForNeedArgsEnvInput struct {
	g.Meta `name:"env" usage:"root env" brief:"root env command" dc:"root env command description" ad:"root env command ad"`
}

type TestObjectForNeedArgsEnvOutput struct{}

type TestObjectForNeedArgsTestInput struct {
	g.Meta `name:"test"`
	Arg1   string `arg:"true" brief:"arg1 for test command"`
	Arg2   string `arg:"true" brief:"arg2 for test command"`
	Name   string `v:"required" short:"n" orphan:"false" brief:"name for test command"`
}

type TestObjectForNeedArgsTestOutput struct {
	Args []string
}

func (TestObjectForNeedArgs) Env(ctx context.Context, in TestObjectForNeedArgsEnvInput) (out *TestObjectForNeedArgsEnvOutput, err error) {
	return
}

func (TestObjectForNeedArgs) Test(ctx context.Context, in TestObjectForNeedArgsTestInput) (out *TestObjectForNeedArgsTestOutput, err error) {
	out = &TestObjectForNeedArgsTestOutput{
		Args: []string{in.Arg1, in.Arg2, in.Name},
	}
	return
}

func Test_Command_NeedArgs(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
		)
		cmd, err := gcmd.NewFromObject(TestObjectForNeedArgs{})
		t.AssertNil(err)

// 将当前命令行参数设置为: ["root", "test", "a", "b", "c", "-h"]
// os.Args = []string{"root", "test", "a", "b", "c", "-h"}
// 调用cmd.RunWithValueError函数并传入ctx作为上下文，获取返回值value和错误信息err
// value, err := cmd.RunWithValueError(ctx)
// 使用t.AssertNil函数验证err是否为nil（即没有错误发生）
// t.AssertNil(err)

		os.Args = []string{"root", "test", "a", "b", "c", "-n=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Args":["a","b","john"]}`)
	})
}

type TestObjectPointerTag struct {
	g.Meta `name:"root" root:"root"`
}

type TestObjectPointerTagEnvInput struct {
	g.Meta `name:"env" usage:"root env" brief:"root env command" dc:"root env command description" ad:"root env command ad"`
}

type TestObjectPointerTagEnvOutput struct{}

type TestObjectPointerTagTestInput struct {
	g.Meta `name:"root"`
	Name   string `v:"required" short:"n" orphan:"false" brief:"name for test command"`
}

type TestObjectPointerTagTestOutput struct {
	Content string
}

func (c *TestObjectPointerTag) Env(ctx context.Context, in TestObjectPointerTagEnvInput) (out *TestObjectPointerTagEnvOutput, err error) {
	return
}

func (c *TestObjectPointerTag) Root(ctx context.Context, in TestObjectPointerTagTestInput) (out *TestObjectPointerTagTestOutput, err error) {
	out = &TestObjectPointerTagTestOutput{
		Content: in.Name,
	}
	return
}

func Test_Command_Pointer(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
		)
		cmd, err := gcmd.NewFromObject(TestObjectPointerTag{})
		t.AssertNil(err)

		os.Args = []string{"root", "-n=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})

	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
		)
		cmd, err := gcmd.NewFromObject(&TestObjectPointerTag{})
		t.AssertNil(err)

		os.Args = []string{"root", "-n=john"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value, `{"Content":"john"}`)
	})
}

type TestCommandOrphan struct {
	g.Meta `name:"root" root:"root"`
}

type TestCommandOrphanIndexInput struct {
	g.Meta  `name:"index"`
	Orphan1 bool `short:"n1" orphan:"true"`
	Orphan2 bool `short:"n2" orphan:"true"`
	Orphan3 bool `short:"n3" orphan:"true"`
}

type TestCommandOrphanIndexOutput struct {
	Orphan1 bool
	Orphan2 bool
	Orphan3 bool
}

func (c *TestCommandOrphan) Index(ctx context.Context, in TestCommandOrphanIndexInput) (out *TestCommandOrphanIndexOutput, err error) {
	out = &TestCommandOrphanIndexOutput{
		Orphan1: in.Orphan1,
		Orphan2: in.Orphan2,
		Orphan3: in.Orphan3,
	}
	return
}

func Test_Command_Orphan_Parameter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var ctx = gctx.New()
		cmd, err := gcmd.NewFromObject(TestCommandOrphan{})
		t.AssertNil(err)

		os.Args = []string{"root", "index", "-n1", "-n2=0", "-n3=1"}
		value, err := cmd.RunWithValueError(ctx)
		t.AssertNil(err)
		t.Assert(value.(*TestCommandOrphanIndexOutput).Orphan1, true)
		t.Assert(value.(*TestCommandOrphanIndexOutput).Orphan2, false)
		t.Assert(value.(*TestCommandOrphanIndexOutput).Orphan3, true)
	})
}
