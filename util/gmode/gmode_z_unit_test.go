// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package gmode_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gmode"
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
