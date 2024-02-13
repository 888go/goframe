// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
package toml类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/encoding/gtoml"
	"github.com/888go/goframe/test/gtest"
)

var tomlStr string = `
# 模板引擎目录
viewpath = "/home/www/templates/"
# MySQL数据库配置
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`

var tomlErr string = `
# 模板引擎目录
viewpath = "/home/www/templates/"
# MySQL数据库配置
[redis]
dd = 11
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`

func TestEncode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := make(map[string]string)
		m["toml"] = tomlStr
		res, err := toml类.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		t.Assert(json类.X创建(res).X取值("toml").String(), tomlStr)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := toml类.Encode(tomlErr)
		if err == nil {
			t.Errorf("encode should be failed. %v", err)
			return
		}
	})
}

func TestDecode(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := make(map[string]string)
		m["toml"] = tomlStr
		res, err := toml类.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		decodeStr, err := toml类.Decode(res)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}

		t.Assert(decodeStr.(map[string]interface{})["toml"], tomlStr)

		decodeStr1 := make(map[string]interface{})
		err = toml类.DecodeTo(res, &decodeStr1)
		if err != nil {
			t.Errorf("decodeTo failed. %v", err)
			return
		}
		t.Assert(decodeStr1["toml"], tomlStr)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := toml类.Decode([]byte(tomlErr))
		if err == nil {
			t.Errorf("decode failed. %v", err)
			return
		}

		decodeStr1 := make(map[string]interface{})
		err = toml类.DecodeTo([]byte(tomlErr), &decodeStr1)
		if err == nil {
			t.Errorf("decodeTo failed. %v", err)
			return
		}
	})
}

func TestToJson(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := make(map[string]string)
		m["toml"] = tomlStr
		res, err := toml类.Encode(m)
		if err != nil {
			t.Errorf("encode failed. %v", err)
			return
		}

		jsonToml, err := toml类.ToJson(res)
		if err != nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}

		p := json类.X创建(res)
		expectJson, err := p.X取json字节集()
		if err != nil {
			t.Errorf("parser ToJson failed. %v", err)
			return
		}
		t.Assert(jsonToml, expectJson)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		_, err := toml类.ToJson([]byte(tomlErr))
		if err == nil {
			t.Errorf("ToJson failed. %v", err)
			return
		}
	})
}
