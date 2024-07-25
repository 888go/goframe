// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package utils_test

import (
	"io"
	"reflect"
	"testing"
	"unsafe"

	"github.com/gogf/gf/v2/internal/utils"
	"github.com/gogf/gf/v2/test/gtest"
)

func Test_ReadCloser(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			n    int
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{})
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			r    []byte
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		r, _ = io.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = io.ReadAll(body)
		t.Assert(r, []byte{})
	})
	gtest.C(t, func(t *gtest.T) {
		var (
			n    int
			r    []byte
			b    = make([]byte, 3)
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, true)
		)
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{1, 2, 3})
		n, _ = body.Read(b)
		t.Assert(b[:n], []byte{4})

		r, _ = io.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = io.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
	})
}

func Test_RemoveSymbols(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(utils.RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`), `abac101`)
		t.Assert(utils.RemoveSymbols(`-a-b我._a c1!@#$%^&*是()_+:帅";'.,哥'01`), `ab我ac1是帅哥01`)
	})
}

func Test_CanCallIsNil(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			iValue         = "gf"
			iChan          = make(chan struct{})
			iFunc          = func() {}
			iMap           = map[string]struct{}{}
			iPtr           = &iValue
			iSlice         = make([]struct{}, 0)
			iUnsafePointer = unsafe.Pointer(&iValue)
		)

		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iValue)), false)
		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iChan)), true)
		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iFunc)), true)
		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iMap)), true)
		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iPtr)), true)
		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iSlice)), true)
		t.Assert(utils.CanCallIsNil(reflect.ValueOf(iUnsafePointer)), true)
	})
}
