// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 资源类_test

import (
	_ "github.com/888go/goframe/os/gres/testdata/data"
	
	"strings"
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gres"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
)

func Test_PackFolderToGoFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath    = 单元测试类.DataPath("files")
			goFilePath = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "testdata.go")
			pkgName    = "testdata"
			err        = 资源类.PackToGoFile(srcPath, goFilePath, pkgName)
		)
		t.AssertNil(err)
		_ = 文件类.X删除(goFilePath)
	})
}

func Test_PackMultiFilesToGoFile(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath    = 单元测试类.DataPath("files")
			goFilePath = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "data.go")
			pkgName    = "data"
			array, err = 文件类.X枚举并含子目录名(srcPath, "*", false)
		)
		t.AssertNil(err)
		err = 资源类.PackToGoFile(strings.Join(array, ","), goFilePath, pkgName)
		t.AssertNil(err)
		defer 文件类.X删除(goFilePath)

		t.AssertNil(文件类.X复制文件(goFilePath, 单元测试类.DataPath("data/data.go")))
	})
}

func Test_Pack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath   = 单元测试类.DataPath("files")
			data, err = 资源类.Pack(srcPath)
		)
		t.AssertNil(err)

		r := 资源类.New()
		err = r.Add(string(data))
		t.AssertNil(err)
		t.Assert(r.Contains("files/"), true)
	})
}

func Test_PackToFile(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath = 单元测试类.DataPath("files")
			dstPath = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err     = 资源类.PackToFile(srcPath, dstPath)
		)
		t.AssertNil(err)

		defer 文件类.X删除(dstPath)

		r := 资源类.New()
		err = r.Load(dstPath)
		t.AssertNil(err)
		t.Assert(r.Contains("files"), true)
	})
}

func Test_PackWithPrefix1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath    = 单元测试类.DataPath("files")
			goFilePath = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "testdata.go")
			pkgName    = "testdata"
			err        = 资源类.PackToGoFile(srcPath, goFilePath, pkgName, "www/gf-site/test")
		)
		t.AssertNil(err)
		_ = 文件类.X删除(goFilePath)
	})
}

func Test_PackWithPrefix2(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath    = 单元测试类.DataPath("files")
			goFilePath = 文件类.X取临时目录(时间类.X取文本时间戳纳秒(), "testdata.go")
			pkgName    = "testdata"
			err        = 资源类.PackToGoFile(srcPath, goFilePath, pkgName, "/var/www/gf-site/test")
		)
		t.AssertNil(err)
		_ = 文件类.X删除(goFilePath)
	})
}

func Test_Unpack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			srcPath    = 单元测试类.DataPath("testdata.txt")
			files, err = 资源类.Unpack(srcPath)
		)
		t.AssertNil(err)
		t.Assert(len(files), 63)
	})
}

func Test_Basic(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(资源类.Get("none"), nil)
		t.Assert(资源类.Contains("none"), false)
		t.Assert(资源类.Contains("dir1"), true)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir1/test1"
		file := 资源类.Get(path)
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

	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir2"
		file := 资源类.Get(path)
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

	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir2/test2"
		file := 资源类.Get(path)
		t.AssertNE(file, nil)
		t.Assert(file.Name(), path)
		t.Assert(file.Content(), "test2 content")
	})
}

func Test_Get(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNE(资源类.Get("dir1/test1"), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		file := 资源类.GetWithIndex("dir1", g.SliceStr别名{"test1"})
		t.AssertNE(file, nil)
		t.Assert(file.Name(), "dir1/test1")
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(资源类.GetContent("dir1"), "")
		t.Assert(资源类.GetContent("dir1/test1"), "test1 content")
	})
}

func Test_ScanDir(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir1"
		files := 资源类.ScanDir(path, "*", false)
		t.AssertNE(files, nil)
		t.Assert(len(files), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir1"
		files := 资源类.ScanDir(path, "*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 3)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir1"
		files := 资源类.ScanDir(path, "*.*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
		t.Assert(files[0].Name(), "dir1/sub/sub-test1.txt")
		t.Assert(files[0].Content(), "sub-test1 content")
	})
}

func Test_ScanDirFile(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir2"
		files := 资源类.ScanDirFile(path, "*", false)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir2"
		files := 资源类.ScanDirFile(path, "*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		path := "dir2"
		files := 资源类.ScanDirFile(path, "*.*", true)
		t.AssertNE(files, nil)
		t.Assert(len(files), 1)
		t.Assert(files[0].Name(), "dir2/sub/sub-test2.txt")
		t.Assert(files[0].Content(), "sub-test2 content")
	})
}

func Test_Export(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res`
			dst = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err = 资源类.Export(src, dst)
		)
		defer 文件类.X删除(dst)
		t.AssertNil(err)
		files, err := 文件类.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 14)

		name := `template-res/index.html`
		t.Assert(文件类.X读文本(文件类.X路径生成(dst, name)), 资源类.GetContent(name))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res`
			dst = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err = 资源类.Export(src, dst, 资源类.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer 文件类.X删除(dst)
		t.AssertNil(err)
		files, err := 文件类.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 13)

		nameInRes := `template-res/index.html`
		nameInSys := `index.html`
		t.Assert(文件类.X读文本(文件类.X路径生成(dst, nameInSys)), 资源类.GetContent(nameInRes))
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res/layout1/container.html`
			dst = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err = 资源类.Export(src, dst, 资源类.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer 文件类.X删除(dst)
		t.AssertNil(err)
		files, err := 文件类.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 2)
	})
}

func Test_IsEmpty(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(资源类.IsEmpty(), false)
	})
}

func TestFile_Name(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res`
		)
		t.Assert(资源类.Get(src).Name(), src)
	})
}

func TestFile_Export(t *testing.T) {
	资源类.Dump()
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res`
			dst = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err = 资源类.Get(src).Export(dst)
		)
		defer 文件类.X删除(dst)
		t.AssertNil(err)
		files, err := 文件类.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 14)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res`
			dst = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err = 资源类.Get(src).Export(dst, 资源类.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer 文件类.X删除(dst)
		t.AssertNil(err)
		files, err := 文件类.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 13)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			src = `template-res/layout1/container.html`
			dst = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
			err = 资源类.Get(src).Export(dst, 资源类.ExportOption{
				RemovePrefix: `template-res`,
			})
		)
		defer 文件类.X删除(dst)
		t.AssertNil(err)
		files, err := 文件类.X枚举并含子目录名(dst, "*", true)
		t.AssertNil(err)
		t.Assert(len(files), 2)
	})
}
