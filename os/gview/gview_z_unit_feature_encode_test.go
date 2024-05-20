// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gview_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
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
