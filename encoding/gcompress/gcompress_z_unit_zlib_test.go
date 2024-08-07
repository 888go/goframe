// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 压缩类_test

import (
	"testing"

	gcompress "github.com/888go/goframe/encoding/gcompress"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_Zlib_UnZlib(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		src := "hello, world\n"
		dst := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
		data, _ := gcompress.Zlib压缩字节集([]byte(src))
		t.Assert(data, dst)

		data, _ = gcompress.Zlib解压字节集(dst)
		t.Assert(data, []byte(src))

		data, _ = gcompress.Zlib压缩字节集(nil)
		t.Assert(data, nil)
		data, _ = gcompress.Zlib解压字节集(nil)
		t.Assert(data, nil)

		data, _ = gcompress.Zlib解压字节集(dst[1:])
		t.Assert(data, nil)
	})
}
