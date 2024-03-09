// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类_test

import (
	"bytes"
	"testing"
	
	"github.com/888go/goframe/gcompress"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_ZipPath(t *testing.T) {
	// file
	gtest.C(t, func(t *gtest.T) {
		srcPath := gtest.DataPath("zip", "path1", "1.txt")
		dstPath := gtest.DataPath("zip", "zip.zip")

		t.Assert(gfile.Exists(dstPath), false)
		t.Assert(压缩类.Zip压缩文件(srcPath, dstPath), nil)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		// 解压缩到临时目录。
		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		t.Assert(gfile.Mkdir(tempDirPath), nil)
		t.Assert(压缩类.Zip解压文件(dstPath, tempDirPath), nil)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "1.txt")),
			gfile.GetContents(srcPath),
		)
	})
	// multiple files
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath1 = gtest.DataPath("zip", "path1", "1.txt")
			srcPath2 = gtest.DataPath("zip", "path2", "2.txt")
			dstPath  = gfile.Temp(gtime.TimestampNanoStr(), "zip.zip")
		)
		if p := gfile.Dir(dstPath); !gfile.Exists(p) {
			t.Assert(gfile.Mkdir(p), nil)
		}

		t.Assert(gfile.Exists(dstPath), false)
		err := 压缩类.Zip压缩文件(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		// 将文件解压缩到另一个临时目录。
		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		t.Assert(gfile.Mkdir(tempDirPath), nil)
		err = 压缩类.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "1.txt")),
			gfile.GetContents(srcPath1),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "2.txt")),
			gfile.GetContents(srcPath2),
		)
	})
	// 一个目录和一个文件。
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath1 = gtest.DataPath("zip", "path1")
			srcPath2 = gtest.DataPath("zip", "path2", "2.txt")
			dstPath  = gfile.Temp(gtime.TimestampNanoStr(), "zip.zip")
		)
		if p := gfile.Dir(dstPath); !gfile.Exists(p) {
			t.Assert(gfile.Mkdir(p), nil)
		}

		t.Assert(gfile.Exists(dstPath), false)
		err := 压缩类.Zip压缩文件(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		// 将文件解压缩到另一个临时目录。
		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		t.Assert(gfile.Mkdir(tempDirPath), nil)
		err = 压缩类.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath1, "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "2.txt")),
			gfile.GetContents(srcPath2),
		)
	})
	// directory.
	gtest.C(t, func(t *gtest.T) {
		srcPath := gtest.DataPath("zip")
		dstPath := gtest.DataPath("zip", "zip.zip")

		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.AssertNil(err)

		t.Assert(gfile.Exists(dstPath), false)
		err = 压缩类.Zip压缩文件(srcPath, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.AssertNil(err)

		err = 压缩类.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "zip", "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "zip", "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
		)
	})
	// 使用字符','连接多个目录路径。
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath  = gtest.DataPath("zip")
			srcPath1 = gtest.DataPath("zip", "path1")
			srcPath2 = gtest.DataPath("zip", "path2")
			dstPath  = gtest.DataPath("zip", "zip.zip")
		)
		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.AssertNil(err)

		t.Assert(gfile.Exists(dstPath), false)
		err = 压缩类.Zip压缩文件(srcPath1+", "+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(gfile.Exists(dstPath), true)
		defer gfile.Remove(dstPath)

		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.AssertNil(err)

		zipContent := gfile.GetBytes(dstPath)
		t.AssertGT(len(zipContent), 0)
		err = 压缩类.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
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
		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.AssertNil(err)

		writer := bytes.NewBuffer(nil)
		t.Assert(writer.Len(), 0)
		err = 压缩类.Zip压缩文件到Writer(srcPath1+", "+srcPath2, writer)
		t.AssertNil(err)
		t.AssertGT(writer.Len(), 0)

		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.AssertNil(err)

		zipContent := writer.Bytes()
		t.AssertGT(len(zipContent), 0)
		err = 压缩类.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
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
		pwd := gfile.Pwd()
		err := gfile.Chdir(srcPath)
		defer gfile.Chdir(pwd)
		t.AssertNil(err)

		tempDirPath := gfile.Temp(gtime.TimestampNanoStr())
		err = gfile.Mkdir(tempDirPath)
		t.AssertNil(err)

		zipContent, err := 压缩类.Zip压缩文件到字节集(srcPath1 + ", " + srcPath2)
		t.AssertGT(len(zipContent), 0)
		err = 压缩类.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer gfile.Remove(tempDirPath)

		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path1", "1.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			gfile.GetContents(gfile.Join(tempDirPath, "path2", "2.txt")),
			gfile.GetContents(gfile.Join(srcPath, "path2", "2.txt")),
		)
	})
}
