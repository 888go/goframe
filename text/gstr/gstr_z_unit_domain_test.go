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

func Test_IsSubDomain(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		main := "goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8080", main), true)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org:80", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})

	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.IsSubDomain("goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org", main), true)
		t.Assert(gstr.IsSubDomain("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.IsSubDomain("s.s.s.goframe.org", main), false)
		t.Assert(gstr.IsSubDomain("johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.johng.cn", main), false)
		t.Assert(gstr.IsSubDomain("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "s.goframe.org"
		t.Assert(gstr.IsSubDomain("goframe.org", main), false)
	})
}
