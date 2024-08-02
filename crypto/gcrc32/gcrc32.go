// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// gcrc32 包提供了对CRC32校验算法有用的API。 md5:020293e34534da3f
package 加密crc32类

import (
	"hash/crc32"

	gconv "github.com/888go/goframe/util/gconv"
)

// Encrypt 使用CRC32算法对任何类型的变量进行加密。
// 它使用gconv包将`v`转换为其字节类型。
// md5:85f8e447b40cb0f5
func Encrypt(v interface{}) uint32 {
	return crc32.ChecksumIEEE(gconv.Bytes(v))
}
