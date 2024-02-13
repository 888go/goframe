// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package json类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func TestJson_UnmarshalJSON(t *testing.T) {
	// Json Array
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			data = []byte(`["a", "b", "c"]`)
			j    = json类.X创建(nil)
			err  = json.UnmarshalUseNumber(data, j)
		)
		t.AssertNil(err)
		t.Assert(j.X取值(".").String(), `["a","b","c"]`)
		t.Assert(j.X取值("2").String(), `c`)
	})
	// Json Array Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			data = []byte(`[{"a":1}, {"b":2}, {"c":3}]`)
			j    = json类.X创建(nil)
			err  = json.UnmarshalUseNumber(data, j)
		)
		t.AssertNil(err)
		t.Assert(j.X取值(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.X取值("2.c").String(), `3`)
	})
	// Json Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			data = []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
			j    = json类.X创建(nil)
			err  = json.UnmarshalUseNumber(data, j)
		)
		t.AssertNil(err)
		t.Assert(j.X取值("n").String(), "123456789")
		t.Assert(j.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(j.X取值("m.k").String(), "v")
		t.Assert(j.X取值("a").Array别名(), g.Slice别名{1, 2, 3})
		t.Assert(j.X取值("a.1").X取整数(), 2)
	})

}

func TestJson_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Json *json类.Json
	}
	// Json Map.
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name": "john",
			"json": []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.X取值("n").String(), "123456789")
		t.Assert(v.Json.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(v.Json.X取值("m.k").String(), "v")
		t.Assert(v.Json.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(v.Json.X取值("a.1").X取整数(), 2)
	})
	// Json Array.
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name": "john",
			"json": `["a", "b", "c"]`,
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.X取值(".").String(), `["a","b","c"]`)
		t.Assert(v.Json.X取值("2").String(), `c`)
	})
	// Json Array Map.
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name": "john",
			"json": `[{"a":1},{"b":2},{"c":3}]`,
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.X取值(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(v.Json.X取值("2.c").String(), `3`)
	})
	// Map
	单元测试类.C(t, func(t *单元测试类.T) {
		var v *V
		err := 转换类.Struct(g.Map{
			"name": "john",
			"json": g.Map{
				"n": 123456789,
				"m": g.Map{"k": "v"},
				"a": g.Slice别名{1, 2, 3},
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.X取值("n").String(), "123456789")
		t.Assert(v.Json.X取值("m").X取Map(), g.Map{"k": "v"})
		t.Assert(v.Json.X取值("m.k").String(), "v")
		t.Assert(v.Json.X取值("a").Slice别名(), g.Slice别名{1, 2, 3})
		t.Assert(v.Json.X取值("a.1").X取整数(), 2)
	})
}
