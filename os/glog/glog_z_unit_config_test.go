// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"bytes"
	"strings"
	"testing"
	
	"github.com/888go/goframe/test/gtest"
)

func Test_SetConfigWithMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		l := New()
		m := map[string]interface{}{
			"path":     "/var/log",
			"level":    "all",
			"stdout":   false,
			"StStatus": 0,
		}
		err := l.SetConfigWithMap(m)
		t.AssertNil(err)
		t.Assert(l.config.Path, m["path"])
		t.Assert(l.config.Level, LEVEL_ALL)
		t.Assert(l.config.StdoutPrint, m["stdout"])
	})
}

func Test_SetConfigWithMap_LevelStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		buffer := bytes.NewBuffer(nil)
		l := New()
		m := map[string]interface{}{
			"level": "all",
		}
		err := l.SetConfigWithMap(m)
		t.AssertNil(err)

		l.SetWriter(buffer)

		l.Debug(ctx, "test")
		l.Warning(ctx, "test")
		t.Assert(strings.Contains(buffer.String(), "DEBU"), true)
		t.Assert(strings.Contains(buffer.String(), "WARN"), true)
	})

	gtest.C(t, func(t *gtest.T) {
		buffer := bytes.NewBuffer(nil)
		l := New()
		m := map[string]interface{}{
			"level": "warn",
		}
		err := l.SetConfigWithMap(m)
		t.AssertNil(err)
		l.SetWriter(buffer)
		l.Debug(ctx, "test")
		l.Warning(ctx, "test")
		t.Assert(strings.Contains(buffer.String(), "DEBU"), false)
		t.Assert(strings.Contains(buffer.String(), "WARN"), true)
	})
}
