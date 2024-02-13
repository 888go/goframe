// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 文本类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_SearchArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a := g.SliceStr别名{"a", "b", "c"}
		t.AssertEQ(文本类.X数组查找(a, "a"), 0)
		t.AssertEQ(文本类.X数组查找(a, "b"), 1)
		t.AssertEQ(文本类.X数组查找(a, "c"), 2)
		t.AssertEQ(文本类.X数组查找(a, "d"), -1)
	})
}

func Test_InArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a := g.SliceStr别名{"a", "b", "c"}
		t.AssertEQ(文本类.X数组是否存在(a, "a"), true)
		t.AssertEQ(文本类.X数组是否存在(a, "b"), true)
		t.AssertEQ(文本类.X数组是否存在(a, "c"), true)
		t.AssertEQ(文本类.X数组是否存在(a, "d"), false)
	})
}

func Test_PrefixArray(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		a := g.SliceStr别名{"a", "b", "c"}
		文本类.X数组加前缀(a, "1-")
		t.AssertEQ(a, g.SliceStr别名{"1-a", "1-b", "1-c"})
	})
}
