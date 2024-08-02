// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 压缩类_test

import (
	"testing"

	gcompress "github.com/888go/goframe/encoding/gcompress"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
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
		data, _ := gcompress.Gzip(arr)
		t.Assert(data, gzip)

		data, _ = gcompress.UnGzip(gzip)
		t.Assert(data, arr)

		data, _ = gcompress.UnGzip(gzip[1:])
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
		err := gcompress.GzipFile(srcPath, dstPath1, 9)
		t.AssertNil(err)
		defer gfile.Remove(dstPath1)
		t.Assert(gfile.Exists(dstPath1), true)

		// Decompress.
		err = gcompress.UnGzipFile(dstPath1, dstPath2)
		t.AssertNil(err)
		defer gfile.Remove(dstPath2)
		t.Assert(gfile.Exists(dstPath2), true)

		t.Assert(gfile.GetContents(srcPath), gfile.GetContents(dstPath2))
	})
}
