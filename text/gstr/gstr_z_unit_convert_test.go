// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文本类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_OctStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X八进制到文本(`\346\200\241`), "怡")
	})
}

func Test_WordWrap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X按字符数量换行("12 34", 2, "<br>"), "12<br>34")
		t.Assert(文本类.X按字符数量换行("12 34", 2, "\n"), "12\n34")
		t.Assert(文本类.X按字符数量换行("我爱 GF", 2, "\n"), "我爱\nGF")
		t.Assert(文本类.X按字符数量换行("A very long woooooooooooooooooord. and something", 7, "<br>"),
			"A very<br>long<br>woooooooooooooooooord.<br>and<br>something")
	})
	// 中文标点符号
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			br      = "                       "
			content = "    DelRouteKeyIPv6    删除VPC内的服务的Route信息;和DelRouteIPv6接口相比，这个接口可以删除满足条件的多条RS\n"
			length  = 120
		)
		wrappedContent := 文本类.X按字符数量换行(content, length, "\n"+br)
		t.Assert(wrappedContent, `    DelRouteKeyIPv6    删除VPC内的服务的Route信息;和DelRouteIPv6接口相比，
                       这个接口可以删除满足条件的多条RS
`)
	})
}
