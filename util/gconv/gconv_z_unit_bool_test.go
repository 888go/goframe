// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gconv_test

import (
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
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
