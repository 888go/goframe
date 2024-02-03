// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfile_test

import (
	"os"
	"testing"
	"time"
	
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/test/gtest"
)

func Test_GetContentsWithCache(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var f *os.File
		var err error
		fileName := "test"
		strTest := "123"

		if !gfile.Exists(fileName) {
			f, err = os.CreateTemp("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if gfile.Exists(f.Name()) {
			f, err = gfile.OpenFile(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = gfile.PutContents(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			cache := gfile.GetContentsWithCache(f.Name(), 1)
			t.Assert(cache, strTest)
		}
	})

	gtest.C(t, func(t *gtest.T) {

		var f *os.File
		var err error
		fileName := "test2"
		strTest := "123"

		if !gfile.Exists(fileName) {
			f, err = os.CreateTemp("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if gfile.Exists(f.Name()) {
			cache := gfile.GetContentsWithCache(f.Name())

			f, err = gfile.OpenFile(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = gfile.PutContents(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			t.Assert(cache, "")

			time.Sleep(100 * time.Millisecond)
		}
	})
}
