// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gspath_test

import (
	"testing"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gspath"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestSPath_Api(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		pwd := gfile.Pwd()
		root := pwd
		gfile.Create(gfile.Join(root, "gf_tmp", "gf.txt"))
		defer gfile.Remove(gfile.Join(root, "gf_tmp"))
		fp, isDir := gspath.Search(root, "gf_tmp")
		t.Assert(fp, gfile.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gspath.Search(root, "gf_tmp", "gf.txt")
		t.Assert(fp, gfile.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)

		fp, isDir = gspath.SearchWithCache(root, "gf_tmp")
		t.Assert(fp, gfile.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gspath.SearchWithCache(root, "gf_tmp", "gf.txt")
		t.Assert(fp, gfile.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
	})
}

func TestSPath_Basic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		pwd := gfile.Pwd()
		root := pwd

		gfile.Create(gfile.Join(root, "gf_tmp", "gf.txt"))
		defer gfile.Remove(gfile.Join(root, "gf_tmp"))
		gsp := gspath.New(root, false)
		realPath, err := gsp.Add(gfile.Join(root, "gf_tmp"))
		t.AssertNil(err)
		t.Assert(realPath, gfile.Join(root, "gf_tmp"))
		realPath, err = gsp.Add("gf_tmp1")
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		realPath, err = gsp.Add(gfile.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(err != nil, true)
		t.Assert(realPath, "")

		gsp.Remove("gf_tmp1")

		t.Assert(gsp.Size(), 2)
		t.Assert(len(gsp.Paths()), 2)
		t.Assert(len(gsp.AllPaths()), 0)
		realPath, err = gsp.Set(gfile.Join(root, "gf_tmp1"))
		t.Assert(err != nil, true)
		t.Assert(realPath, "")
		realPath, err = gsp.Set(gfile.Join(root, "gf_tmp", "gf.txt"))
		t.AssertNE(err, nil)
		t.Assert(realPath, "")

		realPath, err = gsp.Set(root)
		t.AssertNil(err)
		t.Assert(realPath, root)

		fp, isDir := gsp.Search("gf_tmp")
		t.Assert(fp, gfile.Join(root, "gf_tmp"))
		t.Assert(isDir, true)
		fp, isDir = gsp.Search("gf_tmp", "gf.txt")
		t.Assert(fp, gfile.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
		fp, isDir = gsp.Search("/", "gf.txt")
		t.Assert(fp, root)
		t.Assert(isDir, true)

		gsp = gspath.New(root, true)
		realPath, err = gsp.Add(gfile.Join(root, "gf_tmp"))
		t.AssertNil(err)
		t.Assert(realPath, gfile.Join(root, "gf_tmp"))

		gfile.Mkdir(gfile.Join(root, "gf_tmp1"))
		gfile.Rename(gfile.Join(root, "gf_tmp1"), gfile.Join(root, "gf_tmp2"))
		gfile.Rename(gfile.Join(root, "gf_tmp2"), gfile.Join(root, "gf_tmp1"))
		defer gfile.Remove(gfile.Join(root, "gf_tmp1"))
		realPath, err = gsp.Add("gf_tmp1")
		t.Assert(err != nil, false)
		t.Assert(realPath, gfile.Join(root, "gf_tmp1"))

		realPath, err = gsp.Add("gf_tmp3")
		t.Assert(err != nil, true)
		t.Assert(realPath, "")

		gsp.Remove(gfile.Join(root, "gf_tmp"))
		gsp.Remove(gfile.Join(root, "gf_tmp1"))
		gsp.Remove(gfile.Join(root, "gf_tmp3"))
		t.Assert(gsp.Size(), 3)
		t.Assert(len(gsp.Paths()), 3)

		gsp.AllPaths()
		gsp.Set(root)
		fp, isDir = gsp.Search("gf_tmp")
		t.Assert(fp, gfile.Join(root, "gf_tmp"))
		t.Assert(isDir, true)

		fp, isDir = gsp.Search("gf_tmp", "gf.txt")
		t.Assert(fp, gfile.Join(root, "gf_tmp", "gf.txt"))
		t.Assert(isDir, false)
	})
}
