// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 单元测试类_test

import (
	"errors"
	"path/filepath"
	"strconv"
	"testing"
	
	"github.com/888go/goframe/test/gtest"
)

var (
	map1           = map[string]string{"k1": "v1"}
	map1Expect     = map[string]string{"k1": "v1"}
	map2           = map[string]string{"k2": "v2"}
	mapLong1       = map[string]string{"k1": "v1", "k2": "v2"}
	mapLong1Expect = map[string]string{"k2": "v2", "k1": "v1"}
)

func TestC(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(1, 1)
		t.AssertNE(1, 0)
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(float32(123.456), float32(123.456))
		t.Assert(map[string]string{"1": "1"}, map[string]string{"1": "1"})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 == 0")
			}
		}()
		t.Assert(1, 0)
	})
}

func TestCase(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(1, 1)
		t.AssertNE(1, 0)
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(float32(123.456), float32(123.456))
	})
}

func TestAssert(t *testing.T) {

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			nilChan chan struct{}
		)
		t.Assert(1, 1)
		t.Assert(nilChan, nil)
		t.Assert(map1, map1Expect)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT VALUE map["k2"]: == map["k2"]:v2
GIVEN : map[k1:v1]
EXPECT: map[k2:v2]`)
			}
		}()
		t.Assert(map1, map2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT MAP LENGTH 2 == 1`)
			}
		}()
		t.Assert(mapLong1, map2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT VALUE TO BE A MAP, BUT GIVEN "int"`)
			}
		}()
		t.Assert(0, map1)
	})
}

func TestAssertEQ(t *testing.T) {

	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			nilChan chan struct{}
		)
		t.AssertEQ(nilChan, nil)
		t.AssertEQ("0", "0")
		t.AssertEQ(float32(123.456), float32(123.456))
		t.AssertEQ(mapLong1, mapLong1Expect)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 == 0")
			}
		}()
		t.AssertEQ(1, 0)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT TYPE 1[int] == 1[string]")
			}
		}()
		t.AssertEQ(1, "1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT VALUE map["k2"]: == map["k2"]:v2
GIVEN : map[k1:v1]
EXPECT: map[k2:v2]`)
			}
		}()
		t.AssertEQ(map1, map2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT MAP LENGTH 2 == 1`)
			}
		}()
		t.AssertEQ(mapLong1, map2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT VALUE TO BE A MAP, BUT GIVEN "int"`)
			}
		}()
		t.AssertEQ(0, map1)
	})
}

func TestAssertNE(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			c = make(chan struct{}, 1)
		)
		t.AssertNE(nil, c)
		t.AssertNE("0", "1")
		t.AssertNE(float32(123.456), float32(123.4567))
		t.AssertNE(map1, map2)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 != 1")
			}
		}()
		t.AssertNE(1, 1)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, `[ASSERT] EXPECT map[k1:v1] != map[k1:v1]`)
			}
		}()
		t.AssertNE(map1, map1Expect)
	})
}

func TestAssertNQ(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNQ(1, "0")
		t.AssertNQ(float32(123.456), float64(123.4567))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 != 1")
			}
		}()
		t.AssertNQ(1, "1")
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT TYPE 1[int] != 1[int]")
			}
		}()
		t.AssertNQ(1, 1)
	})
}

func TestAssertGT(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGT("b", "a")
		t.AssertGT(1, -1)
		t.AssertGT(uint(1), uint(0))
		t.AssertGT(float32(123.45678), float32(123.4567))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT -1 > 1")
			}
		}()
		t.AssertGT(-1, 1)
	})
}

func TestAssertGE(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGE("b", "a")
		t.AssertGE("a", "a")
		t.AssertGE(1, -1)
		t.AssertGE(1, 1)
		t.AssertGE(uint(1), uint(0))
		t.AssertGE(uint(0), uint(0))
		t.AssertGE(float32(123.45678), float32(123.4567))
		t.AssertGE(float32(123.456), float32(123.456))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT -1(int) >= 1(int)")
			}
		}()
		t.AssertGE(-1, 1)
	})
}

func TestAssertLT(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertLT("a", "b")
		t.AssertLT(-1, 1)
		t.AssertLT(uint(0), uint(1))
		t.AssertLT(float32(123.456), float32(123.4567))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 < -1")
			}
		}()
		t.AssertLT(1, -1)
	})
}

func TestAssertLE(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertLE("a", "b")
		t.AssertLE("a", "a")
		t.AssertLE(-1, 1)
		t.AssertLE(1, 1)
		t.AssertLE(uint(0), uint(1))
		t.AssertLE(uint(0), uint(0))
		t.AssertLE(float32(123.456), float32(123.4567))
		t.AssertLE(float32(123.456), float32(123.456))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 <= -1")
			}
		}()
		t.AssertLE(1, -1)
	})
}

func TestAssertIN(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertIN("a", []string{"a", "b", "c"})
		t.AssertIN(1, []int{1, 2, 3})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] INVALID EXPECT VALUE TYPE: int")
			}
		}()
		t.AssertIN(0, 0)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 4 IN [1 2 3]")
			}
		}()
// t.AssertIN(0, []int{0, 1, 2, 3})
// 检查整数0是否在给定的整数切片([]int)中，该切片包含元素0、1、2和3
// t.AssertIN(0, []int{1, 2, 3})
// 检查整数0是否在给定的整数切片([]int)中，该切片包含元素1、2、3
		t.AssertIN(4, []int{1, 2, 3})
	})
}

func TestAssertNI(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertNI("d", []string{"a", "b", "c"})
		t.AssertNI(4, []int{1, 2, 3})
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] INVALID EXPECT VALUE TYPE: int")
			}
		}()
		t.AssertNI(0, 0)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ASSERT] EXPECT 1 NOT IN [1 2 3]")
			}
		}()
		t.AssertNI(1, []int{1, 2, 3})
	})
}

func TestAssertNil(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			nilChan chan struct{}
		)
		t.AssertNil(nilChan)
		_, err := strconv.ParseInt("123", 10, 64)
		t.AssertNil(err)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "error")
			}
		}()
		t.AssertNil(errors.New("error"))
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.AssertNE(err, nil)
			}
		}()
		t.AssertNil(1)
	})
}

func TestAssertError(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Assert(err, "[ERROR] this is an error")
			}
		}()
		t.Error("this is an error")
	})
}

func TestDataPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(filepath.ToSlash(单元测试类.DataPath("testdata.txt")), `./testdata/testdata.txt`)
	})
}

func TestDataContent(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(单元测试类.DataContent("testdata.txt"), `hello`)
		t.Assert(单元测试类.DataContent(""), "")
	})
}
