// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid_test
import (
	"context"
	"testing"
	
	"github.com/888go/goframe/i18n/gi18n"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gvalid"
	)

func TestValidator_I18n(t *testing.T) {
	var (
		err         gvalid.Error
		i18nManager = gi18n.New(gi18n.Options{Path: gtest.DataPath("i18n")})
		ctxCn       = gi18n.WithLanguage(context.TODO(), "cn")
		validator   = gvalid.New().I18n(i18nManager)
	)
	gtest.C(t, func(t *gtest.T) {
		err = validator.Rules("required").Data("").Run(ctx)
		t.Assert(err.String(), "The field is required")

		err = validator.Rules("required").Data("").Run(ctxCn)
		t.Assert(err.String(), "字段不能为空")
	})
	gtest.C(t, func(t *gtest.T) {
		err = validator.Rules("required").Messages("CustomMessage").Data("").Run(ctxCn)
		t.Assert(err.String(), "自定义错误")
	})
	gtest.C(t, func(t *gtest.T) {
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
