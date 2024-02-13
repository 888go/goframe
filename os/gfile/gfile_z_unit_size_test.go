// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"testing"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Size(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1 string = "/testfile_t1.txt"
			sizes  int64
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = 文件类.X取大小(testpath() + paths1)
		t.Assert(sizes, 14)

		sizes = 文件类.X取大小("")
		t.Assert(sizes, 0)

	})
}

func Test_SizeFormat(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			paths1 = "/testfile_t1.txt"
			sizes  string
		)

		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)

		sizes = 文件类.X取大小并易读格式(testpath() + paths1)
		t.Assert(sizes, "14.00B")

		sizes = 文件类.X取大小并易读格式("")
		t.Assert(sizes, "0.00B")

	})
}

func Test_StrToSize(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文件类.X易读格式转字节长度("0.00B"), 0)
		t.Assert(文件类.X易读格式转字节长度("16.00B"), 16)
		t.Assert(文件类.X易读格式转字节长度("1.00K"), 1024)
		t.Assert(文件类.X易读格式转字节长度("1.00KB"), 1024)
		t.Assert(文件类.X易读格式转字节长度("1.00KiloByte"), 1024)
		t.Assert(文件类.X易读格式转字节长度("15.26M"), 转换类.X取整数64位(15.26*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("15.26MB"), 转换类.X取整数64位(15.26*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("1.49G"), 转换类.X取整数64位(1.49*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("1.49GB"), 转换类.X取整数64位(1.49*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("8.73T"), 转换类.X取整数64位(8.73*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("8.73TB"), 转换类.X取整数64位(8.73*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("8.53P"), 转换类.X取整数64位(8.53*1024*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("8.53PB"), 转换类.X取整数64位(8.53*1024*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("8.01EB"), 转换类.X取整数64位(8.01*1024*1024*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("0.01ZB"), 转换类.X取整数64位(0.01*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("0.01YB"), 转换类.X取整数64位(0.01*1024*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("0.01BB"), 转换类.X取整数64位(0.01*1024*1024*1024*1024*1024*1024*1024*1024*1024))
		t.Assert(文件类.X易读格式转字节长度("0.01AB"), 转换类.X取整数64位(-1))
		t.Assert(文件类.X易读格式转字节长度("123456789"), 123456789)
	})
}

func Test_FormatSize(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文件类.X字节长度转易读格式(0), "0.00B")
		t.Assert(文件类.X字节长度转易读格式(16), "16.00B")

		t.Assert(文件类.X字节长度转易读格式(1024), "1.00K")

		t.Assert(文件类.X字节长度转易读格式(16000000), "15.26M")

		t.Assert(文件类.X字节长度转易读格式(1600000000), "1.49G")

		t.Assert(文件类.X字节长度转易读格式(9600000000000), "8.73T")
		t.Assert(文件类.X字节长度转易读格式(9600000000000000), "8.53P")
	})
}

func Test_ReadableSize(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {

		var (
			paths1 string = "/testfile_t1.txt"
		)
		createTestFile(paths1, "abcdefghijklmn")
		defer delTestFiles(paths1)
		t.Assert(文件类.ReadableSize别名(testpath()+paths1), "14.00B")
		t.Assert(文件类.ReadableSize别名(""), "0.00B")

	})
}
