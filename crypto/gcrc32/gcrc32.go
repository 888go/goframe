// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gcrc32 提供了用于 CRC32 加密算法的有用 API。
package gcrc32
import (
	"hash/crc32"
	
	"github.com/888go/goframe/util/gconv"
	)
// Encrypt 使用CRC32算法对任意类型的变量进行加密。
// 它使用gconv包将`v`转换为字节类型。
func Encrypt(v interface{}) uint32 {
	return crc32.ChecksumIEEE(gconv.Bytes(v))
}
