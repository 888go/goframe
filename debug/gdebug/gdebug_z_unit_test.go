package gdebug_test

import (
	"testing"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func Test_CallerPackage(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(gdebug.CallerPackage(), "github.com/888go/goframe/test/gtest")
	})
}

func Test_CallerFunction(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(gdebug.CallerFunction(), "C")
	})
}

func Test_CallerFilePath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X是否包含(gdebug.CallerFilePath(), "gtest_util.go"), true)
	})
}

func Test_CallerDirectory(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X是否包含(gdebug.CallerDirectory(), "gtest"), true)
	})
}

func Test_CallerFileLine(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X是否包含(gdebug.CallerFileLine(), "gtest_util.go:35"), true)
	})
}

func Test_CallerFileLineShort(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X是否包含(gdebug.CallerFileLineShort(), "gtest_util.go:35"), true)
	})
}

func Test_FuncPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(gdebug.FuncPath(Test_FuncPath), "github.com/888go/goframe/debug/gdebug_test.Test_FuncPath")
	})
}

func Test_FuncName(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(gdebug.FuncName(Test_FuncName), "gdebug_test.Test_FuncName")
	})
}

func Test_PrintStack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		gdebug.PrintStack()
	})
}

func Test_GoroutineId(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGT(gdebug.GoroutineId(), 0)
	})
}

func Test_Stack(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X是否包含(gdebug.Stack(), "gtest_util.go:35"), true)
	})
}

func Test_StackWithFilter(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.Assert(文本类.X是否包含(gdebug.StackWithFilter([]string{"github.com"}), "gtest_util.go:35"), true)
	})
}

func Test_BinVersion(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGT(len(gdebug.BinVersion()), 0)
	})
}

func Test_BinVersionMd5(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		t.AssertGT(len(gdebug.BinVersionMd5()), 0)
	})
}
