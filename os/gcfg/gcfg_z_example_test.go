// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 配置类_test

import (
	"fmt"
	"os"

	"github.com/888go/goframe/frame/g"
	gcfg "github.com/888go/goframe/os/gcfg"
	gcmd "github.com/888go/goframe/os/gcmd"
	gctx "github.com/888go/goframe/os/gctx"
	genv "github.com/888go/goframe/os/genv"
)

func ExampleConfig_GetWithEnv() {
	var (
		key = `ENV_TEST`
		ctx = gctx.X创建()
	)
	v, err := g.Cfg别名().X取值并从环境变量(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("env:%s\n", v)
	if err = genv.X设置值(key, "gf"); err != nil {
		panic(err)
	}
	v, err = g.Cfg别名().X取值并从环境变量(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("env:%s", v)

	// Output:
	// env:
	// env:gf
}

func ExampleConfig_GetWithCmd() {
	var (
		key = `cmd.test`
		ctx = gctx.X创建()
	)
	v, err := g.Cfg别名().X取值并从启动命令(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cmd:%s\n", v)
		// 重新初始化自定义命令参数。 md5:6d308df3285ccc65
	os.Args = append(os.Args, fmt.Sprintf(`--%s=yes`, key))
	gcmd.Init(os.Args...)
		// 再次获取配置和命令选项。 md5:4ac97b4c3c56a003
	v, err = g.Cfg别名().X取值并从启动命令(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cmd:%s", v)

	// Output:
	// cmd:
	// cmd:yes
}

func Example_NewWithAdapter() {
	var (
		ctx          = gctx.X创建()
		content      = `{"a":"b", "c":1}`
		adapter, err = gcfg.NewAdapterContent(content)
	)
	if err != nil {
		panic(err)
	}
	config := gcfg.X创建并按适配器(adapter)
	fmt.Println(config.X取值PANI(ctx, "a"))
	fmt.Println(config.X取值PANI(ctx, "c"))

	// Output:
	// b
	// 1
}
