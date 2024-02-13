// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Structs_WithTag(t *testing.T) {
	type User struct {
		Uid      int    `json:"id"`
		NickName string `json:"name"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []User
		params := g.Slice别名{
			g.Map{
				"id":   1,
				"name": "name1",
			},
			g.Map{
				"id":   2,
				"name": "name2",
			},
		}
		err := 转换类.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		params := g.Slice别名{
			g.Map{
				"id":   1,
				"name": "name1",
			},
			g.Map{
				"id":   2,
				"name": "name2",
			},
		}
		err := 转换类.Structs(params, &users)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []User
		params := g.Slice别名{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := 转换类.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var users []*User
		params := g.Slice别名{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := 转换类.Structs(params, &users)
		t.AssertNil(err)
		t.Assert(len(users), 2)
		t.Assert(users[0].Uid, 1)
		t.Assert(users[0].NickName, "name1")
		t.Assert(users[1].Uid, 2)
		t.Assert(users[1].NickName, "name2")
	})
}

func Test_Structs_SliceParameter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			NickName string
		}
		var users []User
		params := g.Slice别名{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := 转换类.Structs(params, users)
		t.AssertNE(err, nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid      int
			NickName string
		}
		type A struct {
			Users []User
		}
		var a A
		params := g.Slice别名{
			g.Map{
				"uid":       1,
				"nick-name": "name1",
			},
			g.Map{
				"uid":       2,
				"nick-name": "name2",
			},
		}
		err := 转换类.Structs(params, a.Users)
		t.AssertNE(err, nil)
	})
}

func Test_Structs_DirectReflectSet(t *testing.T) {
	type A struct {
		Id   int
		Name string
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			a = []*A{
				{Id: 1, Name: "john"},
				{Id: 2, Name: "smith"},
			}
			b []*A
		)
		err := 转换类.Structs(a, &b)
		t.AssertNil(err)
		t.AssertEQ(a, b)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			a = []A{
				{Id: 1, Name: "john"},
				{Id: 2, Name: "smith"},
			}
			b []A
		)
		err := 转换类.Structs(a, &b)
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
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			array []*B
		)
		err := 转换类.Structs(g.Slice别名{
			g.Map{"id": nil, "name": "john"},
			g.Map{"id": nil, "name": "smith"},
		}, &array)
		t.AssertNil(err)
		t.Assert(len(array), 2)
		t.Assert(array[0].Name, "john")
		t.Assert(array[1].Name, "smith")
	})
}
