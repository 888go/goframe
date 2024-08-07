// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ini类_test

import (
	"testing"

	gini "github.com/888go/goframe/encoding/gini"
	gjson "github.com/888go/goframe/encoding/gjson"
	gtest "github.com/888go/goframe/test/gtest"
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
	gtest.C(t, func(t *gtest.T) {
		res, err := gini.X取Map([]byte(iniContent))
		if err != nil {
			gtest.Fatal(err)
		}
		t.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		t.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		t.Assert(res["addr"].(map[string]interface{})["command"], `/bin/echo "gf=GoFrame"`)
		t.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		t.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")
		t.Assert(res["键"].(map[string]interface{})["呵呵"], "值")
	})

	gtest.C(t, func(t *gtest.T) {
		errContent := `
		a = b
`
		_, err := gini.X取Map([]byte(errContent))
		if err == nil {
			gtest.Fatal(err)
		}
	})
}

func TestEncode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		iniMap, err := gini.X取Map([]byte(iniContent))
		if err != nil {
			gtest.Fatal(err)
		}

		iniStr, err := gini.Map到ini(iniMap)
		if err != nil {
			gtest.Fatal(err)
		}

		res, err := gini.X取Map(iniStr)
		if err != nil {
			gtest.Fatal(err)
		}

		t.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		t.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		t.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		t.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")

	})
}

func TestToJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		jsonStr, err := gini.X取json([]byte(iniContent))
		if err != nil {
			gtest.Fatal(err)
		}

		json, err := gjson.X加载并自动识别格式(jsonStr)
		if err != nil {
			gtest.Fatal(err)
		}

		iniMap, err := gini.X取Map([]byte(iniContent))
		t.AssertNil(err)

		t.Assert(iniMap["addr"].(map[string]interface{})["ip"], json.X取值("addr.ip").String())
		t.Assert(iniMap["addr"].(map[string]interface{})["port"], json.X取值("addr.port").String())
		t.Assert(iniMap["DBINFO"].(map[string]interface{})["user"], json.X取值("DBINFO.user").String())
		t.Assert(iniMap["DBINFO"].(map[string]interface{})["type"], json.X取值("DBINFO.type").String())
	})
}
