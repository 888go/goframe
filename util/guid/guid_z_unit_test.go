// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*"

package uid类_test

import (
	"testing"

	gset "github.com/888go/goframe/container/gset"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_S(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		set := gset.NewStrSet()
		for i := 0; i < 1000000; i++ {
			s := guid.S()
			t.Assert(set.AddIfNotExist(s), true)
			t.Assert(len(s), 32)
		}
	})
}

func Test_S_Data(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(guid.S([]byte("123"))), 32)
	})
}
