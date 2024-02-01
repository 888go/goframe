// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gbuild_test
import (
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gbuild"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
	)

func Test_Info(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gconv.Map(gbuild.Info()), g.Map{
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
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gbuild.Get(`none`), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gbuild.Get(`none`, 1), 1)
	})
}

func Test_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gbuild.Data(), map[string]interface{}{})
	})
}
