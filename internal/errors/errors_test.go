// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package errors_test

import (
	"testing"

	"github.com/888go/goframe/internal/errors"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_IsStackModeBrief(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(errors.IsStackModeBrief(), true)
	})
}
