// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gbuild_test

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
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
