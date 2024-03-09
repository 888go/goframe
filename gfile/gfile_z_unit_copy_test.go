// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"os"
	"testing"
	
	"github.com/888go/goframe/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_Copy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(文件类.X复制(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(文件类.X是否为文件(testpath()+topath), true)
		t.AssertNE(文件类.X复制(paths, ""), nil)
		t.AssertNE(文件类.X复制("", topath), nil)
	})
}

func Test_Copy_File_To_Dir(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			src = gtest.DataPath("dir1", "file1")
			dst = 文件类.X取临时目录(guid.S(), "dir2")
		)
		err := 文件类.X创建目录(dst)
		t.AssertNil(err)
		defer 文件类.X删除(dst)

		err = 文件类.X复制(src, dst)
		t.AssertNil(err)

		expectPath := 文件类.X路径生成(dst, "file1")
		t.Assert(文件类.X读文本(expectPath), 文件类.X读文本(src))
	})
}

func Test_Copy_Dir_To_File(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			src = gtest.DataPath("dir1")
			dst = 文件类.X取临时目录(guid.S(), "file2")
		)
		f, err := 文件类.X创建文件与目录(dst)
		t.AssertNil(err)
		defer f.Close()
		defer 文件类.X删除(dst)

		err = 文件类.X复制(src, dst)
		t.AssertNE(err, nil)
	})
}

func Test_CopyFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(文件类.X复制文件(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(文件类.X是否为文件(testpath()+topath), true)
		t.AssertNE(文件类.X复制文件(paths, ""), nil)
		t.AssertNE(文件类.X复制文件("", topath), nil)
	})
	// 内容替换
	gtest.C(t, func(t *gtest.T) {
		src := 文件类.X取临时目录(gtime.TimestampNanoStr())
		dst := 文件类.X取临时目录(gtime.TimestampNanoStr())
		srcContent := "1"
		dstContent := "1"
		t.Assert(文件类.X写入文本(src, srcContent), nil)
		t.Assert(文件类.X写入文本(dst, dstContent), nil)
		t.Assert(文件类.X读文本(src), srcContent)
		t.Assert(文件类.X读文本(dst), dstContent)

		t.Assert(文件类.X复制文件(src, dst), nil)
		t.Assert(文件类.X读文本(src), srcContent)
		t.Assert(文件类.X读文本(dst), srcContent)
	})
	// Set mode
	gtest.C(t, func(t *gtest.T) {
		var (
			src     = "/testfile_copyfile1.txt"
			dst     = "/testfile_copyfile2.txt"
			dstMode = os.FileMode(0600)
		)
		t.AssertNil(createTestFile(src, ""))
		defer delTestFiles(src)

		t.Assert(文件类.X复制文件(testpath()+src, testpath()+dst, 文件类.CopyOption{Mode: dstMode}), nil)
		defer delTestFiles(dst)

		dstStat, err := 文件类.X取详情(testpath() + dst)
		t.AssertNil(err)
		t.Assert(dstStat.Mode().Perm(), dstMode)
	})
	// 保留源文件的模式
	gtest.C(t, func(t *gtest.T) {
		var (
			src = "/testfile_copyfile1.txt"
			dst = "/testfile_copyfile2.txt"
		)
		t.AssertNil(createTestFile(src, ""))
		defer delTestFiles(src)

		t.Assert(文件类.X复制文件(testpath()+src, testpath()+dst, 文件类.CopyOption{PreserveMode: true}), nil)
		defer delTestFiles(dst)

		srcStat, err := 文件类.X取详情(testpath() + src)
		t.AssertNil(err)
		dstStat, err := 文件类.X取详情(testpath() + dst)
		t.AssertNil(err)
		t.Assert(srcStat.Mode().Perm(), dstStat.Mode().Perm())
	})
}

func Test_CopyDir(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			dirPath1 = "/test-copy-dir1"
			dirPath2 = "/test-copy-dir2"
		)
		haveList := []string{
			"t1.txt",
			"t2.txt",
		}
		createDir(dirPath1)
		for _, v := range haveList {
			t.Assert(createTestFile(dirPath1+"/"+v, ""), nil)
		}
		defer delTestFiles(dirPath1)

		var (
			yfolder  = testpath() + dirPath1
			tofolder = testpath() + dirPath2
		)

		if 文件类.X是否存在目录(tofolder) {
			t.Assert(文件类.X删除(tofolder), nil)
			t.Assert(文件类.X删除(""), nil)
		}

		t.Assert(文件类.X复制目录(yfolder, tofolder), nil)
		defer delTestFiles(tofolder)

		t.Assert(文件类.X是否存在目录(yfolder), true)

		for _, v := range haveList {
			t.Assert(文件类.X是否为文件(yfolder+"/"+v), true)
		}

		t.Assert(文件类.X是否存在目录(tofolder), true)

		for _, v := range haveList {
			t.Assert(文件类.X是否为文件(tofolder+"/"+v), true)
		}

		t.Assert(文件类.X删除(tofolder), nil)
		t.Assert(文件类.X删除(""), nil)
	})
	// 内容替换
	gtest.C(t, func(t *gtest.T) {
		src := 文件类.X取临时目录(gtime.TimestampNanoStr(), gtime.TimestampNanoStr())
		dst := 文件类.X取临时目录(gtime.TimestampNanoStr(), gtime.TimestampNanoStr())
		defer func() {
			文件类.X删除(src)
			文件类.X删除(dst)
		}()
		srcContent := "1"
		dstContent := "1"
		t.Assert(文件类.X写入文本(src, srcContent), nil)
		t.Assert(文件类.X写入文本(dst, dstContent), nil)
		t.Assert(文件类.X读文本(src), srcContent)
		t.Assert(文件类.X读文本(dst), dstContent)

		err := 文件类.X复制目录(文件类.X路径取父目录(src), 文件类.X路径取父目录(dst))
		t.AssertNil(err)
		t.Assert(文件类.X读文本(src), srcContent)
		t.Assert(文件类.X读文本(dst), srcContent)

		t.AssertNE(文件类.X复制目录(文件类.X路径取父目录(src), ""), nil)
		t.AssertNE(文件类.X复制目录("", 文件类.X路径取父目录(dst)), nil)
	})
}
