// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 随机类_test

import (
	"strings"
	"testing"
	"time"
	
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/888go/goframe/grand"
)

func Test_Intn(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 1000000; i++ {
			n := 随机类.X整数(100)
			t.AssertLT(n, 100)
			t.AssertGE(n, 0)
		}
		for i := 0; i < 1000000; i++ {
			n := 随机类.X整数(-100)
			t.AssertLE(n, 0)
			t.Assert(n, -100)
		}
	})
}

func Test_Meet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(随机类.Meet(100, 100), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(随机类.Meet(0, 100), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(随机类.Meet(50, 100), []bool{true, false})
		}
	})
}

func Test_MeetProb(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(随机类.MeetProb(1), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(随机类.MeetProb(0), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(随机类.MeetProb(0.5), []bool{true, false})
		}
	})
}

func Test_N(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(随机类.X区间整数(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(随机类.X区间整数(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(随机类.X区间整数(1, 2), []int{1, 2})
		}
	})
}

func Test_D(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(随机类.X时长(time.Second, time.Second), time.Second)
		}
		for i := 0; i < 100; i++ {
			t.Assert(随机类.X时长(0, 0), time.Duration(0))
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(
				随机类.X时长(1*time.Second, 3*time.Second),
				[]time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second},
			)
		}
	})
}

func Test_Rand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(随机类.X区间整数(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(随机类.X区间整数(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(随机类.X区间整数(1, 2), []int{1, 2})
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(随机类.X区间整数(-1, 2), []int{-1, 0, 1, 2})
		}
	})
}

func Test_S(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X文本(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X文本(5, true)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(随机类.X文本(0)), 0)
	})
}

func Test_B(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			b := 随机类.X字节集(5)
			t.Assert(len(b), 5)
			t.AssertNE(b, make([]byte, 5))
		}
	})
	gtest.C(t, func(t *gtest.T) {
		b := 随机类.X字节集(0)
		t.AssertNil(b)
	})
}

func Test_Str(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X文本(5)), 5)
		}
	})
}

func Test_RandStr(t *testing.T) {
	str := "我爱GoFrame"
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 10; i++ {
			s := 随机类.X从文本生成文本(str, 100000)
			t.Assert(gstr.Contains(s, "我"), true)
			t.Assert(gstr.Contains(s, "爱"), true)
			t.Assert(gstr.Contains(s, "G"), true)
			t.Assert(gstr.Contains(s, "o"), true)
			t.Assert(gstr.Contains(s, "F"), true)
			t.Assert(gstr.Contains(s, "r"), true)
			t.Assert(gstr.Contains(s, "a"), true)
			t.Assert(gstr.Contains(s, "m"), true)
			t.Assert(gstr.Contains(s, "e"), true)
			t.Assert(gstr.Contains(s, "w"), false)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(随机类.X从文本生成文本(str, 0), "")
	})
	gtest.C(t, func(t *gtest.T) {
		list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		str := ""
		for _, s := range list {
			tmp := ""
			for i := 0; i < 15; i++ {
				tmp += tmp + s
			}
			str += tmp
		}
		t.Assert(len(随机类.X从文本生成文本(str, 300)), 300)
	})
}

func Test_Digits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X数字文本(5)), 5)
		}
	})
}

func Test_RandDigits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X数字文本(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(随机类.X数字文本(0)), 0)
	})
}

func Test_Letters(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X字母文本(5)), 5)
		}
	})
}

func Test_RandLetters(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(随机类.X字母文本(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(随机类.X字母文本(0)), 0)
	})
}

func Test_Perm(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.AssertIN(随机类.X整数数组(5), []int{0, 1, 2, 3, 4})
		}
	})
}

func Test_Symbols(t *testing.T) {
	symbols := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			syms := []byte(随机类.X特殊字符文本(5))
			for _, sym := range syms {
				t.AssertNE(strings.Index(symbols, string(sym)), -1)
			}
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(随机类.X特殊字符文本(0), "")
	})
}
