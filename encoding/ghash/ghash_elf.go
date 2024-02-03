// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghash

// ELF 实现了适用于 32 位的经典的 ELF 哈希算法。
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

// ELF64 实现了适用于 64 位的经典的 ELF 哈希算法。
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
