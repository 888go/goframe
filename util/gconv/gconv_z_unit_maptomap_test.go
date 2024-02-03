// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_MapToMap1(t *testing.T) {
// 将map[int]int 转换为 map[string]string
// 创建一个空的原始映射
// 这段Go代码注释翻译成中文注释为：
// ```go
// 将整数到整数的映射（map[int]int）转换为字符串到字符串的映射（map[string]string）
// 初始化一个空的原生映射
	gtest.C(t, func(t *gtest.T) {
		m1 := g.MapIntInt{}
		m2 := g.MapStrStr{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(len(m1), len(m2))
	})
	// map[int]int -> map[string]string
	gtest.C(t, func(t *gtest.T) {
		m1 := g.MapIntInt{
			1: 100,
			2: 200,
		}
		m2 := g.MapStrStr{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["1"], m1[1])
		t.Assert(m2["2"], m1[2])
	})
	// 将 map[string]interface{} 类型转换为 map[string]string 类型
// 这段注释表明了代码的功能是将一个键为字符串、值为接口类型的映射（map）转换为键同样为字符串但值为字符串类型的映射。
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := g.MapStrStr{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// 将字符串到字符串的映射转换为字符串到接口的映射
	gtest.C(t, func(t *gtest.T) {
		m1 := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := g.Map{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// 将 map[string]interface{} 转换为 map[interface{}]interface{}
// 这段注释表明，该代码片段的功能是将键类型为字符串、值类型为接口的映射（map）转换为键和值类型都为接口的映射。
	gtest.C(t, func(t *gtest.T) {
		m1 := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		m2 := g.MapAnyAny{}
		t.Assert(gconv.MapToMap(m1, &m2), nil)
		t.Assert(m2["k1"], m1["k1"])
		t.Assert(m2["k2"], m1["k2"])
	})
	// 将字符串转换为 map[string]interface{}
	gtest.C(t, func(t *gtest.T) {
		jsonStr := `{"id":100, "name":"john"}`

		m1 := g.MapStrAny{}
		t.Assert(gconv.MapToMap(jsonStr, &m1), nil)
		t.Assert(m1["id"], 100)

		m2 := g.MapStrAny{}
		t.Assert(gconv.MapToMap([]byte(jsonStr), &m2), nil)
		t.Assert(m2["id"], 100)
	})
}

func Test_MapToMap2(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	params := g.Map{
		"key": g.Map{
			"id":   1,
			"name": "john",
		},
	}
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]User)
		err := gconv.MapToMap(params, &m)
		t.AssertNil(err)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		m := (map[string]User)(nil)
		err := gconv.MapToMap(params, &m)
		t.AssertNil(err)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		m := make(map[string]*User)
		err := gconv.MapToMap(params, &m)
		t.AssertNil(err)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		m := (map[string]*User)(nil)
		err := gconv.MapToMap(params, &m)
		t.AssertNil(err)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
}

func Test_MapToMapDeep(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	params := g.Map{
		"key": g.Map{
			"id":   1,
			"name": "john",
		},
	}
	gtest.C(t, func(t *gtest.T) {
		m := (map[string]*User)(nil)
		err := gconv.MapToMap(params, &m)
		t.AssertNil(err)
		t.Assert(len(m), 1)
		t.Assert(m["key"].Id, 1)
		t.Assert(m["key"].Name, "john")
	})
}

func Test_MapToMaps(t *testing.T) {
	params := g.Slice{
		g.Map{"id": 1, "name": "john"},
		g.Map{"id": 2, "name": "smith"},
	}
	gtest.C(t, func(t *gtest.T) {
		var s []g.Map
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s, params)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []*g.Map
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
		t.Assert(s, params)
	})
	gtest.C(t, func(t *gtest.T) {
		jsonStr := `[{"id":100, "name":"john"},{"id":200, "name":"smith"}]`

		var m1 []g.Map
		t.Assert(gconv.MapToMaps(jsonStr, &m1), nil)
		t.Assert(m1[0]["id"], 100)
		t.Assert(m1[1]["id"], 200)

		t.Assert(gconv.MapToMaps([]byte(jsonStr), &m1), nil)
		t.Assert(m1[0]["id"], 100)
		t.Assert(m1[1]["id"], 200)
	})
}

func Test_MapToMaps_StructParams(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	params := g.Slice{
		User{1, "name1"},
		User{2, "name2"},
	}
	gtest.C(t, func(t *gtest.T) {
		var s []g.Map
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		var s []*g.Map
		err := gconv.MapToMaps(params, &s)
		t.AssertNil(err)
		t.Assert(len(s), 2)
	})
}
