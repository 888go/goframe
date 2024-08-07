// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 环境变量类_test

import (
	"os"
	"testing"

	"github.com/888go/goframe/frame/g"
	gcmd "github.com/888go/goframe/os/gcmd"
	genv "github.com/888go/goframe/os/genv"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func Test_GEnv_All(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(os.Environ(), genv.X取全部())
	})
}

func Test_GEnv_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.Assert(genv.X取Map()[key], "TEST")
	})
}

func Test_GEnv_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(genv.X取值(key).String(), "TEST")
	})
}

func Test_GEnv_GetVar(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(genv.X取值(key).String(), "TEST")
	})
}

func Test_GEnv_Contains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(genv.X是否存在(key), true)
		t.AssertEQ(genv.X是否存在("none"), false)
	})
}

func Test_GEnv_Set(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := genv.X设置值(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "TEST")
	})
}

func Test_GEnv_SetMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := genv.X设置Map值(g.MapStrStr{
			"K1": "TEST1",
			"K2": "TEST2",
		})
		t.AssertNil(err)
		t.AssertEQ(os.Getenv("K1"), "TEST1")
		t.AssertEQ(os.Getenv("K2"), "TEST2")
	})
}

func Test_GEnv_Build(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := genv.Map到切片(map[string]string{
			"k1": "v1",
			"k2": "v2",
		})
		t.AssertIN("k1=v1", s)
		t.AssertIN("k2=v2", s)
	})
}

func Test_GEnv_Remove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		err = genv.X删除(key)
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "")
	})
}

func Test_GetWithCmd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gcmd.Init("-test", "2")
		t.Assert(genv.X取值或命令行("TEST"), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		genv.X设置值("TEST", "1")
		defer genv.X删除("TEST")
		gcmd.Init("-test", "2")
		t.Assert(genv.X取值或命令行("test"), 1)
	})
}

func Test_MapFromEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := genv.X切片到Map([]string{"a=1", "b=2"})
		t.Assert(m, g.Map{"a": 1, "b": 2})
	})
}

func Test_MapToEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := genv.MapToEnv别名(g.MapStrStr{"a": "1"})
		t.Assert(s, []string{"a=1"})
	})
}

func Test_Filter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := genv.X切片去重([]string{"a=1", "a=3"})
		t.Assert(s, []string{"a=3"})
	})
}
