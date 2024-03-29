// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gutil"
)

func Test_MapCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := 工具类.MapCopy(m1)
		m2["k2"] = "v2"

		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], "v1")
		t.Assert(m2["k2"], "v2")
	})
}

func Test_MapContains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		t.Assert(工具类.MapContains(m1, "k1"), true)
		t.Assert(工具类.MapContains(m1, "K1"), false)
		t.Assert(工具类.MapContains(m1, "k2"), false)
		m2 := g.Map{}
		t.Assert(工具类.MapContains(m2, "k1"), false)
	})
}

func Test_MapDelete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		工具类.MapDelete(m1, "k1")
		工具类.MapDelete(m1, "K1")
		m2 := g.Map{}
		工具类.MapDelete(m2, "k1")
	})
}

func Test_MapMerge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		工具类.MapMerge(m1, m2, m3, nil)
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], "v2")
		t.Assert(m1["k3"], "v3")
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
		工具类.MapMerge(nil)
	})
}

func Test_MapMergeCopy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m1 := g.Map{
			"k1": "v1",
		}
		m2 := g.Map{
			"k2": "v2",
		}
		m3 := g.Map{
			"k3": "v3",
		}
		m := 工具类.MapMergeCopy(m1, m2, m3, nil)
		t.Assert(m["k1"], "v1")
		t.Assert(m["k2"], "v2")
		t.Assert(m["k3"], "v3")
		t.Assert(m1["k1"], "v1")
		t.Assert(m1["k2"], nil)
		t.Assert(m2["k1"], nil)
		t.Assert(m3["k1"], nil)
	})
}

func Test_MapPossibleItemByKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		k, v := 工具类.MapPossibleItemByKey(m, "NAME")
		t.Assert(k, "name")
		t.Assert(v, "guo")

		k, v = 工具类.MapPossibleItemByKey(m, "nick name")
		t.Assert(k, "NickName")
		t.Assert(v, "john")

		k, v = 工具类.MapPossibleItemByKey(m, "none")
		t.Assert(k, "")
		t.Assert(v, nil)
	})
}

func Test_MapContainsPossibleKey(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"name":     "guo",
			"NickName": "john",
		}
		t.Assert(工具类.MapContainsPossibleKey(m, "name"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "NAME"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "nickname"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "nick name"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "nick_name"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "nick-name"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "nick.name"), true)
		t.Assert(工具类.MapContainsPossibleKey(m, "none"), false)
	})
}

func Test_MapOmitEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"k1": "john",
			"e1": "",
			"e2": 0,
			"e3": nil,
			"k2": "smith",
		}
		工具类.MapOmitEmpty(m)
		t.Assert(len(m), 2)
		t.AssertNE(m["k1"], nil)
		t.AssertNE(m["k2"], nil)
		m1 := g.Map{}
		工具类.MapOmitEmpty(m1)
		t.Assert(len(m1), 0)
	})
}

func Test_MapToSlice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := g.Map{
			"k1": "v1",
			"k2": "v2",
		}
		s := 工具类.MapToSlice(m)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[1], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[2], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[3], g.Slice{"k1", "k2", "v1", "v2"})
		s1 := 工具类.MapToSlice(&m)
		t.Assert(len(s1), 4)
		t.AssertIN(s1[0], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[1], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[2], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s1[3], g.Slice{"k1", "k2", "v1", "v2"})
	})
	gtest.C(t, func(t *gtest.T) {
		m := g.MapStrStr{
			"k1": "v1",
			"k2": "v2",
		}
		s := 工具类.MapToSlice(m)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[1], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[2], g.Slice{"k1", "k2", "v1", "v2"})
		t.AssertIN(s[3], g.Slice{"k1", "k2", "v1", "v2"})
	})
	gtest.C(t, func(t *gtest.T) {
		m := g.MapStrStr{}
		s := 工具类.MapToSlice(m)
		t.Assert(len(s), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		s := 工具类.MapToSlice(1)
		t.Assert(s, nil)
	})
}
