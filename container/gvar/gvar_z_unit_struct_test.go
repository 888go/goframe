// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 泛型类_test

import (
	"testing"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func TestVar_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type StTest struct {
			Test int
		}

		Kv := make(map[string]int, 1)
		Kv["Test"] = 100

		testObj := &StTest{}

		objOne := gvar.X创建(Kv, true)

		objOne.Struct(testObj)

		t.Assert(testObj.Test, Kv["Test"])
	})
	gtest.C(t, func(t *gtest.T) {
		type StTest struct {
			Test int8
		}
		o := &StTest{}
		v := gvar.X创建(g.Slice别名{"Test", "-25"})
		v.Struct(o)
		t.Assert(o.Test, -25)
	})
}

func TestVar_Var_Attribute_Struct(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Uid  int
			Name string
		}
		user := new(User)
		err := gconv.Struct(
			g.Map{
				"uid":  gvar.X创建(1),
				"name": gvar.X创建("john"),
			}, user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
	gtest.C(t, func(t *gtest.T) {
		type User struct {
			Uid  int
			Name string
		}
		var user *User
		err := gconv.Struct(
			g.Map{
				"uid":  gvar.X创建(1),
				"name": gvar.X创建("john"),
			}, &user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
}
