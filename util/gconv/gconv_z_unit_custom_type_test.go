// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gconv_test

import (
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

type Duration time.Duration

// UnmarshalText 将文本解析为持续时间。 md5:1e2a68aed8e2fc35
func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

func Test_Struct_CustomTimeDuration_Attribute(t *testing.T) {
	type A struct {
		Name    string
		Timeout Duration
	}
	gtest.C(t, func(t *gtest.T) {
		var a A
		err := gconv.Struct(g.Map{
			"name":    "john",
			"timeout": "1s",
		}, &a)
		t.AssertNil(err)
	})
}
