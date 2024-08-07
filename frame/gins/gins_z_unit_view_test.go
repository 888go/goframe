// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gins

import (
	"context"
	"fmt"
	"testing"

	"github.com/888go/goframe/internal/instance"
	gcfg "github.com/888go/goframe/os/gcfg"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_View(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(View(), nil)
		b, e := View().ParseContent(context.TODO(), `{{"我是中国人" | substr 2 -1}}`, nil)
		t.Assert(e, nil)
		t.Assert(b, "中国")
	})
	gtest.C(t, func(t *gtest.T) {
		tpl := "t.tpl"
		err := gfile.X写入文本(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.AssertNil(err)
		defer gfile.X删除(tpl)

		b, e := View().Parse(context.TODO(), "t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国")
	})
	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.X取临时目录(), gtime.X取时间戳纳秒())
		tpl := fmt.Sprintf(`%s/%s`, path, "t.tpl")
		err := gfile.X写入文本(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.AssertNil(err)
		defer gfile.X删除(tpl)
		err = View().AddPath(path)
		t.AssertNil(err)

		b, e := View().Parse(context.TODO(), "t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国")
	})
}

func Test_View_Config(t *testing.T) {
	var ctx = context.TODO()
	// view1 test1
	gtest.C(t, func(t *gtest.T) {
		dirPath := gtest.DataPath("view1")
		Config().X取适配器().(*gcfg.AdapterFile).SetContent(gfile.X读文本(gfile.X路径生成(dirPath, "config.toml")))
		defer Config().X取适配器().(*gcfg.AdapterFile).ClearContent()
		defer instance.Clear()

		view := View("test1")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(ctx, str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello test1,version:1.9.0")

		result, err = view.ParseDefault(ctx)
		t.AssertNil(err)
		t.Assert(result, "test1:test1")
	})
	// view1 test2
	gtest.C(t, func(t *gtest.T) {
		dirPath := gtest.DataPath("view1")
		Config().X取适配器().(*gcfg.AdapterFile).SetContent(gfile.X读文本(gfile.X路径生成(dirPath, "config.toml")))
		defer Config().X取适配器().(*gcfg.AdapterFile).ClearContent()
		defer instance.Clear()

		view := View("test2")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.AssertNil(err)

		str := `hello #{.name},version:#{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello test2,version:1.9.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "test2:test2")
	})
	// view2
	gtest.C(t, func(t *gtest.T) {
		dirPath := gtest.DataPath("view2")
		Config().X取适配器().(*gcfg.AdapterFile).SetContent(gfile.X读文本(gfile.X路径生成(dirPath, "config.toml")))
		defer Config().X取适配器().(*gcfg.AdapterFile).ClearContent()
		defer instance.Clear()

		view := View()
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.AssertNil(err)

		str := `hello {.name},version:{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello test,version:1.9.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "test:test")
	})
	// view2
	gtest.C(t, func(t *gtest.T) {
		dirPath := gtest.DataPath("view2")
		Config().X取适配器().(*gcfg.AdapterFile).SetContent(gfile.X读文本(gfile.X路径生成(dirPath, "config.toml")))
		defer Config().X取适配器().(*gcfg.AdapterFile).ClearContent()
		defer instance.Clear()

		view := View("test100")
		t.AssertNE(view, nil)
		err := view.AddPath(dirPath)
		t.AssertNil(err)

		str := `hello {.name},version:{.version}`
		view.Assigns(map[string]interface{}{"version": "1.9.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello test,version:1.9.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "test:test")
	})
}
