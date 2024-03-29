// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类_test

import (
	"testing"
	
	"github.com/888go/goframe/gcompress"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
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

	gtest.C(t, func(t *gtest.T) {
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
		srcPath  = gtest.DataPath("gzip", "file.txt")
		dstPath1 = gfile.Temp(gtime.TimestampNanoStr(), "gzip.zip")
		dstPath2 = gfile.Temp(gtime.TimestampNanoStr(), "file.txt")
	)

	// Compress.
	gtest.C(t, func(t *gtest.T) {
		err := 压缩类.Gzip压缩文件(srcPath, dstPath1, 9)
		t.AssertNil(err)
		defer gfile.Remove(dstPath1)
		t.Assert(gfile.Exists(dstPath1), true)

		// Decompress.
		err = 压缩类.Gzip解压文件(dstPath1, dstPath2)
		t.AssertNil(err)
		defer gfile.Remove(dstPath2)
		t.Assert(gfile.Exists(dstPath2), true)

		t.Assert(gfile.GetContents(srcPath), gfile.GetContents(dstPath2))
	})
}
