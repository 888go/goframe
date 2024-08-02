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
