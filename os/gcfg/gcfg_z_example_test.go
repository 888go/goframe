// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcfg_test

import (
	"fmt"
	"os"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/genv"
)

func ExampleConfig_GetWithEnv() {
	var (
		key = `ENV_TEST`
		ctx = gctx.New()
	)
	v, err := g.Cfg().GetWithEnv(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("env:%s\n", v)
	if err = genv.Set(key, "gf"); err != nil {
		panic(err)
	}
	v, err = g.Cfg().GetWithEnv(ctx, key)
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
		ctx = gctx.New()
	)
	v, err := g.Cfg().GetWithCmd(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cmd:%s\n", v)
	// 重新初始化自定义命令参数。
	os.Args = append(os.Args, fmt.Sprintf(`--%s=yes`, key))
	gcmd.Init(os.Args...)
	// 再次获取配置和命令选项。
	v, err = g.Cfg().GetWithCmd(ctx, key)
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
		ctx          = gctx.New()
		content      = `{"a":"b", "c":1}`
		adapter, err = gcfg.NewAdapterContent(content)
	)
	if err != nil {
		panic(err)
	}
	config := gcfg.NewWithAdapter(adapter)
	fmt.Println(config.MustGet(ctx, "a"))
	fmt.Println(config.MustGet(ctx, "c"))

	// Output:
	// b
	// 1
}
