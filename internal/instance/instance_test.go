// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package instance_test

import (
	"testing"
	
	"github.com/888go/goframe/internal/instance"
	"github.com/888go/goframe/test/gtest"
)

func Test_SetGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		instance.Set("test-user", 1)
		t.Assert(instance.Get("test-user"), 1)
		t.Assert(instance.Get("none-exists"), nil)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(instance.GetOrSet("test-1", 1), 1)
		t.Assert(instance.Get("test-1"), 1)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(instance.GetOrSetFunc("test-2", func() interface{} {
			return 2
		}), 2)
		t.Assert(instance.Get("test-2"), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(instance.GetOrSetFuncLock("test-3", func() interface{} {
			return 3
		}), 3)
		t.Assert(instance.Get("test-3"), 3)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(instance.SetIfNotExist("test-4", 4), true)
		t.Assert(instance.Get("test-4"), 4)
		t.Assert(instance.SetIfNotExist("test-4", 5), false)
		t.Assert(instance.Get("test-4"), 4)
	})
}
