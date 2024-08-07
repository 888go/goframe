// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件类_test

import (
	"os"
	"testing"
	"time"

	gfile "github.com/888go/goframe/os/gfile"
	gtest "github.com/888go/goframe/test/gtest"
)

func Test_GetContentsWithCache(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var f *os.File
		var err error
		fileName := "test"
		strTest := "123"

		if !gfile.X是否存在(fileName) {
			f, err = os.CreateTemp("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if gfile.X是否存在(f.Name()) {
			f, err = gfile.X打开(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = gfile.X写入文本(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			cache := gfile.X缓存读文本(f.Name(), 1)
			t.Assert(cache, strTest)
		}
	})

	gtest.C(t, func(t *gtest.T) {

		var f *os.File
		var err error
		fileName := "test2"
		strTest := "123"

		if !gfile.X是否存在(fileName) {
			f, err = os.CreateTemp("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if gfile.X是否存在(f.Name()) {
			cache := gfile.X缓存读文本(f.Name())

			f, err = gfile.X打开(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = gfile.X写入文本(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			t.Assert(cache, "")

			time.Sleep(100 * time.Millisecond)
		}
	})
}
