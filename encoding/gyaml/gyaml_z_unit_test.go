// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package yaml类_test

import (
	"testing"

	gjson "github.com/888go/goframe/encoding/gjson"
	gyaml "github.com/888go/goframe/encoding/gyaml"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	gtest "github.com/888go/goframe/test/gtest"
)

var yamlStr string = `
#即表示url属性值；
url: https://goframe.org

#数组，即表示server为[a,b,c]
server:
    - 120.168.117.21
    - 120.168.117.22
#常量
pi: 3.14   #定义一个数值3.14
hasChild: true  #定义一个boolean值
name: '你好YAML'   #定义一个字符串
`

var yamlErr string = `
[redis]
dd = 11
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`

func Test_Encode(t *testing.T) {
	// Map.
	gtest.C(t, func(t *gtest.T) {
		b, err := gyaml.Encode(g.Map{
			"k": "v",
		})
		t.AssertNil(err)
		t.Assert(string(b), `k: v
`)
	})
	// Array.
	gtest.C(t, func(t *gtest.T) {
		b, err := gyaml.Encode([]string{"a", "b", "c"})
		t.AssertNil(err)
		t.Assert(string(b), `- a
- b
- c
`)
	})
}

func Test_EncodeIndent(t *testing.T) {
	// Array.
	gtest.C(t, func(t *gtest.T) {
		b, err := gyaml.EncodeIndent([]string{"a", "b", "c"}, "####")
		t.AssertNil(err)
		t.Assert(string(b), `####- a
####- b
####- c
`)
	})
}

func Test_Decode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		result, err := gyaml.Decode([]byte(yamlStr))
		t.AssertNil(err)

		t.Assert(result, map[string]interface{}{
			"url":      "https://goframe.org",
			"server":   g.Slice别名{"120.168.117.21", "120.168.117.22"},
			"pi":       3.14,
			"hasChild": true,
			"name":     "你好YAML",
		})
	})
}

func Test_DecodeTo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		result := make(map[string]interface{})
		err := gyaml.DecodeTo([]byte(yamlStr), &result)
		t.AssertNil(err)
		t.Assert(result, map[string]interface{}{
			"url":      "https://goframe.org",
			"server":   g.Slice别名{"120.168.117.21", "120.168.117.22"},
			"pi":       3.14,
			"hasChild": true,
			"name":     "你好YAML",
		})
	})
}

func Test_DecodeError(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_, err := gyaml.Decode([]byte(yamlErr))
		t.AssertNE(err, nil)

		result := make(map[string]interface{})
		err = gyaml.DecodeTo([]byte(yamlErr), &result)
		t.AssertNE(err, nil)
	})
}

func Test_DecodeMapToJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := []byte(`
m:
 k: v
    `)
		v, err := gyaml.Decode(data)
		t.AssertNil(err)
		b, err := json.Marshal(v)
		t.AssertNil(err)
		t.Assert(b, `{"m":{"k":"v"}}`)
	})
}

func Test_ToJson(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]string)
		m["yaml"] = yamlStr
		res, err := gyaml.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		jsonyaml, err := gyaml.ToJson(res)
		if err != nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}

		p := gjson.X创建(res)
		if err != nil {
			t.Errorf("parser failed. %v", err)
			return
		}
		expectJson, err := p.X取json字节集()
		if err != nil {
			t.Errorf("parser ToJson failed. %v", err)
			return
		}
		t.Assert(jsonyaml, expectJson)
	})

	gtest.C(t, func(t *gtest.T) {
		_, err := gyaml.ToJson([]byte(yamlErr))
		if err == nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}
	})
}
