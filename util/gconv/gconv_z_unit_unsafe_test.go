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
