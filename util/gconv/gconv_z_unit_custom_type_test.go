// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

type Duration time.Duration

// UnmarshalText 将文本反序列化为持续时间。
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
	单元测试类.C(t, func(t *单元测试类.T) {
		var a A
		err := 转换类.Struct(g.Map{
			"name":    "john",
			"timeout": "1s",
		}, &a)
		t.AssertNil(err)
	})
}
