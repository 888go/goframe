// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gbinary provides useful API for handling binary/bytes data.
//
// Note that package gbinary encodes the data using LittleEndian in default.
package gbinary


// ff:
// values:
func Encode(values ...interface{}) []byte {
	return LeEncode(values...)
}


// ff:
// values:
// length:
func EncodeByLength(length int, values ...interface{}) []byte {
	return LeEncodeByLength(length, values...)
}


// ff:
// values:
// b:
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
