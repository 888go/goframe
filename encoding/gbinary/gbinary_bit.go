// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gbinary

// 注意：这是一个实验性功能！. md5:f06e54ac5db3ae47

// Bit Binary bit (0 | 1)
type Bit int8

// EncodeBits 对bits进行编码，返回默认的编码结果. md5:0e00e3ef37578d60
func EncodeBits(bits []Bit, i int, l int) []Bit {
	return EncodeBitsWithUint(bits, uint(i), l)
}

// 将Uint与bits数组进行位运算合并，并占用length个比特位
// （注意：二进制的0和1位存储在uis数组中） md5:d607be5613f78556
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit {
	a := make([]Bit, l)
	for i := l - 1; i >= 0; i-- {
		a[i] = Bit(ui & 1)
		ui >>= 1
	}
	if bits != nil {
		return append(bits, a...)
	}
	return a
}

// EncodeBitsToBytes 将位编码为字节
// 将位转换为[]byte，从左到右进行编码，并在末尾添加不足1字节的0。 md5:0f20302eaa5194b7
func EncodeBitsToBytes(bits []Bit) []byte {
	if len(bits)%8 != 0 {
		for i := 0; i < len(bits)%8; i++ {
			bits = append(bits, 0)
		}
	}
	b := make([]byte, 0)
	for i := 0; i < len(bits); i += 8 {
		b = append(b, byte(DecodeBitsToUint(bits[i:i+8])))
	}
	return b
}

// DecodeBits 解码位为整数
// 转换为整数 md5:120b81a9c7484ad9
func DecodeBits(bits []Bit) int {
	v := 0
	for _, i := range bits {
		v = v<<1 | int(i)
	}
	return v
}

// DecodeBitsToUint 解码位并转换为无符号整数. md5:6f9a7b3d400472c8
func DecodeBitsToUint(bits []Bit) uint {
	v := uint(0)
	for _, i := range bits {
		v = v<<1 | uint(i)
	}
	return v
}

// DecodeBytesToBits 将字节切片解析为字符数组 []uint8. md5:bf0570931fd34b7c
func DecodeBytesToBits(bs []byte) []Bit {
	bits := make([]Bit, 0)
	for _, b := range bs {
		bits = EncodeBitsWithUint(bits, uint(b), 8)
	}
	return bits
}
