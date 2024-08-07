// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文本类_test

import (
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_OctStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X八进制到文本(`\346\200\241`), "怡")
	})
}

func Test_WordWrap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X按字符数量换行("12 34", 2, "<br>"), "12<br>34")
		t.Assert(gstr.X按字符数量换行("12 34", 2, "\n"), "12\n34")
		t.Assert(gstr.X按字符数量换行("我爱 GF", 2, "\n"), "我爱\nGF")
		t.Assert(gstr.X按字符数量换行("A very long woooooooooooooooooord. and something", 7, "<br>"),
			"A very<br>long<br>woooooooooooooooooord.<br>and<br>something")
	})
	// Chinese Punctuations.
	gtest.C(t, func(t *gtest.T) {
		var (
			br      = "                       "
			content = "    DelRouteKeyIPv6    删除VPC内的服务的Route信息;和DelRouteIPv6接口相比，这个接口可以删除满足条件的多条RS\n"
			length  = 120
		)
		wrappedContent := gstr.X按字符数量换行(content, length, "\n"+br)
		t.Assert(wrappedContent, `    DelRouteKeyIPv6    删除VPC内的服务的Route信息;和DelRouteIPv6接口相比，
                       这个接口可以删除满足条件的多条RS
`)
	})
}
