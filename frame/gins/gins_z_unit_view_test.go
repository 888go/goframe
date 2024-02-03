// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gins

import (
	"context"
	"fmt"
	"testing"
	
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
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
		err := gfile.PutContents(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.AssertNil(err)
		defer gfile.Remove(tpl)

		b, e := View().Parse(context.TODO(), "t.tpl", nil)
		t.Assert(e, nil)
		t.Assert(b, "中国")
	})
	gtest.C(t, func(t *gtest.T) {
		path := fmt.Sprintf(`%s/%d`, gfile.Temp(), gtime.TimestampNano())
		tpl := fmt.Sprintf(`%s/%s`, path, "t.tpl")
		err := gfile.PutContents(tpl, `{{"我是中国人" | substr 2 -1}}`)
		t.AssertNil(err)
		defer gfile.Remove(tpl)
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
		Config().GetAdapter().(*gcfg.AdapterFile).SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer Config().GetAdapter().(*gcfg.AdapterFile).ClearContent()
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
		Config().GetAdapter().(*gcfg.AdapterFile).SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer Config().GetAdapter().(*gcfg.AdapterFile).ClearContent()
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
		Config().GetAdapter().(*gcfg.AdapterFile).SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer Config().GetAdapter().(*gcfg.AdapterFile).ClearContent()
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
		Config().GetAdapter().(*gcfg.AdapterFile).SetContent(gfile.GetContents(gfile.Join(dirPath, "config.toml")))
		defer Config().GetAdapter().(*gcfg.AdapterFile).ClearContent()
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
