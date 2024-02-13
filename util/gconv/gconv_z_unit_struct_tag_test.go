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

func Test_StructTag(t *testing.T) {
	type User struct {
		Uid   int
		Name  string
		Pass1 string `orm:"password1"`
		Pass2 string `orm:"password2"`
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"password1": "123",
			"password2": "456",
		}
		if err := 转换类.Struct(params1, user); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "",
			Pass2: "",
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"password1": "123",
			"password2": "456",
		}
		if err := 转换类.StructTag(params1, user, "orm"); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "123",
			Pass2: "456",
		})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		user := new(User)
		params2 := g.Map{
			"uid":       2,
			"name":      "smith",
			"password1": "111",
			"password2": "222",
		}
		if err := 转换类.StructTag(params2, user, "orm"); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   2,
			Name:  "smith",
			Pass1: "111",
			Pass2: "222",
		})
	})
}
