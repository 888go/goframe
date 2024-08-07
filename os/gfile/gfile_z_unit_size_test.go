// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类_test

import (
	"testing"

	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_Size(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			paths1 string = "/testfile_t1.txt"
			sizes  int64
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = gfile.X取大小(testpath() + paths1)
		t.Assert(sizes, 14)

		sizes = gfile.X取大小("")
		t.Assert(sizes, 0)

	})
}

func Test_SizeFormat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			paths1 = "/testfile_t1.txt"
			sizes  string
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = gfile.X取大小并易读格式(testpath() + paths1)
		t.Assert(sizes, "14.00B")

		sizes = gfile.X取大小并易读格式("")
		t.Assert(sizes, "0.00B")

	})
}

func Test_StrToSize(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gfile.X易读格式转字节长度("0.00B"), 0)
		t.Assert(gfile.X易读格式转字节长度("16.00B"), 16)
		t.Assert(gfile.X易读格式转字节长度("1.00K"), 1024)
		t.Assert(gfile.X易读格式转字节长度("1.00KB"), 1024)
		t.Assert(gfile.X易读格式转字节长度("1.00KiloByte"), 1024)
		t.Assert(gfile.X易读格式转字节长度("15.26M"), gconv.X取整数64位(15.26*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("15.26MB"), gconv.X取整数64位(15.26*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("1.49G"), gconv.X取整数64位(1.49*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("1.49GB"), gconv.X取整数64位(1.49*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("8.73T"), gconv.X取整数64位(8.73*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("8.73TB"), gconv.X取整数64位(8.73*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("8.53P"), gconv.X取整数64位(8.53*1024*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("8.53PB"), gconv.X取整数64位(8.53*1024*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("8.01EB"), gconv.X取整数64位(8.01*1024*1024*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("0.01ZB"), gconv.X取整数64位(0.01*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("0.01YB"), gconv.X取整数64位(0.01*1024*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("0.01BB"), gconv.X取整数64位(0.01*1024*1024*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(gfile.X易读格式转字节长度("0.01AB"), gconv.X取整数64位(-1))
		t.Assert(gfile.X易读格式转字节长度("123456789"), 123456789)
	})
}

func Test_FormatSize(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gfile.X字节长度转易读格式(0), "0.00B")
		t.Assert(gfile.X字节长度转易读格式(16), "16.00B")

		t.Assert(gfile.X字节长度转易读格式(1024), "1.00K")

		t.Assert(gfile.X字节长度转易读格式(16000000), "15.26M")

		t.Assert(gfile.X字节长度转易读格式(1600000000), "1.49G")

		t.Assert(gfile.X字节长度转易读格式(9600000000000), "8.73T")
		t.Assert(gfile.X字节长度转易读格式(9600000000000000), "8.53P")
	})
}

func Test_ReadableSize(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		var (
			paths1 string = "/testfile_t1.txt"
		)
		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)
		t.Assert(gfile.ReadableSize别名(testpath()+paths1), "14.00B")
		t.Assert(gfile.ReadableSize别名(""), "0.00B")

	})
}
