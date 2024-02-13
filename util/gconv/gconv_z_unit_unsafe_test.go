// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 转换类_test

import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_Unsafe(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "I love 小泽玛利亚"
		t.AssertEQ(转换类.X文本到字节集_非安全(s), []byte(s))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		b := []byte("I love 小泽玛利亚")
		t.AssertEQ(转换类.X字节集到文本_非安全(b), string(b))
	})
}
