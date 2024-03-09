// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类_test

import (
	"testing"
	
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/888go/goframe/gutil"
)

func Test_ComparatorString(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较文本(1, 1), 0)
		t.Assert(工具类.X比较文本(1, 2), -1)
		t.Assert(工具类.X比较文本(2, 1), 1)
	})
}

func Test_ComparatorInt(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较整数(1, 1), 0)
		t.Assert(工具类.X比较整数(1, 2), -1)
		t.Assert(工具类.X比较整数(2, 1), 1)
	})
}

func Test_ComparatorInt8(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较整数8位(1, 1), 0)
		t.Assert(工具类.X比较整数8位(1, 2), -1)
		t.Assert(工具类.X比较整数8位(2, 1), 1)
	})
}

func Test_ComparatorInt16(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较整数16位(1, 1), 0)
		t.Assert(工具类.X比较整数16位(1, 2), -1)
		t.Assert(工具类.X比较整数16位(2, 1), 1)
	})
}

func Test_ComparatorInt32(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较整数32位(1, 1), 0)
		t.Assert(工具类.X比较整数32位(1, 2), -1)
		t.Assert(工具类.X比较整数32位(2, 1), 1)
	})
}

func Test_ComparatorInt64(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较整数64位(1, 1), 0)
		t.Assert(工具类.X比较整数64位(1, 2), -1)
		t.Assert(工具类.X比较整数64位(2, 1), 1)
	})
}

func Test_ComparatorUint(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较正整数(1, 1), 0)
		t.Assert(工具类.X比较正整数(1, 2), -1)
		t.Assert(工具类.X比较正整数(2, 1), 1)
	})
}

func Test_ComparatorUint8(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较正整数8位(1, 1), 0)
		t.Assert(工具类.X比较正整数8位(2, 6), 252)
		t.Assert(工具类.X比较正整数8位(2, 1), 1)
	})
}

func Test_ComparatorUint16(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较正整数16位(1, 1), 0)
		t.Assert(工具类.X比较正整数16位(1, 2), 65535)
		t.Assert(工具类.X比较正整数16位(2, 1), 1)
	})
}

func Test_ComparatorUint32(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较正整数32位(1, 1), 0)
		t.Assert(工具类.X比较正整数32位(-1000, 2147483640), 2147482656)
		t.Assert(工具类.X比较正整数32位(2, 1), 1)
	})
}

func Test_ComparatorUint64(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较正整数64位(1, 1), 0)
		t.Assert(工具类.X比较正整数64位(1, 2), -1)
		t.Assert(工具类.X比较正整数64位(2, 1), 1)
	})
}

func Test_ComparatorFloat32(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较小数32位(1, 1), 0)
		t.Assert(工具类.X比较小数32位(1, 2), -1)
		t.Assert(工具类.X比较小数32位(2, 1), 1)
	})
}

func Test_ComparatorFloat64(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较小数64位(1, 1), 0)
		t.Assert(工具类.X比较小数64位(1, 2), -1)
		t.Assert(工具类.X比较小数64位(2, 1), 1)
	})
}

func Test_ComparatorByte(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较字节(1, 1), 0)
		t.Assert(工具类.X比较字节(1, 2), 255)
		t.Assert(工具类.X比较字节(2, 1), 1)
	})
}

func Test_ComparatorRune(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较字符(1, 1), 0)
		t.Assert(工具类.X比较字符(1, 2), -1)
		t.Assert(工具类.X比较字符(2, 1), 1)
	})
}

func Test_ComparatorTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		j := 工具类.X比较时间("2019-06-14", "2019-06-14")
		t.Assert(j, 0)

		k := 工具类.X比较时间("2019-06-15", "2019-06-14")
		t.Assert(k, 1)

		l := 工具类.X比较时间("2019-06-13", "2019-06-14")
		t.Assert(l, -1)
	})
}

func Test_ComparatorFloat32OfFixed(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较小数32位(0.1, 0.1), 0)
		t.Assert(工具类.X比较小数32位(1.1, 2.1), -1)
		t.Assert(工具类.X比较小数32位(2.1, 1.1), 1)
	})
}

func Test_ComparatorFloat64OfFixed(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		t.Assert(工具类.X比较小数64位(0.1, 0.1), 0)
		t.Assert(工具类.X比较小数64位(1.1, 2.1), -1)
		t.Assert(工具类.X比较小数64位(2.1, 1.1), 1)
	})
}
