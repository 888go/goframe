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

// BeEncode 使用大端字节序将一个或多个`values`编码为字节。它通过检查`values`中每个值的类型，并调用相应的转换函数来进行内部字节转换。
// 
// 它支持常见的变量类型断言，最后如果无法确定类型，会使用`fmt.Sprintf`将值转换为字符串，然后再转换为字节。
// md5:9b191604817267a1
func BeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}

		switch value := values[i].(type) {
		case int:
			buf.Write(BeEncodeInt(value))
		case int8:
			buf.Write(BeEncodeInt8(value))
		case int16:
			buf.Write(BeEncodeInt16(value))
		case int32:
			buf.Write(BeEncodeInt32(value))
		case int64:
			buf.Write(BeEncodeInt64(value))
		case uint:
			buf.Write(BeEncodeUint(value))
		case uint8:
			buf.Write(BeEncodeUint8(value))
		case uint16:
			buf.Write(BeEncodeUint16(value))
		case uint32:
			buf.Write(BeEncodeUint32(value))
		case uint64:
			buf.Write(BeEncodeUint64(value))
		case bool:
			buf.Write(BeEncodeBool(value))
		case string:
			buf.Write(BeEncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(BeEncodeFloat32(value))
		case float64:
			buf.Write(BeEncodeFloat64(value))
		default:
			if err := binary.Write(buf, binary.BigEndian, value); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
				buf.Write(BeEncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

func BeEncodeByLength(length int, values ...interface{}) []byte {
	b := BeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

func BeDecode(b []byte, values ...interface{}) error {
	var (
		err error
		buf = bytes.NewBuffer(b)
	)
	for i := 0; i < len(values); i++ {
		if err = binary.Read(buf, binary.BigEndian, values[i]); err != nil {
			err = gerror.X多层错误(err, `binary.Read failed`)
			return err
		}
	}
	return nil
}

func BeEncodeString(s string) []byte {
	return []byte(s)
}

func BeDecodeToString(b []byte) string {
	return string(b)
}

func BeEncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

func BeEncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return BeEncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return BeEncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return BeEncodeInt32(int32(i))
	} else {
		return BeEncodeInt64(int64(i))
	}
}

func BeEncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return BeEncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return BeEncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return BeEncodeUint32(uint32(i))
	} else {
		return BeEncodeUint64(uint64(i))
	}
}

func BeEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func BeEncodeUint8(i uint8) []byte {
	return []byte{i}
}

func BeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

func BeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func BeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func BeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func BeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func BeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func BeEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, bits)
	return b
}

func BeEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, bits)
	return b
}

func BeDecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(BeDecodeToUint8(b))
	} else if len(b) < 3 {
		return int(BeDecodeToUint16(b))
	} else if len(b) < 5 {
		return int(BeDecodeToUint32(b))
	} else {
		return int(BeDecodeToUint64(b))
	}
}

func BeDecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(BeDecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(BeDecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(BeDecodeToUint32(b))
	} else {
		return uint(BeDecodeToUint64(b))
	}
}

func BeDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Equal(b, make([]byte, len(b))) {
		return false
	}
	return true
}

func BeDecodeToInt8(b []byte) int8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return int8(b[0])
}

func BeDecodeToUint8(b []byte) uint8 {
	if len(b) == 0 {
		panic(`empty slice given`)
	}
	return b[0]
}

func BeDecodeToInt16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(BeFillUpSize(b, 2)))
}

func BeDecodeToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(BeFillUpSize(b, 2))
}

func BeDecodeToInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(BeFillUpSize(b, 4)))
}

func BeDecodeToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(BeFillUpSize(b, 4))
}

func BeDecodeToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(BeFillUpSize(b, 8)))
}

func BeDecodeToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(BeFillUpSize(b, 8))
}

func BeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(BeFillUpSize(b, 4)))
}

func BeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(BeFillUpSize(b, 8)))
}

// BeFillUpSize 使用大端字节序填充字节切片 `b` 到给定长度 `l`。
// 
// 注意，它通过复制原始字节切片来创建一个新的字节切片，以避免修改原始参数字节。
// md5:ae17f54c414e3a97
func BeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c[l-len(b):], b)
	return c
}
