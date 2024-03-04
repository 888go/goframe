// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gview_test

import (
	"context"
	"testing"
	
	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/888go/goframe/gview"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_I18n(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gtest.DataPath("i18n"))

		g.I18n().SetLanguage("zh-CN")
		result1, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)

		g.I18n().SetLanguage("ja")
		result2, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result2, expect2)

		g.I18n().SetLanguage("none")
		result3, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result3, expect3)
	})
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")

		result1, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "zh-CN",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)

		result2, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "ja",
		})
		t.AssertNil(err)
		t.Assert(result2, expect2)

		result3, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "none",
		})
		t.AssertNil(err)
		t.Assert(result3, expect3)
	})
	// gi18n 经理为空
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")

		view := gview.New()
		view.SetI18n(nil)
		result1, err := view.ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "zh-CN",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)
	})
	// 在上下文中设置语言
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		ctx := gctx.New()
		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")
		ctx = gi18n.WithLanguage(ctx, "zh-CN")
		t.Log(gi18n.LanguageFromCtx(ctx))

		view := gview.New()

		result1, err := view.ParseContent(ctx, content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)
	})

}
