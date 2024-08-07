// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类_test

import (
	"testing"

	garray "github.com/888go/goframe/container/garray"
	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_ScanDir(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.X枚举并含子目录名(teatPath, "*", false)
		t.AssertNil(err)
		t.AssertIN(teatPath+gfile.Separator+"dir1", files)
		t.AssertIN(teatPath+gfile.Separator+"dir2", files)
		t.AssertNE(teatPath+gfile.Separator+"dir1"+gfile.Separator+"file1", files)
	})
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.X枚举并含子目录名(teatPath, "*", true)
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
		files, err := gfile.X枚举并含子目录名_函数(teatPath, "*", true, func(path string) string {
			if gfile.X路径取文件名且不含扩展名(path) != "file1" {
				return ""
			}
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(gfile.X路径取文件名且不含扩展名(files[0]), "file1")
	})
}

func Test_ScanDirFile(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.X枚举(teatPath, "*", false)
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		files, err := gfile.X枚举(teatPath, "*", true)
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
		array := garray.X创建()
		files, err := gfile.X枚举_函数(teatPath, "*", false, func(path string) string {
			array.Append别名(1)
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 0)
		t.Assert(array.X取长度(), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		array := garray.X创建()
		files, err := gfile.X枚举_函数(teatPath, "*", true, func(path string) string {
			array.Append别名(1)
			if gfile.X路径取文件名(path) == "file1" {
				return path
			}
			return ""
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(array.X取长度(), 3)
	})
}
