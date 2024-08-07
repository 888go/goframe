package gdebug_test

import (
	"fmt"
	"testing"

	"github.com/888go/goframe/debug/gdebug"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

func Test_CallerPackage(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gdebug.CallerPackage(), "github.com/888go/goframe/test/gtest")
	})
}

func Test_CallerFunction(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gdebug.CallerFunction(), "C")
	})
}

func Test_CallerFilePath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含(gdebug.CallerFilePath(), "gtest_util.go"), true)
	})
}

func Test_CallerDirectory(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含(gdebug.CallerDirectory(), "gtest"), true)
	})
}

func Test_CallerFileLine(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		fmt.Println(gdebug.CallerFileLine())
		t.Assert(gstr.X是否包含(gdebug.CallerFileLine(), "gtest_util.go:36"), true)
	})
}

func Test_CallerFileLineShort(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含(gdebug.CallerFileLineShort(), "gtest_util.go:36"), true)
	})
}

func Test_FuncPath(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gdebug.FuncPath(Test_FuncPath), "github.com/888go/goframe/debug/gdebug_test.Test_FuncPath")
	})
}

func Test_FuncName(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gdebug.FuncName(Test_FuncName), "gdebug_test.Test_FuncName")
	})
}

func Test_PrintStack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		gdebug.PrintStack()
	})
}

func Test_GoroutineId(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGT(gdebug.GoroutineId(), 0)
	})
}

func Test_Stack(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含(gdebug.Stack(), "gtest_util.go:36"), true)
	})
}

func Test_StackWithFilter(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(gstr.X是否包含(gdebug.StackWithFilter([]string{"github.com"}), "gtest_util.go:36"), true)
	})
}

func Test_BinVersion(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGT(len(gdebug.BinVersion()), 0)
	})
}

func Test_BinVersionMd5(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.AssertGT(len(gdebug.BinVersionMd5()), 0)
	})
}
