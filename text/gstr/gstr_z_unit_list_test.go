// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package gstr_test
import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
	)

func Test_List2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.List2("1:2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.List2("1:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.List2("1", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.List2("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.List2("1:2:3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2:3")
	})
}

func Test_ListAndTrim2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.ListAndTrim2("1::2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.ListAndTrim2("1::", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.ListAndTrim2("1:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.ListAndTrim2("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2 := gstr.ListAndTrim2("1::2::3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2:3")
	})
}

func Test_List3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("1:2:3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "3")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("1:2:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("1:2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("1:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("1", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.List3("1:2:3:4", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "3:4")
	})
}

func Test_ListAndTrim3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.ListAndTrim3("1::2:3", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "3")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.ListAndTrim3("1::2:", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.ListAndTrim3("1::2", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "2")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.ListAndTrim3("1::", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.ListAndTrim3("1::", ":")
		t.Assert(p1, "1")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
	gtest.C(t, func(t *gtest.T) {
		p1, p2, p3 := gstr.ListAndTrim3("", ":")
		t.Assert(p1, "")
		t.Assert(p2, "")
		t.Assert(p3, "")
	})
}
