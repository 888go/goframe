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

func Test_IsGNUVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(文本类.X版本号是否有效(""), false)
		t.AssertEQ(文本类.X版本号是否有效("v"), false)
		t.AssertEQ(文本类.X版本号是否有效("v0"), true)
		t.AssertEQ(文本类.X版本号是否有效("v0."), false)
		t.AssertEQ(文本类.X版本号是否有效("v1."), false)
		t.AssertEQ(文本类.X版本号是否有效("v1.1"), true)
		t.AssertEQ(文本类.X版本号是否有效("v1.1.0"), true)
		t.AssertEQ(文本类.X版本号是否有效("v1.1."), false)
		t.AssertEQ(文本类.X版本号是否有效("v1.1.0.0"), false)
		t.AssertEQ(文本类.X版本号是否有效("v0.0.0"), true)
		t.AssertEQ(文本类.X版本号是否有效("v1.1.-1"), false)
		t.AssertEQ(文本类.X版本号是否有效("v1.1.+1"), false)
	})
}

func Test_CompareVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(文本类.X版本号比较GNU格式("1", ""), 1)
		t.AssertEQ(文本类.X版本号比较GNU格式("", ""), 0)
		t.AssertEQ(文本类.X版本号比较GNU格式("", "v0.1"), -1)
		t.AssertEQ(文本类.X版本号比较GNU格式("1", "v0.99"), 1)
		t.AssertEQ(文本类.X版本号比较GNU格式("v1.0", "v0.99"), 1)
		t.AssertEQ(文本类.X版本号比较GNU格式("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(文本类.X版本号比较GNU格式("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(文本类.X版本号比较GNU格式("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(文本类.X版本号比较GNU格式("1.0.0", "v1.0.0"), 0)
	})
}

func Test_CompareVersionGo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(文本类.X版本号比较GO格式("1", ""), 1)
		t.AssertEQ(文本类.X版本号比较GO格式("", ""), 0)
		t.AssertEQ(文本类.X版本号比较GO格式("", "v0.1"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(文本类.X版本号比较GO格式("1.0.0", "v1.0.0"), 0)
		t.AssertEQ(文本类.X版本号比较GO格式("1.0.0", "v1.0"), 0)
		t.AssertEQ(文本类.X版本号比较GO格式("v0.0.0-20190626092158-b2ccc519800e", "0.0.0-20190626092158"), 0)
		t.AssertEQ(文本类.X版本号比较GO格式("v0.0.0-20190626092159-b2ccc519800e", "0.0.0-20190626092158"), 1)

// 特别地在 Golang 中：
// "v1.12.2-0.20200413154443-b17e3a6804fa" 小于 "v1.12.2"
// "v1.12.3-0.20200413154443-b17e3a6804fa" 大于 "v1.12.2"
// 这段代码注释是关于 Golang 中版本字符串比较的特殊规则：
// 在 Golang 中，对于包含预发布版本号（如 "-0.20200413154443-b17e3a6804fa"）的版本字符串，在进行字符串比较时，主版本号、次版本号和补丁版本号部分会被优先比较。当这部分相同时，带有预发布版本号的版本会认为小于不带预发布版本号的版本。
// 因此，尽管 "v1.12.2-0.20200413154443-b17e3a6804fa" 的主要部分与 "v1.12.2" 相同，但由于其附加了预发布标识，所以在比较中它被认为小于 "v1.12.2"。
// 同样，"v1.12.3-0.20200413154443-b17e3a6804fa" 由于其主版本号部分高于 "v1.12.2"，所以即使它也有预发布版本号，依然会在比较中大于 "v1.12.2"。
		t.AssertEQ(文本类.X版本号比较GO格式("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("v1.12.2", "v1.12.2-0.20200413154443-b17e3a6804fa"), 1)
		t.AssertEQ(文本类.X版本号比较GO格式("v1.12.3-0.20200413154443-b17e3a6804fa", "v1.12.2"), 1)
		t.AssertEQ(文本类.X版本号比较GO格式("v1.12.2", "v1.12.3-0.20200413154443-b17e3a6804fa"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("v1.12.2-0.20200413154443-b17e3a6804fa", "v0.0.0-20190626092158-b2ccc519800e"), 1)
		t.AssertEQ(文本类.X版本号比较GO格式("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2-0.20200413154444-b2ccc519800e"), -1)

// 特别在 Golang 中：
// "v4.20.1+incompatible" < "v4.20.1"
// （译注：这里表示在 Go 语言中，带后缀 "+incompatible" 的版本字符串比较时，会被视为小于不带此后缀的相同主版本号、次版本号和补丁号的版本。这是因为在 Go 语言模块系统中，“+incompatible”表示该版本并非遵循 Go 语义化版本规范，因此在排序时会排在符合规范的版本之后。）
		t.AssertEQ(文本类.X版本号比较GO格式("v4.20.0+incompatible", "4.20.0"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("4.20.0", "v4.20.0+incompatible"), 1)
		t.AssertEQ(文本类.X版本号比较GO格式("v4.20.0+incompatible", "4.20.1"), -1)
		t.AssertEQ(文本类.X版本号比较GO格式("v4.20.0+incompatible", "v4.20.0+incompatible"), 0)

	})
}
