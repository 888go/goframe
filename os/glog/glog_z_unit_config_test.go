	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package glog

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
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
