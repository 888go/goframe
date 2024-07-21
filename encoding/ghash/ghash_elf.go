// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// ELF实现了经典的32位ELF哈希算法。 md5:fafebee32654c802
// ff:
// str:
func ELF(str []byte) uint32 {
	var (
		hash uint32
		x    uint32
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint32(str[i])
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x + 1
		}
	}
	return hash
}

// ELF64 实现经典的ELF（Executable and Linkable Format）64位哈希算法。 md5:0afc74082404f23f
// ff:
// str:
func ELF64(str []byte) uint64 {
	var (
		hash uint64
		x    uint64
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint64(str[i])
		if x = hash & 0xF000000000000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x + 1
		}
	}
	return hash
}
