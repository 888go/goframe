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

// Uint32 is a struct for concurrent-safe operation for type uint32.
type Uint32 struct {
	value uint32
}

// NewUint32 creates and returns a concurrent-safe object for uint32 type,
// with given initial value `value`.

// ff:
// value:
func NewUint32(value ...uint32) *Uint32 {
	if len(value) > 0 {
		return &Uint32{
			value: value[0],
		}
	}
	return &Uint32{}
}

// Clone clones and returns a new concurrent-safe object for uint32 type.

// ff:
func (v *Uint32) Clone() *Uint32 {
	return NewUint32(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.

// ff:设置值
// old:
// value:
func (v *Uint32) Set(value uint32) (old uint32) {
	return atomic.SwapUint32(&v.value, value)
}

// Val atomically loads and returns t.value.

// ff:取值
func (v *Uint32) Val() uint32 {
	return atomic.LoadUint32(&v.value)
}

// Add atomically adds `delta` to t.value and returns the new value.

// ff:
// new:
// delta:
func (v *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&v.value, delta)
}

// Cas executes the compare-and-swap operation for value.

// ff:
// swapped:
// new:
// old:
func (v *Uint32) Cas(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&v.value, old, new)
}

// String implements String interface for string printing.

// ff:
func (v *Uint32) String() string {
	return strconv.FormatUint(uint64(v.Val()), 10)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.

// ff:
func (v Uint32) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.Val()), 10)), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

// ff:
// b:
func (v *Uint32) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Uint32(string(b)))
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.

// ff:
// value:
func (v *Uint32) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Uint32(value))
	return nil
}

// DeepCopy implements interface for deep copy of current type.

// ff:
func (v *Uint32) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewUint32(v.Val())
}
