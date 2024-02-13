// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 模板类_test

import (
	"context"
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gview"
	"github.com/888go/goframe/test/gtest"
)

func Test_Encode_Parse(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New()
		v.SetPath(单元测试类.DataPath("tpl"))
		v.SetAutoEncode(true)
		result, err := v.Parse(context.TODO(), "encode.tpl", g.Map{
			"title": "<b>my title</b>",
		})
		t.AssertNil(err)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}

func Test_Encode_ParseContent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		v := 模板类.New()
		tplContent := 文件类.X读文本(单元测试类.DataPath("tpl", "encode.tpl"))
		v.SetAutoEncode(true)
		result, err := v.ParseContent(context.TODO(), tplContent, g.Map{
			"title": "<b>my title</b>",
		})
		t.AssertNil(err)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}
