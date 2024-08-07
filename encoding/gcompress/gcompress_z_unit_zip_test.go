// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 压缩类_test

import (
	"bytes"
	"testing"

	gcompress "github.com/888go/goframe/encoding/gcompress"
	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_ZipPath(t *testing.T) {
	// file
	gtest.C(t, func(t *gtest.T) {
		srcPath := gtest.DataPath("zip", "path1", "1.txt")
		dstPath := gtest.DataPath("zip", "zip.zip")

		t.Assert(gfile.X是否存在(dstPath), false)
		t.Assert(gcompress.Zip压缩文件(srcPath, dstPath), nil)
		t.Assert(gfile.X是否存在(dstPath), true)
		defer gfile.X删除(dstPath)

				// 解压缩到临时目录。 md5:dca515f4cd33b4f1
		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		t.Assert(gfile.X创建目录(tempDirPath), nil)
		t.Assert(gcompress.Zip解压文件(dstPath, tempDirPath), nil)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "1.txt")),
			gfile.X读文本(srcPath),
		)
	})
	// multiple files
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath1 = gtest.DataPath("zip", "path1", "1.txt")
			srcPath2 = gtest.DataPath("zip", "path2", "2.txt")
			dstPath  = gfile.X取临时目录(gtime.X取文本时间戳纳秒(), "zip.zip")
		)
		if p := gfile.X路径取父目录(dstPath); !gfile.X是否存在(p) {
			t.Assert(gfile.X创建目录(p), nil)
		}

		t.Assert(gfile.X是否存在(dstPath), false)
		err := gcompress.Zip压缩文件(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.X是否存在(dstPath), true)
		defer gfile.X删除(dstPath)

				// 解压缩到另一个临时目录。 md5:33401796f8abba9e
		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		t.Assert(gfile.X创建目录(tempDirPath), nil)
		err = gcompress.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "1.txt")),
			gfile.X读文本(srcPath1),
		)
		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "2.txt")),
			gfile.X读文本(srcPath2),
		)
	})
	// one dir and one file.
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath1 = gtest.DataPath("zip", "path1")
			srcPath2 = gtest.DataPath("zip", "path2", "2.txt")
			dstPath  = gfile.X取临时目录(gtime.X取文本时间戳纳秒(), "zip.zip")
		)
		if p := gfile.X路径取父目录(dstPath); !gfile.X是否存在(p) {
			t.Assert(gfile.X创建目录(p), nil)
		}

		t.Assert(gfile.X是否存在(dstPath), false)
		err := gcompress.Zip压缩文件(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.X是否存在(dstPath), true)
		defer gfile.X删除(dstPath)

				// 解压缩到另一个临时目录。 md5:33401796f8abba9e
		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		t.Assert(gfile.X创建目录(tempDirPath), nil)
		err = gcompress.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path1", "1.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath1, "1.txt")),
		)
		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "2.txt")),
			gfile.X读文本(srcPath2),
		)
	})
	// directory.
	gtest.C(t, func(t *gtest.T) {
		srcPath := gtest.DataPath("zip")
		dstPath := gtest.DataPath("zip", "zip.zip")

		pwd := gfile.X取当前工作目录()
		err := gfile.X设置当前工作目录(srcPath)
		defer gfile.X设置当前工作目录(pwd)
		t.AssertNil(err)

		t.Assert(gfile.X是否存在(dstPath), false)
		err = gcompress.Zip压缩文件(srcPath, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.X是否存在(dstPath), true)
		defer gfile.X删除(dstPath)

		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(tempDirPath)
		t.AssertNil(err)

		err = gcompress.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "zip", "path1", "1.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "zip", "path2", "2.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
		// 使用字符'，'连接多个目录路径。 md5:d801a18d5afe6f27
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath  = gtest.DataPath("zip")
			srcPath1 = gtest.DataPath("zip", "path1")
			srcPath2 = gtest.DataPath("zip", "path2")
			dstPath  = gtest.DataPath("zip", "zip.zip")
		)
		pwd := gfile.X取当前工作目录()
		err := gfile.X设置当前工作目录(srcPath)
		defer gfile.X设置当前工作目录(pwd)
		t.AssertNil(err)

		t.Assert(gfile.X是否存在(dstPath), false)
		err = gcompress.Zip压缩文件(srcPath1+", "+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.X是否存在(dstPath), true)
		defer gfile.X删除(dstPath)

		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(tempDirPath)
		t.AssertNil(err)

		zipContent := gfile.X读字节集(dstPath)
		t.AssertGT(len(zipContent), 0)
		err = gcompress.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path1", "1.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path2", "2.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathWriter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath  = gtest.DataPath("zip")
			srcPath1 = gtest.DataPath("zip", "path1")
			srcPath2 = gtest.DataPath("zip", "path2")
		)
		pwd := gfile.X取当前工作目录()
		err := gfile.X设置当前工作目录(srcPath)
		defer gfile.X设置当前工作目录(pwd)
		t.AssertNil(err)

		writer := bytes.NewBuffer(nil)
		t.Assert(writer.Len(), 0)
		err = gcompress.Zip压缩文件到Writer(srcPath1+", "+srcPath2, writer)
		t.AssertNil(err)
		t.AssertGT(writer.Len(), 0)

		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(tempDirPath)
		t.AssertNil(err)

		zipContent := writer.Bytes()
		t.AssertGT(len(zipContent), 0)
		err = gcompress.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path1", "1.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path2", "2.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath  = gtest.DataPath("zip")
			srcPath1 = gtest.DataPath("zip", "path1")
			srcPath2 = gtest.DataPath("zip", "path2")
		)
		pwd := gfile.X取当前工作目录()
		err := gfile.X设置当前工作目录(srcPath)
		defer gfile.X设置当前工作目录(pwd)
		t.AssertNil(err)

		tempDirPath := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err = gfile.X创建目录(tempDirPath)
		t.AssertNil(err)

		zipContent, err := gcompress.Zip压缩文件到字节集(srcPath1 + ", " + srcPath2)
		t.AssertGT(len(zipContent), 0)
		err = gcompress.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer gfile.X删除(tempDirPath)

		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path1", "1.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.X读文本(gfile.X路径生成(tempDirPath, "path2", "2.txt")),
			gfile.X读文本(gfile.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
}
