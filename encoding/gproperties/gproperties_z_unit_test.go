// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gproperties_test

import (
	"fmt"
	"strings"
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/encoding/gproperties"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
)

var pStr string = `
# 模板引擎目录
viewpath = "/home/www/templates/"
# redis数据库配置
redis.disk  = "127.0.0.1:6379,0"
redis.cache = "127.0.0.1:6379,1"
#SQL配置
sql.mysql.0.type = mysql
sql.mysql.0.ip = 127.0.0.1
sql.mysql.0.user = root
`

var errorTests = []struct {
	input, msg string
}{
	// unicode literals
	{"key\\u1 = value", "invalid unicode literal"},
	{"key\\u12 = value", "invalid unicode literal"},
	{"key\\u123 = value", "invalid unicode literal"},
	{"key\\u123g = value", "invalid unicode literal"},
	{"key\\u123", "invalid unicode literal"},

	// 循环引用
	{"key=${key}", `circular reference in:\nkey=\$\{key\}`},
	{"key1=${key2}\nkey2=${key1}", `circular reference in:\n(key1=\$\{key2\}\nkey2=\$\{key1\}|key2=\$\{key1\}\nkey1=\$\{key2\})`},

	// 不规范的表达式
	{"key=${ke", "malformed expression"},
	{"key=valu${ke", "malformed expression"},
}

func TestDecode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := make(map[string]interface{})
		m["properties"] = pStr
		res, err := gproperties.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}
		decodeMap, err := gproperties.Decode(res)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(decodeMap["properties"], pStr)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		for _, v := range errorTests {
			_, err := gproperties.Decode(([]byte)(v.input))
			if err == nil {
				t.Errorf("encode should be failed. %v", err)
				return
			}
			t.AssertIN(`Lib magiconair load Properties data failed.`, strings.Split(err.Error(), ":"))
		}
	})
}

func TestEncode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := make(map[string]interface{})
		m["properties"] = pStr
		res, err := gproperties.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}
		decodeMap, err := gproperties.Decode(res)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		t.Assert(decodeMap["properties"], pStr)
	})
}

func TestToJson(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		res, err := gproperties.Encode(map[string]interface{}{
			"sql": g.Map{
				"userName": "admin",
				"password": "123456",
			},
			"user": "admin",
			"no":   123,
		})
		fmt.Print(string(res))
		jsonPr, err := gproperties.ToJson(res)
		if err != nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}
		fmt.Print(string(jsonPr))

		p := json类.X创建(res)
		expectJson, err := p.X取json字节集()
		if err != nil {
			t.Errorf("parser ToJson failed. %v", err)
			return
		}
		t.Assert(jsonPr, expectJson)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		for _, v := range errorTests {
			_, err := gproperties.ToJson(([]byte)(v.input))
			if err == nil {
				t.Errorf("encode should be failed. %v", err)
				return
			}
			t.AssertIN(`Lib magiconair load Properties data failed.`, strings.Split(err.Error(), ":"))
		}
	})
}
