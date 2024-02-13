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

type boolStruct struct{}

func Test_Bool(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var any interface{} = nil
		t.AssertEQ(转换类.X取布尔(any), false)
		t.AssertEQ(转换类.X取布尔(false), false)
		t.AssertEQ(转换类.X取布尔(nil), false)
		t.AssertEQ(转换类.X取布尔(0), false)
		t.AssertEQ(转换类.X取布尔("0"), false)
		t.AssertEQ(转换类.X取布尔(""), false)
		t.AssertEQ(转换类.X取布尔("false"), false)
		t.AssertEQ(转换类.X取布尔("off"), false)
		t.AssertEQ(转换类.X取布尔([]byte{}), false)
		t.AssertEQ(转换类.X取布尔([]string{}), false)
		t.AssertEQ(转换类.X取布尔([]interface{}{}), false)
		t.AssertEQ(转换类.X取布尔([]map[int]int{}), false)

		t.AssertEQ(转换类.X取布尔("1"), true)
		t.AssertEQ(转换类.X取布尔("on"), true)
		t.AssertEQ(转换类.X取布尔(1), true)
		t.AssertEQ(转换类.X取布尔(123.456), true)
		t.AssertEQ(转换类.X取布尔(boolStruct{}), true)
		t.AssertEQ(转换类.X取布尔(&boolStruct{}), true)
	})
}
