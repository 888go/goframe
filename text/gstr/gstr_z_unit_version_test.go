// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package gstr_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
)

func Test_IsGNUVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.IsGNUVersion(""), false)
		t.AssertEQ(gstr.IsGNUVersion("v"), false)
		t.AssertEQ(gstr.IsGNUVersion("v0"), true)
		t.AssertEQ(gstr.IsGNUVersion("v0."), false)
		t.AssertEQ(gstr.IsGNUVersion("v1."), false)
		t.AssertEQ(gstr.IsGNUVersion("v1.1"), true)
		t.AssertEQ(gstr.IsGNUVersion("v1.1.0"), true)
		t.AssertEQ(gstr.IsGNUVersion("v1.1."), false)
		t.AssertEQ(gstr.IsGNUVersion("v1.1.0.0"), false)
		t.AssertEQ(gstr.IsGNUVersion("v0.0.0"), true)
		t.AssertEQ(gstr.IsGNUVersion("v1.1.-1"), false)
		t.AssertEQ(gstr.IsGNUVersion("v1.1.+1"), false)
	})
}

func Test_CompareVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.CompareVersion("1", ""), 1)
		t.AssertEQ(gstr.CompareVersion("", ""), 0)
		t.AssertEQ(gstr.CompareVersion("", "v0.1"), -1)
		t.AssertEQ(gstr.CompareVersion("1", "v0.99"), 1)
		t.AssertEQ(gstr.CompareVersion("v1.0", "v0.99"), 1)
		t.AssertEQ(gstr.CompareVersion("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersion("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersion("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(gstr.CompareVersion("1.0.0", "v1.0.0"), 0)
	})
}

func Test_CompareVersionGo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.CompareVersionGo("1", ""), 1)
		t.AssertEQ(gstr.CompareVersionGo("", ""), 0)
		t.AssertEQ(gstr.CompareVersionGo("", "v0.1"), -1)
		t.AssertEQ(gstr.CompareVersionGo("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersionGo("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.CompareVersionGo("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(gstr.CompareVersionGo("1.0.0", "v1.0.0"), 0)
		t.AssertEQ(gstr.CompareVersionGo("1.0.0", "v1.0"), 0)
		t.AssertEQ(gstr.CompareVersionGo("v0.0.0-20190626092158-b2ccc519800e", "0.0.0-20190626092158"), 0)
		t.AssertEQ(gstr.CompareVersionGo("v0.0.0-20190626092159-b2ccc519800e", "0.0.0-20190626092158"), 1)

// 特别是在Go语言中：
// 特别是：
// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2" // 表示 v1.12.2-0.20200413154443-b17e3a6804fa 版本早于 v1.12.2
// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2" // 表示 v1.12.3-0.20200413154443-b17e3a6804fa 版本晚于 v1.12.2
// md5:685fe05f97473463
		t.AssertEQ(gstr.CompareVersionGo("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2"), -1)
		t.AssertEQ(gstr.CompareVersionGo("v1.12.2", "v1.12.2-0.20200413154443-b17e3a6804fa"), 1)
		t.AssertEQ(gstr.CompareVersionGo("v1.12.3-0.20200413154443-b17e3a6804fa", "v1.12.2"), 1)
		t.AssertEQ(gstr.CompareVersionGo("v1.12.2", "v1.12.3-0.20200413154443-b17e3a6804fa"), -1)
		t.AssertEQ(gstr.CompareVersionGo("v1.12.2-0.20200413154443-b17e3a6804fa", "v0.0.0-20190626092158-b2ccc519800e"), 1)
		t.AssertEQ(gstr.CompareVersionGo("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2-0.20200413154444-b2ccc519800e"), -1)

// 特别是在 Go 语言中：
// 特别是在 v4.20.1+incompatible < v4.20.1 的情况下
// md5:dc6d93f041f6a414
		t.AssertEQ(gstr.CompareVersionGo("v4.20.0+incompatible", "4.20.0"), -1)
		t.AssertEQ(gstr.CompareVersionGo("4.20.0", "v4.20.0+incompatible"), 1)
		t.AssertEQ(gstr.CompareVersionGo("v4.20.0+incompatible", "4.20.1"), -1)
		t.AssertEQ(gstr.CompareVersionGo("v4.20.0+incompatible", "v4.20.0+incompatible"), 0)

	})
}
