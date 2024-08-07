// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

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

		t.Assert(gfile.X取修改时间秒(testpath()+file1), fileobj.ModTime())
		t.Assert(gfile.X取修改时间秒(""), "")
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
			gfile.X取修改时间戳毫秒(testpath()+file1),
			fileobj.ModTime().UnixNano()/1000000,
		)
		t.Assert(gfile.X取修改时间戳毫秒(""), -1)
	})
}
