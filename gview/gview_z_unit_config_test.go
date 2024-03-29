// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类_test

import (
	"context"
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/888go/goframe/gview/internal/command"
	"github.com/888go/goframe/gview"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Config(t *testing.T) {
	// show error print
	command.Init("-gf.gview.errorprint=true")
	gtest.C(t, func(t *gtest.T) {
		config := 模板类.Config{
			Paths: []string{gtest.DataPath("config")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := 模板类.New()
		err := view.SetConfig(config)
		t.AssertNil(err)

		view.SetI18n(gi18n.New())

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello gf,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:gf")

		t.Assert(view.GetDefaultFile(), "test.html")
	})
	// 设置配置文件路径失败：notexist
	gtest.C(t, func(t *gtest.T) {
		config := 模板类.Config{
			Paths: []string{"notexist", gtest.DataPath("config/test.html")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := 模板类.New()
		err := view.SetConfig(config)
		t.AssertNE(err, nil)
	})
	// SetConfig 设置配置文件路径失败：设置文件路径
	gtest.C(t, func(t *gtest.T) {
		config := 模板类.Config{
			Paths: []string{gtest.DataPath("config/test.html")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := 模板类.New()
		err := view.SetConfig(config)
		t.AssertNE(err, nil)
	})
}

func Test_ConfigWithMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		view := 模板类.New()
		err := view.SetConfigWithMap(g.Map{
			"Paths":       []string{gtest.DataPath("config")},
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "gf",
			},
		})
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello gf,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:gf")
	})
	// path as paths
	gtest.C(t, func(t *gtest.T) {
		view := 模板类.New()
		err := view.SetConfigWithMap(g.Map{
			"Path":        gtest.DataPath("config"),
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "gf",
			},
		})
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello gf,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:gf")
	})
	// path as paths
	gtest.C(t, func(t *gtest.T) {
		view := 模板类.New()
		err := view.SetConfigWithMap(g.Map{
			"Path":        []string{gtest.DataPath("config")},
			"DefaultFile": "test.html",
			"Delimiters":  []string{"${", "}"},
			"Data": g.Map{
				"name": "gf",
			},
		})
		t.AssertNil(err)

		str := `hello ${.name},version:${.version}`
		view.Assigns(g.Map{"version": "1.7.0"})
		result, err := view.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, "hello gf,version:1.7.0")

		result, err = view.ParseDefault(context.TODO())
		t.AssertNil(err)
		t.Assert(result, "name:gf")
	})
	// map is nil
	gtest.C(t, func(t *gtest.T) {
		view := 模板类.New()
		err := view.SetConfigWithMap(nil)
		t.AssertNE(err, nil)
	})
}
