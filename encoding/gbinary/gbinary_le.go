// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 字节集类

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"

	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
)

// LeEncode 使用小端字节序编码一个或多个 `values` 成字节。
// 它通过类型断言检查 `values` 中每个值的类型，并在内部调用相应的转换函数来进行字节转换。
//
// 它支持常见的变量类型断言，最终使用 fmt.Sprintf 将值转换为字符串，再转换为字节。
// md5:4210ce2214f05499
func LeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}
		switch value := values[i].(type) {
		case int:
			buf.Write(LeEncodeInt(value))
		case int8:
			buf.Write(LeEncodeInt8(value))
		case int16:
			buf.Write(LeEncodeInt16(value))
		case int32:
			buf.Write(LeEncodeInt32(value))
		case int64:
			buf.Write(LeEncodeInt64(value))
		case uint:
			buf.Write(LeEncodeUint(value))
		case uint8:
			buf.Write(LeEncodeUint8(value))
		case uint16:
			buf.Write(LeEncodeUint16(value))
		case uint32:
			buf.Write(LeEncodeUint32(value))
		case uint64:
			buf.Write(LeEncodeUint64(value))
		case bool:
			buf.Write(LeEncodeBool(value))
		case string:
			buf.Write(LeEncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(LeEncodeFloat32(value))
		case float64:
			buf.Write(LeEncodeFloat64(value))

		default:
			if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
				buf.Write(LeEncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

func LeEncodeByLength(length int, values ...interface{}) []byte {
	b := LeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

func LeDecode(b []byte, values ...interface{}) error {
	var (
		err error
		buf = bytes.NewBuffer(b)
	)
	for i := 0; i < len(values); i++ {
		if err = binary.Read(buf, binary.LittleEndian, values[i]); err != nil {
			err = gerror.X多层错误(err, `binary.Read failed`)
			return err
		}
	}
	return nil
}

func LeEncodeString(s string) []byte {
	return []byte(s)
}

func LeDecodeToString(b []byte) string {
	return string(b)
}

func LeEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func LeEncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return EncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return EncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return EncodeInt32(int32(i))
	} else {
		return EncodeInt64(int64(i))
	}
}

func LeEncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return EncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return EncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return EncodeUint32(uint32(i))
	} else {
		return EncodeUint64(uint64(i))
	}
}

func LeEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func LeEncodeUint8(i uint8) []byte {
	return []byte{i}
}

func LeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(i))
	return b
}

func LeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return b
}

func LeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}

func LeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func LeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

func LeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func LeEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func LeEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}

func LeDecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(LeDecodeToUint8(b))
	} else if len(b) < 3 {
		return int(LeDecodeToUint16(b))
	} else if len(b) < 5 {
		return int(LeDecodeToUint32(b))
	} else {
		return int(LeDecodeToUint64(b))
	}
}

func LeDecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(LeDecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(LeDecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(LeDecodeToUint32(b))
	} else {
		return uint(LeDecodeToUint64(b))
	}
}

func LeDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Equal(b, make([]byte, len(b))) {
		return false
	}
	return true
}

func LeDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

func LeDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

func LeDecodeToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(LeFillUpSize(b, 2)))
}

func LeDecodeToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(LeFillUpSize(b, 2))
}

func LeDecodeToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(LeFillUpSize(b, 4)))
}

func LeDecodeToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(LeFillUpSize(b, 4))
}

func LeDecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

func LeDecodeToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(LeFillUpSize(b, 8))
}

func LeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(LeFillUpSize(b, 4)))
}

func LeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(LeFillUpSize(b, 8)))
}

// LeFillUpSize 使用LittleEndian字节序将字节`b`填充到给定长度`l`。
//
// 请注意，它通过复制原始字节切片创建一个新的字节切片，以避免修改原始参数字节。
// md5:9ac3ba6f3d5c0177
func LeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c, b)
	return c
}
