// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghash

// PJW 实现了适用于32位的经典的PJW哈希算法。
func PJW(str []byte) uint32 {
	var (
		BitsInUnsignedInt uint32 = 32 // 4 * 8
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint32 = (0xFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
		hash              uint32
		test              uint32
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint32(str[i])
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits + 1)
		}
	}
	return hash
}

// PJW64 实现了适用于 64 位的经典 PJW 哈希算法。
func PJW64(str []byte) uint64 {
	var (
		BitsInUnsignedInt uint64 = 32 // 4 * 8
		ThreeQuarters            = (BitsInUnsignedInt * 3) / 4
		OneEighth                = BitsInUnsignedInt / 8
		HighBits          uint64 = (0xFFFFFFFFFFFFFFFF) << (BitsInUnsignedInt - OneEighth)
		hash              uint64
		test              uint64
	)
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint64(str[i])
		if test = hash & HighBits; test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits + 1)
		}
	}
	return hash
}
