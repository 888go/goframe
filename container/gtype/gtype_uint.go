// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtype

import (
	"strconv"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// Uint is a struct for concurrent-safe operation for type uint.
type Uint struct {
	value uint64
}

// NewUint creates and returns a concurrent-safe object for uint type,
// with given initial value `value`.

// ff:
// value:
func NewUint(value ...uint) *Uint {
	if len(value) > 0 {
		return &Uint{
			value: uint64(value[0]),
		}
	}
	return &Uint{}
}

// Clone clones and returns a new concurrent-safe object for uint type.

// ff:
func (v *Uint) Clone() *Uint {
	return NewUint(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.

// ff:设置值
// yx:true
// old:
// value:
func (v *Uint) Set(value uint) (old uint) {
	return uint(atomic.SwapUint64(&v.value, uint64(value)))
}

// Val atomically loads and returns t.value.

// ff:取值
// yx:true
func (v *Uint) Val() uint {
	return uint(atomic.LoadUint64(&v.value))
}

// Add atomically adds `delta` to t.value and returns the new value.

// ff:
// new:
// delta:
func (v *Uint) Add(delta uint) (new uint) {
	return uint(atomic.AddUint64(&v.value, uint64(delta)))
}

// Cas executes the compare-and-swap operation for value.

// ff:
// swapped:
// new:
// old:
func (v *Uint) Cas(old, new uint) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, uint64(old), uint64(new))
}

// String implements String interface for string printing.

// ff:
func (v *Uint) String() string {
	return strconv.FormatUint(uint64(v.Val()), 10)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.

// ff:
func (v Uint) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.Val()), 10)), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

// ff:
// b:
func (v *Uint) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Uint(string(b)))
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.

// ff:
// value:
func (v *Uint) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Uint(value))
	return nil
}

// DeepCopy implements interface for deep copy of current type.

// ff:
func (v *Uint) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewUint(v.Val())
}
