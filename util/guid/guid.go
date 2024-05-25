// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// guid包提供了简单且高性能的唯一ID生成功能。. md5:22d1fe7516a2dff2
package guid

import (
	"os"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/util/grand"
)

const (
	sequenceMax   = uint32(46655)                          // Sequence max("zzz").
	randomStrBase = "0123456789abcdefghijklmnopqrstuvwxyz" // 随机字符字符串（36个字节）。. md5:0e81d3c5a56e7cf2
)

var (
	sequence     gtype.Uint32 // 当前进程独有的序列。. md5:2e6129c144b94c7b
	macAddrStr   = "0000000"  // MAC地址的哈希结果为7字节。. md5:99aa7e69b289dd55
	processIdStr = "0000"     // Process id in 4 bytes.
)

// init 函数用于初始化几个固定的局部变量。. md5:3e44426e20423c37
func init() {
	// MAC地址的哈希结果为7字节。. md5:99aa7e69b289dd55
	macs, _ := gipv4.GetMacArray()
	if len(macs) > 0 {
		var macAddrBytes []byte
		for _, mac := range macs {
			macAddrBytes = append(macAddrBytes, []byte(mac)...)
		}
		b := []byte{'0', '0', '0', '0', '0', '0', '0'}
		s := strconv.FormatUint(uint64(ghash.SDBM(macAddrBytes)), 36)
		copy(b, s)
		macAddrStr = string(b)
	}
	// Process id in 4 bytes.
	{
		b := []byte{'0', '0', '0', '0'}
		s := strconv.FormatInt(int64(os.Getpid()), 36)
		copy(b, s)
		processIdStr = string(b)
	}
}

// S 函数创建并返回一个32字节的全局唯一字符串，它满足大多数常见的使用需求，但不严格遵循UUID算法。如果没有提供`data`，则返回默认的唯一字符串。
//
// 指定的`data`最多可以有2个部分。无论每个`data`的长度多长，它们都将被哈希成7字节作为结果的一部分。如果给定的`data`部分少于2个，结果字节的剩余部分将由随机字符串填充。
//
// 返回的字符串由以下组成：
// 1. 默认：MAC哈希(7) + 进程ID(4) + 时间戳纳秒(12) + 序列号(3) + 随机字符串(6)
// 2. 自定义数据：Data哈希(7/14) + 时间戳纳秒(12) + 序列号(3) + 随机字符串(3/10)
//
// 注意：
//  1. 为了性能考虑，返回的长度固定为32字节。
//  2. 自定义参数`data`组合的内容在你的业务场景中应具有唯一性。
// md5:b09b2d34d56e1344
func S(data ...[]byte) string {
	var (
		b       = make([]byte, 32)
		nanoStr = strconv.FormatInt(time.Now().UnixNano(), 36)
	)
	if len(data) == 0 {
		copy(b, macAddrStr)
		copy(b[7:], processIdStr)
		copy(b[11:], nanoStr)
		copy(b[23:], getSequence())
		copy(b[26:], getRandomStr(6))
	} else if len(data) <= 2 {
		n := 0
		for i, v := range data {
			// 忽略空数据项字节。. md5:653aa2fb92e185e8
			if len(v) > 0 {
				copy(b[i*7:], getDataHashStr(v))
				n += 7
			}
		}
		copy(b[n:], nanoStr)
		copy(b[n+12:], getSequence())
		copy(b[n+12+3:], getRandomStr(32-n-12-3))
	} else {
		panic(gerror.NewCode(
			gcode.CodeInvalidParameter,
			"too many data parts, it should be no more than 2 parts",
		))
	}
	return string(b)
}

// getSequence 递增并返回一个以3个字节表示的序列字符串。
// 序列小于"zzz"(46655)。
// md5:742b11b09412718d
func getSequence() []byte {
	b := []byte{'0', '0', '0'}
	s := strconv.FormatUint(uint64(sequence.Add(1)%sequenceMax), 36)
	copy(b, s)
	return b
}

// getRandomStr 从 randomStrBase 中随机选取并返回 `n` 个字符。. md5:fbef2b139ac9b42f
func getRandomStr(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	var (
		b           = make([]byte, n)
		numberBytes = grand.B(n)
	)
	for i := range b {
		b[i] = randomStrBase[numberBytes[i]%36]
	}
	return b
}

// getDataHashStr 根据给定的数据字节创建并返回7字节的哈希字节。. md5:8947e4208efac1a4
func getDataHashStr(data []byte) []byte {
	b := []byte{'0', '0', '0', '0', '0', '0', '0'}
	s := strconv.FormatUint(uint64(ghash.SDBM(data)), 36)
	copy(b, s)
	return b
}
