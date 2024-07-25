// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_MapToMap1(t *testing.T) {
	// int到int的映射 -> string到string的映射
	// 清空原始映射。 md5:53ade5c68bd0aad0
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
	// 将map[string]interface{}类型的值转换为map[string]string类型. md5:273bd8baf5a0dc6f
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
	// 将字符串到字符串的映射转换为字符串到接口的映射. md5:47bac1ad94816db2
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
	// map[string]interface{} 转换为 map[interface{}]interface{}. md5:2e0e68b112586507
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
	// 字符串 -> 映射（string为键，interface{}为值）. md5:962827d1d4fb0447
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
