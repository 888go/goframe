// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package utils_test

import (
	"io"
	"testing"
	
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/test/gtest"
)

func Test_ReadCloser(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			r    []byte
			body = utils.NewReadCloser([]byte{1, 2, 3, 4}, false)
		)
		r, _ = io.ReadAll(body)
		t.Assert(r, []byte{1, 2, 3, 4})
		r, _ = io.ReadAll(body)
		t.Assert(r, []byte{})
	})
	单元测试类.C(t, func(t *单元测试类.T) {
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
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(utils.RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`), `abac101`)
		t.Assert(utils.RemoveSymbols(`-a-b我._a c1!@#$%^&*是()_+:帅";'.,哥'01`), `ab我ac1是帅哥01`)
	})
}
