// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtype

import (
	"math"
	"strconv"
	"sync/atomic"

	"github.com/gogf/gf/v2/util/gconv"
)

// Float64 is a struct for concurrent-safe operation for type float64.
type Float64 struct {
	value uint64
}

// NewFloat64 creates and returns a concurrent-safe object for float64 type,
// with given initial value `value`.
// ff:
// value:
func NewFloat64(value ...float64) *Float64 {
	if len(value) > 0 {
		return &Float64{
			value: math.Float64bits(value[0]),
		}
	}
	return &Float64{}
}

// Clone clones and returns a new concurrent-safe object for float64 type.
// ff:
// v:
func (v *Float64) Clone() *Float64 {
	return NewFloat64(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.
// yx:true
// ff:设置值
// v:
// value:
// old:
func (v *Float64) Set(value float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&v.value, math.Float64bits(value)))
}

// Val atomically loads and returns t.value.
// yx:true
// ff:取值
// v:
func (v *Float64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(&v.value))
}

// Add atomically adds `delta` to t.value and returns the new value.
// ff:
// v:
// delta:
// new:
func (v *Float64) Add(delta float64) (new float64) {
	for {
		old := math.Float64frombits(v.value)
		new = old + delta
		if atomic.CompareAndSwapUint64(
			&v.value,
			math.Float64bits(old),
			math.Float64bits(new),
		) {
			break
		}
	}
	return
}

// Cas executes the compare-and-swap operation for value.
// ff:
// v:
// old:
// new:
// swapped:
func (v *Float64) Cas(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&v.value, math.Float64bits(old), math.Float64bits(new))
}

// String implements String interface for string printing.
// ff:
// v:
func (v *Float64) String() string {
	return strconv.FormatFloat(v.Val(), 'g', -1, 64)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
// ff:
// v:
func (v Float64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(v.Val(), 'g', -1, 64)), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
// ff:
// v:
// b:
func (v *Float64) UnmarshalJSON(b []byte) error {
	v.Set(gconv.Float64(string(b)))
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.
// ff:
// v:
// value:
func (v *Float64) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Float64(value))
	return nil
}

// DeepCopy implements interface for deep copy of current type.
// ff:
// v:
func (v *Float64) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewFloat64(v.Val())
}
