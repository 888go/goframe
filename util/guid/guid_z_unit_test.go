// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package guid_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
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
