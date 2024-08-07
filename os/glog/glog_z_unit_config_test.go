// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"bytes"
	"strings"
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
)

func Test_SetConfigWithMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		l := X创建()
		m := map[string]interface{}{
			"path":     "/var/log",
			"level":    "all",
			"stdout":   false,
			"StStatus": 0,
		}
		err := l.X设置配置Map(m)
		t.AssertNil(err)
		t.Assert(l.config.Path, m["path"])
		t.Assert(l.config.Level, LEVEL_ALL)
		t.Assert(l.config.StdoutPrint, m["stdout"])
	})
}

func Test_SetConfigWithMap_LevelStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
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

	gtest.C(t, func(t *gtest.T) {
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
