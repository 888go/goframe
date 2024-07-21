// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package ghash

// PJW实现了经典的32位PJW哈希算法。 md5:3822bc8ce34ff2e9
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

// PJW64 实现了经典的 PJW 算法，用于生成64位的哈希值。 md5:5c34a99d5f76f106
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
