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

	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_List2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份("1:2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份("1:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份("1", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份("1:2:3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2:3")
	})
}

func Test_ListAndTrim2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份并忽略空值("1::2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份并忽略空值("1::", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份并忽略空值("1:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份并忽略空值("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.X分割2份并忽略空值("1::2::3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2:3")
	})
}

func Test_List3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("1:2:3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "3")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("1:2:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("1:2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("1:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("1", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份("1:2:3:4", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "3:4")
	})
}

func Test_ListAndTrim3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份并忽略空值("1::2:3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "3")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份并忽略空值("1::2:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份并忽略空值("1::2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份并忽略空值("1::", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份并忽略空值("1::", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.X分割3份并忽略空值("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
}
