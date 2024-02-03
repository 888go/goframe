// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试所有.go文件，并执行基准测试（-bench=".*"），同时显示内存使用情况统计（-benchmem）

package gcmd_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/test/gtest"
)

func Test_Default(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gcmd.Init([]string{"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root"}...)
		t.Assert(len(gcmd.GetArgAll()), 2)
		t.Assert(gcmd.GetArg(1), "path")
		t.Assert(gcmd.GetArg(100, "test"), "test")
		t.Assert(gcmd.GetOpt("force"), "remove")
		t.Assert(gcmd.GetOpt("n"), "root")
		t.Assert(gcmd.GetOpt("fq").IsNil(), false)
		t.Assert(gcmd.GetOpt("p").IsNil(), false)
		t.Assert(gcmd.GetOpt("none").IsNil(), true)
		t.Assert(gcmd.GetOpt("none", "value"), "value")
	})
	gtest.C(t, func(t *gtest.T) {
		gcmd.Init([]string{"gf", "gen", "-h"}...)
		t.Assert(len(gcmd.GetArgAll()), 2)
		t.Assert(gcmd.GetOpt("h"), "")
		t.Assert(gcmd.GetOpt("h").IsNil(), false)
	})
}

func Test_BuildOptions(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := gcmd.BuildOptions(g.MapStrStr{
			"n": "john",
		})
		t.Assert(s, "-n=john")
	})

	gtest.C(t, func(t *gtest.T) {
		s := gcmd.BuildOptions(g.MapStrStr{
			"n": "john",
		}, "-test")
		t.Assert(s, "-testn=john")
	})

	gtest.C(t, func(t *gtest.T) {
		s := gcmd.BuildOptions(g.MapStrStr{
			"n1": "john",
			"n2": "huang",
		})
		t.Assert(strings.Contains(s, "-n1=john"), true)
		t.Assert(strings.Contains(s, "-n2=huang"), true)
	})
}

func Test_GetWithEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		genv.Set("TEST", "1")
		defer genv.Remove("TEST")
		t.Assert(gcmd.GetOptWithEnv("test"), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		genv.Set("TEST", "1")
		defer genv.Remove("TEST")
		gcmd.Init("-test", "2")
		t.Assert(gcmd.GetOptWithEnv("test"), 2)
	})
}

func Test_Command(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
			err error
		)
		commandRoot := &gcmd.Command{
			Name: "gf",
		}
		// env
		commandEnv := &gcmd.Command{
			Name: "env",
			Func: func(ctx context.Context, parser *gcmd.Parser) error {
				fmt.Println("env")
				return nil
			},
		}
		// test
		commandTest := &gcmd.Command{
			Name:        "test",
			Brief:       "test brief",
			Description: "test description current Golang environment variables",
			Examples: `
gf get github.com/gogf/gf
gf get github.com/gogf/gf@latest
gf get github.com/gogf/gf@master
gf get golang.org/x/sys
`,
			Arguments: []gcmd.Argument{
				{
					Name:   "my-option",
					Short:  "o",
					Brief:  "It's my custom option",
					Orphan: true,
				},
				{
					Name:   "another",
					Short:  "a",
					Brief:  "It's my another custom option",
					Orphan: true,
				},
			},
			Func: func(ctx context.Context, parser *gcmd.Parser) error {
				fmt.Println("test")
				return nil
			},
		}
		err = commandRoot.AddCommand(
			commandEnv,
		)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		err = commandRoot.AddObject(
			commandTest,
		)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}

		if err = commandRoot.RunWithError(ctx); err != nil {
			if gerror.Code(err) == gcode.CodeNotFound {
				commandRoot.Print()
			}
		}
	})
}

func Test_Command_Print(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.New()
			err error
		)
		c := &gcmd.Command{
			Name:        "gf",
			Description: `GoFrame Command Line Interface, which is your helpmate for building GoFrame application with convenience.`,
			Additional: `
Use 'gf help COMMAND' or 'gf COMMAND -h' for detail about a command, which has '...' in the tail of their comments.`,
		}
		// env
		commandEnv := &gcmd.Command{
			Name:        "env",
			Brief:       "show current Golang environment variables, long brief.long brief.long brief.long brief.long brief.long brief.long brief.long brief.",
			Description: "show current Golang environment variables",
			Func: func(ctx context.Context, parser *gcmd.Parser) error {
				return nil
			},
		}
		if err = c.AddCommand(commandEnv); err != nil {
			g.Log().Fatal(ctx, err)
		}
		// get
		commandGet := &gcmd.Command{
			Name:        "get",
			Brief:       "install or update GF to system in default...",
			Description: "show current Golang environment variables",

			Examples: `
gf get github.com/gogf/gf
gf get github.com/gogf/gf@latest
gf get github.com/gogf/gf@master
gf get golang.org/x/sys
`,
			Func: func(ctx context.Context, parser *gcmd.Parser) error {
				return nil
			},
		}
		if err = c.AddCommand(commandGet); err != nil {
			g.Log().Fatal(ctx, err)
		}
// 构建
//-n, --name       输出二进制文件名称
//-v, --version    输出二进制文件版本信息
//-a, --arch       输出二进制文件架构，多个架构使用','分隔
//-s, --system     输出二进制文件系统，多个操作系统使用','分隔
//-o, --output     输出二进制文件路径，仅在构建单个二进制文件时使用
//-p, --path       输出二进制目录路径，默认为'./bin'
//-e, --extra      额外的自定义 "go build" 选项
//-m, --mod        类似于 "go build" 命令中的 "-mod" 选项，使用 "-m none" 禁用 Go 模块功能
//-c, --cgo        启用或禁用 cgo 功能，默认是禁用状态
// 以上代码是对golang构建命令行参数的注释翻译，这些参数用于控制构建过程中的各种行为，如指定输出文件名、版本、架构、操作系统、输出路径等，并可对Go模块和cgo特性进行控制。

		commandBuild := gcmd.Command{
			Name:  "build",
			Usage: "gf build FILE [OPTION]",
			Brief: "cross-building go project for lots of platforms...",
			Description: `
The "build" command is most commonly used command, which is designed as a powerful wrapper for
"go build" command for convenience cross-compiling usage.
It provides much more features for building binary:
1. Cross-Compiling for many platforms and architectures.
2. Configuration file support for compiling.
3. Build-In Variables.
`,
			Examples: `
gf build main.go
gf build main.go --swagger
gf build main.go --pack public,template
gf build main.go --cgo
gf build main.go -m none 
gf build main.go -n my-app -a all -s all
gf build main.go -n my-app -a amd64,386 -s linux -p .
gf build main.go -n my-app -v 1.0 -a amd64,386 -s linux,windows,darwin -p ./docker/bin
`,
			Func: func(ctx context.Context, parser *gcmd.Parser) error {
				return nil
			},
		}
		if err = c.AddCommand(&commandBuild); err != nil {
			g.Log().Fatal(ctx, err)
		}
		_ = c.RunWithError(ctx)
	})
}

func Test_Command_NotFound(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		c0 := &gcmd.Command{
			Name: "c0",
		}
		c1 := &gcmd.Command{
			Name: "c1",
			FuncWithValue: func(ctx context.Context, parser *gcmd.Parser) (interface{}, error) {
				return nil, nil
			},
		}
		c21 := &gcmd.Command{
			Name: "c21",
			FuncWithValue: func(ctx context.Context, parser *gcmd.Parser) (interface{}, error) {
				return nil, nil
			},
		}
		c22 := &gcmd.Command{
			Name: "c22",
			FuncWithValue: func(ctx context.Context, parser *gcmd.Parser) (interface{}, error) {
				return nil, nil
			},
		}
		t.AssertNil(c0.AddCommand(c1))
		t.AssertNil(c1.AddCommand(c21, c22))

		os.Args = []string{"c0", "c1", "c23", `--test="abc"`}
		err := c0.RunWithError(gctx.New())
		t.Assert(err.Error(), `command "c1 c23" not found for command "c0", command line: c0 c1 c23 --test="abc"`)
	})
}
