// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 哈希类

// RS 实现了适用于 32 位的经典 RS 哈希算法。
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

// RS64 实现了适用于 64 位的经典的 RS 哈希算法。
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
