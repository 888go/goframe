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

func Test_Trim(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.Trim(" 123456\n "), "123456")
		t.Assert(gstr.Trim("#123456#;", "#;"), "123456")
	})
}

func Test_TrimStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimStr("gogo我爱gogo", "go"), "我爱")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimStr("gogo我爱gogo", "go", 1), "go我爱go")
		t.Assert(gstr.TrimStr("gogo我爱gogo", "go", 2), "我爱")
		t.Assert(gstr.TrimStr("gogo我爱gogo", "go", -1), "我爱")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimStr("啊我爱中国人啊", "啊"), "我爱中国人")
	})
}

func Test_TrimRight(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimRight(" 123456\n "), " 123456")
		t.Assert(gstr.TrimRight("#123456#;", "#;"), "#123456")
	})
}

func Test_TrimRightStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimRightStr("gogo我爱gogo", "go"), "gogo我爱")
		t.Assert(gstr.TrimRightStr("gogo我爱gogo", "go我爱gogo"), "go")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimRightStr("gogo我爱gogo", "go", 1), "gogo我爱go")
		t.Assert(gstr.TrimRightStr("gogo我爱gogo", "go", 2), "gogo我爱")
		t.Assert(gstr.TrimRightStr("gogo我爱gogo", "go", -1), "gogo我爱")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimRightStr("我爱中国人", "人"), "我爱中国")
		t.Assert(gstr.TrimRightStr("我爱中国人", "爱中国人"), "我")
	})
}

func Test_TrimLeft(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimLeft(" \r123456\n "), "123456\n ")
		t.Assert(gstr.TrimLeft("#;123456#;", "#;"), "123456#;")
	})
}

func Test_TrimLeftStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimLeftStr("gogo我爱gogo", "go"), "我爱gogo")
		t.Assert(gstr.TrimLeftStr("gogo我爱gogo", "gogo我爱go"), "go")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimLeftStr("gogo我爱gogo", "go", 1), "go我爱gogo")
		t.Assert(gstr.TrimLeftStr("gogo我爱gogo", "go", 2), "我爱gogo")
		t.Assert(gstr.TrimLeftStr("gogo我爱gogo", "go", -1), "我爱gogo")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimLeftStr("我爱中国人", "我爱"), "中国人")
		t.Assert(gstr.TrimLeftStr("我爱中国人", "我爱中国"), "人")
	})
}

func Test_TrimAll(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimAll("gogo我go\n爱gogo\n", "go"), "我爱")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimAll("gogo\n我go爱gogo", "go"), "我爱")
		t.Assert(gstr.TrimAll("gogo\n我go爱gogo\n", "go"), "我爱")
		t.Assert(gstr.TrimAll("gogo\n我go\n爱gogo", "go"), "我爱")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.TrimAll("啊我爱\n啊中国\n人啊", "啊"), "我爱中国人")
	})
}
