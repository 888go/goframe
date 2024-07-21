// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// RS 实现了经典的 32 位 RS 哈希算法。 md5:b79898a4563914ab
func RS(str []byte) uint32 {
	var (
		b    uint32 = 378551
		a    uint32 = 63689
		hash uint32 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*a + uint32(str[i])
		a *= b
	}
	return hash
}

// RS64实现了经典RS哈希算法用于64位。 md5:26281dc7803cc7e0
func RS64(str []byte) uint64 {
	var (
		b    uint64 = 378551
		a    uint64 = 63689
		hash uint64 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*a + uint64(str[i])
		a *= b
	}
	return hash
}
