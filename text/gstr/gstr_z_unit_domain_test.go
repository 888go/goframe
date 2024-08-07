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

func Test_IsSubDomain(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		main := "goframe.org"
		t.Assert(gstr.X是否为子域名("goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org:8080", main), true)
		t.Assert(gstr.X是否为子域名("johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.goframe.org"
		t.Assert(gstr.X是否为子域名("goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.goframe.org:80", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org", main), false)
		t.Assert(gstr.X是否为子域名("johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org"
		t.Assert(gstr.X是否为子域名("goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.X是否为子域名("s.s.s.goframe.org", main), false)
		t.Assert(gstr.X是否为子域名("johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.X是否为子域名("goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.X是否为子域名("s.s.s.goframe.org", main), false)
		t.Assert(gstr.X是否为子域名("johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.s.johng.cn", main), false)
	})

	gtest.C(t, func(t *gtest.T) {
		main := "*.*.goframe.org:8080"
		t.Assert(gstr.X是否为子域名("goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org", main), true)
		t.Assert(gstr.X是否为子域名("s.s.goframe.org:8000", main), true)
		t.Assert(gstr.X是否为子域名("s.s.s.goframe.org", main), false)
		t.Assert(gstr.X是否为子域名("johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.johng.cn", main), false)
		t.Assert(gstr.X是否为子域名("s.s.johng.cn", main), false)
	})
	gtest.C(t, func(t *gtest.T) {
		main := "s.goframe.org"
		t.Assert(gstr.X是否为子域名("goframe.org", main), false)
	})
}
