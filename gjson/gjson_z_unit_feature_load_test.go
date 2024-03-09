// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"testing"
	
	"github.com/888go/goframe/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Load_JSON1(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	// JSON
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
	// JSON
	gtest.C(t, func(t *gtest.T) {
		errData := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)
		_, err := json类.X加载并按格式("json", errData, true)
		t.AssertNE(err, nil)
	})
	// JSON
	gtest.C(t, func(t *gtest.T) {
		path := "test.json"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := json类.X加载文件(path, true)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
}

func Test_Load_JSON2(t *testing.T) {
	data := []byte(`{"n":123456789000000000000, "m":{"k":"v"}, "a":[1,2,3]}`)
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789000000000000")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
}

func Test_Load_XML(t *testing.T) {
	data := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n></doc>`)
	// XML
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("doc.n").String(), "123456789")
		t.Assert(j.X取值("doc.m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("doc.m.k").String(), "v")
		t.Assert(j.X取值("doc.a").Slice(), g.Slice{"1", "2", "3"})
		t.Assert(j.X取值("doc.a.1").Int(), 2)
	})
	// XML
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载xml(data, true)
		t.AssertNil(err)
		t.Assert(j.X取值("doc.n").String(), "123456789")
		t.Assert(j.X取值("doc.m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("doc.m.k").String(), "v")
		t.Assert(j.X取值("doc.a").Slice(), g.Slice{"1", "2", "3"})
		t.Assert(j.X取值("doc.a.1").Int(), 2)
	})
	// XML
	gtest.C(t, func(t *gtest.T) {
		errData := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n><doc>`)
		_, err := json类.X加载并按格式("xml", errData, true)
		t.AssertNE(err, nil)
	})
	// XML
	gtest.C(t, func(t *gtest.T) {
		path := "test.xml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := json类.X加载文件(path)
		t.AssertNil(err)
		t.Assert(j.X取值("doc.n").String(), "123456789")
		t.Assert(j.X取值("doc.m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("doc.m.k").String(), "v")
		t.Assert(j.X取值("doc.a").Array(), g.Slice{"1", "2", "3"})
		t.Assert(j.X取值("doc.a.1").Int(), 2)
	})

	// XML
	gtest.C(t, func(t *gtest.T) {
		xml := `<?xml version="1.0"?>

	<Output type="o">
	<itotalSize>0</itotalSize>
	<ipageSize>1</ipageSize>
	<ipageIndex>2</ipageIndex>
	<itotalRecords>GF框架</itotalRecords>
	<nworkOrderDtos/>
	<nworkOrderFrontXML/>
	</Output>`
		j, err := json类.X加载并自动识别格式(xml)
		t.AssertNil(err)
		t.Assert(j.X取值("Output.ipageIndex"), "2")
		t.Assert(j.X取值("Output.itotalRecords"), "GF框架")
	})
}

func Test_Load_YAML1(t *testing.T) {
	data := []byte(`
a:
- 1
- 2
- 3
m:
 k: v
"n": 123456789
    `)
	// YAML
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
	// YAML
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载Yaml(data, true)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
	// YAML
	gtest.C(t, func(t *gtest.T) {
		path := "test.yaml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := json类.X加载文件(path)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
}

func Test_Load_YAML2(t *testing.T) {
	data := []byte("i : 123456789")
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("i"), "123456789")
	})
	gtest.C(t, func(t *gtest.T) {
		errData := []byte("i # 123456789")
		_, err := json类.X加载并按格式("yaml", errData, true)
		t.AssertNE(err, nil)
	})
}

func Test_Load_TOML1(t *testing.T) {
	data := []byte(`
a = ["1", "2", "3"]
n = 123456789

[m]
  k = "v"
`)
	// TOML
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{"1", "2", "3"})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
	// TOML
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载Toml(data, true)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{"1", "2", "3"})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
	// TOML
	gtest.C(t, func(t *gtest.T) {
		path := "test.toml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := json类.X加载文件(path)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice(), g.Slice{"1", "2", "3"})
		t.Assert(j.X取值("a.1").Int(), 2)
	})
}

func Test_Load_TOML2(t *testing.T) {
	data := []byte("i=123456789")
	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("i"), "123456789")
	})
	gtest.C(t, func(t *gtest.T) {
		errData := []byte("i : 123456789")
		_, err := json类.X加载并按格式("toml", errData, true)
		t.AssertNE(err, nil)
	})
}

func Test_Load_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := json类.X创建(nil)
		t.Assert(j.Interface(), nil)
		_, err := json类.Json格式到变量(nil)
		t.AssertNE(err, nil)
		_, err = json类.X解码到json(nil)
		t.AssertNE(err, nil)
		j, err = json类.X加载并自动识别格式(nil)
		t.AssertNil(err)
		t.Assert(j.Interface(), nil)

		j, err = json类.X加载并自动识别格式(`{"name": "gf"}`)
		t.AssertNil(err)

		j, err = json类.X加载并自动识别格式(`{"name": "gf"""}`)
		t.AssertNE(err, nil)

		j = json类.X创建(&g.Map{"name": "gf"})
		t.Assert(j.X取值("name").String(), "gf")

	})
}

func Test_Load_Ini(t *testing.T) {
	var data = `

;注释

[addr]
ip = 127.0.0.1
port=9001
enable=true

	[DBINFO]
	type=mysql
	user=root
	password=password

`

	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		if err != nil {
			gtest.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")

		_, err = j.X取ini字节集()
		if err != nil {
			gtest.Fatal(err)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载ini(data, true)
		if err != nil {
			gtest.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")
	})

	gtest.C(t, func(t *gtest.T) {
		errData := []byte("i : 123456789")
		_, err := json类.X加载并按格式("ini", errData, true)
		t.AssertNE(err, nil)
	})
}

func Test_Load_YamlWithV3(t *testing.T) {
	content := `
# CLI tool, only in development environment.
# https://goframe.org/pages/viewpage.action?pageId=3673173
gfcli:
  gen:
    dao:
    - path            : "../../pkg/oss/oss/internal"
      group           : "oss"
      stdTime         : true
      descriptionTag  : true
      noJsonTag       : true
      noModelComment  : true
      overwriteDao    : true
      modelFileForDao : "model_dao.go"
      tablesEx        : |
        bpmn_info,
        dlocker,
        dlocker_detail,
        message_table,
        monitor_data,
        resource_param_info,
        version_info,
        version_topology_info,
        work_flow,
        work_flow_step_info,
        work_flow_undo_step_info

    - path            : "../../pkg/oss/workflow/internal"
      group           : "workflow"
      stdTime         : true
      descriptionTag  : true
      noJsonTag       : true
      noModelComment  : true
      overwriteDao    : true
      modelFileForDao : "model_dao.go"
`
	gtest.C(t, func(t *gtest.T) {
		_, err := json类.X加载并自动识别格式(content)
		t.AssertNil(err)
	})
}

func Test_Load_Properties(t *testing.T) {
	var data = `

#注释


addr.ip = 127.0.0.1
addr.port=9001
addr.enable=true
DBINFO.type=mysql
DBINFO.user=root
DBINFO.password=password

`

	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载并自动识别格式(data)
		if err != nil {
			gtest.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")

		_, err = j.X取properties字节集()
		if err != nil {
			gtest.Fatal(err)
		}
	})

	gtest.C(t, func(t *gtest.T) {
		j, err := json类.X加载Properties(data, true)
		if err != nil {
			gtest.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")
	})

	gtest.C(t, func(t *gtest.T) {
		errData := []byte("i\\u1 : 123456789")
		_, err := json类.X加载并按格式("properties", errData, true)
		t.AssertNE(err, nil)
	})
}
