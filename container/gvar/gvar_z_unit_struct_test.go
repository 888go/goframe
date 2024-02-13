// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 泛型类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func TestVar_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type StTest struct {
			Test int
		}

		Kv := make(map[string]int, 1)
		Kv["Test"] = 100

		testObj := &StTest{}

		objOne := 泛型类.X创建(Kv, true)

		objOne.Struct(testObj)

		t.Assert(testObj.Test, Kv["Test"])
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type StTest struct {
			Test int8
		}
		o := &StTest{}
		v := 泛型类.X创建(g.Slice别名{"Test", "-25"})
		v.Struct(o)
		t.Assert(o.Test, -25)
	})
}

func TestVar_Var_Attribute_Struct(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  int
			Name string
		}
		user := new(User)
		err := 转换类.Struct(
			g.Map{
				"uid":  泛型类.X创建(1),
				"name": 泛型类.X创建("john"),
			}, user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type User struct {
			Uid  int
			Name string
		}
		var user *User
		err := 转换类.Struct(
			g.Map{
				"uid":  泛型类.X创建(1),
				"name": 泛型类.X创建("john"),
			}, &user)
		t.AssertNil(err)
		t.Assert(user.Uid, 1)
		t.Assert(user.Name, "john")
	})
}
