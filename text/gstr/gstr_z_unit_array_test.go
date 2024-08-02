// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

package 文本类_test

import (
	"testing"

	"github.com/888go/goframe/frame/g"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_SearchArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr{"a", "b", "c"}
		t.AssertEQ(gstr.SearchArray(a, "a"), 0)
		t.AssertEQ(gstr.SearchArray(a, "b"), 1)
		t.AssertEQ(gstr.SearchArray(a, "c"), 2)
		t.AssertEQ(gstr.SearchArray(a, "d"), -1)
	})
}

func Test_InArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr{"a", "b", "c"}
		t.AssertEQ(gstr.InArray(a, "a"), true)
		t.AssertEQ(gstr.InArray(a, "b"), true)
		t.AssertEQ(gstr.InArray(a, "c"), true)
		t.AssertEQ(gstr.InArray(a, "d"), false)
	})
}

func Test_PrefixArray(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		a := g.SliceStr{"a", "b", "c"}
		gstr.PrefixArray(a, "1-")
		t.AssertEQ(a, g.SliceStr{"1-a", "1-b", "1-c"})
	})
}
