// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类

import (
	"bytes"
	"strings"
	"testing"
	
	"github.com/888go/goframe/test/gtest"
)

func Test_SetConfigWithMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := X创建()
		m := map[string]interface{}{
			"path":     "/var/log",
			"level":    "all",
			"stdout":   false,
			"StStatus": 0,
		}
		err := l.X设置配置Map(m)
		t.AssertNil(err)
		t.Assert(l.config.X文件路径, m["path"])
		t.Assert(l.config.X级别, LEVEL_ALL)
		t.Assert(l.config.X是否同时输出到终端, m["stdout"])
	})
}

func Test_SetConfigWithMap_LevelStr(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		buffer := bytes.NewBuffer(nil)
		l := X创建()
		m := map[string]interface{}{
			"level": "all",
		}
		err := l.X设置配置Map(m)
		t.AssertNil(err)

		l.X设置Writer(buffer)

		l.X输出DEBU(ctx, "test")
		l.X输出WARN(ctx, "test")
		t.Assert(strings.Contains(buffer.String(), "DEBU"), true)
		t.Assert(strings.Contains(buffer.String(), "WARN"), true)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		buffer := bytes.NewBuffer(nil)
		l := X创建()
		m := map[string]interface{}{
			"level": "warn",
		}
		err := l.X设置配置Map(m)
		t.AssertNil(err)
		l.X设置Writer(buffer)
		l.X输出DEBU(ctx, "test")
		l.X输出WARN(ctx, "test")
		t.Assert(strings.Contains(buffer.String(), "DEBU"), false)
		t.Assert(strings.Contains(buffer.String(), "WARN"), true)
	})
}
