// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 效验类_test

import (
	"context"
	"testing"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gvalid "github.com/888go/goframe/util/gvalid"
)

func Test_CheckMap1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		data := map[string]interface{}{
			"id":   "0",
			"name": "john",
		}
		rules := map[string]string{
			"id":   "required|between:1,100",
			"name": "required|length:6,16",
		}
		if m := g.X效验类().Data(data).Rules(rules).Run(context.TODO()); m == nil {
			t.Error("CheckMap校验失败")
		} else {
			t.Assert(len(m.Maps()), 2)
			t.Assert(m.Maps()["id"]["between"], "The id value `0` must be between 1 and 100")
			t.Assert(m.Maps()["name"]["length"], "The name value `john` length must be between 6 and 16")
		}
	})
}

func Test_CheckMap2(t *testing.T) {
	var params interface{}
	gtest.C(t, func(t *gtest.T) {
		if err := g.X效验类().Data(params).Run(context.TODO()); err == nil {
			t.AssertNil(err)
		}
	})

	kvmap := map[string]interface{}{
		"id":   "0",
		"name": "john",
	}
	rules := map[string]string{
		"id":   "required|between:1,100",
		"name": "required|length:6,16",
	}
	msgs := gvalid.CustomMsg{
		"id": "ID不能为空|ID范围应当为{min}到{max}",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为{min}到{max}个字符",
		},
	}
	if m := g.X效验类().Data(kvmap).Rules(rules).Messages(msgs).Run(context.TODO()); m == nil {
		t.Error("CheckMap校验失败")
	}

	kvmap = map[string]interface{}{
		"id":   "1",
		"name": "john",
	}
	rules = map[string]string{
		"id":   "required|between:1,100",
		"name": "required|length:4,16",
	}
	msgs = map[string]interface{}{
		"id": "ID不能为空|ID范围应当为{min}到{max}",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为{min}到{max}个字符",
		},
	}
	if m := g.X效验类().Data(kvmap).Rules(rules).Messages(msgs).Run(context.TODO()); m != nil {
		t.Error(m)
	}

	kvmap = map[string]interface{}{
		"id":   "1",
		"name": "john",
	}
	rules = map[string]string{
		"id":   "",
		"name": "",
	}
	msgs = map[string]interface{}{
		"id": "ID不能为空|ID范围应当为{min}到{max}",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为{min}到{max}个字符",
		},
	}
	if m := g.X效验类().Data(kvmap).Rules(rules).Messages(msgs).Run(context.TODO()); m != nil {
		t.Error(m)
	}

	kvmap = map[string]interface{}{
		"id":   "1",
		"name": "john",
	}
	rules2 := []string{
		"@required|between:1,100",
		"@required|length:4,16",
	}
	msgs = map[string]interface{}{
		"id": "ID不能为空|ID范围应当为{min}到{max}",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为{min}到{max}个字符",
		},
	}
	if m := g.X效验类().Data(kvmap).Rules(rules2).Messages(msgs).Run(context.TODO()); m != nil {
		t.Error(m)
	}

	kvmap = map[string]interface{}{
		"id":   "1",
		"name": "john",
	}
	rules2 = []string{
		"id@required|between:1,100",
		"name@required|length:4,16#名称不能为空|",
	}
	msgs = map[string]interface{}{
		"id": "ID不能为空|ID范围应当为{min}到{max}",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为{min}到{max}个字符",
		},
	}
	if m := g.X效验类().Data(kvmap).Rules(rules2).Messages(msgs).Run(context.TODO()); m != nil {
		t.Error(m)
	}

	kvmap = map[string]interface{}{
		"id":   "1",
		"name": "john",
	}
	rules2 = []string{
		"id@required|between:1,100",
		"name@required|length:4,16#名称不能为空",
	}
	msgs = map[string]interface{}{
		"id": "ID不能为空|ID范围应当为{min}到{max}",
		"name": map[string]string{
			"required": "名称不能为空",
			"length":   "名称长度为{min}到{max}个字符",
		},
	}
	if m := g.X效验类().Data(kvmap).Rules(rules2).Messages(msgs).Run(context.TODO()); m != nil {
		t.Error(m)
	}
}

func Test_CheckMapWithNilAndNotRequiredField(t *testing.T) {
	data := map[string]interface{}{
		"id": "1",
	}
	rules := map[string]string{
		"id":   "required",
		"name": "length:4,16",
	}
	if m := g.X效验类().Data(data).Rules(rules).Run(context.TODO()); m != nil {
		t.Error(m)
	}
}

func Test_Sequence(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		params := map[string]interface{}{
			"passport":  "",
			"password":  "123456",
			"password2": "1234567",
		}
		rules := []string{
			"passport@required|length:6,16#账号不能为空|账号长度应当在{min}到{max}之间",
			"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在{min}到{max}之间|两次密码输入不相等",
			"password2@required|length:6,16#",
		}
		err := g.X效验类().Data(params).Rules(rules).Run(context.TODO())
		t.AssertNE(err, nil)
		t.Assert(len(err.Map()), 2)
		t.Assert(err.Map()["required"], "账号不能为空")
		t.Assert(err.Map()["length"], "账号长度应当在6到16之间")
		t.Assert(len(err.Maps()), 2)

		t.Assert(len(err.Items()), 2)
		t.Assert(err.Items()[0]["passport"]["length"], "账号长度应当在6到16之间")
		t.Assert(err.Items()[0]["passport"]["required"], "账号不能为空")
		t.Assert(err.Items()[1]["password"]["same"], "两次密码输入不相等")

		t.Assert(err.String(), "账号不能为空; 账号长度应当在6到16之间; 两次密码输入不相等")
		t.Assert(err.Strings(), []string{"账号不能为空", "账号长度应当在6到16之间", "两次密码输入不相等"})

		k, m := err.FirstItem()
		t.Assert(k, "passport")
		t.Assert(m, err.Map())

		r, s := err.FirstRule()
		t.Assert(r, "required")
		t.Assert(s, "账号不能为空")

		t.Assert(gerror.X取当前错误(err), "账号不能为空")
	})
}

func Test_Map_Bail(t *testing.T) {
	// global bail
	gtest.C(t, func(t *gtest.T) {
		params := map[string]interface{}{
			"passport":  "",
			"password":  "123456",
			"password2": "1234567",
		}
		rules := []string{
			"passport@required|length:6,16#账号不能为空|账号长度应当在{min}到{max}之间",
			"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在{min}到{max}之间|两次密码输入不相等",
			"password2@required|length:6,16#",
		}
		err := g.X效验类().Bail().Rules(rules).Data(params).Run(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.String(), "账号不能为空")
	})
			// 全局退出规则：bail. md5:2f9b3353f582cdf5
	gtest.C(t, func(t *gtest.T) {
		params := map[string]interface{}{
			"passport":  "",
			"password":  "123456",
			"password2": "1234567",
		}
		rules := []string{
			"passport@bail|required|length:6,16#|账号不能为空|账号长度应当在{min}到{max}之间",
			"password@required|length:6,16|same:password2#密码不能为空|密码长度应当在{min}到{max}之间|两次密码输入不相等",
			"password2@required|length:6,16#",
		}
		err := g.X效验类().Bail().Rules(rules).Data(params).Run(ctx)
		t.AssertNE(err, nil)
		t.Assert(err.String(), "账号不能为空")
	})
}
