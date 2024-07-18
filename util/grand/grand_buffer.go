// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package grand

import (
	"crypto/rand"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

const (
	// 用于uint32随机数的缓冲区大小。 md5:7a9c359f249cc1ff
	bufferChanSize = 10000
)

var (
// bufferChan 是用于随机字节的缓冲区，
// 每个元素存储4个字节。
// md5:4d20aff4f4a62f55
	bufferChan = make(chan []byte, bufferChanSize)
)

func init() {
	go asyncProducingRandomBufferBytesLoop()
}

// asyncProducingRandomBufferBytes 是一个命名的 Goroutine，它使用异步 Goroutine
// 来生成随机字节，并通过一个缓冲通道来存储这些随机字节。
// 因此，它具有高性能地生成随机数的特性。
// md5:33995801ded223e5
func asyncProducingRandomBufferBytesLoop() {
	var step int
	for {
		buffer := make([]byte, 1024)
		if n, err := rand.Read(buffer); err != nil {
			panic(gerror.WrapCode(gcode.CodeInternalError, err, `error reading random buffer from system`))
		} else {
// 系统提供的随机缓冲区非常昂贵，
// 因此通过改变步长并使用不同的数字来完全重用随机缓冲区，
// 可以大大提高性能。
// 对于 _step 在整数切片 []int{4, 5, 6, 7} 中的每个元素：
// md5:dc55adf04393f3bc
			for _, step = range []int{4} {
				for i := 0; i <= n-4; i += step {
					bufferChan <- buffer[i : i+4]
				}
			}
		}
	}
}
