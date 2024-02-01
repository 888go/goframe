// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghash

// DJB 实现了适用于32位的经典的DJB哈希算法。
func DJB(str []byte) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint32(str[i])
	}
	return hash
}

// DJB64 实现了适用于 64 位的经典 DJB 哈希算法。
func DJB64(str []byte) uint64 {
	var hash uint64 = 5381
	for i := 0; i < len(str); i++ {
		hash += (hash << 5) + uint64(str[i])
	}
	return hash
}
