// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gview_test

import (
	"context"
	"testing"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/888go/goframe/gview"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Encode_Parse(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		v.SetPath(gtest.DataPath("tpl"))
		v.SetAutoEncode(true)
		result, err := v.Parse(context.TODO(), "encode.tpl", g.Map{
			"title": "<b>my title</b>",
		})
		t.AssertNil(err)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}

func Test_Encode_ParseContent(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := gview.New()
		tplContent := gfile.GetContents(gtest.DataPath("tpl", "encode.tpl"))
		v.SetAutoEncode(true)
		result, err := v.ParseContent(context.TODO(), tplContent, g.Map{
			"title": "<b>my title</b>",
		})
		t.AssertNil(err)
		t.Assert(result, "<div>&lt;b&gt;my title&lt;/b&gt;</div>")
	})
}
