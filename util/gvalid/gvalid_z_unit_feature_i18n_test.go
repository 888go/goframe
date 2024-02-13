// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 效验类_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gvalid"
)

func TestValidator_I18n(t *testing.T) {
	var (
		err         效验类.Error
		i18nManager = gi18n.New(gi18n.Options{Path: 单元测试类.DataPath("i18n")})
		ctxCn       = gi18n.WithLanguage(context.TODO(), "cn")
		validator   = 效验类.New().I18n(i18nManager)
	)
	单元测试类.C(t, func(t *单元测试类.T) {
		err = validator.Rules("required").Data("").Run(ctx)
		t.Assert(err.String(), "The field is required")

		err = validator.Rules("required").Data("").Run(ctxCn)
		t.Assert(err.String(), "字段不能为空")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		err = validator.Rules("required").Messages("CustomMessage").Data("").Run(ctxCn)
		t.Assert(err.String(), "自定义错误")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		type Params struct {
			Page      int `v:"required|min:1         # page is required"`
			Size      int `v:"required|between:1,100 # size is required"`
			ProjectId int `v:"between:1,10000        # project id must between {min}, {max}"`
		}
		obj := &Params{
			Page: 1,
			Size: 10,
		}
		err = validator.Data(obj).Run(ctxCn)
		t.Assert(err.String(), "项目ID必须大于等于1并且要小于等于10000")
	})
}
