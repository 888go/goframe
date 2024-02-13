// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类_test

import (
	"bytes"
	"testing"
	
	"github.com/888go/goframe/encoding/gcompress"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_ZipPath(t *testing.T) {
	// file
	单元测试类.C(t, func(t *单元测试类.T) {
		srcPath := 单元测试类.DataPath("zip", "path1", "1.txt")
		dstPath := 单元测试类.DataPath("zip", "zip.zip")

		t.Assert(文件类.X是否存在(dstPath), false)
		t.Assert(压缩类.Zip压缩文件(srcPath, dstPath), nil)
		t.Assert(文件类.X是否存在(dstPath), true)
		defer 文件类.X删除(dstPath)

		// 解压缩到临时目录。
		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		t.Assert(文件类.X创建目录(tempDirPath), nil)
		t.Assert(压缩类.Zip解压文件(dstPath, tempDirPath), nil)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "1.txt")),
			文件类.X读文本(srcPath),
		)
	})
	// multiple files
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath1 = 单元测试类.DataPath("zip", "path1", "1.txt")
			srcPath2 = 单元测试类.DataPath("zip", "path2", "2.txt")
			dstPath  = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "zip.zip")
		)
		if p := 文件类.X路径取父目录(dstPath); !文件类.X是否存在(p) {
			t.Assert(文件类.X创建目录(p), nil)
		}

		t.Assert(文件类.X是否存在(dstPath), false)
		err := 压缩类.Zip压缩文件(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(文件类.X是否存在(dstPath), true)
		defer 文件类.X删除(dstPath)

		// 将文件解压缩到另一个临时目录。
		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		t.Assert(文件类.X创建目录(tempDirPath), nil)
		err = 压缩类.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "1.txt")),
			文件类.X读文本(srcPath1),
		)
		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "2.txt")),
			文件类.X读文本(srcPath2),
		)
	})
	// 一个目录和一个文件。
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath1 = 单元测试类.DataPath("zip", "path1")
			srcPath2 = 单元测试类.DataPath("zip", "path2", "2.txt")
			dstPath  = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "zip.zip")
		)
		if p := 文件类.X路径取父目录(dstPath); !文件类.X是否存在(p) {
			t.Assert(文件类.X创建目录(p), nil)
		}

		t.Assert(文件类.X是否存在(dstPath), false)
		err := 压缩类.Zip压缩文件(srcPath1+","+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(文件类.X是否存在(dstPath), true)
		defer 文件类.X删除(dstPath)

		// 将文件解压缩到另一个临时目录。
		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		t.Assert(文件类.X创建目录(tempDirPath), nil)
		err = 压缩类.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path1", "1.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath1, "1.txt")),
		)
		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "2.txt")),
			文件类.X读文本(srcPath2),
		)
	})
	// directory.
	单元测试类.C(t, func(t *单元测试类.T) {
		srcPath := 单元测试类.DataPath("zip")
		dstPath := 单元测试类.DataPath("zip", "zip.zip")

		pwd := 文件类.X取当前工作目录()
		err := 文件类.X设置当前工作目录(srcPath)
		defer 文件类.X设置当前工作目录(pwd)
		t.AssertNil(err)

		t.Assert(文件类.X是否存在(dstPath), false)
		err = 压缩类.Zip压缩文件(srcPath, dstPath)
		t.AssertNil(err)
		t.Assert(文件类.X是否存在(dstPath), true)
		defer 文件类.X删除(dstPath)

		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(tempDirPath)
		t.AssertNil(err)

		err = 压缩类.Zip解压文件(dstPath, tempDirPath)
		t.AssertNil(err)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "zip", "path1", "1.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "zip", "path2", "2.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
	// 使用字符','连接多个目录路径。
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath  = 单元测试类.DataPath("zip")
			srcPath1 = 单元测试类.DataPath("zip", "path1")
			srcPath2 = 单元测试类.DataPath("zip", "path2")
			dstPath  = 单元测试类.DataPath("zip", "zip.zip")
		)
		pwd := 文件类.X取当前工作目录()
		err := 文件类.X设置当前工作目录(srcPath)
		defer 文件类.X设置当前工作目录(pwd)
		t.AssertNil(err)

		t.Assert(文件类.X是否存在(dstPath), false)
		err = 压缩类.Zip压缩文件(srcPath1+", "+srcPath2, dstPath)
		t.AssertNil(err)
		t.Assert(文件类.X是否存在(dstPath), true)
		defer 文件类.X删除(dstPath)

		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(tempDirPath)
		t.AssertNil(err)

		zipContent := 文件类.X读字节集(dstPath)
		t.AssertGT(len(zipContent), 0)
		err = 压缩类.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path1", "1.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path2", "2.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathWriter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath  = 单元测试类.DataPath("zip")
			srcPath1 = 单元测试类.DataPath("zip", "path1")
			srcPath2 = 单元测试类.DataPath("zip", "path2")
		)
		pwd := 文件类.X取当前工作目录()
		err := 文件类.X设置当前工作目录(srcPath)
		defer 文件类.X设置当前工作目录(pwd)
		t.AssertNil(err)

		writer := bytes.NewBuffer(nil)
		t.Assert(writer.Len(), 0)
		err = 压缩类.Zip压缩文件到Writer(srcPath1+", "+srcPath2, writer)
		t.AssertNil(err)
		t.AssertGT(writer.Len(), 0)

		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(tempDirPath)
		t.AssertNil(err)

		zipContent := writer.Bytes()
		t.AssertGT(len(zipContent), 0)
		err = 压缩类.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path1", "1.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path2", "2.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
}

func Test_ZipPathContent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath  = 单元测试类.DataPath("zip")
			srcPath1 = 单元测试类.DataPath("zip", "path1")
			srcPath2 = 单元测试类.DataPath("zip", "path2")
		)
		pwd := 文件类.X取当前工作目录()
		err := 文件类.X设置当前工作目录(srcPath)
		defer 文件类.X设置当前工作目录(pwd)
		t.AssertNil(err)

		tempDirPath := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err = 文件类.X创建目录(tempDirPath)
		t.AssertNil(err)

		zipContent, err := 压缩类.Zip压缩文件到字节集(srcPath1 + ", " + srcPath2)
		t.AssertGT(len(zipContent), 0)
		err = 压缩类.Zip解压字节集(zipContent, tempDirPath)
		t.AssertNil(err)
		defer 文件类.X删除(tempDirPath)

		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path1", "1.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path1", "1.txt")),
		)
		t.Assert(
			文件类.X读文本(文件类.X路径生成(tempDirPath, "path2", "2.txt")),
			文件类.X读文本(文件类.X路径生成(srcPath, "path2", "2.txt")),
		)
	})
}
