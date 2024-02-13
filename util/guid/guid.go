// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package guid 提供了简单且高性能的唯一标识符生成功能。
package uid类

import (
	"os"
	"strconv"
	"time"
	
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/encoding/ghash"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/net/gipv4"
	"github.com/888go/goframe/util/grand"
)

const (
	sequenceMax   = uint32(46655)                          // 获取序列中最大值，示例为"zzz"。
	randomStrBase = "0123456789abcdefghijklmnopqrstuvwxyz" // 随机字符字符串（36字节）
)

var (
	sequence     安全变量类.Uint32 // 用于当前进程唯一目的的序列号
	macAddrStr   = "0000000"  // MAC地址哈希结果为7字节。
	processIdStr = "0000"     // 4字节表示的进程ID
)

// init 初始化几个固定的局部变量。
func init() {
	// MAC地址哈希结果为7字节。
	macs, _ := ipv4类.GetMacArray()
	if len(macs) > 0 {
		var macAddrBytes []byte
		for _, mac := range macs {
			macAddrBytes = append(macAddrBytes, []byte(mac)...)
		}
		b := []byte{'0', '0', '0', '0', '0', '0', '0'}
		s := strconv.FormatUint(uint64(哈希类.SDBM(macAddrBytes)), 36)
		copy(b, s)
		macAddrStr = string(b)
	}
	// 4字节表示的进程ID
	{
		b := []byte{'0', '0', '0', '0'}
		s := strconv.FormatInt(int64(os.Getpid()), 36)
		copy(b, s)
		processIdStr = string(b)
	}
}

// S creates and returns a global unique string in 32 bytes that meets most common
// usages without strict UUID algorithm. It returns a unique string using default
// unique algorithm if no `data` is given.
//
// The specified `data` can be no more than 2 parts. No matter how long each of the
// `data` size is, each of them will be hashed into 7 bytes as part of the result.
// If given `data` parts is less than 2, the leftover size of the result bytes will
// be token by random string.
//
// The returned string is composed with:
// 1. Default:    MACHash(7) + PID(4) + TimestampNano(12) + Sequence(3) + RandomString(6)
// 2. CustomData: DataHash(7/14) + TimestampNano(12) + Sequence(3) + RandomString(3/10)
//
// Note that：
//  1. The returned length is fixed to 32 bytes for performance purpose.
//  2. The custom parameter `data` composed should have unique attribute in your
//     business scenario.
func X生成(参数 ...[]byte) string {
	var (
		b       = make([]byte, 32)
		nanoStr = strconv.FormatInt(time.Now().UnixNano(), 36)
	)
	if len(参数) == 0 {
		copy(b, macAddrStr)
		copy(b[7:], processIdStr)
		copy(b[11:], nanoStr)
		copy(b[23:], getSequence())
		copy(b[26:], getRandomStr(6))
	} else if len(参数) <= 2 {
		n := 0
		for i, v := range 参数 {
			// 忽略空数据项字节。
			if len(v) > 0 {
				copy(b[i*7:], getDataHashStr(v))
				n += 7
			}
		}
		copy(b[n:], nanoStr)
		copy(b[n+12:], getSequence())
		copy(b[n+12+3:], getRandomStr(32-n-12-3))
	} else {
		panic(错误类.X创建错误码(
			错误码类.CodeInvalidParameter,
			"too many data parts, it should be no more than 2 parts",
		))
	}
	return string(b)
}

// getSequence 以3字节的方式递增并返回序列字符串。
// 序列小于"zzz"(46655)。
func getSequence() []byte {
	b := []byte{'0', '0', '0'}
	s := strconv.FormatUint(uint64(sequence.Add(1)%sequenceMax), 36)
	copy(b, s)
	return b
}

// getRandomStr 随机地从 randomStrBase 中选取并返回 `n` 个字符。
func getRandomStr(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	var (
		b           = make([]byte, n)
		numberBytes = 随机类.X字节集(n)
	)
	for i := range b {
		b[i] = randomStrBase[numberBytes[i]%36]
	}
	return b
}

// getDataHashStr 根据给定的数据字节创建并返回7字节的哈希值。
func getDataHashStr(data []byte) []byte {
	b := []byte{'0', '0', '0', '0', '0', '0', '0'}
	s := strconv.FormatUint(uint64(哈希类.SDBM(data)), 36)
	copy(b, s)
	return b
}
