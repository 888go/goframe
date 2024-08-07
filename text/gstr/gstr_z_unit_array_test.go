// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 文本类_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_SearchArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr别名{"a", "b", "c"}
		t.AssertEQ(gstr.X切片查找(a, "a"), 0)
		t.AssertEQ(gstr.X切片查找(a, "b"), 1)
		t.AssertEQ(gstr.X切片查找(a, "c"), 2)
		t.AssertEQ(gstr.X切片查找(a, "d"), -1)
	})
}

func Test_InArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr别名{"a", "b", "c"}
		t.AssertEQ(gstr.X切片是否存在(a, "a"), true)
		t.AssertEQ(gstr.X切片是否存在(a, "b"), true)
		t.AssertEQ(gstr.X切片是否存在(a, "c"), true)
		t.AssertEQ(gstr.X切片是否存在(a, "d"), false)
	})
}

func Test_PrefixArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr别名{"a", "b", "c"}
		gstr.X切片加前缀(a, "1-")
		t.AssertEQ(a, g.SliceStr别名{"1-a", "1-b", "1-c"})
	})
}
