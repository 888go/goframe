// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 压缩类_test

import (
	"testing"
	
	"github.com/888go/goframe/gcompress"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Zlib_UnZlib(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := 压缩类.Zlib压缩字节集([]byte(src))
		t.Assert(data, dst)

		data, _ = 压缩类.Zlib解压字节集(dst)
		t.Assert(data, []byte(src))

		data, _ = 压缩类.Zlib压缩字节集(nil)
		t.Assert(data, nil)
		data, _ = 压缩类.Zlib解压字节集(nil)
		t.Assert(data, nil)

		data, _ = 压缩类.Zlib解压字节集(dst[1:])
		t.Assert(data, nil)
	})
}
