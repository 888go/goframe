// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 使用`go test`命令，对所有`.go`文件进行测试，指定运行基准测试（Benchmark）中的所有模式（".*"），同时输出内存使用情况（-benchmem）。 md5:81db3d7bd1ed4da8

package gcmd_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/test/gtest"
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
		//-v, --version    输出二进制文件版本
		//-a, --arch       输出二进制文件架构，多个架构用逗号分隔
		//-s, --system     输出二进制文件系统，多个操作系统用逗号分隔
		//-o, --output     输出二进制文件路径，当构建单个二进制文件时使用
		//-p, --path       输出二进制文件目录路径，默认为'./bin'
		//-e, --extra      额外自定义的 "go build" 选项
		//-m, --mod        类似 "go build" 的 "-mod" 选项，使用 "-m none" 来禁用 Go 模块
		//-c, --cgo        启用或禁用 Cgo 功能，默认禁用 md5:4d71bb397da31ab5

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
