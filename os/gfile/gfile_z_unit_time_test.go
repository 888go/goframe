// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package 文件类_test

import (
	"os"
	"testing"
	"time"

	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_MTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.AssertNil(err)

		t.Assert(gfile.MTime(testpath()+file1), fileobj.ModTime())
		t.Assert(gfile.MTime(""), "")
	})
}

func Test_MTimeMillisecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.AssertNil(err)

		time.Sleep(time.Millisecond * 100)
		t.AssertGE(
			gfile.MTimestampMilli(testpath()+file1),
			fileobj.ModTime().UnixNano()/1000000,
		)
		t.Assert(gfile.MTimestampMilli(""), -1)
	})
}
