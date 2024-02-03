// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gvalid_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gvalid"
)

type UserCreateReq struct {
	g.Meta `v:"UserCreateReq"`
	Name   string
	Pass   string
}

func RuleUserCreateReq(ctx context.Context, in gvalid.RuleFuncInput) error {
	var req *UserCreateReq
	if err := in.Data.Scan(&req); err != nil {
		return gerror.Wrap(err, `Scan data to UserCreateReq failed`)
	}
	return gerror.Newf(`The name "%s" is already token by others`, req.Name)
}

func Test_Meta(t *testing.T) {
	var user = &UserCreateReq{
		Name: "john",
		Pass: "123456",
	}

	gtest.C(t, func(t *gtest.T) {
		err := g.Validator().RuleFunc("UserCreateReq", RuleUserCreateReq).
			Data(user).
			Assoc(g.Map{
				"Name": "john smith",
			}).Run(ctx)
		t.Assert(err.String(), `The name "john smith" is already token by others`)
	})
}
