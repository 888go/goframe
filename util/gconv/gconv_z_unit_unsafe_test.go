// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gconv_test
import (
	"testing"
	
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
	)

func Test_Unsafe(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "I love 小泽玛利亚"
		t.AssertEQ(gconv.UnsafeStrToBytes(s), []byte(s))
	})

	gtest.C(t, func(t *gtest.T) {
		b := []byte("I love 小泽玛利亚")
		t.AssertEQ(gconv.UnsafeBytesToStr(b), string(b))
	})
}
