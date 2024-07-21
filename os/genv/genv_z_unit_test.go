		// 版权归GoFrame作者(https:		//goframe.org)所有。保留所有权利。
		//
		// 本源代码形式受MIT许可证条款约束。
		// 如果未随本文件一同分发MIT许可证副本，
		// 您可以在https:		//github.com/gogf/gf处获取。
		// md5:a9832f33b234e3f3

package genv_test//bm:环境变量类_test

import (
	"os"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

func Test_GEnv_All(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(os.Environ(), genv.All())
	})
}

func Test_GEnv_Map(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.Assert(genv.Map()[key], "TEST")
	})
}

func Test_GEnv_Get(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(genv.Get(key).String(), "TEST")
	})
}

func Test_GEnv_GetVar(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(genv.Get(key).String(), "TEST")
	})
}

func Test_GEnv_Contains(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(genv.Contains(key), true)
		t.AssertEQ(genv.Contains("none"), false)
	})
}

func Test_GEnv_Set(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.TimestampNano())
		key := "TEST_ENV_" + value
		err := genv.Set(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "TEST")
	})
}

func Test_GEnv_SetMap(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := genv.SetMap(g.MapStrStr{
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
		s := genv.Build(map[string]string{
			"k1": "v1",
			"k2": "v2",
		})
		t.AssertIN("k1=v1", s)
		t.AssertIN("k2=v2", s)
	})
}

func Test_GEnv_Remove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gconv.String(gtime.TimestampNano())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		err = genv.Remove(key)
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "")
	})
}

func Test_GetWithCmd(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gcmd.Init("-test", "2")
		t.Assert(genv.GetWithCmd("TEST"), 2)
	})
	gtest.C(t, func(t *gtest.T) {
		genv.Set("TEST", "1")
		defer genv.Remove("TEST")
		gcmd.Init("-test", "2")
		t.Assert(genv.GetWithCmd("test"), 1)
	})
}

func Test_MapFromEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m := genv.MapFromEnv([]string{"a=1", "b=2"})
		t.Assert(m, g.Map{"a": 1, "b": 2})
	})
}

func Test_MapToEnv(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := genv.MapToEnv(g.MapStrStr{"a": "1"})
		t.Assert(s, []string{"a=1"})
	})
}

func Test_Filter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := genv.Filter([]string{"a=1", "a=3"})
		t.Assert(s, []string{"a=3"})
	})
}
