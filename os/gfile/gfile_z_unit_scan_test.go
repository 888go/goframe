// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"testing"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
)

func Test_ScanDir(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.ScanDir(teatPath, "*", false)
		t.AssertNil(err)
		t.AssertIN(teatPath+gfile.Separator+"dir1", files)
		t.AssertIN(teatPath+gfile.Separator+"dir2", files)
		t.AssertNE(teatPath+gfile.Separator+"dir1"+gfile.Separator+"file1", files)
	})
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.ScanDir(teatPath, "*", true)
		t.AssertNil(err)
		t.AssertIN(teatPath+gfile.Separator+"dir1", files)
		t.AssertIN(teatPath+gfile.Separator+"dir2", files)
		t.AssertIN(teatPath+gfile.Separator+"dir1"+gfile.Separator+"file1", files)
		t.AssertIN(teatPath+gfile.Separator+"dir2"+gfile.Separator+"file2", files)
	})
}

func Test_ScanDirFunc(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.ScanDirFunc(teatPath, "*", true, func(path string) string {
			if gfile.Name(path) != "file1" {
				return ""
			}
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(gfile.Name(files[0]), "file1")
	})
}

func Test_ScanDirFile(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.ScanDirFile(teatPath, "*", false)
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.ScanDirFile(teatPath, "*", true)
		t.AssertNil(err)
		t.AssertNI(teatPath+gfile.Separator+"dir1", files)
		t.AssertNI(teatPath+gfile.Separator+"dir2", files)
		t.AssertIN(teatPath+gfile.Separator+"dir1"+gfile.Separator+"file1", files)
		t.AssertIN(teatPath+gfile.Separator+"dir2"+gfile.Separator+"file2", files)
	})
}

func Test_ScanDirFileFunc(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		array := garray.New()
		files, err := gfile.ScanDirFileFunc(teatPath, "*", false, func(path string) string {
			array.Append(1)
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 0)
		t.Assert(array.Len(), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		array := garray.New()
		files, err := gfile.ScanDirFileFunc(teatPath, "*", true, func(path string) string {
			array.Append(1)
			if gfile.Basename(path) == "file1" {
				return path
			}
			return ""
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(array.Len(), 3)
	})
}
