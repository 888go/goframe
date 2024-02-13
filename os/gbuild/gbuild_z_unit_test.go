// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 构建信息类_test

import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gbuild"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Info(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(转换类.X取Map(构建信息类.Info()), g.Map{
			"GoFrame": "",
			"Golang":  "",
			"Git":     "",
			"Time":    "",
			"Version": "",
			"Data":    g.Map{},
		})
	})
}

func Test_Get(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(构建信息类.Get(`none`), nil)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(构建信息类.Get(`none`, 1), 1)
	})
}

func Test_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(构建信息类.Data(), map[string]interface{}{})
	})
}
