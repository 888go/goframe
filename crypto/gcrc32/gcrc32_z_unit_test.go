// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 使用go test命令运行当前目录下所有.go文件的性能测试，模式为匹配所有函数. md5:b546d3aaffaebd06

package 加密crc32类_test

import (
	"testing"

	gcrc32 "github.com/888go/goframe/crypto/gcrc32"
	gmd5 "github.com/888go/goframe/crypto/gmd5"
	gtest "github.com/888go/goframe/test/gtest"
)

func TestEncrypt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		s := "pibigstar"
		result := 693191136
		encrypt1 := gcrc32.X加密(s)
		encrypt2 := gcrc32.X加密([]byte(s))
		t.AssertEQ(int(encrypt1), result)
		t.AssertEQ(int(encrypt2), result)

		strmd5, _ := gmd5.X加密(s)
		test1 := gcrc32.X加密(strmd5)
		test2 := gcrc32.X加密([]byte(strmd5))
		t.AssertEQ(test2, test1)
	})
}
