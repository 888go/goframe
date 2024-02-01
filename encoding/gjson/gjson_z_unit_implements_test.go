// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson_test
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
	gtest.C(t, func(t *gtest.T) {
		var (
			data = []byte(`["a", "b", "c"]`)
			j    = gjson.New(nil)
			err  = json.UnmarshalUseNumber(data, j)
		)
		t.AssertNil(err)
		t.Assert(j.Get(".").String(), `["a","b","c"]`)
		t.Assert(j.Get("2").String(), `c`)
	})
	// Json Array Map
	gtest.C(t, func(t *gtest.T) {
		var (
			data = []byte(`[{"a":1}, {"b":2}, {"c":3}]`)
			j    = gjson.New(nil)
			err  = json.UnmarshalUseNumber(data, j)
		)
		t.AssertNil(err)
		t.Assert(j.Get(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(j.Get("2.c").String(), `3`)
	})
	// Json Map
	gtest.C(t, func(t *gtest.T) {
		var (
			data = []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
			j    = gjson.New(nil)
			err  = json.UnmarshalUseNumber(data, j)
		)
		t.AssertNil(err)
		t.Assert(j.Get("n").String(), "123456789")
		t.Assert(j.Get("m").Map(), g.Map{"k": "v"})
		t.Assert(j.Get("m.k").String(), "v")
		t.Assert(j.Get("a").Array(), g.Slice{1, 2, 3})
		t.Assert(j.Get("a.1").Int(), 2)
	})

}

func TestJson_UnmarshalValue(t *testing.T) {
	type V struct {
		Name string
		Json *gjson.Json
	}
	// Json Map.
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"json": []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`),
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get("n").String(), "123456789")
		t.Assert(v.Json.Get("m").Map(), g.Map{"k": "v"})
		t.Assert(v.Json.Get("m.k").String(), "v")
		t.Assert(v.Json.Get("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(v.Json.Get("a.1").Int(), 2)
	})
	// Json Array.
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"json": `["a", "b", "c"]`,
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get(".").String(), `["a","b","c"]`)
		t.Assert(v.Json.Get("2").String(), `c`)
	})
	// Json Array Map.
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"json": `[{"a":1},{"b":2},{"c":3}]`,
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get(".").String(), `[{"a":1},{"b":2},{"c":3}]`)
		t.Assert(v.Json.Get("2.c").String(), `3`)
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v *V
		err := gconv.Struct(g.Map{
			"name": "john",
			"json": g.Map{
				"n": 123456789,
				"m": g.Map{"k": "v"},
				"a": g.Slice{1, 2, 3},
			},
		}, &v)
		t.AssertNil(err)
		t.Assert(v.Name, "john")
		t.Assert(v.Json.Get("n").String(), "123456789")
		t.Assert(v.Json.Get("m").Map(), g.Map{"k": "v"})
		t.Assert(v.Json.Get("m.k").String(), "v")
		t.Assert(v.Json.Get("a").Slice(), g.Slice{1, 2, 3})
		t.Assert(v.Json.Get("a.1").Int(), 2)
	})
}
