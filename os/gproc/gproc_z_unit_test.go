// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*" -benchmem

package gproc_test

import (
	"testing"

	"coding.net/gogit/go/goframe/os/gctx"
	"coding.net/gogit/go/goframe/os/gproc"
	"coding.net/gogit/go/goframe/test/gtest"
)

func Test_ShellExec(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s, err := gproc.ShellExec(gctx.New(), `echo 123`)
		t.AssertNil(err)
		t.Assert(s, "123\n")
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		_, err := gproc.ShellExec(gctx.New(), `NoneExistCommandCall`)
		t.AssertNE(err, nil)
	})
}
