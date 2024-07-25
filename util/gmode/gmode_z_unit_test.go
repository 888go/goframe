// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package gmode_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gmode"
)

func Test_AutoCheckSourceCodes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gmode.IsDevelop(), true)
	})
}

func Test_Set(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		oldMode := gmode.Mode()
		defer gmode.Set(oldMode)
		gmode.SetDevelop()
		t.Assert(gmode.IsDevelop(), true)
		t.Assert(gmode.IsTesting(), false)
		t.Assert(gmode.IsStaging(), false)
		t.Assert(gmode.IsProduct(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		oldMode := gmode.Mode()
		defer gmode.Set(oldMode)
		gmode.SetTesting()
		t.Assert(gmode.IsDevelop(), false)
		t.Assert(gmode.IsTesting(), true)
		t.Assert(gmode.IsStaging(), false)
		t.Assert(gmode.IsProduct(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		oldMode := gmode.Mode()
		defer gmode.Set(oldMode)
		gmode.SetStaging()
		t.Assert(gmode.IsDevelop(), false)
		t.Assert(gmode.IsTesting(), false)
		t.Assert(gmode.IsStaging(), true)
		t.Assert(gmode.IsProduct(), false)
	})
	gtest.C(t, func(t *gtest.T) {
		oldMode := gmode.Mode()
		defer gmode.Set(oldMode)
		gmode.SetProduct()
		t.Assert(gmode.IsDevelop(), false)
		t.Assert(gmode.IsTesting(), false)
		t.Assert(gmode.IsStaging(), false)
		t.Assert(gmode.IsProduct(), true)
	})
}
