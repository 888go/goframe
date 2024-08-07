// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

package 文本类_test

import (
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_IsGNUVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.X版本号是否有效(""), false)
		t.AssertEQ(gstr.X版本号是否有效("v"), false)
		t.AssertEQ(gstr.X版本号是否有效("v0"), true)
		t.AssertEQ(gstr.X版本号是否有效("v0."), false)
		t.AssertEQ(gstr.X版本号是否有效("v1."), false)
		t.AssertEQ(gstr.X版本号是否有效("v1.1"), true)
		t.AssertEQ(gstr.X版本号是否有效("v1.1.0"), true)
		t.AssertEQ(gstr.X版本号是否有效("v1.1."), false)
		t.AssertEQ(gstr.X版本号是否有效("v1.1.0.0"), false)
		t.AssertEQ(gstr.X版本号是否有效("v0.0.0"), true)
		t.AssertEQ(gstr.X版本号是否有效("v1.1.-1"), false)
		t.AssertEQ(gstr.X版本号是否有效("v1.1.+1"), false)
	})
}

func Test_CompareVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.X版本号比较GNU格式("1", ""), 1)
		t.AssertEQ(gstr.X版本号比较GNU格式("", ""), 0)
		t.AssertEQ(gstr.X版本号比较GNU格式("", "v0.1"), -1)
		t.AssertEQ(gstr.X版本号比较GNU格式("1", "v0.99"), 1)
		t.AssertEQ(gstr.X版本号比较GNU格式("v1.0", "v0.99"), 1)
		t.AssertEQ(gstr.X版本号比较GNU格式("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.X版本号比较GNU格式("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.X版本号比较GNU格式("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(gstr.X版本号比较GNU格式("1.0.0", "v1.0.0"), 0)
	})
}

func Test_CompareVersionGo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(gstr.X版本号比较GO格式("1", ""), 1)
		t.AssertEQ(gstr.X版本号比较GO格式("", ""), 0)
		t.AssertEQ(gstr.X版本号比较GO格式("", "v0.1"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("v1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("1.0.1", "v1.1.0"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("1.0.0", "v0.1.0"), 1)
		t.AssertEQ(gstr.X版本号比较GO格式("1.0.0", "v1.0.0"), 0)
		t.AssertEQ(gstr.X版本号比较GO格式("1.0.0", "v1.0"), 0)
		t.AssertEQ(gstr.X版本号比较GO格式("v0.0.0-20190626092158-b2ccc519800e", "0.0.0-20190626092158"), 0)
		t.AssertEQ(gstr.X版本号比较GO格式("v0.0.0-20190626092159-b2ccc519800e", "0.0.0-20190626092158"), 1)

		// Specially in Golang:
		// "v1.12.2-0.20200413154443-b17e3a6804fa" < "v1.12.2"
		// "v1.12.3-0.20200413154443-b17e3a6804fa" > "v1.12.2"
		t.AssertEQ(gstr.X版本号比较GO格式("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("v1.12.2", "v1.12.2-0.20200413154443-b17e3a6804fa"), 1)
		t.AssertEQ(gstr.X版本号比较GO格式("v1.12.3-0.20200413154443-b17e3a6804fa", "v1.12.2"), 1)
		t.AssertEQ(gstr.X版本号比较GO格式("v1.12.2", "v1.12.3-0.20200413154443-b17e3a6804fa"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("v1.12.2-0.20200413154443-b17e3a6804fa", "v0.0.0-20190626092158-b2ccc519800e"), 1)
		t.AssertEQ(gstr.X版本号比较GO格式("v1.12.2-0.20200413154443-b17e3a6804fa", "v1.12.2-0.20200413154444-b2ccc519800e"), -1)

		// Specially in Golang:
		// "v4.20.1+incompatible" < "v4.20.1"
		t.AssertEQ(gstr.X版本号比较GO格式("v4.20.0+incompatible", "4.20.0"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("4.20.0", "v4.20.0+incompatible"), 1)
		t.AssertEQ(gstr.X版本号比较GO格式("v4.20.0+incompatible", "4.20.1"), -1)
		t.AssertEQ(gstr.X版本号比较GO格式("v4.20.0+incompatible", "v4.20.0+incompatible"), 0)

	})
}
