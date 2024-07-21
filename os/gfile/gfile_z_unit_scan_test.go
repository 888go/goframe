	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package gfile_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
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
