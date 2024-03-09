// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 文本类_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gstr"
)

func Test_IsSubDomain(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		main := "goframe.org"
		t.Assert(文本类.X是否为子域名("goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org:8080", main), true)
		t.Assert(文本类.X是否为子域名("johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.goframe.org"
		t.Assert(文本类.X是否为子域名("goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.goframe.org:80", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org", main), false)
		t.Assert(文本类.X是否为子域名("johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org"
		t.Assert(文本类.X是否为子域名("goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org:8000", main), true)
		t.Assert(文本类.X是否为子域名("s.s.s.goframe.org", main), false)
		t.Assert(文本类.X是否为子域名("johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(文本类.X是否为子域名("goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org:8000", main), true)
		t.Assert(文本类.X是否为子域名("s.s.s.goframe.org", main), false)
		t.Assert(文本类.X是否为子域名("johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.s.johng.cn", main), false)
	})

	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(文本类.X是否为子域名("goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(文本类.X是否为子域名("s.s.goframe.org:8000", main), true)
		t.Assert(文本类.X是否为子域名("s.s.s.goframe.org", main), false)
		t.Assert(文本类.X是否为子域名("johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.johng.cn", main), false)
		t.Assert(文本类.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "s.goframe.org"
		t.Assert(文本类.X是否为子域名("goframe.org", main), false)
	})
}
