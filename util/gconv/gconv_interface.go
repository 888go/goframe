// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import "github.com/gogf/gf/v2/os/gtime"

// iVal is used for type assert api for String().
type iVal interface {
	Val() interface{} //qm:取值 cz:Val() interface{} yx:true     
}

// iString is used for type assert api for String().
type iString interface {
	String() string
}

// iBool is used for type assert api for Bool().
type iBool interface {
	Bool() bool //qm:取布尔 cz:Bool() bool yx:true     
}

// iInt64 is used for type assert api for Int64().
type iInt64 interface {
	Int64() int64 //qm:取整数64位 cz:Int64() int64 yx:true     
}

// iUint64 is used for type assert api for Uint64().
type iUint64 interface {
	Uint64() uint64 //qm:取正整数64位 cz:Uint64() uint64 yx:true     
}

// iFloat32 is used for type assert api for Float32().
type iFloat32 interface {
	Float32() float32 //qm:取小数32位 cz:Float32() float32 yx:true     
}

// iFloat64 is used for type assert api for Float64().
type iFloat64 interface {
	Float64() float64 //qm:取小数64位 cz:Float64() float64 yx:true     
}

// iError is used for type assert api for Error().
type iError interface {
	Error() string
}

// iBytes is used for type assert api for Bytes().
type iBytes interface {
	Bytes() []byte //qm:取字节集 cz:Bytes() []byte yx:true     
}

// iInterface is used for type assert api for Interface().
type iInterface interface {
	Interface() interface{}
}

// iInterfaces is used for type assert api for Interfaces().
type iInterfaces interface {
	Interfaces() []interface{} //qm:取any数组 cz:Interfaces() []interface{} yx:true     
}

// iFloats is used for type assert api for Floats().
type iFloats interface {
	Floats() []float64 //qm:取小数数组 cz:Floats() []float64 yx:true     
}

// iInts is used for type assert api for Ints().
type iInts interface {
	Ints() []int //qm:取整数数组 cz:Ints() []int yx:true     
}

// iStrings is used for type assert api for Strings().
type iStrings interface {
	Strings() []string //qm:取文本数组 cz:Strings() []string yx:true     
}

// iUints is used for type assert api for Uints().
type iUints interface {
	Uints() []uint //qm:取正整数数组 cz:Uints() []uint yx:true     
}

// iMapStrAny is the interface support for converting struct parameter to map.
type iMapStrAny interface {
	MapStrAny() map[string]interface{} //qm:取MapStrAny cz:MapStrAny() map[string]interface{} yx:true     
}

// iUnmarshalValue is the interface for custom defined types customizing value assignment.
// Note that only pointer can implement interface iUnmarshalValue.
type iUnmarshalValue interface {
	UnmarshalValue(interface{}) error
}

// iUnmarshalText is the interface for custom defined types customizing value assignment.
// Note that only pointer can implement interface iUnmarshalText.
type iUnmarshalText interface {
	UnmarshalText(text []byte) error
}

// iUnmarshalText is the interface for custom defined types customizing value assignment.
// Note that only pointer can implement interface iUnmarshalJSON.
type iUnmarshalJSON interface {
	UnmarshalJSON(b []byte) error
}

// iSet is the interface for custom value assignment.
type iSet interface {
	Set(value interface{}) (old interface{}) //qm:设置值 cz:Set( yx:true     
}

// iGTime is the interface for gtime.Time converting.
type iGTime interface {
	GTime(format ...string) *gtime.Time //qm:取gtime时间类 cz:GTime( yx:true     
}
