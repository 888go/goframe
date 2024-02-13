// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 文本类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_Trim(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首尾符并含空白(" 123456\n "), "123456")
		t.Assert(文本类.X过滤首尾符并含空白("#123456#;", "#;"), "123456")
	})
}

func Test_TrimStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首尾("gogo我爱gogo", "go"), "我爱")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首尾("gogo我爱gogo", "go", 1), "go我爱go")
		t.Assert(文本类.X过滤首尾("gogo我爱gogo", "go", 2), "我爱")
		t.Assert(文本类.X过滤首尾("gogo我爱gogo", "go", -1), "我爱")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首尾("啊我爱中国人啊", "啊"), "我爱中国人")
	})
}

func Test_TrimRight(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤尾字符并含空白(" 123456\n "), " 123456")
		t.Assert(文本类.X过滤尾字符并含空白("#123456#;", "#;"), "#123456")
	})
}

func Test_TrimRightStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤尾字符("gogo我爱gogo", "go"), "gogo我爱")
		t.Assert(文本类.X过滤尾字符("gogo我爱gogo", "go我爱gogo"), "go")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤尾字符("gogo我爱gogo", "go", 1), "gogo我爱go")
		t.Assert(文本类.X过滤尾字符("gogo我爱gogo", "go", 2), "gogo我爱")
		t.Assert(文本类.X过滤尾字符("gogo我爱gogo", "go", -1), "gogo我爱")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤尾字符("我爱中国人", "人"), "我爱中国")
		t.Assert(文本类.X过滤尾字符("我爱中国人", "爱中国人"), "我")
	})
}

func Test_TrimLeft(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首字符并含空白(" \r123456\n "), "123456\n ")
		t.Assert(文本类.X过滤首字符并含空白("#;123456#;", "#;"), "123456#;")
	})
}

func Test_TrimLeftStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首字符("gogo我爱gogo", "go"), "我爱gogo")
		t.Assert(文本类.X过滤首字符("gogo我爱gogo", "gogo我爱go"), "go")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首字符("gogo我爱gogo", "go", 1), "go我爱gogo")
		t.Assert(文本类.X过滤首字符("gogo我爱gogo", "go", 2), "我爱gogo")
		t.Assert(文本类.X过滤首字符("gogo我爱gogo", "go", -1), "我爱gogo")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤首字符("我爱中国人", "我爱"), "中国人")
		t.Assert(文本类.X过滤首字符("我爱中国人", "我爱中国"), "人")
	})
}

func Test_TrimAll(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤所有字符并含空白("gogo我go\n爱gogo\n", "go"), "我爱")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤所有字符并含空白("gogo\n我go爱gogo", "go"), "我爱")
		t.Assert(文本类.X过滤所有字符并含空白("gogo\n我go爱gogo\n", "go"), "我爱")
		t.Assert(文本类.X过滤所有字符并含空白("gogo\n我go\n爱gogo", "go"), "我爱")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X过滤所有字符并含空白("啊我爱\n啊中国\n人啊", "啊"), "我爱中国人")
	})
}
