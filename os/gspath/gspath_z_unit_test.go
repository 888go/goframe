// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件搜索类_test

import (
	"testing"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gspath"
	"github.com/888go/goframe/test/gtest"
)

func TestSPath_Api(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		pwd := 文件类.X取当前工作目录()
		root := pwd
		文件类.X创建文件与目录(文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		defer 文件类.X删除(文件类.X路径生成(root, "gf_tmp"))
		fp, isDir := 文件搜索类.Search(root, "gf_tmp")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = 文件搜索类.Search(root, "gf_tmp", "gf.txt")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)

		fp, isDir = 文件搜索类.SearchWithCache(root, "gf_tmp")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = 文件搜索类.SearchWithCache(root, "gf_tmp", "gf.txt")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
	})
}

func TestSPath_Basic(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		pwd := 文件类.X取当前工作目录()
		root := pwd

		文件类.X创建文件与目录(文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		defer 文件类.X删除(文件类.X路径生成(root, "gf_tmp"))
		gsp := 文件搜索类.New(root, false)
		realPath, err := gsp.Add(文件类.X路径生成(root, "gf_tmp"))
		t.AssertNil(err)
		t.Assert(realPath, 文件类.X路径生成(root, "gf_tmp"))
		realPath, err = gsp.Add("gf_tmp1")
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		realPath, err = gsp.Add(文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		t.Assert(err != nil, true)
		t.Assert(realPath, "")

		gsp.Remove("gf_tmp1")

		t.Assert(gsp.Size(), 2)
		t.Assert(len(gsp.Paths()), 2)
		t.Assert(len(gsp.AllPaths()), 0)
		realPath, err = gsp.X设置值(文件类.X路径生成(root, "gf_tmp1"))
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		realPath, err = gsp.X设置值(文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		t.AssertNE(err, nil)
		t.Assert(realPath, "")

		realPath, err = gsp.X设置值(root)
		t.AssertNil(err)
		t.Assert(realPath, root)

		fp, isDir := gsp.Search("gf_tmp")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gsp.Search("gf_tmp", "gf.txt")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
		fp, isDir = gsp.Search("/", "gf.txt")
		t.Assert(fp, root)
		t.Assert(isDir, true)

		gsp = 文件搜索类.New(root, true)
		realPath, err = gsp.Add(文件类.X路径生成(root, "gf_tmp"))
		t.AssertNil(err)
		t.Assert(realPath, 文件类.X路径生成(root, "gf_tmp"))

		文件类.X创建目录(文件类.X路径生成(root, "gf_tmp1"))
		文件类.Rename别名(文件类.X路径生成(root, "gf_tmp1"), 文件类.X路径生成(root, "gf_tmp2"))
		文件类.Rename别名(文件类.X路径生成(root, "gf_tmp2"), 文件类.X路径生成(root, "gf_tmp1"))
		defer 文件类.X删除(文件类.X路径生成(root, "gf_tmp1"))
		realPath, err = gsp.Add("gf_tmp1")
		t.Assert(err != nil, false)
		t.Assert(realPath, 文件类.X路径生成(root, "gf_tmp1"))

		realPath, err = gsp.Add("gf_tmp3")
		t.Assert(err != nil, true)
		t.Assert(realPath, "")

		gsp.Remove(文件类.X路径生成(root, "gf_tmp"))
		gsp.Remove(文件类.X路径生成(root, "gf_tmp1"))
		gsp.Remove(文件类.X路径生成(root, "gf_tmp3"))
		t.Assert(gsp.Size(), 3)
		t.Assert(len(gsp.Paths()), 3)

		gsp.AllPaths()
		gsp.X设置值(root)
		fp, isDir = gsp.Search("gf_tmp")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp"))
		t.Assert(isDir, true)

		fp, isDir = gsp.Search("gf_tmp", "gf.txt")
		t.Assert(fp, 文件类.X路径生成(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
	})
}
