// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gview_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/internal/command"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Config(t *testing.T) {
	// show error print
	command.Init("-gf.gview.errorprint=true")
	gtest.C(t, func(t *gtest.T) {
		config := gview.Config{
			Paths: []string{gtest.DataPath("config")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := gview.New()
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
	// 设置配置路径失败：不存在. md5:173c2e3f1a4ba9d3
	gtest.C(t, func(t *gtest.T) {
		config := gview.Config{
			Paths: []string{"notexist", gtest.DataPath("config/test.html")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := gview.New()
		err := view.SetConfig(config)
		t.AssertNE(err, nil)
	})
	// 设置配置路径失败：设置文件路径. md5:8a10375be44f8d09
	gtest.C(t, func(t *gtest.T) {
		config := gview.Config{
			Paths: []string{gtest.DataPath("config/test.html")},
			Data: g.Map{
				"name": "gf",
			},
			DefaultFile: "test.html",
			Delimiters:  []string{"${", "}"},
		}

		view := gview.New()
		err := view.SetConfig(config)
		t.AssertNE(err, nil)
	})
}

func Test_ConfigWithMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		view := gview.New()
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
		view := gview.New()
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
		view := gview.New()
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
		view := gview.New()
		err := view.SetConfigWithMap(nil)
		t.AssertNE(err, nil)
	})
}
