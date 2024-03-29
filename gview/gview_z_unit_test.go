// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/888go/goframe/gview"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/gogf/gf/v2/util/guid"
)

func init() {
	os.Setenv("GF_GVIEW_ERRORPRINT", "false")
}

func Test_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		str := `hello {{.name}},version:{{.version}};hello {{GetName}},version:{{GetVersion}};{{.other}}`
		pwd := gfile.Pwd()
		view := 模板类.New()
		view.SetDelimiters("{{", "}}")
		view.AddPath(pwd)
		view.SetPath(pwd)
		view.Assign("name", "gf")
		view.Assigns(g.Map{"version": "1.7.0"})
		view.BindFunc("GetName", func() string { return "gf" })
		view.BindFuncMap(模板类.FuncMap{"GetVersion": func() string { return "1.7.0" }})
		result, err := view.ParseContent(context.TODO(), str, g.Map{"other": "that's all"})
		t.Assert(err != nil, false)
		t.Assert(result, "hello gf,version:1.7.0;hello gf,version:1.7.0;that's all")

		// 测试api方法
		str = `hello {{.name}}`
		result, err = 模板类.ParseContent(context.TODO(), str, g.Map{"name": "gf"})
		t.Assert(err != nil, false)
		t.Assert(result, "hello gf")

		// 测试instance方法
		result, err = 模板类.Instance().ParseContent(context.TODO(), str, g.Map{"name": "gf"})
		t.Assert(err != nil, false)
		t.Assert(result, "hello gf")
	})
}

func Test_Func(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		str := `{{eq 1 1}};{{eq 1 2}};{{eq "A" "B"}}`
		result, err := 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `true;false;false`)

		str = `{{ne 1 2}};{{ne 1 1}};{{ne "A" "B"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `true;false;true`)

		str = `{{lt 1 2}};{{lt 1 1}};{{lt 1 0}};{{lt "A" "B"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `true;false;false;true`)

		str = `{{le 1 2}};{{le 1 1}};{{le 1 0}};{{le "A" "B"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `true;true;false;true`)

		str = `{{gt 1 2}};{{gt 1 1}};{{gt 1 0}};{{gt "A" "B"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `false;false;true;false`)

		str = `{{ge 1 2}};{{ge 1 1}};{{ge 1 0}};{{ge "A" "B"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `false;true;true;false`)

		str = `{{"<div>测试</div>"|text}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `测试`)

		str = `{{"<div>测试</div>"|html}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `&lt;div&gt;测试&lt;/div&gt;`)

		str = `{{"<div>测试</div>"|htmlencode}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `&lt;div&gt;测试&lt;/div&gt;`)

		str = `{{"&lt;div&gt;测试&lt;/div&gt;"|htmldecode}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `<div>测试</div>`)

		str = `{{"https://goframe.org"|url}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `https%3A%2F%2Fgoframe.org`)

		str = `{{"https://goframe.org"|urlencode}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `https%3A%2F%2Fgoframe.org`)

		str = `{{"https%3A%2F%2Fgoframe.org"|urldecode}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `https://goframe.org`)
		str = `{{"https%3NA%2F%2Fgoframe.org"|urldecode}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(gstr.Contains(result, "invalid URL escape"), true)

		str = `{{1540822968 | date "Y-m-d"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `2018-10-29`)
		str = `{{date "Y-m-d"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)

		str = `{{"我是中国人" | substr 2 -1}};{{"我是中国人" | substr 2  2}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `中国;中国`)

		str = `{{"我是中国人" | strlimit 2  "..."}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `我是...`)

		str = `{{"I'm中国人" | replace "I'm" "我是"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `我是中国人`)

		str = `{{compare "A" "B"}};{{compare "1" "2"}};{{compare 2 1}};{{compare 1 1}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `-1;-1;1;0`)

		str = `{{"热爱GF热爱生活" | hidestr 20  "*"}};{{"热爱GF热爱生活" | hidestr 50  "*"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `热爱GF*爱生活;热爱****生活`)

		str = `{{"热爱GF热爱生活" | highlight "GF" "red"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `热爱<span style="color:red;">GF</span>热爱生活`)

		str = `{{"gf" | toupper}};{{"GF" | tolower}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `GF;gf`)

		str = `{{concat "I" "Love" "GoFrame"}}`
		result, err = 模板类.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, `ILoveGoFrame`)
	})
	// eq: 多个值。
	gtest.C(t, func(t *gtest.T) {
		str := `{{eq 1 2 1 3 4 5}}`
		result, err := 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `true`)
	})
	gtest.C(t, func(t *gtest.T) {
		str := `{{eq 6 2 1 3 4 5}}`
		result, err := 模板类.ParseContent(context.TODO(), str, nil)
		t.Assert(err != nil, false)
		t.Assert(result, `false`)
	})
}

func Test_FuncNl2Br(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		str := `{{"Go\nFrame" | nl2br}}`
		result, err := 模板类.ParseContent(context.TODO(), str, nil)
		t.AssertNil(err)
		t.Assert(result, `Go<br>Frame`)
	})
	gtest.C(t, func(t *gtest.T) {
		s := ""
		for i := 0; i < 3000; i++ {
			s += "Go\nFrame\n中文"
		}
		str := `{{.content | nl2br}}`
		result, err := 模板类.ParseContent(context.TODO(), str, g.Map{
			"content": s,
		})
		t.AssertNil(err)
		t.Assert(result, strings.Replace(s, "\n", "<br>", -1))
	})
}

func Test_FuncInclude(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			header = `<h1>HEADER</h1>`
			main   = `<h1>hello gf</h1>`
			footer = `<h1>FOOTER</h1>`
			layout = `{{include "header.html" .}}
{{include "main.html" .}}
{{include "footer.html" .}}
{{include "footer_not_exist.html" .}}
{{include "" .}}`
			templatePath = gfile.Temp(guid.S())
		)

		gfile.Mkdir(templatePath)
		defer gfile.Remove(templatePath)

		t.AssertNil(gfile.PutContents(gfile.Join(templatePath, `header.html`), header))
		t.AssertNil(gfile.PutContents(gfile.Join(templatePath, `main.html`), main))
		t.AssertNil(gfile.PutContents(gfile.Join(templatePath, `footer.html`), footer))
		t.AssertNil(gfile.PutContents(gfile.Join(templatePath, `layout.html`), layout))

		view := 模板类.New(templatePath)
		result, err := view.Parse(context.TODO(), "notfound.html")
		t.AssertNE(err, nil)
		t.Assert(result, ``)

		result, err = view.Parse(context.TODO(), "layout.html")
		t.AssertNil(err)
		t.Assert(result, `<h1>HEADER</h1>
<h1>hello gf</h1>
<h1>FOOTER</h1>
template file "footer_not_exist.html" not found
`)

		t.AssertNil(gfile.PutContents(gfile.Join(templatePath, `notfound.html`), "notfound"))
		result, err = view.Parse(context.TODO(), "notfound.html")
		t.AssertNil(err)
		t.Assert(result, `notfound`)
	})
}

func Test_SetPath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		view := 模板类.Instance("addpath")
		err := view.AddPath("tmp")
		t.AssertNE(err, nil)

		err = view.AddPath("gview.go")
		t.AssertNE(err, nil)

		os.Setenv("GF_GVIEW_PATH", "tmp")
		view = 模板类.Instance("setpath")
		err = view.SetPath("tmp")
		t.AssertNE(err, nil)

		err = view.SetPath("gview.go")
		t.AssertNE(err, nil)

		view = 模板类.New(gfile.Pwd())
		err = view.SetPath("tmp")
		t.AssertNE(err, nil)

		err = view.SetPath("gview.go")
		t.AssertNE(err, nil)

		os.Setenv("GF_GVIEW_PATH", "template")
		gfile.Mkdir(gfile.Pwd() + gfile.Separator + "template")
		view = 模板类.New()
	})
}

func Test_ParseContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		str := `{{.name}}`
		view := 模板类.New()
		result, err := view.ParseContent(context.TODO(), str, g.Map{"name": func() {}})
		t.Assert(err != nil, true)
		t.Assert(result, ``)
	})
}

func Test_HotReload(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		dirPath := gfile.Join(
			gfile.Temp(),
			"testdata",
			"template-"+gconv.String(gtime.TimestampNano()),
		)
		defer gfile.Remove(dirPath)
		filePath := gfile.Join(dirPath, "test.html")

		// Initialize data.
		err := gfile.PutContents(filePath, "test:{{.var}}")
		t.AssertNil(err)

		view := 模板类.New(dirPath)

		time.Sleep(100 * time.Millisecond)
		result, err := view.Parse(context.TODO(), "test.html", g.Map{
			"var": "1",
		})
		t.AssertNil(err)
		t.Assert(result, `test:1`)

		// Update data.
		err = gfile.PutContents(filePath, "test2:{{.var}}")
		t.AssertNil(err)

		time.Sleep(100 * time.Millisecond)
		result, err = view.Parse(context.TODO(), "test.html", g.Map{
			"var": "2",
		})
		t.AssertNil(err)
		t.Assert(result, `test2:2`)
	})
}

func Test_XSS(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		s := "<br>"
		r, err := v.ParseContent(context.TODO(), "{{.v}}", g.Map{
			"v": s,
		})
		t.AssertNil(err)
		t.Assert(r, s)
	})
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.SetAutoEncode(true)
		s := "<br>"
		r, err := v.ParseContent(context.TODO(), "{{.v}}", g.Map{
			"v": s,
		})
		t.AssertNil(err)
		t.Assert(r, ghtml.Entities(s))
	})
	// Tag "if".
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.SetAutoEncode(true)
		s := "<br>"
		r, err := v.ParseContent(context.TODO(), "{{if eq 1 1}}{{.v}}{{end}}", g.Map{
			"v": s,
		})
		t.AssertNil(err)
		t.Assert(r, ghtml.Entities(s))
	})
}

type TypeForBuildInFuncMap struct {
	Name  string
	Score float32
}

func (t *TypeForBuildInFuncMap) Test() (*TypeForBuildInFuncMap, error) {
	return &TypeForBuildInFuncMap{"john", 99.9}, nil
}

func Test_BuildInFuncMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", new(TypeForBuildInFuncMap))
		r, err := v.ParseContent(context.TODO(), "{{range $k, $v := map .v.Test}} {{$k}}:{{$v}} {{end}}")
		t.AssertNil(err)
		t.Assert(gstr.Contains(r, "Name:john"), true)
		t.Assert(gstr.Contains(r, "Score:99.9"), true)
	})

	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(context.TODO(), "{{range $k, $v := map }} {{$k}}:{{$v}} {{end}}")
		t.AssertNil(err)
		t.Assert(gstr.Contains(r, "Name:john"), false)
		t.Assert(gstr.Contains(r, "Score:99.9"), false)
	})
}

type TypeForBuildInFuncMaps struct {
	Name  string
	Score float32
}

func (t *TypeForBuildInFuncMaps) Test() ([]*TypeForBuildInFuncMaps, error) {
	return []*TypeForBuildInFuncMaps{
		{"john", 99.9},
		{"smith", 100},
	}, nil
}

func Test_BuildInFuncMaps(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", new(TypeForBuildInFuncMaps))
		r, err := v.ParseContent(context.TODO(), "{{range $k, $v := maps .v.Test}} {{$k}}:{{$v.Name}} {{$v.Score}} {{end}}")
		t.AssertNil(err)
		t.Assert(r, ` 0:john 99.9  1:smith 100 `)
	})

	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", new(TypeForBuildInFuncMaps))
		r, err := v.ParseContent(context.TODO(), "{{range $k, $v := maps }} {{$k}}:{{$v.Name}} {{$v.Score}} {{end}}")
		t.AssertNil(err)
		t.Assert(r, ``)
	})
}

func Test_BuildInFuncDump(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name":  "john",
			"score": 100,
		})
		r, err := v.ParseContent(context.TODO(), "{{dump .}}")
		t.AssertNil(err)
		fmt.Println(r)
		t.Assert(gstr.Contains(r, `"name":  "john"`), true)
		t.Assert(gstr.Contains(r, `"score": 100`), true)
	})

	gtest.C(t, func(t *gtest.T) {
		mode := gmode.Mode()
		gmode.SetTesting()
		defer gmode.Set(mode)
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name":  "john",
			"score": 100,
		})
		r, err := v.ParseContent(context.TODO(), "{{dump .}}")
		t.AssertNil(err)
		fmt.Println(r)
		t.Assert(gstr.Contains(r, `"name":  "john"`), false)
		t.Assert(gstr.Contains(r, `"score": 100`), false)
	})
}

func Test_BuildInFuncJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name": "john",
		})
		r, err := v.ParseContent(context.TODO(), "{{json .v}}")
		t.AssertNil(err)
		t.Assert(r, `{"name":"john"}`)
	})
}

func Test_BuildInFuncXml(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name": "john",
		})
		r, err := v.ParseContent(context.TODO(), "{{xml .v}}")
		t.AssertNil(err)
		t.Assert(r, `<name>john</name>`)
	})
}

func Test_BuildInFuncIni(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name": "john",
		})
		r, err := v.ParseContent(context.TODO(), "{{ini .v}}")
		t.AssertNil(err)
		t.Assert(r, `name=john
`)
	})
}

func Test_BuildInFuncYaml(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name": "john",
		})
		r, err := v.ParseContent(context.TODO(), "{{yaml .v}}")
		t.AssertNil(err)
		t.Assert(r, `name: john
`)
	})
}

func Test_BuildInFuncYamlIndent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name": "john",
		})
		r, err := v.ParseContent(context.TODO(), `{{yamli .v "####"}}`)
		t.AssertNil(err)
		t.Assert(r, `####name: john
`)
	})
}

func Test_BuildInFuncToml(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.Assign("v", g.Map{
			"name": "john",
		})
		r, err := v.ParseContent(context.TODO(), "{{toml .v}}")
		t.AssertNil(err)
		t.Assert(r, `name = "john"
`)
	})
}

func Test_BuildInFuncPlus(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{plus 1 2 3}}")
		t.AssertNil(err)
		t.Assert(r, `6`)
	})
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{1| plus 2}}")
		t.AssertNil(err)
		t.Assert(r, `3`)
	})
}

func Test_BuildInFuncMinus(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{minus 1 2 3}}")
		t.AssertNil(err)
		t.Assert(r, `-4`)
	})
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{2 | minus 3}}")
		t.AssertNil(err)
		t.Assert(r, `1`)
	})
}

func Test_BuildInFuncTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{times 1 2 3 4}}")
		t.AssertNil(err)
		t.Assert(r, `24`)
	})
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{2 | times 3}}")
		t.AssertNil(err)
		t.Assert(r, `6`)
	})
}

func Test_BuildInFuncDivide(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{divide 8 2 2}}")
		t.AssertNil(err)
		t.Assert(r, `2`)
	})
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{2 | divide 4}}")
		t.AssertNil(err)
		t.Assert(r, `2`)
	})
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		r, err := v.ParseContent(gctx.New(), "{{divide 8 0}}")
		t.AssertNil(err)
		t.Assert(r, `0`)
	})
}

func Test_Issue1416(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		err := v.SetPath(gtest.DataPath("issue1416"))
		t.AssertNil(err)
		r, err := v.ParseOption(context.TODO(), 模板类.Option{
			File:   "gview.tpl",
			Orphan: true,
			Params: map[string]interface{}{
				"hello": "world",
			},
		})
		t.AssertNil(err)
		t.Assert(r, `test.tpl content, vars: world`)
	})
}

// template/gview_test.html
// name:{{.name}} // （在模板中）名称：{{.name}}
func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBIFA0LY0ACEgycDCWpuQU5iSWp+ullmanl8SWpxSV6GSW5OaEhrAyM5o1fk095n/HdumrdNeaLW7c2MDAw/P8f4M3OoZ+9QESIgYGBj4GBAWYBA0MTmgUcSBaADSxt/JoM0o6sKMCbkUmEGeFCZKNBLoSBbY0gkqB7EcZhdw8ECDD8d0xEMg7JdaxsIAVMDEwMfQwMDAvAygEBAAD//0d6jptEAQAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}

func Test_GviewInGres(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		v := 模板类.New()
		v.SetPath("template")
		result, err := v.Parse(context.TODO(), "gview_test.html", g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result, "name:john")
	})
}
