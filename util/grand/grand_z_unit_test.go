// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 随机类_test

import (
	"strings"
	"testing"
	"time"

	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
	grand "github.com/888go/goframe/util/grand"
)

func Test_Intn(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 1000000; i++ {
			n := grand.X整数(100)
			t.AssertLT(n, 100)
			t.AssertGE(n, 0)
		}
		for i := 0; i < 1000000; i++ {
			n := grand.X整数(-100)
			t.AssertLE(n, 0)
			t.Assert(n, -100)
		}
	})
}

func Test_Meet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.Meet(100, 100), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.Meet(0, 100), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.Meet(50, 100), []bool{true, false})
		}
	})
}

func Test_MeetProb(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.MeetProb(1), true)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.MeetProb(0), false)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.MeetProb(0.5), []bool{true, false})
		}
	})
}

func Test_N(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.X区间整数(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.X区间整数(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.X区间整数(1, 2), []int{1, 2})
		}
	})
}

func Test_D(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.X时长(time.Second, time.Second), time.Second)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.X时长(0, 0), time.Duration(0))
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(
				grand.X时长(1*time.Second, 3*time.Second),
				[]time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second},
			)
		}
	})
}

func Test_Rand(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(grand.X区间整数(1, 1), 1)
		}
		for i := 0; i < 100; i++ {
			t.Assert(grand.X区间整数(0, 0), 0)
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.X区间整数(1, 2), []int{1, 2})
		}
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.X区间整数(-1, 2), []int{-1, 0, 1, 2})
		}
	})
}

func Test_S(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X文本(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X文本(5, true)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(grand.X文本(0)), 0)
	})
}

func Test_B(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			b := grand.X字节集(5)
			t.Assert(len(b), 5)
			t.AssertNE(b, make([]byte, 5))
		}
	})
	gtest.C(t, func(t *gtest.T) {
		b := grand.X字节集(0)
		t.AssertNil(b)
	})
}

func Test_Str(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X文本(5)), 5)
		}
	})
}

func Test_RandStr(t *testing.T) {
	str := "我爱GoFrame"
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 10; i++ {
			s := grand.X从文本生成文本(str, 100000)
			t.Assert(gstr.X是否包含(s, "我"), true)
			t.Assert(gstr.X是否包含(s, "爱"), true)
			t.Assert(gstr.X是否包含(s, "G"), true)
			t.Assert(gstr.X是否包含(s, "o"), true)
			t.Assert(gstr.X是否包含(s, "F"), true)
			t.Assert(gstr.X是否包含(s, "r"), true)
			t.Assert(gstr.X是否包含(s, "a"), true)
			t.Assert(gstr.X是否包含(s, "m"), true)
			t.Assert(gstr.X是否包含(s, "e"), true)
			t.Assert(gstr.X是否包含(s, "w"), false)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(grand.X从文本生成文本(str, 0), "")
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
		t.Assert(len(grand.X从文本生成文本(str, 300)), 300)
	})
}

func Test_Digits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X数字文本(5)), 5)
		}
	})
}

func Test_RandDigits(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X数字文本(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(grand.X数字文本(0)), 0)
	})
}

func Test_Letters(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X字母文本(5)), 5)
		}
	})
}

func Test_RandLetters(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.Assert(len(grand.X字母文本(5)), 5)
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(len(grand.X字母文本(0)), 0)
	})
}

func Test_Perm(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			t.AssertIN(grand.X整数切片(5), []int{0, 1, 2, 3, 4})
		}
	})
}

func Test_Symbols(t *testing.T) {
	symbols := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			syms := []byte(grand.X特殊字符文本(5))
			for _, sym := range syms {
				t.AssertNE(strings.Index(symbols, string(sym)), -1)
			}
		}
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(grand.X特殊字符文本(0), "")
	})
}
