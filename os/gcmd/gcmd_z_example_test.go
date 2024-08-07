// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package cmd类_test

import (
	"context"
	"fmt"
	"os"

	"github.com/888go/goframe/frame/g"
	gcmd "github.com/888go/goframe/os/gcmd"
	gctx "github.com/888go/goframe/os/gctx"
	genv "github.com/888go/goframe/os/genv"
)

func ExampleInit() {
	gcmd.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(`%#v`, gcmd.GetArgAll())

	// Output:
	// []string{"gf", "build", "main.go"}
}

func ExampleGetArg() {
	gcmd.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(
		`Arg[0]: "%v", Arg[1]: "%v", Arg[2]: "%v", Arg[3]: "%v"`,
		gcmd.GetArg(0), gcmd.GetArg(1), gcmd.GetArg(2), gcmd.GetArg(3),
	)

	// Output:
	// Arg[0]: "gf", Arg[1]: "build", Arg[2]: "main.go", Arg[3]: ""
}

func ExampleGetArgAll() {
	gcmd.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(`%#v`, gcmd.GetArgAll())

	// Output:
	// []string{"gf", "build", "main.go"}
}

func ExampleGetOpt() {
	gcmd.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(
		`Opt["o"]: "%v", Opt["y"]: "%v", Opt["d"]: "%v"`,
		gcmd.GetOpt("o"), gcmd.GetOpt("y"), gcmd.GetOpt("d", "default value"),
	)

	// Output:
	// Opt["o"]: "gf.exe", Opt["y"]: "", Opt["d"]: "default value"
}

func ExampleGetOpt_Def() {
	gcmd.Init("gf", "build", "main.go", "-o=gf.exe", "-y")

	fmt.Println(gcmd.GetOpt("s", "Def").String())

	// Output:
	// Def
}

func ExampleGetOptAll() {
	gcmd.Init("gf", "build", "main.go", "-o=gf.exe", "-y")
	fmt.Printf(`%#v`, gcmd.GetOptAll())

	// May Output:
	// map[string]string{"o":"gf.exe", "y":""}
}

func ExampleGetOptWithEnv() {
	fmt.Printf("Opt[gf.test]:%s\n", gcmd.GetOptWithEnv("gf.test"))
	_ = genv.X设置值("GF_TEST", "YES")
	fmt.Printf("Opt[gf.test]:%s\n", gcmd.GetOptWithEnv("gf.test"))

	// Output:
	// Opt[gf.test]:
	// Opt[gf.test]:YES
}

func ExampleParse() {
	os.Args = []string{"gf", "build", "main.go", "-o=gf.exe", "-y"}
	p, err := gcmd.Parse(g.MapStrBool{
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
		command = gcmd.Command{
			Name: "start",
		}
	)

	ctx := context.WithValue(gctx.X创建(), gcmd.CtxKeyCommand, &command)
	unAddCtx := context.WithValue(gctx.X创建(), gcmd.CtxKeyCommand, &gcmd.Command{})
	nonKeyCtx := context.WithValue(gctx.X创建(), "Testkey", &gcmd.Command{})

	fmt.Println(gcmd.CommandFromCtx(ctx).Name)
	fmt.Println(gcmd.CommandFromCtx(unAddCtx).Name)
	fmt.Println(gcmd.CommandFromCtx(nonKeyCtx) == nil)

	// Output:
	// start
	//
	// true
}

func ExampleCommand_AddCommand() {
	commandRoot := &gcmd.Command{
		Name: "gf",
	}
	commandRoot.AddCommand(&gcmd.Command{
		Name: "start",
	}, &gcmd.Command{})

	commandRoot.Print()

	// Output:
	// USAGE
	//     gf COMMAND [OPTION]
	//
	// COMMAND
	//     start
}

func ExampleCommand_AddCommand_Repeat() {
	commandRoot := &gcmd.Command{
		Name: "gf",
	}
	err := commandRoot.AddCommand(&gcmd.Command{
		Name: "start",
	}, &gcmd.Command{
		Name: "stop",
	}, &gcmd.Command{
		Name: "start",
	})

	fmt.Println(err)

	// Output:
	// command "start" is already added to command "gf"
}

func ExampleCommand_AddObject() {
	var (
		command = gcmd.Command{
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
		command = gcmd.Command{
			Name: "start",
		}
	)

	err := command.AddObject(&[]string{"Test"})

	fmt.Println(err)

	// Output:
	// input object should be type of struct, but got "*[]string"
}

func ExampleCommand_Print() {
	commandRoot := &gcmd.Command{
		Name: "gf",
	}
	commandRoot.AddCommand(&gcmd.Command{
		Name: "start",
	}, &gcmd.Command{})

	commandRoot.Print()

	// Output:
	// USAGE
	//     gf COMMAND [OPTION]
	//
	// COMMAND
	//     start
}

func ExampleScan() {
	fmt.Println(gcmd.Scan("gf scan"))

	// Output:
	// gf scan
}

func ExampleScanf() {
	fmt.Println(gcmd.Scanf("gf %s", "scanf"))

	// Output:
	// gf scanf
}

func ExampleParserFromCtx() {
	parser, _ := gcmd.Parse(nil)

	ctx := context.WithValue(gctx.X创建(), gcmd.CtxKeyParser, parser)
	nilCtx := context.WithValue(gctx.X创建(), "NilCtxKeyParser", parser)

	fmt.Println(gcmd.ParserFromCtx(ctx).GetArgAll())
	fmt.Println(gcmd.ParserFromCtx(nilCtx) == nil)

	// Output:
	// [gf build main.go]
	// true
}

func ExampleParseArgs() {
	p, _ := gcmd.ParseArgs([]string{
		"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root",
	}, nil)

	fmt.Println(p.GetArgAll())
	fmt.Println(p.GetOptAll())

	// Output:
	// [gf path]
	// map[force:remove fq: n:root p:www]
}

func ExampleParser_GetArg() {
	p, _ := gcmd.ParseArgs([]string{
		"gf", "--force", "remove", "-fq", "-p=www", "path", "-n", "root",
	}, nil)

	fmt.Println(p.GetArg(-1, "Def").String())
	fmt.Println(p.GetArg(-1) == nil)

	// Output:
	// Def
	// true
}
