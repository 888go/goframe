// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 哈希类

// SDBM 实现了适用于 32 位的经典的 SDBM 哈希算法。
func SDBM(str []byte) uint32 {
	var hash uint32
	for i := 0; i < len(str); i++ {
		// 等价于：hash = 65599 * hash + uint32(str[i])
		hash = uint32(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// SDBM64 实现了适用于 64 位的经典的 SDBM 哈希算法。
func SDBM64(str []byte) uint64 {
	var hash uint64
	for i := 0; i < len(str); i++ {
		// equivalent to: hash = 65599*hash + uint32(str[i])
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
