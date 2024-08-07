// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 资源类_test

import (
	_ "github.com/888go/goframe/os/gres/testdata/data"

	"strings"
	"testing"

	"github.com/888go/goframe/frame/g"
	gfile "github.com/888go/goframe/os/gfile"
	gres "github.com/888go/goframe/os/gres"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_PackFolderToGoFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath    = gtest.DataPath("files")
			goFilePath = gfile.X取临时目录(gtime.X取文本时间戳纳秒(), "testdata.go")
			pkgName    = "testdata"
			err        = gres.PackToGoFile(srcPath, goFilePath, pkgName)
		)
		t.AssertNil(err)
		_ = gfile.X删除(goFilePath)
	})
}

func Test_PackMultiFilesToGoFile(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath    = gtest.DataPath("files")
			goFilePath = gfile.X取临时目录(gtime.X取文本时间戳纳秒(), "data.go")
			pkgName    = "data"
			array, err = gfile.X枚举并含子目录名(srcPath, "*", false)
		)
		t.AssertNil(err)
		err = gres.PackToGoFile(strings.Join(array, ","), goFilePath, pkgName)
		t.AssertNil(err)
		defer gfile.X删除(goFilePath)

		t.AssertNil(gfile.X复制文件(goFilePath, gtest.DataPath("data/data.go")))
	})
}

func Test_Pack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath   = gtest.DataPath("files")
			data, err = gres.Pack(srcPath)
		)
		t.AssertNil(err)

		r := gres.New()
		err = r.Add(string(data))
		t.AssertNil(err)
		t.Assert(r.Contains("files/"), true)
	})
}

func Test_PackToFile(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath = gtest.DataPath("files")
			dstPath = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err     = gres.PackToFile(srcPath, dstPath)
		)
		t.AssertNil(err)

		defer gfile.X删除(dstPath)

		r := gres.New()
		err = r.Load(dstPath)
		t.AssertNil(err)
		t.Assert(r.Contains("files"), true)
	})
}

func Test_PackWithPrefix1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath    = gtest.DataPath("files")
			goFilePath = gfile.X取临时目录(gtime.X取文本时间戳纳秒(), "testdata.go")
			pkgName    = "testdata"
			err        = gres.PackToGoFile(srcPath, goFilePath, pkgName, "www/gf-site/test")
		)
		t.AssertNil(err)
		_ = gfile.X删除(goFilePath)
	})
}

func Test_PackWithPrefix2(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath    = gtest.DataPath("files")
			goFilePath = gfile.X取临时目录(gtime.X取文本时间戳纳秒(), "testdata.go")
			pkgName    = "testdata"
			err        = gres.PackToGoFile(srcPath, goFilePath, pkgName, "/var/www/gf-site/test")
		)
		t.AssertNil(err)
		_ = gfile.X删除(goFilePath)
	})
}

func Test_Unpack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			srcPath    = gtest.DataPath("testdata.txt")
			files, err = gres.Unpack(srcPath)
		)
		t.AssertNil(err)
		t.Assert(len(files), 63)
	})
}

func Test_Basic(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gres.Get("none"), nil)
		t.Assert(gres.Contains("none"), false)
		t.Assert(gres.Contains("dir1"), true)
	})

	gtest.C(t, func(t *gtest.T) {
		path := "dir1/test1"
		file := gres.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)

		info := file.FileInfo()
		t.AssertNE(info, nil)
		t.Assert(info.IsDir(), false)
		t.Assert(info.Name(), "test1")

		rc, err := file.Open()
		t.AssertNil(err)
		defer rc.Close()

		b := make([]byte, 5)
		n, err := rc.Read(b)
		t.Assert(n, 5)
		t.AssertNil(err)
		t.Assert(string(b), "test1")

		t.Assert(file.Content(), "test1 content")
	})

	gtest.C(t, func(t *gtest.T) {
		path := "dir2"
		file := gres.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)

		info := file.FileInfo()
		t.AssertNE(info, nil)
		t.Assert(info.IsDir(), true)
		t.Assert(info.Name(), "dir2")

		rc, err := file.Open()
		t.AssertNil(err)
		defer rc.Close()

		t.Assert(file.Content(), nil)
	})

	gtest.C(t, func(t *gtest.T) {
		path := "dir2/test2"
		file := gres.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)
		t.Assert(file.Content(), "test2 content")
	})
}

func Test_Get(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		t.AssertNE(gres.Get("dir1/test1"), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		file := gres.GetWithIndex("dir1", g.SliceStr别名{"test1"})
		t.AssertNE(file, nil)
		t.Assert(file.Name(), "dir1/test1")
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gres.GetContent("dir1"), "")
		t.Assert(gres.GetContent("dir1/test1"), "test1 content")
	})
}

func Test_ScanDir(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		path := "dir1"
		files := gres.ScanDir(path, "*", false)
		t.AssertNE(files, nil)
		t.Assert(len(files), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		path := "dir1"
		files := gres.ScanDir(path, "*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 3)
	})

	gtest.C(t, func(t *gtest.T) {
		path := "dir1"
		files := gres.ScanDir(path, "*.*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
		t.Assert(files[0].Name(), "dir1/sub/sub-test1.txt")
		t.Assert(files[0].Content(), "sub-test1 content")
	})
}

func Test_ScanDirFile(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		path := "dir2"
		files := gres.ScanDirFile(path, "*", false)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		path := "dir2"
		files := gres.ScanDirFile(path, "*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 2)
	})

	gtest.C(t, func(t *gtest.T) {
		path := "dir2"
		files := gres.ScanDirFile(path, "*.*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
		t.Assert(files[0].Name(), "dir2/sub/sub-test2.txt")
		t.Assert(files[0].Content(), "sub-test2 content")
	})
}

func Test_Export(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res`
			dst = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err = gres.Export(src, dst)
		)
		defer gfile.X删除(dst)
		t.AssertNil(err)
		files, err := gfile.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 14)

		name := `template-res/index.html`
		t.Assert(gfile.X读文本(gfile.X路径生成(dst, name)), gres.GetContent(name))
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res`
			dst = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err = gres.Export(src, dst, gres.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer gfile.X删除(dst)
		t.AssertNil(err)
		files, err := gfile.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 13)

		nameInRes := `template-res/index.html`
		nameInSys := `index.html`
		t.Assert(gfile.X读文本(gfile.X路径生成(dst, nameInSys)), gres.GetContent(nameInRes))
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res/layout1/container.html`
			dst = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err = gres.Export(src, dst, gres.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer gfile.X删除(dst)
		t.AssertNil(err)
		files, err := gfile.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 2)
	})
}

func Test_IsEmpty(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gres.IsEmpty(), false)
	})
}

func TestFile_Name(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res`
		)
		t.Assert(gres.Get(src).Name(), src)
	})
}

func TestFile_Export(t *testing.T) {
	gres.Dump()
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res`
			dst = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err = gres.Get(src).Export(dst)
		)
		defer gfile.X删除(dst)
		t.AssertNil(err)
		files, err := gfile.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 14)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res`
			dst = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err = gres.Get(src).Export(dst, gres.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer gfile.X删除(dst)
		t.AssertNil(err)
		files, err := gfile.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 13)
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			src = `template-res/layout1/container.html`
			dst = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			err = gres.Get(src).Export(dst, gres.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer gfile.X删除(dst)
		t.AssertNil(err)
		files, err := gfile.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 2)
	})
}
