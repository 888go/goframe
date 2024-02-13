// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 随机类

import (
	"crypto/rand"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
)

const (
	// Buffer大小用于uint32随机数。
	bufferChanSize = 10000
)

var (
// bufferChan 是用于存储随机字节的缓冲通道，
// 每个元素存储 4 字节。
	bufferChan = make(chan []byte, bufferChanSize)
)

func init() {
	go asyncProducingRandomBufferBytesLoop()
}

// asyncProducingRandomBufferBytes 是一个命名的 goroutine，它使用异步 goroutine 生成随机字节，并通过一个缓冲通道（chan）来存储这些随机字节。
// 因此，它在生成随机数时具有较高的性能。
func asyncProducingRandomBufferBytesLoop() {
	var step int
	for {
		buffer := make([]byte, 1024)
		if n, err := rand.Read(buffer); err != nil {
			panic(错误类.X多层错误码(错误码类.CodeInternalError, err, `error reading random buffer from system`))
		} else {
// 系统提供的随机缓冲区代价非常高昂，
// 因此，通过改变步进值（使用不同的数字）来充分复用随机缓冲区，
// 可以显著提升性能。
// 对以下整数数组进行遍历：[4, 5, 6, 7] {
			for _, step = range []int{4} {
				for i := 0; i <= n-4; i += step {
					bufferChan <- buffer[i : i+4]
				}
			}
		}
	}
}
