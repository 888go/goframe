// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 转换类_test

import (
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

type boolStruct struct{}

func Test_Bool(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var any interface{} = nil
		t.AssertEQ(gconv.X取布尔(any), false)
		t.AssertEQ(gconv.X取布尔(false), false)
		t.AssertEQ(gconv.X取布尔(nil), false)
		t.AssertEQ(gconv.X取布尔(0), false)
		t.AssertEQ(gconv.X取布尔("0"), false)
		t.AssertEQ(gconv.X取布尔(""), false)
		t.AssertEQ(gconv.X取布尔("false"), false)
		t.AssertEQ(gconv.X取布尔("off"), false)
		t.AssertEQ(gconv.X取布尔([]byte{}), false)
		t.AssertEQ(gconv.X取布尔([]string{}), false)
		t.AssertEQ(gconv.X取布尔([]interface{}{}), false)
		t.AssertEQ(gconv.X取布尔([]map[int]int{}), false)

		t.AssertEQ(gconv.X取布尔("1"), true)
		t.AssertEQ(gconv.X取布尔("on"), true)
		t.AssertEQ(gconv.X取布尔(1), true)
		t.AssertEQ(gconv.X取布尔(123.456), true)
		t.AssertEQ(gconv.X取布尔(boolStruct{}), true)
		t.AssertEQ(gconv.X取布尔(&boolStruct{}), true)
	})
}
