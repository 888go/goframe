// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gi18n_test

import (
	"time"
	
	"github.com/888go/goframe/encoding/gbase64"
	"github.com/888go/goframe/os/gctx"
	
	"context"
	"testing"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: 单元测试类.DataPath("i18n"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
		i18n.SetDelimiters("{$", "}")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")
		t.Assert(i18n.T(context.Background(), "{$hello}{$world}"), "你好世界")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")
		t.Assert(i18n.T(context.Background(), "{$你好} {$世界}"), "hello world")
		// 未定义的变量。
		t.Assert(i18n.T(context.Background(), "{$你好1}{$世界1}"), "{$你好1}{$世界1}")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: 单元测试类.DataPath("i18n-file"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
		t.Assert(i18n.T(context.Background(), "{#你好} {#世界}"), "hello world")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: gdebug.CallerDirectory() + 文件类.Separator + "testdata" + 文件类.Separator + "i18n-dir",
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})
}

func Test_TranslateFormat(t *testing.T) {
	// Tf
	单元测试类.C(t, func(t *单元测试类.T) {
		i18n := gi18n.New(gi18n.Options{
			Path: 单元测试类.DataPath("i18n"),
		})
		i18n.SetLanguage("none")
		t.Assert(i18n.Tf(context.Background(), "{#hello}{#world} %d", 2020), "{#hello}{#world} 2020")

		i18n.SetLanguage("ja")
		t.Assert(i18n.Tf(context.Background(), "{#hello}{#world} %d", 2020), "こんにちは世界 2020")
	})
}

func Test_DefaultManager(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := gi18n.SetPath(单元测试类.DataPath("i18n"))
		t.AssertNil(err)

		gi18n.SetLanguage("none")
		t.Assert(gi18n.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		gi18n.SetLanguage("ja")
		t.Assert(gi18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		gi18n.SetLanguage("zh-CN")
		t.Assert(gi18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		err := gi18n.SetPath(gdebug.CallerDirectory() + 文件类.Separator + "testdata" + 文件类.Separator + "i18n-dir")
		t.AssertNil(err)

		gi18n.SetLanguage("none")
		t.Assert(gi18n.Translate(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		gi18n.SetLanguage("ja")
		t.Assert(gi18n.Translate(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		gi18n.SetLanguage("zh-CN")
		t.Assert(gi18n.Translate(context.Background(), "{#hello}{#world}"), "你好世界")
	})
}

func Test_Instance(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := gi18n.Instance()
		err := m.SetPath(单元测试类.DataPath("i18n-dir"))
		t.AssertNil(err)
		m.SetLanguage("zh-CN")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "你好世界")
		t.Assert(m.T(context.Background(), "{#你好} {#世界}"), "hello world")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		m := gi18n.Instance()
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(g.X多语言类().T(context.Background(), "{#hello}{#world}"), "你好世界")
	})
	// 默认语言为：en
	单元测试类.C(t, func(t *单元测试类.T) {
		m := gi18n.Instance(转换类.String(时间类.X取时间戳纳秒()))
		m.SetPath(单元测试类.DataPath("i18n-dir"))
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "HelloWorld")
	})
}

func Test_Resource(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := g.X多语言类("resource")
		err := m.SetPath(单元测试类.DataPath("i18n-dir"))
		t.AssertNil(err)

		m.SetLanguage("none")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "{#hello}{#world}")

		m.SetLanguage("ja")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")

		m.SetLanguage("zh-CN")
		t.Assert(m.T(context.Background(), "{#hello}{#world}"), "你好世界")
	})
}

func Test_SetCtxLanguage(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		ctx := 上下文类.X创建()
		t.Assert(gi18n.LanguageFromCtx(ctx), "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(gi18n.LanguageFromCtx(nil), "")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		ctx := 上下文类.X创建()
		ctx = gi18n.WithLanguage(ctx, "zh-CN")
		t.Assert(gi18n.LanguageFromCtx(ctx), "zh-CN")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		ctx := gi18n.WithLanguage(context.Background(), "zh-CN")
		t.Assert(gi18n.LanguageFromCtx(ctx), "zh-CN")
	})

}

func Test_GetContent(t *testing.T) {
	i18n := gi18n.New(gi18n.Options{
		Path: 单元测试类.DataPath("i18n-file"),
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(i18n.GetContent(context.Background(), "hello"), "Hello")

		ctx := gi18n.WithLanguage(context.Background(), "zh-CN")
		t.Assert(i18n.GetContent(ctx, "hello"), "你好")

		ctx = gi18n.WithLanguage(context.Background(), "unknown")
		t.Assert(i18n.GetContent(ctx, "hello"), "")
	})
}

func Test_PathInResource(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		binContent, err := 资源类.Pack(单元测试类.DataPath("i18n"))
		t.AssertNil(err)
		err = 资源类.Add(编码base64类.X字节集编码到文本(binContent))
		t.AssertNil(err)

		i18n := gi18n.New()
		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")

		err = i18n.SetPath("i18n")
		t.Assert(err, nil)
		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")
	})
}

func Test_PathInNormal(t *testing.T) {
	// 将国际化文件复制到当前目录。
	文件类.X复制目录(单元测试类.DataPath("i18n"), 文件类.X路径生成(gdebug.CallerDirectory(), "manifest/i18n"))
	// 在测试后删除复制的文件。
	defer 文件类.X删除(文件类.X路径生成(gdebug.CallerDirectory(), "manifest"))

	i18n := gi18n.New()

	单元测试类.C(t, func(t *单元测试类.T) {
		i18n.SetLanguage("zh-CN")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "你好世界")
		// 设置不存在的路径。
		err := i18n.SetPath("i18n-not-exist")
		t.AssertNE(err, nil)
		err = i18n.SetPath("")
		t.AssertNE(err, nil)
		i18n.SetLanguage("ja")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}"), "こんにちは世界")
	})

	// 更改语言文件内容
	单元测试类.C(t, func(t *单元测试类.T) {
		i18n.SetLanguage("en")
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}{#name}"), "HelloWorld{#name}")
		err := 文件类.X追加文本(文件类.X路径生成(gdebug.CallerDirectory(), "manifest/i18n/en.toml"), "\nname = \"GoFrame\"")
		t.Assert(err, nil)
		// 等待文件修改时间发生变化。
		time.Sleep(10 * time.Millisecond)
		t.Assert(i18n.T(context.Background(), "{#hello}{#world}{#name}"), "HelloWorldGoFrame")
	})

	// Add new language
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 文件类.X写入文本(文件类.X路径生成(gdebug.CallerDirectory(), "manifest/i18n/en-US.toml"), "lang = \"en-US\"")
		t.Assert(err, nil)
		// 等待文件修改时间发生变化。
		time.Sleep(10 * time.Millisecond)
		i18n.SetLanguage("en-US")
		t.Assert(i18n.T(context.Background(), "{#lang}"), "en-US")
	})
}
