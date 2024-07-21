// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// DJB 实现了经典的 DJB 哈希算法，适用于32位。 md5:5643f3da8f083639
// ff:
// str:
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint32(str[i])
	}
	return hash
}

// DJB64 实现了经典的 DJB 哈希算法，用于64位。 md5:a3db0c4ac3099e2b
// ff:
// str:
func DJB64(str []byte) uint64 {
	var hash uint64 = 5381
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint64(str[i])
	}
	return hash
}
