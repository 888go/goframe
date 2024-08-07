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

	gfile "github.com/888go/goframe/os/gfile"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	guid "github.com/888go/goframe/util/guid"
)

func Test_Copy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			paths  = "/testfile_copyfile1.txt"
			topath = "/testfile_copyfile2.txt"
		)

		createTestFile(paths, "")
		defer delTestFiles(paths)

		t.Assert(gfile.X复制(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(gfile.X是否为文件(testpath()+topath), true)
		t.AssertNE(gfile.X复制(paths, ""), nil)
		t.AssertNE(gfile.X复制("", topath), nil)
	})
}

func Test_Copy_File_To_Dir(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			src = gtest.DataPath("dir1", "file1")
			dst = gfile.X取临时目录(guid.X生成(), "dir2")
		)
		err := gfile.X创建目录(dst)
		t.AssertNil(err)
		defer gfile.X删除(dst)

		err = gfile.X复制(src, dst)
		t.AssertNil(err)

		expectPath := gfile.X路径生成(dst, "file1")
		t.Assert(gfile.X读文本(expectPath), gfile.X读文本(src))
	})
}

func Test_Copy_Dir_To_File(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			src = gtest.DataPath("dir1")
			dst = gfile.X取临时目录(guid.X生成(), "file2")
		)
		f, err := gfile.X创建文件与目录(dst)
		t.AssertNil(err)
		defer f.Close()
		defer gfile.X删除(dst)

		err = gfile.X复制(src, dst)
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

		t.Assert(gfile.X复制文件(testpath()+paths, testpath()+topath), nil)
		defer delTestFiles(topath)

		t.Assert(gfile.X是否为文件(testpath()+topath), true)
		t.AssertNE(gfile.X复制文件(paths, ""), nil)
		t.AssertNE(gfile.X复制文件("", topath), nil)
	})
	// Content replacement.
	gtest.C(t, func(t *gtest.T) {
		src := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		dst := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		srcContent := "1"
		dstContent := "1"
		t.Assert(gfile.X写入文本(src, srcContent), nil)
		t.Assert(gfile.X写入文本(dst, dstContent), nil)
		t.Assert(gfile.X读文本(src), srcContent)
		t.Assert(gfile.X读文本(dst), dstContent)

		t.Assert(gfile.X复制文件(src, dst), nil)
		t.Assert(gfile.X读文本(src), srcContent)
		t.Assert(gfile.X读文本(dst), srcContent)
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

		t.Assert(gfile.X复制文件(testpath()+src, testpath()+dst, gfile.CopyOption{Mode: dstMode}), nil)
		defer delTestFiles(dst)

		dstStat, err := gfile.X取详情(testpath() + dst)
		t.AssertNil(err)
		t.Assert(dstStat.Mode().Perm(), dstMode)
	})
		// 保留src文件的模式. md5:a4c68c56c9abe9c4
	gtest.C(t, func(t *gtest.T) {
		var (
			src = "/testfile_copyfile1.txt"
			dst = "/testfile_copyfile2.txt"
		)
		t.AssertNil(createTestFile(src, ""))
		defer delTestFiles(src)

		t.Assert(gfile.X复制文件(testpath()+src, testpath()+dst, gfile.CopyOption{PreserveMode: true}), nil)
		defer delTestFiles(dst)

		srcStat, err := gfile.X取详情(testpath() + src)
		t.AssertNil(err)
		dstStat, err := gfile.X取详情(testpath() + dst)
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

		if gfile.X是否存在目录(tofolder) {
			t.Assert(gfile.X删除(tofolder), nil)
			t.Assert(gfile.X删除(""), nil)
		}

		t.Assert(gfile.X复制目录(yfolder, tofolder), nil)
		defer delTestFiles(tofolder)

		t.Assert(gfile.X是否存在目录(yfolder), true)

		for _, v := range haveList {
			t.Assert(gfile.X是否为文件(yfolder+"/"+v), true)
		}

		t.Assert(gfile.X是否存在目录(tofolder), true)

		for _, v := range haveList {
			t.Assert(gfile.X是否为文件(tofolder+"/"+v), true)
		}

		t.Assert(gfile.X删除(tofolder), nil)
		t.Assert(gfile.X删除(""), nil)
	})
	// Content replacement.
	gtest.C(t, func(t *gtest.T) {
		src := gfile.X取临时目录(gtime.X取文本时间戳纳秒(), gtime.X取文本时间戳纳秒())
		dst := gfile.X取临时目录(gtime.X取文本时间戳纳秒(), gtime.X取文本时间戳纳秒())
		defer func() {
			gfile.X删除(src)
			gfile.X删除(dst)
		}()
		srcContent := "1"
		dstContent := "1"
		t.Assert(gfile.X写入文本(src, srcContent), nil)
		t.Assert(gfile.X写入文本(dst, dstContent), nil)
		t.Assert(gfile.X读文本(src), srcContent)
		t.Assert(gfile.X读文本(dst), dstContent)

		err := gfile.X复制目录(gfile.X路径取父目录(src), gfile.X路径取父目录(dst))
		t.AssertNil(err)
		t.Assert(gfile.X读文本(src), srcContent)
		t.Assert(gfile.X读文本(dst), srcContent)

		t.AssertNE(gfile.X复制目录(gfile.X路径取父目录(src), ""), nil)
		t.AssertNE(gfile.X复制目录("", gfile.X路径取父目录(dst)), nil)
	})
}
