// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件类_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	
	"github.com/888go/goframe/debug/gdebug"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

func createTestFile(filename, content string) error {
	TempDir := testpath()
	err := os.WriteFile(TempDir+filename, []byte(content), 0666)
	return err
}

func delTestFiles(filenames string) {
	os.RemoveAll(testpath() + filenames)
}

func createDir(paths string) {
	TempDir := testpath()
	os.Mkdir(TempDir+paths, 0777)
}

func formatpaths(paths []string) []string {
	for k, v := range paths {
		paths[k] = filepath.ToSlash(v)
		paths[k] = strings.Replace(paths[k], "./", "/", 1)
	}

	return paths
}

func formatpath(paths string) string {
	paths = filepath.ToSlash(paths)
	paths = strings.Replace(paths, "./", "/", 1)
	return paths
}

func testpath() string {
	return 文本类.X过滤尾字符并含空白(os.TempDir(), "\\/")
}

func Test_GetContents(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {

		var (
			filepaths string = "/testfile_t1.txt"
		)
		createTestFile(filepaths, "my name is jroam")
		defer delTestFiles(filepaths)

		t.Assert(文件类.X读文本(testpath()+filepaths), "my name is jroam")
		t.Assert(文件类.X读文本(""), "")

	})
}

func Test_GetBinContents(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths1  = "/testfile_t1.txt"
			filepaths2  = testpath() + "/testfile_t1_no.txt"
			readcontent []byte
			str1        = "my name is jroam"
		)
		createTestFile(filepaths1, str1)
		defer delTestFiles(filepaths1)
		readcontent = 文件类.X读字节集(testpath() + filepaths1)
		t.Assert(readcontent, []byte(str1))

		readcontent = 文件类.X读字节集(filepaths2)
		t.Assert(string(readcontent), "")

		t.Assert(string(文件类.X读字节集(filepaths2)), "")

	})
}

func Test_Truncate(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths1 = "/testfile_GetContentsyyui.txt"
			err        error
			files      *os.File
		)
		createTestFile(filepaths1, "abcdefghijkmln")
		defer delTestFiles(filepaths1)
		err = 文件类.X截断(testpath()+filepaths1, 10)
		t.AssertNil(err)

		files, err = os.Open(testpath() + filepaths1)
		t.AssertNil(err)
		defer files.Close()
		fileinfo, err2 := files.Stat()
		t.Assert(err2, nil)
		t.Assert(fileinfo.Size(), 10)

		err = 文件类.X截断("", 10)
		t.AssertNE(err, nil)

	})
}

func Test_PutContents(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)

		err = 文件类.X写入文本(testpath()+filepaths, "test!")
		t.AssertNil(err)

		readcontent, err = os.ReadFile(testpath() + filepaths)
		t.AssertNil(err)
		t.Assert(string(readcontent), "test!")

		err = 文件类.X写入文本("", "test!")
		t.AssertNE(err, nil)

	})
}

func Test_PutContentsAppend(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)

		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)
		err = 文件类.X追加文本(testpath()+filepaths, "hello")
		t.AssertNil(err)

		readcontent, err = os.ReadFile(testpath() + filepaths)
		t.AssertNil(err)
		t.Assert(string(readcontent), "ahello")

		err = 文件类.X追加文本("", "hello")
		t.AssertNE(err, nil)

	})

}

func Test_PutBinContents(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)

		err = 文件类.X写入字节集(testpath()+filepaths, []byte("test!!"))
		t.AssertNil(err)

		readcontent, err = os.ReadFile(testpath() + filepaths)
		t.AssertNil(err)
		t.Assert(string(readcontent), "test!!")

		err = 文件类.X写入字节集("", []byte("test!!"))
		t.AssertNE(err, nil)

	})
}

func Test_PutBinContentsAppend(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths   = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "test!!")
		defer delTestFiles(filepaths)
		err = 文件类.X追加字节集(testpath()+filepaths, []byte("word"))
		t.AssertNil(err)

		readcontent, err = os.ReadFile(testpath() + filepaths)
		t.AssertNil(err)
		t.Assert(string(readcontent), "test!!word")

		err = 文件类.X追加字节集("", []byte("word"))
		t.AssertNE(err, nil)

	})
}

func Test_GetBinContentsByTwoOffsetsByPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths   = "/testfile_GetContents.txt"
			readcontent []byte
		)

		createTestFile(filepaths, "abcdefghijk")
		defer delTestFiles(filepaths)
		readcontent = 文件类.X取文件字节集按范围(testpath()+filepaths, 2, 5)

		t.Assert(string(readcontent), "cde")

		readcontent = 文件类.X取文件字节集按范围("", 2, 5)
		t.Assert(len(readcontent), 0)

	})

}

func Test_GetNextCharOffsetByPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			filepaths  = "/testfile_GetContents.txt"
			localindex int64
		)
		createTestFile(filepaths, "abcdefghijk")
		defer delTestFiles(filepaths)
		localindex = 文件类.X取文件字符偏移位置(testpath()+filepaths, 'd', 1)
		t.Assert(localindex, 3)

		localindex = 文件类.X取文件字符偏移位置("", 'd', 1)
		t.Assert(localindex, -1)

	})
}

func Test_GetNextCharOffset(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			localindex int64
		)
		reader := strings.NewReader("helloword")

		localindex = 文件类.X取字符偏移位置(reader, 'w', 1)
		t.Assert(localindex, 5)

		localindex = 文件类.X取字符偏移位置(reader, 'j', 1)
		t.Assert(localindex, -1)

	})
}

func Test_GetBinContentsByTwoOffsets(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			reads []byte
		)
		reader := strings.NewReader("helloword")

		reads = 文件类.X取字节集按范围(reader, 1, 3)
		t.Assert(string(reads), "el")

		reads = 文件类.X取字节集按范围(reader, 10, 30)
		t.Assert(string(reads), "")

	})
}

func Test_GetBinContentsTilChar(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			reads  []byte
			indexs int64
		)
		reader := strings.NewReader("helloword")

		reads, _ = 文件类.X取字节集按字符位置(reader, 'w', 2)
		t.Assert(string(reads), "llow")

		_, indexs = 文件类.X取字节集按字符位置(reader, 'w', 20)
		t.Assert(indexs, -1)

	})
}

func Test_GetBinContentsTilCharByPath(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			reads     []byte
			indexs    int64
			filepaths = "/testfile_GetContents.txt"
		)

		createTestFile(filepaths, "abcdefghijklmn")
		defer delTestFiles(filepaths)

		reads, _ = 文件类.X取文件字节集按字符位置(testpath()+filepaths, 'c', 2)
		t.Assert(string(reads), "c")

		reads, _ = 文件类.X取文件字节集按字符位置(testpath()+filepaths, 'y', 1)
		t.Assert(string(reads), "")

		_, indexs = 文件类.X取文件字节集按字符位置(testpath()+filepaths, 'x', 1)
		t.Assert(indexs, -1)

	})
}

func Test_Home(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			reads string
			err   error
		)

		reads, err = 文件类.X取用户目录("a", "b")
		t.AssertNil(err)
		t.AssertNE(reads, "")
	})
}

func Test_NotFound(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		teatFile := 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/readline/error.log"
		callback := func(line string) error {
			return nil
		}
		err := 文件类.X逐行读文本_函数(teatFile, callback)
		t.AssertNE(err, nil)
	})
}

func Test_ReadLines(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			expectList = []string{"a", "b", "c", "d", "e"}
			getList    = make([]string, 0)
			callback   = func(line string) error {
				getList = append(getList, line)
				return nil
			}
			teatFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/readline/file.log"
		)
		err := 文件类.X逐行读文本_函数(teatFile, callback)
		t.AssertEQ(getList, expectList)
		t.AssertEQ(err, nil)
	})
}

func Test_ReadLines_Error(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			callback = func(line string) error {
				return 错误类.X创建("custom error")
			}
			teatFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/readline/file.log"
		)
		err := 文件类.X逐行读文本_函数(teatFile, callback)
		t.AssertEQ(err.Error(), "custom error")
	})
}

func Test_ReadLinesBytes(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			expectList = [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d"), []byte("e")}
			getList    = make([][]byte, 0)
			callback   = func(line []byte) error {
				getList = append(getList, line)
				return nil
			}
			teatFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/readline/file.log"
		)
		err := 文件类.X逐行读字节集_函数(teatFile, callback)
		t.AssertEQ(getList, expectList)
		t.AssertEQ(err, nil)
	})
}

func Test_ReadLinesBytes_Error(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			callback = func(line []byte) error {
				return 错误类.X创建("custom error")
			}
			teatFile = 文件类.X路径取父目录(gdebug.CallerFilePath()) + 文件类.Separator + "testdata/readline/file.log"
		)
		err := 文件类.X逐行读字节集_函数(teatFile, callback)
		t.AssertEQ(err.Error(), "custom error")
	})
}
