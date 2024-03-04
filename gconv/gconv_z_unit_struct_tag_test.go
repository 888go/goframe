// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gconv"
)

func Test_StructTag(t *testing.T) {
	type User struct {
		Uid   int
		Name  string
		Pass1 string `orm:"password1"`
		Pass2 string `orm:"password2"`
	}
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"password1": "123",
			"password2": "456",
		}
		if err := gconv.Struct(params1, user); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "",
			Pass2: "",
		})
	})
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		params1 := g.Map{
			"uid":       1,
			"Name":      "john",
			"password1": "123",
			"password2": "456",
		}
		if err := gconv.StructTag(params1, user, "orm"); err != nil {
			t.Error(err)
		}
		t.Assert(user, &User{
			Uid:   1,
			Name:  "john",
			Pass1: "123",
			Pass2: "456",
		})
	})
	gtest.C(t, func(t *gtest.T) {
		user := new(User)
		params2 := g.Map{
			"uid":       2,
			"name":      "smith",
			"password1": "111",
			"password2": "222",
		}
		if err := gconv.StructTag(params2, user, "orm"); err != nil {
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
