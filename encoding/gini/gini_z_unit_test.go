// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ini类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gini"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/test/gtest"
)

var iniContent = `

;注释
aa=bb
[addr] 
#注释
ip = 127.0.0.1
port=9001
enable=true
command=/bin/echo "gf=GoFrame"

	[DBINFO]
	type=mysql
	user=root
	password=password
[键]
呵呵=值

`

func TestDecode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		res, err := ini类.X取Map([]byte(iniContent))
		if err != nil {
			单元测试类.Fatal(err)
		}
		t.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		t.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		t.Assert(res["addr"].(map[string]interface{})["command"], `/bin/echo "gf=GoFrame"`)
		t.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		t.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")
		t.Assert(res["键"].(map[string]interface{})["呵呵"], "值")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		errContent := `
		a = b
`
		_, err := ini类.X取Map([]byte(errContent))
		if err == nil {
			单元测试类.Fatal(err)
		}
	})
}

func TestEncode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		iniMap, err := ini类.X取Map([]byte(iniContent))
		if err != nil {
			单元测试类.Fatal(err)
		}

		iniStr, err := ini类.Map到ini(iniMap)
		if err != nil {
			单元测试类.Fatal(err)
		}

		res, err := ini类.X取Map(iniStr)
		if err != nil {
			单元测试类.Fatal(err)
		}

		t.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		t.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		t.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		t.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")

	})
}

func TestToJson(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		jsonStr, err := ini类.X取json([]byte(iniContent))
		if err != nil {
			单元测试类.Fatal(err)
		}

		json, err := json类.X加载并自动识别格式(jsonStr)
		if err != nil {
			单元测试类.Fatal(err)
		}

		iniMap, err := ini类.X取Map([]byte(iniContent))
		t.AssertNil(err)

		t.Assert(iniMap["addr"].(map[string]interface{})["ip"], json.X取值("addr.ip").String())
		t.Assert(iniMap["addr"].(map[string]interface{})["port"], json.X取值("addr.port").String())
		t.Assert(iniMap["DBINFO"].(map[string]interface{})["user"], json.X取值("DBINFO.user").String())
		t.Assert(iniMap["DBINFO"].(map[string]interface{})["type"], json.X取值("DBINFO.type").String())
	})
}
