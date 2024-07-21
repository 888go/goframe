// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

	// Package gbinary 提供了处理二进制/字节数据的有用 API。
	// 
	// 注意，gbinary 包默认使用 LittleEndian 方式编码数据。
	// md5:20767570ab1713ca
package gbinary

// ff:
// values:
func Encode(values ...interface{}) []byte {
	return LeEncode(values...)
}

// ff:
// length:
// values:
func EncodeByLength(length int, values ...interface{}) []byte {
	return LeEncodeByLength(length, values...)
}

// ff:
// b:
// values:
func Decode(b []byte, values ...interface{}) error {
	return LeDecode(b, values...)
}

// ff:
// s:
func EncodeString(s string) []byte {
	return LeEncodeString(s)
}

// ff:
// b:
func DecodeToString(b []byte) string {
	return LeDecodeToString(b)
}

// ff:
// b:
func EncodeBool(b bool) []byte {
	return LeEncodeBool(b)
}

// ff:
// i:
func EncodeInt(i int) []byte {
	return LeEncodeInt(i)
}

// ff:
// i:
func EncodeUint(i uint) []byte {
	return LeEncodeUint(i)
}

// ff:
// i:
func EncodeInt8(i int8) []byte {
	return LeEncodeInt8(i)
}

// ff:
// i:
func EncodeUint8(i uint8) []byte {
	return LeEncodeUint8(i)
}

// ff:
// i:
func EncodeInt16(i int16) []byte {
	return LeEncodeInt16(i)
}

// ff:
// i:
func EncodeUint16(i uint16) []byte {
	return LeEncodeUint16(i)
}

// ff:
// i:
func EncodeInt32(i int32) []byte {
	return LeEncodeInt32(i)
}

// ff:
// i:
func EncodeUint32(i uint32) []byte {
	return LeEncodeUint32(i)
}

// ff:
// i:
func EncodeInt64(i int64) []byte {
	return LeEncodeInt64(i)
}

// ff:
// i:
func EncodeUint64(i uint64) []byte {
	return LeEncodeUint64(i)
}

// ff:
// f:
func EncodeFloat32(f float32) []byte {
	return LeEncodeFloat32(f)
}

// ff:
// f:
func EncodeFloat64(f float64) []byte {
	return LeEncodeFloat64(f)
}

// ff:
// b:
func DecodeToInt(b []byte) int {
	return LeDecodeToInt(b)
}

// ff:
// b:
func DecodeToUint(b []byte) uint {
	return LeDecodeToUint(b)
}

// ff:
// b:
func DecodeToBool(b []byte) bool {
	return LeDecodeToBool(b)
}

// ff:
// b:
func DecodeToInt8(b []byte) int8 {
	return LeDecodeToInt8(b)
}

// ff:
// b:
func DecodeToUint8(b []byte) uint8 {
	return LeDecodeToUint8(b)
}

// ff:
// b:
func DecodeToInt16(b []byte) int16 {
	return LeDecodeToInt16(b)
}

// ff:
// b:
func DecodeToUint16(b []byte) uint16 {
	return LeDecodeToUint16(b)
}

// ff:
// b:
func DecodeToInt32(b []byte) int32 {
	return LeDecodeToInt32(b)
}

// ff:
// b:
func DecodeToUint32(b []byte) uint32 {
	return LeDecodeToUint32(b)
}

// ff:
// b:
func DecodeToInt64(b []byte) int64 {
	return LeDecodeToInt64(b)
}

// ff:
// b:
func DecodeToUint64(b []byte) uint64 {
	return LeDecodeToUint64(b)
}

// ff:
// b:
func DecodeToFloat32(b []byte) float32 {
	return LeDecodeToFloat32(b)
}

// ff:
// b:
func DecodeToFloat64(b []byte) float64 {
	return LeDecodeToFloat64(b)
}
