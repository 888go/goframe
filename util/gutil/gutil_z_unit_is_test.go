// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类_test

import (
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gutil "github.com/888go/goframe/util/gutil"
)

func Test_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gutil.IsEmpty(1), false)
	})
}

func Test_IsTypeOf(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gutil.IsTypeOf(1, 0), true)
		t.Assert(gutil.IsTypeOf(1.1, 0.1), true)
		t.Assert(gutil.IsTypeOf(1.1, 1), false)
		t.Assert(gutil.IsTypeOf(true, false), true)
		t.Assert(gutil.IsTypeOf(true, 1), false)
	})
	gtest.C(t, func(t *gtest.T) {
		type A struct {
			Name string
		}
		type B struct {
			Name string
		}
		t.Assert(gutil.IsTypeOf(1, A{}), false)
		t.Assert(gutil.IsTypeOf(A{}, B{}), false)
		t.Assert(gutil.IsTypeOf(A{Name: "john"}, &A{Name: "john"}), false)
		t.Assert(gutil.IsTypeOf(A{Name: "john"}, A{Name: "john"}), true)
		t.Assert(gutil.IsTypeOf(A{Name: "john"}, A{}), true)
		t.Assert(gutil.IsTypeOf(&A{Name: "john"}, &A{}), true)
		t.Assert(gutil.IsTypeOf(&A{Name: "john"}, &B{}), false)
		t.Assert(gutil.IsTypeOf(A{Name: "john"}, B{Name: "john"}), false)
	})
}
