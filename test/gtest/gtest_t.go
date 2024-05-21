// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtest

import (
	"testing"
)

// T is the testing unit case management object.
type T struct {
	*testing.T
}

// Assert checks `value` and `expect` EQUAL.

// ff:
// expect:
// value:
func (t *T) Assert(value, expect interface{}) {
	Assert(value, expect)
}

// AssertEQ checks `value` and `expect` EQUAL, including their TYPES.

// ff:
// expect:
// value:
func (t *T) AssertEQ(value, expect interface{}) {
	AssertEQ(value, expect)
}

// AssertNE checks `value` and `expect` NOT EQUAL.

// ff:
// expect:
// value:
func (t *T) AssertNE(value, expect interface{}) {
	AssertNE(value, expect)
}

// AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.

// ff:
// expect:
// value:
func (t *T) AssertNQ(value, expect interface{}) {
	AssertNQ(value, expect)
}

// AssertGT checks `value` is GREATER THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGT,
// others are invalid.

// ff:
// expect:
// value:
func (t *T) AssertGT(value, expect interface{}) {
	AssertGT(value, expect)
}

// AssertGE checks `value` is GREATER OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGTE,
// others are invalid.

// ff:
// expect:
// value:
func (t *T) AssertGE(value, expect interface{}) {
	AssertGE(value, expect)
}

// AssertLT checks `value` is LESS EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLT,
// others are invalid.

// ff:
// expect:
// value:
func (t *T) AssertLT(value, expect interface{}) {
	AssertLT(value, expect)
}

// AssertLE checks `value` is LESS OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLTE,
// others are invalid.

// ff:
// expect:
// value:
func (t *T) AssertLE(value, expect interface{}) {
	AssertLE(value, expect)
}

// AssertIN checks `value` is IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.

// ff:
// expect:
// value:
func (t *T) AssertIN(value, expect interface{}) {
	AssertIN(value, expect)
}

// AssertNI checks `value` is NOT IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.

// ff:
// expect:
// value:
func (t *T) AssertNI(value, expect interface{}) {
	AssertNI(value, expect)
}

// AssertNil asserts `value` is nil.

// ff:
// value:
func (t *T) AssertNil(value interface{}) {
	AssertNil(value)
}

// Error panics with given `message`.

// ff:
// message:
func (t *T) Error(message ...interface{}) {
	Error(message...)
}

// Fatal prints `message` to stderr and exit the process.

// ff:
// message:
func (t *T) Fatal(message ...interface{}) {
	Fatal(message...)
}
