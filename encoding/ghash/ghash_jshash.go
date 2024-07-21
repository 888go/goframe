// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// JS 实现了经典的 JS 哈希算法，用于32位。 md5:a87e9c7d1a1f7232
func JS(str []byte) uint32 {
	var hash uint32 = 1315423911
	for i := 0; i < len(str); i++ {
		hash ^= (hash << 5) + uint32(str[i]) + (hash >> 2)
	}
	return hash
}

// JS64实现了经典的64位JS哈希算法。 md5:8248afb9bdc18117
func JS64(str []byte) uint64 {
	var hash uint64 = 1315423911
	for i := 0; i < len(str); i++ {
		hash ^= (hash << 5) + uint64(str[i]) + (hash >> 2)
	}
	return hash
}
