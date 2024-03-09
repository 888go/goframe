// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/container/garray"
	"github.com/888go/goframe/gfile"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_ScanDir(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := 文件类.X枚举并含子目录名(teatPath, "*", false)
		t.AssertNil(err)
		t.AssertIN(teatPath+文件类.Separator+"dir1", files)
		t.AssertIN(teatPath+文件类.Separator+"dir2", files)
		t.AssertNE(teatPath+文件类.Separator+"dir1"+文件类.Separator+"file1", files)
	})
	gtest.C(t, func(t *gtest.T) {
		files, err := 文件类.X枚举并含子目录名(teatPath, "*", true)
		t.AssertNil(err)
		t.AssertIN(teatPath+文件类.Separator+"dir1", files)
		t.AssertIN(teatPath+文件类.Separator+"dir2", files)
		t.AssertIN(teatPath+文件类.Separator+"dir1"+文件类.Separator+"file1", files)
		t.AssertIN(teatPath+文件类.Separator+"dir2"+文件类.Separator+"file2", files)
	})
}

func Test_ScanDirFunc(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := 文件类.X枚举并含子目录名_函数(teatPath, "*", true, func(path string) string {
			if 文件类.X路径取文件名且不含扩展名(path) != "file1" {
				return ""
			}
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(文件类.X路径取文件名且不含扩展名(files[0]), "file1")
	})
}

func Test_ScanDirFile(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		files, err := 文件类.X枚举(teatPath, "*", false)
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		files, err := 文件类.X枚举(teatPath, "*", true)
		t.AssertNil(err)
		t.AssertNI(teatPath+文件类.Separator+"dir1", files)
		t.AssertNI(teatPath+文件类.Separator+"dir2", files)
		t.AssertIN(teatPath+文件类.Separator+"dir1"+文件类.Separator+"file1", files)
		t.AssertIN(teatPath+文件类.Separator+"dir2"+文件类.Separator+"file2", files)
	})
}

func Test_ScanDirFileFunc(t *testing.T) {
	teatPath := gtest.DataPath()
	gtest.C(t, func(t *gtest.T) {
		array := garray.New()
		files, err := 文件类.X枚举_函数(teatPath, "*", false, func(path string) string {
			array.Append(1)
			return path
		})
		t.AssertNil(err)
		t.Assert(len(files), 0)
		t.Assert(array.Len(), 0)
	})
	gtest.C(t, func(t *gtest.T) {
		array := garray.New()
		files, err := 文件类.X枚举_函数(teatPath, "*", true, func(path string) string {
			array.Append(1)
			if 文件类.X路径取文件名(path) == "file1" {
				return path
			}
			return ""
		})
		t.AssertNil(err)
		t.Assert(len(files), 1)
		t.Assert(array.Len(), 3)
	})
}
