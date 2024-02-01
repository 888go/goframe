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
type boolStruct struct{}

func Test_Bool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var any interface{} = nil
		t.AssertEQ(gconv.Bool(any), false)
		t.AssertEQ(gconv.Bool(false), false)
		t.AssertEQ(gconv.Bool(nil), false)
		t.AssertEQ(gconv.Bool(0), false)
		t.AssertEQ(gconv.Bool("0"), false)
		t.AssertEQ(gconv.Bool(""), false)
		t.AssertEQ(gconv.Bool("false"), false)
		t.AssertEQ(gconv.Bool("off"), false)
		t.AssertEQ(gconv.Bool([]byte{}), false)
		t.AssertEQ(gconv.Bool([]string{}), false)
		t.AssertEQ(gconv.Bool([]interface{}{}), false)
		t.AssertEQ(gconv.Bool([]map[int]int{}), false)

		t.AssertEQ(gconv.Bool("1"), true)
		t.AssertEQ(gconv.Bool("on"), true)
		t.AssertEQ(gconv.Bool(1), true)
		t.AssertEQ(gconv.Bool(123.456), true)
		t.AssertEQ(gconv.Bool(boolStruct{}), true)
		t.AssertEQ(gconv.Bool(&boolStruct{}), true)
	})
}
