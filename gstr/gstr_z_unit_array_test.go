// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package gstr_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gstr"
)

func Test_SearchArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr{"a", "b", "c"}
		t.AssertEQ(gstr.SearchArray(a, "a"), 0)
		t.AssertEQ(gstr.SearchArray(a, "b"), 1)
		t.AssertEQ(gstr.SearchArray(a, "c"), 2)
		t.AssertEQ(gstr.SearchArray(a, "d"), -1)
	})
}

func Test_InArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr{"a", "b", "c"}
		t.AssertEQ(gstr.InArray(a, "a"), true)
		t.AssertEQ(gstr.InArray(a, "b"), true)
		t.AssertEQ(gstr.InArray(a, "c"), true)
		t.AssertEQ(gstr.InArray(a, "d"), false)
	})
}

func Test_PrefixArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr{"a", "b", "c"}
		gstr.PrefixArray(a, "1-")
		t.AssertEQ(a, g.SliceStr{"1-a", "1-b", "1-c"})
	})
}
