// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtype

import (
	"bytes"
	"encoding/base64"
	"sync/atomic"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

// Bytes is a struct for concurrent-safe operation for type []byte.
type Bytes struct {
	value atomic.Value
}

// NewBytes creates and returns a concurrent-safe object for []byte type,
// with given initial value `value`.

// ff:
// value:
func NewBytes(value ...[]byte) *Bytes {
	t := &Bytes{}
	if len(value) > 0 {
		t.value.Store(value[0])
	}
	return t
}

// Clone clones and returns a new shallow copy object for []byte type.

// ff:
func (v *Bytes) Clone() *Bytes {
	return NewBytes(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.
// Note: The parameter `value` cannot be nil.

// ff:设置值
// old:
// value:
func (v *Bytes) Set(value []byte) (old []byte) {
	old = v.Val()
	v.value.Store(value)
	return
}

// Val atomically loads and returns t.value.

// ff:取值
func (v *Bytes) Val() []byte {
	if s := v.value.Load(); s != nil {
		return s.([]byte)
	}
	return nil
}

// String implements String interface for string printing.

// ff:
func (v *Bytes) String() string {
	return string(v.Val())
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.

// ff:
func (v Bytes) MarshalJSON() ([]byte, error) {
	val := v.Val()
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(val)))
	base64.StdEncoding.Encode(dst, val)
	return []byte(`"` + string(dst) + `"`), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

// ff:
// b:
func (v *Bytes) UnmarshalJSON(b []byte) error {
	var (
		src    = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
		n, err = base64.StdEncoding.Decode(src, bytes.Trim(b, `"`))
	)
	if err != nil {
		err = gerror.Wrap(err, `base64.StdEncoding.Decode failed`)
		return err
	}
	v.Set(src[:n])
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.

// ff:
// value:
func (v *Bytes) UnmarshalValue(value interface{}) error {
	v.Set(gconv.Bytes(value))
	return nil
}

// DeepCopy implements interface for deep copy of current type.

// ff:
func (v *Bytes) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	oldBytes := v.Val()
	newBytes := make([]byte, len(oldBytes))
	copy(newBytes, oldBytes)
	return NewBytes(newBytes)
}
