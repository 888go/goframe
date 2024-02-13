// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package cmd类_test

import (
	"context"
	"fmt"
	"os"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/genv"
)

func ExampleInit() {
	cmd类.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(`%#v`, cmd类.GetArgAll())

	// Output:
	// []string{"gf", "build", "main.go"}
}

func ExampleGetArg() {
	cmd类.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(
		`Arg[0]: "%v", Arg[1]: "%v", Arg[2]: "%v", Arg[3]: "%v"`,
		cmd类.GetArg(0), cmd类.GetArg(1), cmd类.GetArg(2), cmd类.GetArg(3),
	)

	// Output:
	// Arg[0]: "gf", Arg[1]: "build", Arg[2]: "main.go", Arg[3]: ""
}

func ExampleGetArgAll() {
	cmd类.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(`%#v`, cmd类.GetArgAll())

	// Output:
	// []string{"gf", "build", "main.go"}
}

func ExampleGetOpt() {
	cmd类.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(
		`Opt["o"]: "%v", Opt["y"]: "%v", Opt["d"]: "%v"`,
		cmd类.GetOpt("o"), cmd类.GetOpt("y"), cmd类.GetOpt("d", "default value"),
	)

	// Output:
	// Opt["o"]: "gf.exe", Opt["y"]: "", Opt["d"]: "default value"
}

func ExampleGetOpt_Def() {
	cmd类.Init("gf", "build", "main.go", "-o=gf.exe", "-y")

	fmt.Println(cmd类.GetOpt("s", "Def").String())

	// Output:
	// Def
}

func ExampleGetOptAll() {
	cmd类.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(`%#v`, cmd类.GetOptAll())

	// May Output:
	// map[string]string{"o":"gf.exe", "y":""}
}

func ExampleGetOptWithEnv() {
	fmt.Printf("Opt[gf.test]:%s\n", cmd类.GetOptWithEnv("gf.test"))
	_ = 环境变量类.X设置值("GF_TEST", "YES")
	fmt.Printf("Opt[gf.test]:%s\n", cmd类.GetOptWithEnv("gf.test"))

	// Output:
	// Opt[gf.test]:
	// Opt[gf.test]:YES
}

func ExampleParse() {
	os.Args = []string{"gf", "build", "main.go", "-o=gf.exe", "-y"}
	p, err := cmd类.Parse(g.MapStrBool{
		"o,output": true,
		"y,yes":    false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(p.GetOpt("o"))
	fmt.Println(p.GetOpt("output"))
	fmt.Println(p.GetOpt("y") != nil)
	fmt.Println(p.GetOpt("yes") != nil)
	fmt.Println(p.GetOpt("none") != nil)
	fmt.Println(p.GetOpt("none", "Def"))

	// Output:
	// gf.exe
	// gf.exe
	// true
	// true
	// false
	// Def
}

func ExampleCommandFromCtx() {
	var (
		command = cmd类.Command{
			Name: "start",
		}
	)

	ctx := context.WithValue(上下文类.X创建(), cmd类.CtxKeyCommand, &command)
	unAddCtx := context.WithValue(上下文类.X创建(), cmd类.CtxKeyCommand, &cmd类.Command{})
	nonKeyCtx := context.WithValue(上下文类.X创建(), "Testkey", &cmd类.Command{})

	fmt.Println(cmd类.CommandFromCtx(ctx).Name)
	fmt.Println(cmd类.CommandFromCtx(unAddCtx).Name)
	fmt.Println(cmd类.CommandFromCtx(nonKeyCtx) == nil)

	// Output:
	// start
	//
	// true
}

func ExampleCommand_AddCommand() {
	commandRoot := &cmd类.Command{
		Name: "gf",
	}
	commandRoot.AddCommand(&cmd类.Command{
		Name: "start",
	}, &cmd类.Command{})

	commandRoot.Print()

	// Output:
	// USAGE
	//     gf COMMAND [OPTION]
	//
	// COMMAND
	//     start
}

func ExampleCommand_AddCommand_Repeat() {
	commandRoot := &cmd类.Command{
		Name: "gf",
	}
	err := commandRoot.AddCommand(&cmd类.Command{
		Name: "start",
	}, &cmd类.Command{
		Name: "stop",
	}, &cmd类.Command{
		Name: "start",
	})

	fmt.Println(err)

	// Output:
	// command "start" is already added to command "gf"
}

func ExampleCommand_AddObject() {
	var (
		command = cmd类.Command{
			Name: "start",
		}
	)

	command.AddObject(&TestCmdObject{})

	command.Print()

	// Output:
	// USAGE
	//     start COMMAND [OPTION]
	//
	// COMMAND
	//     root    root env command
}

func ExampleCommand_AddObject_Error() {
	var (
		command = cmd类.Command{
			Name: "start",
		}
	)

	err := command.AddObject(&[]string{"Test"})

	fmt.Println(err)

	// Output:
	// input object should be type of struct, but got "*[]string"
}

func ExampleCommand_Print() {
	commandRoot := &cmd类.Command{
		Name: "gf",
	}
	commandRoot.AddCommand(&cmd类.Command{
		Name: "start",
	}, &cmd类.Command{})

	commandRoot.Print()

	// Output:
	// USAGE
	//     gf COMMAND [OPTION]
	//
	// COMMAND
	//     start
}

func ExampleScan() {
	fmt.Println(cmd类.Scan("gf scan"))

	// Output:
	// gf scan
}

func ExampleScanf() {
	fmt.Println(cmd类.Scanf("gf %s", "scanf"))

	// Output:
	// gf scanf
}

func ExampleParserFromCtx() {
	parser, _ := cmd类.Parse(nil)

	ctx := context.WithValue(上下文类.X创建(), cmd类.CtxKeyParser, parser)
	nilCtx := context.WithValue(上下文类.X创建(), "NilCtxKeyParser", parser)

	fmt.Println(cmd类.ParserFromCtx(ctx).GetArgAll())
	fmt.Println(cmd类.ParserFromCtx(nilCtx) == nil)

	// Output:
	// [gf build main.go]
	// true
}

func ExampleParseArgs() {
	p, _ := cmd类.ParseArgs([]string{
		"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root",
	}, nil)

	fmt.Println(p.GetArgAll())
	fmt.Println(p.GetOptAll())

	// Output:
	// [gf path]
	// map[force:remove fq: n:root p:www]
}

func ExampleParser_GetArg() {
	p, _ := cmd类.ParseArgs([]string{
		"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root",
	}, nil)

	fmt.Println(p.GetArg(-1, "Def").String())
	fmt.Println(p.GetArg(-1) == nil)

	// Output:
	// Def
	// true
}
