// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package uid类_test

import (
	"testing"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/guid"
)

func Test_S(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		set := 集合类.X创建文本()
		for i := 0; i < 1000000; i++ {
			s := uid类.X生成()
			t.Assert(set.X加入值并跳过已存在(s), true)
			t.Assert(len(s), 32)
		}
	})
}

func Test_S_Data(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(len(uid类.X生成([]byte("123"))), 32)
	})
}
