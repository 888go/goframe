// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package guid_test

import (
	"testing"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/guid"
)

func Test_S(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		set := gset.NewStrSet()
		for i := 0; i < 1000000; i++ {
			s := guid.S()
			t.Assert(set.AddIfNotExist(s), true)
			t.Assert(len(s), 32)
		}
	})
}

func Test_S_Data(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(guid.S([]byte("123"))), 32)
	})
}
