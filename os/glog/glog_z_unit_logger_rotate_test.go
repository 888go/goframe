// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类_test

import (
	"context"
	"testing"
	"time"

	"github.com/888go/goframe/frame/g"
	gfile "github.com/888go/goframe/os/gfile"
	glog "github.com/888go/goframe/os/glog"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gstr "github.com/888go/goframe/text/gstr"
)

var (
	ctx = context.TODO()
)

func Test_Rotate_Size(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		l := glog.X创建()
		p := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err := l.X设置配置Map(g.Map{
			"Path":                 p,
			"File":                 "access.log",
			"StdoutPrint":          false,
			"RotateSize":           10,
			"RotateBackupLimit":    2,
			"RotateBackupExpire":   5 * time.Second,
			"RotateBackupCompress": 9,
			"RotateCheckInterval":  time.Second, // For unit testing only.
		})
		t.AssertNil(err)
		defer gfile.X删除(p)

		s := "1234567890abcdefg"
		for i := 0; i < 8; i++ {
			l.X输出(ctx, s)
			time.Sleep(time.Second)
		}

		logFiles, err := gfile.X枚举(p, "access*")
		t.AssertNil(err)

		for _, v := range logFiles {
			content := gfile.X读文本(v)
			t.AssertIN(gstr.X统计次数(content, s), []int{1, 2})
		}

		time.Sleep(time.Second * 3)

		files, err := gfile.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 2)

		time.Sleep(time.Second * 5)
		files, err = gfile.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
}

func Test_Rotate_Expire(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		l := glog.X创建()
		p := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		err := l.X设置配置Map(g.Map{
			"Path":                 p,
			"File":                 "access.log",
			"StdoutPrint":          false,
			"RotateExpire":         time.Second,
			"RotateBackupLimit":    2,
			"RotateBackupExpire":   5 * time.Second,
			"RotateBackupCompress": 9,
			"RotateCheckInterval":  time.Second, // For unit testing only.
		})
		t.AssertNil(err)
		defer gfile.X删除(p)

		s := "1234567890abcdefg"
		for i := 0; i < 10; i++ {
			l.X输出(ctx, s)
		}

		files, err := gfile.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 0)

		t.Assert(gstr.X统计次数(gfile.X读文本(gfile.X路径生成(p, "access.log")), s), 10)

		time.Sleep(time.Second * 3)

		filenames, err := gfile.X枚举(p, "*")
		t.Log(filenames, err)
		files, err = gfile.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 1)

		t.Assert(gstr.X统计次数(gfile.X读文本(gfile.X路径生成(p, "access.log")), s), 0)

		time.Sleep(time.Second * 5)
		files, err = gfile.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
}
