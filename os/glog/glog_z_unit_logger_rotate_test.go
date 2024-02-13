// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 日志类_test

import (
	"context"
	"testing"
	"time"
	
	"github.com/888go/goframe/frame/g"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/text/gstr"
)

var (
	ctx = context.TODO()
)

func Test_Rotate_Size(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		p := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err := l.X设置配置Map(g.Map{
			"Path":                 p,
			"File":                 "access.log",
			"StdoutPrint":          false,
			"RotateSize":           10,
			"RotateBackupLimit":    2,
			"RotateBackupExpire":   5 * time.Second,
			"RotateBackupCompress": 9,
			"RotateCheckInterval":  time.Second, // 仅用于单元测试。
		})
		t.AssertNil(err)
		defer 文件类.X删除(p)

		s := "1234567890abcdefg"
		for i := 0; i < 8; i++ {
			l.X输出(ctx, s)
			time.Sleep(time.Second)
		}

		logFiles, err := 文件类.X枚举(p, "access*")
		t.AssertNil(err)

		for _, v := range logFiles {
			content := 文件类.X读文本(v)
			t.AssertIN(文本类.X统计次数(content, s), []int{1, 2})
		}

		time.Sleep(time.Second * 3)

		files, err := 文件类.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 2)

		time.Sleep(time.Second * 5)
		files, err = 文件类.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
}

func Test_Rotate_Expire(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		l := 日志类.X创建()
		p := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		err := l.X设置配置Map(g.Map{
			"Path":                 p,
			"File":                 "access.log",
			"StdoutPrint":          false,
			"RotateExpire":         time.Second,
			"RotateBackupLimit":    2,
			"RotateBackupExpire":   5 * time.Second,
			"RotateBackupCompress": 9,
			"RotateCheckInterval":  time.Second, // 仅用于单元测试。
		})
		t.AssertNil(err)
		defer 文件类.X删除(p)

		s := "1234567890abcdefg"
		for i := 0; i < 10; i++ {
			l.X输出(ctx, s)
		}

		files, err := 文件类.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 0)

		t.Assert(文本类.X统计次数(文件类.X读文本(文件类.X路径生成(p, "access.log")), s), 10)

		time.Sleep(time.Second * 3)

		filenames, err := 文件类.X枚举(p, "*")
		t.Log(filenames, err)
		files, err = 文件类.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 1)

		t.Assert(文本类.X统计次数(文件类.X读文本(文件类.X路径生成(p, "access.log")), s), 0)

		time.Sleep(time.Second * 5)
		files, err = 文件类.X枚举(p, "*.gz")
		t.AssertNil(err)
		t.Assert(len(files), 0)
	})
}
