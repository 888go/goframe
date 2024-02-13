// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 环境类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gmode"
)

func Test_AutoCheckSourceCodes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(环境类.IsDevelop(), true)
	})
}

func Test_Set(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		oldMode := 环境类.Mode()
		defer 环境类.X设置值(oldMode)
		环境类.SetDevelop()
		t.Assert(环境类.IsDevelop(), true)
		t.Assert(环境类.IsTesting(), false)
		t.Assert(环境类.IsStaging(), false)
		t.Assert(环境类.IsProduct(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		oldMode := 环境类.Mode()
		defer 环境类.X设置值(oldMode)
		环境类.SetTesting()
		t.Assert(环境类.IsDevelop(), false)
		t.Assert(环境类.IsTesting(), true)
		t.Assert(环境类.IsStaging(), false)
		t.Assert(环境类.IsProduct(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		oldMode := 环境类.Mode()
		defer 环境类.X设置值(oldMode)
		环境类.SetStaging()
		t.Assert(环境类.IsDevelop(), false)
		t.Assert(环境类.IsTesting(), false)
		t.Assert(环境类.IsStaging(), true)
		t.Assert(环境类.IsProduct(), false)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		oldMode := 环境类.Mode()
		defer 环境类.X设置值(oldMode)
		环境类.SetProduct()
		t.Assert(环境类.IsDevelop(), false)
		t.Assert(环境类.IsTesting(), false)
		t.Assert(环境类.IsStaging(), false)
		t.Assert(环境类.IsProduct(), true)
	})
}
