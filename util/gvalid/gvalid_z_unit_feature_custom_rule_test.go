// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 效验类_test

import (
	"context"
	"errors"
	"testing"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
	"github.com/888go/goframe/util/gvalid"
)

func Test_CustomRule1(t *testing.T) {
	rule := "custom"
	效验类.RegisterRule(
		rule,
		func(ctx context.Context, in 效验类.RuleFuncInput) error {
			pass := in.Value.String()
			if len(pass) != 6 {
				return errors.New(in.Message)
			}
			m := in.Data.X取Map()
			if m["data"] != pass {
				return errors.New(in.Message)
			}
			return nil
		},
	)

	单元测试类.C(t, func(t *单元测试类.T) {
		err := g.X效验类().Data("123456").Rules(rule).Messages("custom message").Run(ctx)
		t.Assert(err.String(), "custom message")
		err = g.X效验类().Data("123456").Assoc(g.Map{"data": "123456"}).Rules(rule).Messages("custom message").Run(ctx)
		t.AssertNil(err)
	})
	// 结构体验证错误
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@custom#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "123",
			Data:  "123456",
		}
		err := g.X效验类().Data(st).Run(ctx)
		t.Assert(err.String(), "自定义错误")
	})
	// 结构体验证无错误。
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@custom#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "123456",
			Data:  "123456",
		}
		err := g.X效验类().Data(st).Run(ctx)
		t.AssertNil(err)
	})
}

func Test_CustomRule2(t *testing.T) {
	rule := "required-map"
	效验类.RegisterRule(rule, func(ctx context.Context, in 效验类.RuleFuncInput) error {
		m := in.Value.X取Map()
		if len(m) == 0 {
			return errors.New(in.Message)
		}
		return nil
	})
	// Check.
	单元测试类.C(t, func(t *单元测试类.T) {
		errStr := "data map should not be empty"
		t.Assert(g.X效验类().Data(g.Map{}).Messages(errStr).Rules(rule).Run(ctx), errStr)
		t.Assert(g.X效验类().Data(g.Map{"k": "v"}).Rules(rule).Messages(errStr).Run(ctx), nil)
	})
	// 结构体验证错误
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value map[string]string `v:"uid@required-map#自定义错误"`
			Data  string            `p:"data"`
		}
		st := &T{
			Value: map[string]string{},
			Data:  "123456",
		}
		err := g.X效验类().Data(st).Run(ctx)
		t.Assert(err.String(), "自定义错误")
	})
	// 结构体验证无错误。
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value map[string]string `v:"uid@required-map#自定义错误"`
			Data  string            `p:"data"`
		}
		st := &T{
			Value: map[string]string{"k": "v"},
			Data:  "123456",
		}
		err := g.X效验类().Data(st).Run(ctx)
		t.AssertNil(err)
	})
}

func Test_CustomRule_AllowEmpty(t *testing.T) {
	rule := "allow-empty-str"
	效验类.RegisterRule(rule, func(ctx context.Context, in 效验类.RuleFuncInput) error {
		s := in.Value.String()
		if len(s) == 0 || s == "gf" {
			return nil
		}
		return errors.New(in.Message)
	})
	// Check.
	单元测试类.C(t, func(t *单元测试类.T) {
		errStr := "error"
		t.Assert(g.X效验类().Data("").Rules(rule).Messages(errStr).Run(ctx), "")
		t.Assert(g.X效验类().Data("gf").Rules(rule).Messages(errStr).Run(ctx), "")
		t.Assert(g.X效验类().Data("gf2").Rules(rule).Messages(errStr).Run(ctx), errStr)
	})
	// 结构体验证错误
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@allow-empty-str#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "",
			Data:  "123456",
		}
		err := g.X效验类().Data(st).Run(ctx)
		t.AssertNil(err)
	})
	// 结构体验证无错误。
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@allow-empty-str#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "john",
			Data:  "123456",
		}
		err := g.X效验类().Data(st).Run(ctx)
		t.Assert(err.String(), "自定义错误")
	})
}

func TestValidator_RuleFunc(t *testing.T) {
	ruleName := "custom_1"
	ruleFunc := func(ctx context.Context, in 效验类.RuleFuncInput) error {
		pass := in.Value.String()
		if len(pass) != 6 {
			return errors.New(in.Message)
		}
		if m := in.Data.X取Map(); m["data"] != pass {
			return errors.New(in.Message)
		}
		return nil
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		err := g.X效验类().Rules(ruleName).
			Messages("custom message").
			RuleFunc(ruleName, ruleFunc).
			Data("123456").
			Run(ctx)
		t.Assert(err.String(), "custom message")
		err = g.X效验类().
			Rules(ruleName).
			Messages("custom message").
			Data("123456").Assoc(g.Map{"data": "123456"}).
			RuleFunc(ruleName, ruleFunc).
			Run(ctx)
		t.AssertNil(err)
	})
	// 结构体验证错误
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@custom_1#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "123",
			Data:  "123456",
		}
		err := g.X效验类().RuleFunc(ruleName, ruleFunc).Data(st).Run(ctx)
		t.Assert(err.String(), "自定义错误")
	})
	// 结构体验证无错误。
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@custom_1#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "123456",
			Data:  "123456",
		}
		err := g.X效验类().RuleFunc(ruleName, ruleFunc).Data(st).Run(ctx)
		t.AssertNil(err)
	})
}

func TestValidator_RuleFuncMap(t *testing.T) {
	ruleName := "custom_1"
	ruleFunc := func(ctx context.Context, in 效验类.RuleFuncInput) error {
		pass := in.Value.String()
		if len(pass) != 6 {
			return errors.New(in.Message)
		}
		if m := in.Data.X取Map(); m["data"] != pass {
			return errors.New(in.Message)
		}
		return nil
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		err := g.X效验类().
			Rules(ruleName).
			Messages("custom message").
			RuleFuncMap(map[string]效验类.RuleFunc{
				ruleName: ruleFunc,
			}).Data("123456").Run(ctx)
		t.Assert(err.String(), "custom message")
		err = g.X效验类().
			Rules(ruleName).
			Messages("custom message").
			Data("123456").Assoc(g.Map{"data": "123456"}).
			RuleFuncMap(map[string]效验类.RuleFunc{
				ruleName: ruleFunc,
			}).Run(ctx)
		t.AssertNil(err)
	})
	// 结构体验证错误
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@custom_1#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "123",
			Data:  "123456",
		}
		err := g.X效验类().
			RuleFuncMap(map[string]效验类.RuleFunc{
				ruleName: ruleFunc,
			}).Data(st).Run(ctx)
		t.Assert(err.String(), "自定义错误")
	})
	// 结构体验证无错误。
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@custom_1#自定义错误"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "123456",
			Data:  "123456",
		}
		err := g.X效验类().
			RuleFuncMap(map[string]效验类.RuleFunc{
				ruleName: ruleFunc,
			}).Data(st).Run(ctx)
		t.AssertNil(err)
	})
}

func Test_CustomRule_Overwrite(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var rule = "custom-" + uid类.X生成()
		效验类.RegisterRule(rule, func(ctx context.Context, in 效验类.RuleFuncInput) error {
			return 错误类.X创建("1")
		})
		t.Assert(g.X效验类().Rules(rule).Data(1).Run(ctx), "1")
		效验类.RegisterRule(rule, func(ctx context.Context, in 效验类.RuleFuncInput) error {
			return 错误类.X创建("2")
		})
		t.Assert(g.X效验类().Rules(rule).Data(1).Run(ctx), "2")
	})
	g.X调试输出(效验类.GetRegisteredRuleMap())
}

func Test_Issue2499(t *testing.T) {
	ruleName := "required"
	ruleFunc := func(ctx context.Context, in 效验类.RuleFuncInput) error {
		return errors.New(in.Message)
	}
	单元测试类.C(t, func(t *单元测试类.T) {
		type T struct {
			Value string `v:"uid@required"`
			Data  string `p:"data"`
		}
		st := &T{
			Value: "",
			Data:  "123456",
		}
		err := g.X效验类().
			RuleFuncMap(map[string]效验类.RuleFunc{
				ruleName: ruleFunc,
			}).Data(st).Run(ctx)
		t.Assert(err.String(), `The uid field is required`)
	})
}
