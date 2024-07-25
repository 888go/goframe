// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_Structs_WithTag(t *testing.T) {
	type User struct {
		Uid      int    `json:"id"`
		NickName string `json:"name"`
	}
	gtest.C(t, func(t *gtest.T) {
		var users []User
		params := g.Slice{
			g.Map{
				"id":   1,
				"name": "name1",
			},
			g.Map{
				"id":   2,
				"name": "name2",
			},
		}
		err := gconv.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []*User
		params := g.Slice{
			g.Map{
				"id":   1,
				"name": "name1",
			},
			g.Map{
				"id":   2,
				"name": "name2",
			},
		}
		err := gconv.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
}

func Test_Structs_WithoutTag(t *testing.T) {
	type User struct {
		Uid      int
		NickName string
	}
	gtest.C(t, func(t *gtest.T) {
		var users []User
		params := g.Slice{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := gconv.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
	gtest.C(t, func(t *gtest.T) {
		var users []*User
		params := g.Slice{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := gconv.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
}

func Test_Structs_SliceParameter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Uid      int
			NickName string
		}
		var users []User
		params := g.Slice{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := gconv.Structs(params, users)
		t.AssertNE(err, nil)
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Uid      int
			NickName string
		}
		type A struct {
			Users []User
		}
		var a A
		params := g.Slice{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := gconv.Structs(params, a.Users)
		t.AssertNE(err, nil)
	})
}

func Test_Structs_DirectReflectSet(t *testing.T) {
	type A struct {
		Id   int
		Name string
	}
	gtest.C(t, func(t *gtest.T) {
		var (
			a = []*A{
				{Id: 1, Name: "john"},
				{Id: 2, Name: "smith"},
			}
			b []*A
		)
		err := gconv.Structs(a, &b)
		t.AssertNil(err)
		t.AssertEQ(a, b)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			a = []A{
				{Id: 1, Name: "john"},
				{Id: 2, Name: "smith"},
			}
			b []A
		)
		err := gconv.Structs(a, &b)
		t.AssertNil(err)
		t.AssertEQ(a, b)
	})
}

func Test_Structs_IntSliceAttribute(t *testing.T) {
	type A struct {
		Id []int
	}
	type B struct {
		*A
		Name string
	}
	gtest.C(t, func(t *gtest.T) {
		var (
			array []*B
		)
		err := gconv.Structs(g.Slice{
			g.Map{"id": nil, "name": "john"},
			g.Map{"id": nil, "name": "smith"},
		}, &array)
		t.AssertNil(err)
		t.Assert(len(array), 2)
		t.Assert(array[0].Name, "john")
		t.Assert(array[1].Name, "smith")
	})
}
