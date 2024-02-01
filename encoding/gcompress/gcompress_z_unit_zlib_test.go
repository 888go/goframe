// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcompress_test
import (
	"testing"
	
	"github.com/888go/goframe/encoding/gcompress"
	"github.com/888go/goframe/test/gtest"
	)

func Test_Zlib_UnZlib(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := gcompress.Zlib([]byte(src))
		t.Assert(data, dst)

		data, _ = gcompress.UnZlib(dst)
		t.Assert(data, []byte(src))

		data, _ = gcompress.Zlib(nil)
		t.Assert(data, nil)
		data, _ = gcompress.UnZlib(nil)
		t.Assert(data, nil)

		data, _ = gcompress.UnZlib(dst[1:])
		t.Assert(data, nil)
	})
}
