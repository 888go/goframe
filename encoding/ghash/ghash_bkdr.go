// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 哈希类

// BKDR 实现了适用于32位的经典的BKDR哈希算法。
func BKDR(str []byte) uint32 {
	var (
		seed uint32 = 131 // 以下是按规律生成一系列数字的注释：
// 31, 131, 1313, 13131, 131313 等等...
// 该代码段没有提供具体的Go语言代码，但根据注释内容推测，这可能是一个按照某种规律（如在每个数中间插入1和3）生成序列的说明。
		hash uint32 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint32(str[i])
	}
	return hash
}

// BKDR64 实现了适用于64位的经典的 BKDR 哈希算法。
func BKDR64(str []byte) uint64 {
	var (
		seed uint64 = 131 // 以下是按规律生成一系列数字的注释：
// 31, 131, 1313, 13131, 131313 等等...
// 该代码段没有提供具体的Go语言代码，但根据注释内容推测，这可能是一个按照某种规律（如在每个数中间插入1和3）生成序列的说明。
		hash uint64 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return hash
}
