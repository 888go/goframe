// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
)

func Test_Load_JSON1(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		errData := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]`)
		_, err := json类.X加载并按格式("json", errData, true)
		t.AssertNE(err, nil)
	})
	// JSON
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "test.json"
		文件类.X写入字节集(path, data)
		defer 文件类.X删除(path)
		j, err := json类.X加载文件(path, true)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
}

func Test_Load_JSON2(t *testing.T) {
	data := []byte(`{"n":123456789000000000000, "m":{"k":"v"}, "a":[1,2,3]}`)
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789000000000000")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
}

func Test_Load_XML(t *testing.T) {
	data := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n></doc>`)
	// XML
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("doc.n").String(), "123456789")
		t.Assert(j.X取值("doc.m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("doc.m.k").String(), "v")
		t.Assert(j.X取值("doc.a").Slice别名(), g.Slice别名{"1", "2", "3"})
		t.Assert(j.X取值("doc.a.1").X取整数(), 2)
	})
	// XML
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载xml(data, true)
		t.AssertNil(err)
		t.Assert(j.X取值("doc.n").String(), "123456789")
		t.Assert(j.X取值("doc.m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("doc.m.k").String(), "v")
		t.Assert(j.X取值("doc.a").Slice别名(), g.Slice别名{"1", "2", "3"})
		t.Assert(j.X取值("doc.a.1").X取整数(), 2)
	})
	// XML
	单元测试类.C(t, func(t *单元测试类.T) {
		errData := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n><doc>`)
		_, err := json类.X加载并按格式("xml", errData, true)
		t.AssertNE(err, nil)
	})
	// XML
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "test.xml"
		文件类.X写入字节集(path, data)
		defer 文件类.X删除(path)
		j, err := json类.X加载文件(path)
		t.AssertNil(err)
		t.Assert(j.X取值("doc.n").String(), "123456789")
		t.Assert(j.X取值("doc.m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("doc.m.k").String(), "v")
		t.Assert(j.X取值("doc.a").Array别名(), g.Slice别名{"1", "2", "3"})
		t.Assert(j.X取值("doc.a.1").X取整数(), 2)
	})

	// XML
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
	// YAML
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载Yaml(data, true)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
	// YAML
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "test.yaml"
		文件类.X写入字节集(path, data)
		defer 文件类.X删除(path)
		j, err := json类.X加载文件(path)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
}

func Test_Load_YAML2(t *testing.T) {
	data := []byte("i : 123456789")
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("i"), "123456789")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{"1", "2", "3"})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
	// TOML
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载Toml(data, true)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{"1", "2", "3"})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
	// TOML
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "test.toml"
		文件类.X写入字节集(path, data)
		defer 文件类.X删除(path)
		j, err := json类.X加载文件(path)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Slice别名(), g.Slice别名{"1", "2", "3"})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})
}

func Test_Load_TOML2(t *testing.T) {
	data := []byte("i=123456789")
	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		t.AssertNil(err)
		t.Assert(j.X取值("i"), "123456789")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		errData := []byte("i : 123456789")
		_, err := json类.X加载并按格式("toml", errData, true)
		t.AssertNE(err, nil)
	})
}

func Test_Load_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		if err != nil {
			单元测试类.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")

		_, err = j.X取ini字节集()
		if err != nil {
			单元测试类.Fatal(err)
		}
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载ini(data, true)
		if err != nil {
			单元测试类.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
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

	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载并自动识别格式(data)
		if err != nil {
			单元测试类.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")

		_, err = j.X取properties字节集()
		if err != nil {
			单元测试类.Fatal(err)
		}
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		j, err := json类.X加载Properties(data, true)
		if err != nil {
			单元测试类.Fatal(err)
		}

		t.Assert(j.X取值("addr.ip").String(), "127.0.0.1")
		t.Assert(j.X取值("addr.port").String(), "9001")
		t.Assert(j.X取值("addr.enable").String(), "true")
		t.Assert(j.X取值("DBINFO.type").String(), "mysql")
		t.Assert(j.X取值("DBINFO.user").String(), "root")
		t.Assert(j.X取值("DBINFO.password").String(), "password")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		errData := []byte("i\\u1 : 123456789")
		_, err := json类.X加载并按格式("properties", errData, true)
		t.AssertNE(err, nil)
	})
}
