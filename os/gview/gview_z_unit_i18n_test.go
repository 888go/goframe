// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/os/gctx"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/test/gtest"
)

func Test_I18n(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.X多语言类().SetPath(单元测试类.DataPath("i18n"))

		g.X多语言类().SetLanguage("zh-CN")
		result1, err := g.X模板类().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)

		g.X多语言类().SetLanguage("ja")
		result2, err := g.X模板类().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result2, expect2)

		g.X多语言类().SetLanguage("none")
		result3, err := g.X模板类().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result3, expect3)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.X多语言类().SetPath(gdebug.CallerDirectory() + 文件类.Separator + "testdata" + 文件类.Separator + "i18n")

		result1, err := g.X模板类().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "zh-CN",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)

		result2, err := g.X模板类().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "ja",
		})
		t.AssertNil(err)
		t.Assert(result2, expect2)

		result3, err := g.X模板类().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "none",
		})
		t.AssertNil(err)
		t.Assert(result3, expect3)
	})
	// gi18n 经理为空
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "{#hello}{#world}!"`

		g.X多语言类().SetPath(gdebug.CallerDirectory() + 文件类.Separator + "testdata" + 文件类.Separator + "i18n")

		view := 模板类.New()
		view.SetI18n(nil)
		result1, err := view.ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "zh-CN",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)
	})
	// 在上下文中设置语言
	单元测试类.C(t, func(t *单元测试类.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		ctx := 上下文类.X创建()
		g.X多语言类().SetPath(gdebug.CallerDirectory() + 文件类.Separator + "testdata" + 文件类.Separator + "i18n")
		ctx = gi18n.WithLanguage(ctx, "zh-CN")
		t.Log(gi18n.LanguageFromCtx(ctx))

		view := 模板类.New()

		result1, err := view.ParseContent(ctx, content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)
	})

}
