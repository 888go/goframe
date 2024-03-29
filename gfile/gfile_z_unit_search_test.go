// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"path/filepath"
	"testing"
	
	"github.com/888go/goframe/gfile"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Search(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			paths1  string = "/testfiless"
			paths2  string = "./testfile/dirfiles_no"
			tpath   string
			tpath2  string
			tempstr string
			ypaths1 string
			err     error
		)

		createDir(paths1)
		defer delTestFiles(paths1)
		ypaths1 = paths1

		tpath, err = 文件类.X查找(testpath() + paths1)
		t.AssertNil(err)

		tpath = filepath.ToSlash(tpath)

		// 自定义优先路径
		tpath2, err = 文件类.X查找(testpath() + paths1)
		t.AssertNil(err)
		tpath2 = filepath.ToSlash(tpath2)

		tempstr = testpath()
		paths1 = tempstr + paths1
		paths1 = filepath.ToSlash(paths1)

		t.Assert(tpath, paths1)

		t.Assert(tpath2, tpath)

		// 测试给定目录
		tpath2, err = 文件类.X查找(paths1, "testfiless")
		tpath2 = filepath.ToSlash(tpath2)
		tempss := filepath.ToSlash(paths1)
		t.Assert(tpath2, tempss)

		// 测试当前目录
		tempstr, _ = filepath.Abs("./")
		tempstr = testpath()
		paths1 = tempstr + ypaths1
		paths1 = filepath.ToSlash(paths1)

		t.Assert(tpath2, paths1)

		// 测试目录不存在时
		_, err = 文件类.X查找(paths2)
		t.AssertNE(err, nil)

	})
}
