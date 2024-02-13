// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// 运行go test命令，测试当前目录下所有.go文件，并执行所有benchmark测试

package 加密crc32类_test

import (
	"testing"
	
	"github.com/888go/goframe/crypto/gcrc32"
	"github.com/888go/goframe/crypto/gmd5"
	"github.com/888go/goframe/test/gtest"
)

func TestEncrypt(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		s := "pibigstar"
		result := 693191136
		encrypt1 := 加密crc32类.X加密(s)
		encrypt2 := 加密crc32类.X加密([]byte(s))
		t.AssertEQ(int(encrypt1), result)
		t.AssertEQ(int(encrypt2), result)

		strmd5, _ := 加密md5类.X加密(s)
		test1 := 加密crc32类.X加密(strmd5)
		test2 := 加密crc32类.X加密([]byte(strmd5))
		t.AssertEQ(test2, test1)
	})
}
