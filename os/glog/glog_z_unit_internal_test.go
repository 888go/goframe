// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"bytes"
	"context"
	"testing"

	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

var (
	ctx = context.TODO()
)

func Test_Print(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := X创建并按writer(w)
		l.X输出(ctx, 1, 2, 3)
		l.X输出并格式化(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), "["), 0)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Debug(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := X创建并按writer(w)
		l.X输出DEBU(ctx, 1, 2, 3)
		l.X输出并格式化DEBU(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_DEBU]), 2)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Info(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := X创建并按writer(w)
		l.X输出INFO(ctx, 1, 2, 3)
		l.X输出并格式化INFO(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_INFO]), 2)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Notice(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := X创建并按writer(w)
		l.X输出NOTI(ctx, 1, 2, 3)
		l.X输出并格式化NOTI(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_NOTI]), 2)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Warning(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := X创建并按writer(w)
		l.X输出WARN(ctx, 1, 2, 3)
		l.X输出并格式化WARN(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_WARN]), 2)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_Error(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		w := bytes.NewBuffer(nil)
		l := X创建并按writer(w)
		l.Error(ctx, 1, 2, 3)
		l.X输出并格式化ERR(ctx, "%d %d %d", 1, 2, 3)
		t.Assert(gstr.X统计次数(w.String(), defaultLevelPrefixes[LEVEL_ERRO]), 2)
		t.Assert(gstr.X统计次数(w.String(), "1 2 3"), 2)
	})
}

func Test_LevelPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		l := X创建()
		t.Assert(l.X取级别前缀(LEVEL_DEBU), defaultLevelPrefixes[LEVEL_DEBU])
		t.Assert(l.X取级别前缀(LEVEL_INFO), defaultLevelPrefixes[LEVEL_INFO])
		t.Assert(l.X取级别前缀(LEVEL_NOTI), defaultLevelPrefixes[LEVEL_NOTI])
		t.Assert(l.X取级别前缀(LEVEL_WARN), defaultLevelPrefixes[LEVEL_WARN])
		t.Assert(l.X取级别前缀(LEVEL_ERRO), defaultLevelPrefixes[LEVEL_ERRO])
		t.Assert(l.X取级别前缀(LEVEL_CRIT), defaultLevelPrefixes[LEVEL_CRIT])
		l.X设置级别前缀(LEVEL_DEBU, "debug")
		t.Assert(l.X取级别前缀(LEVEL_DEBU), "debug")
		l.X设置级别前缀Map(map[int]string{
			LEVEL_CRIT: "critical",
		})
		t.Assert(l.X取级别前缀(LEVEL_DEBU), "debug")
		t.Assert(l.X取级别前缀(LEVEL_INFO), defaultLevelPrefixes[LEVEL_INFO])
		t.Assert(l.X取级别前缀(LEVEL_NOTI), defaultLevelPrefixes[LEVEL_NOTI])
		t.Assert(l.X取级别前缀(LEVEL_WARN), defaultLevelPrefixes[LEVEL_WARN])
		t.Assert(l.X取级别前缀(LEVEL_ERRO), defaultLevelPrefixes[LEVEL_ERRO])
		t.Assert(l.X取级别前缀(LEVEL_CRIT), "critical")
	})
	gtest.C(t, func(t *gtest.T) {
		buffer := bytes.NewBuffer(nil)
		l := X创建()
		l.X设置Writer(buffer)
		l.X输出DEBU(ctx, "test1")
		t.Assert(gstr.X是否包含(buffer.String(), defaultLevelPrefixes[LEVEL_DEBU]), true)

		buffer.Reset()

		l.X设置级别前缀(LEVEL_DEBU, "debug")
		l.X输出DEBU(ctx, "test2")
		t.Assert(gstr.X是否包含(buffer.String(), defaultLevelPrefixes[LEVEL_DEBU]), false)
		t.Assert(gstr.X是否包含(buffer.String(), "debug"), true)

		buffer.Reset()
		l.X设置级别前缀Map(map[int]string{
			LEVEL_ERRO: "error",
		})
		l.Error(ctx, "test3")
		t.Assert(gstr.X是否包含(buffer.String(), defaultLevelPrefixes[LEVEL_ERRO]), false)
		t.Assert(gstr.X是否包含(buffer.String(), "error"), true)
	})
}
