// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类_test

import (
	"testing"
	
	"github.com/888go/goframe/encoding/gcompress"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_Gzip_UnGzip(t *testing.T) {
	var (
		src  = "Hello World!!"
		gzip = []byte{
			0x1f, 0x8b, 0x08, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0xff,
			0xf2, 0x48, 0xcd, 0xc9, 0xc9,
			0x57, 0x08, 0xcf, 0x2f, 0xca,
			0x49, 0x51, 0x54, 0x04, 0x04,
			0x00, 0x00, 0xff, 0xff, 0x9d,
			0x24, 0xa8, 0xd1, 0x0d, 0x00,
			0x00, 0x00,
		}
	)

	单元测试类.C(t, func(t *单元测试类.T) {
		arr := []byte(src)
		data, _ := 压缩类.Gzip压缩字节集(arr)
		t.Assert(data, gzip)

		data, _ = 压缩类.Gzip解压字节集(gzip)
		t.Assert(data, arr)

		data, _ = 压缩类.Gzip解压字节集(gzip[1:])
		t.Assert(data, nil)
	})
}

func Test_Gzip_UnGzip_File(t *testing.T) {
	var (
		srcPath  = 单元测试类.DataPath("gzip", "file.txt")
		dstPath1 = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "gzip.zip")
		dstPath2 = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "file.txt")
	)

	// Compress.
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 压缩类.Gzip压缩文件(srcPath, dstPath1, 9)
		t.AssertNil(err)
		defer 文件类.X删除(dstPath1)
		t.Assert(文件类.X是否存在(dstPath1), true)

		// Decompress.
		err = 压缩类.Gzip解压文件(dstPath1, dstPath2)
		t.AssertNil(err)
		defer 文件类.X删除(dstPath2)
		t.Assert(文件类.X是否存在(dstPath2), true)

		t.Assert(文件类.X读文本(srcPath), 文件类.X读文本(dstPath2))
	})
}
