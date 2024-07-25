// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package ghash

// AP 实现了经典的 32 位 AP 哈希算法。 md5:191dd9c40a2bc4c0
func AP(str []byte) uint32 {
	var hash uint32
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}

// AP64 实现了经典的64位AP哈希算法。 md5:90dfb2b5d6382bbe
func AP64(str []byte) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint64(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint64(str[i]) ^ (hash >> 5)) + 1
		}
	}
	return hash
}
