// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 环境变量类_test

import (
	"os"
	"testing"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gcmd"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func Test_GEnv_All(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(os.Environ(), 环境变量类.X取全部())
	})
}

func Test_GEnv_Map(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 转换类.String(时间类.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.Assert(环境变量类.X取Map()[key], "TEST")
	})
}

func Test_GEnv_Get(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 转换类.String(时间类.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(环境变量类.X取值(key).String(), "TEST")
	})
}

func Test_GEnv_GetVar(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 转换类.String(时间类.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(环境变量类.X取值(key).String(), "TEST")
	})
}

func Test_GEnv_Contains(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 转换类.String(时间类.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(环境变量类.X是否存在(key), true)
		t.AssertEQ(环境变量类.X是否存在("none"), false)
	})
}

func Test_GEnv_Set(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 转换类.String(时间类.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := 环境变量类.X设置值(key, "TEST")
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "TEST")
	})
}

func Test_GEnv_SetMap(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		err := 环境变量类.X设置Map值(g.MapStrStr{
			"K1": "TEST1",
			"K2": "TEST2",
		})
		t.AssertNil(err)
		t.AssertEQ(os.Getenv("K1"), "TEST1")
		t.AssertEQ(os.Getenv("K2"), "TEST2")
	})
}

func Test_GEnv_Build(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 环境变量类.Map到数组(map[string]string{
			"k1": "v1",
			"k2": "v2",
		})
		t.AssertIN("k1=v1", s)
		t.AssertIN("k2=v2", s)
	})
}

func Test_GEnv_Remove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 转换类.String(时间类.X取时间戳纳秒())
		key := "TEST_ENV_" + value
		err := os.Setenv(key, "TEST")
		t.AssertNil(err)
		err = 环境变量类.X删除(key)
		t.AssertNil(err)
		t.AssertEQ(os.Getenv(key), "")
	})
}

func Test_GetWithCmd(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		cmd类.Init("-test", "2")
		t.Assert(环境变量类.X取值或命令行("TEST"), 2)
	})
	单元测试类.C(t, func(t *单元测试类.T) {
		环境变量类.X设置值("TEST", "1")
		defer 环境变量类.X删除("TEST")
		cmd类.Init("-test", "2")
		t.Assert(环境变量类.X取值或命令行("test"), 1)
	})
}

func Test_MapFromEnv(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		m := 环境变量类.X数组到Map([]string{"a=1", "b=2"})
		t.Assert(m, g.Map{"a": 1, "b": 2})
	})
}

func Test_MapToEnv(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 环境变量类.MapToEnv别名(g.MapStrStr{"a": "1"})
		t.Assert(s, []string{"a=1"})
	})
}

func Test_Filter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := 环境变量类.X数组去重([]string{"a=1", "a=3"})
		t.Assert(s, []string{"a=3"})
	})
}
