// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// BKDR 实现了经典的 BKDR 哈希算法，用于生成 32 位哈希值。 md5:b92e10073a79eea4
func BKDR(str []byte) uint32 {
	var (
		seed uint32 = 131 // 31 131 1313 13131 131313 等等.... md5:41340a6a56b6dc66
		hash uint32 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint32(str[i])
	}
	return hash
}

// BKDR64 实现了经典的64位BKDR哈希算法。 md5:3767ac6eac21d31d
func BKDR64(str []byte) uint64 {
	var (
		seed uint64 = 131 // 31 131 1313 13131 131313 等等.... md5:41340a6a56b6dc66
		hash uint64 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return hash
}
