// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 字节集类

// 注意：这是一个实验性功能！

// 二进制位 (0 | 1)
type Bit int8

// EncodeBits 对比特进行编码，并返回编码后的比特，默认采用默认编码方式
func EncodeBits(bits []Bit, i int, l int) []Bit {
	return EncodeBitsWithUint(bits, uint(i), l)
}

// EncodeBitsWithUint . 将ui的按位与bits数组进行合并，并占用length位
// （注意：二进制的0 | 1数字存储在uis数组中）
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

// EncodeBitsToBytes . 将位编码为字节
// 将位转换为[] byte，从左到右进行编码，并在末尾不足1字节时添加0填充。
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

// DecodeBits .将位解码为整数
// 解析为整数
func DecodeBits(bits []Bit) int {
	v := 0
	for _, i := range bits {
		v = v<<1 | int(i)
	}
	return v
}

// DecodeBitsToUint 解码bits并转换为uint类型
func DecodeBitsToUint(bits []Bit) uint {
	v := uint(0)
	for _, i := range bits {
		v = v<<1 | uint(i)
	}
	return v
}

// DecodeBytesToBits .将字节切片解析为字符数组（uint8切片）
func DecodeBytesToBits(bs []byte) []Bit {
	bits := make([]Bit, 0)
	for _, b := range bs {
		bits = EncodeBitsWithUint(bits, uint(b), 8)
	}
	return bits
}
