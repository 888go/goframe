// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// SDBM实现了经典的32位SDBM哈希算法。. md5:bd39a0f984820134
func SDBM(str []byte) uint32 {
	var hash uint32
	for i := 0; i < len(str); i++ {
		// 等同于：hash = 65599 * hash + uint32(str[i])；. md5:0315395c252a9dac
		hash = uint32(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// SDBM64 实现了经典的 64 位 SDBM 哈希算法。. md5:ef45ecc578ca9c54
func SDBM64(str []byte) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		// equivalent to: hash = 65599*hash + uint32(str[i])
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
