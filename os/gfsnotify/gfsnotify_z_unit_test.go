// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件监控类_test

import (
	"testing"
	"time"
	
	"github.com/888go/goframe/container/garray"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/os/gfile"
	"github.com/888go/goframe/os/gfsnotify"
	"github.com/888go/goframe/os/gtime"
	"github.com/888go/goframe/test/gtest"
	"github.com/888go/goframe/util/gconv"
)

func TestWatcher_AddOnce(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		value := 安全变量类.New()
		path := 文件类.X取临时目录(转换类.String(时间类.X取时间戳纳秒()))
		err := 文件类.X写入文本(path, "init")
		t.AssertNil(err)
		defer 文件类.X删除(path)

		time.Sleep(100 * time.Millisecond)
		callback1, err := 文件监控类.AddOnce("mywatch", path, func(event *文件监控类.Event) {
			value.X设置值(1)
		})
		t.AssertNil(err)
		callback2, err := 文件监控类.AddOnce("mywatch", path, func(event *文件监控类.Event) {
			value.X设置值(2)
		})
		t.AssertNil(err)
		t.Assert(callback2, nil)

		err = 文件类.X写入文本(path, "1")
		t.AssertNil(err)

		time.Sleep(100 * time.Millisecond)
		t.Assert(value, 1)

		err = 文件监控类.RemoveCallback(callback1.Id)
		t.AssertNil(err)

		err = 文件类.X写入文本(path, "2")
		t.AssertNil(err)

		time.Sleep(100 * time.Millisecond)
		t.Assert(value, 1)
	})
}

func TestWatcher_AddRemove(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path1 := 文件类.X取临时目录() + 文件类.Separator + 转换类.String(时间类.X取时间戳纳秒())
		path2 := 文件类.X取临时目录() + 文件类.Separator + 转换类.String(时间类.X取时间戳纳秒()) + "2"
		文件类.X写入文本(path1, "1")
		defer func() {
			文件类.X删除(path1)
			文件类.X删除(path2)
		}()
		v := 安全变量类.NewInt(1)
		callback, err := 文件监控类.Add(path1, func(event *文件监控类.Event) {
			if event.IsWrite() {
				v.X设置值(2)
				return
			}
			if event.IsRename() {
				v.X设置值(3)
				文件监控类.Exit()
				return
			}
		})
		t.AssertNil(err)
		t.AssertNE(callback, nil)

		文件类.X写入文本(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 2)

		文件类.Rename别名(path1, path2)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 3)
	})

	单元测试类.C(t, func(t *单元测试类.T) {
		path1 := 文件类.X取临时目录() + 文件类.Separator + 转换类.String(时间类.X取时间戳纳秒())
		文件类.X写入文本(path1, "1")
		defer func() {
			文件类.X删除(path1)
		}()
		v := 安全变量类.NewInt(1)
		callback, err := 文件监控类.Add(path1, func(event *文件监控类.Event) {
			if event.IsWrite() {
				v.X设置值(2)
				return
			}
			if event.IsRemove() {
				v.X设置值(4)
				return
			}
		})
		t.AssertNil(err)
		t.AssertNE(callback, nil)

		文件类.X写入文本(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 2)

		文件类.X删除(path1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 4)

		文件类.X写入文本(path1, "1")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 4)
	})
}

func TestWatcher_Callback1(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		path1 := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		文件类.X写入文本(path1, "1")
		defer func() {
			文件类.X删除(path1)
		}()
		v := 安全变量类.NewInt(1)
		callback, err := 文件监控类.Add(path1, func(event *文件监控类.Event) {
			if event.IsWrite() {
				v.X设置值(2)
				return
			}
		})
		t.AssertNil(err)
		t.AssertNE(callback, nil)

		文件类.X写入文本(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 2)

		v.X设置值(3)
		文件监控类.RemoveCallback(callback.Id)
		文件类.X写入文本(path1, "3")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 3)
	})
}

func TestWatcher_Callback2(t *testing.T) {
	// 多个回调函数
	单元测试类.C(t, func(t *单元测试类.T) {
		path1 := 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		t.Assert(文件类.X写入文本(path1, "1"), nil)
		defer func() {
			文件类.X删除(path1)
		}()
		v1 := 安全变量类.NewInt(1)
		v2 := 安全变量类.NewInt(1)
		callback1, err1 := 文件监控类.Add(path1, func(event *文件监控类.Event) {
			if event.IsWrite() {
				v1.X设置值(2)
				return
			}
		})
		callback2, err2 := 文件监控类.Add(path1, func(event *文件监控类.Event) {
			if event.IsWrite() {
				v2.X设置值(2)
				return
			}
		})
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.AssertNE(callback1, nil)
		t.AssertNE(callback2, nil)

		t.Assert(文件类.X写入文本(path1, "2"), nil)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v1.X取值(), 2)
		t.Assert(v2.X取值(), 2)

		v1.X设置值(3)
		v2.X设置值(3)
		文件监控类.RemoveCallback(callback1.Id)
		t.Assert(文件类.X写入文本(path1, "3"), nil)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v1.X取值(), 3)
		t.Assert(v2.X取值(), 2)
	})
}

func TestWatcher_WatchFolderWithoutRecursively(t *testing.T) {
	单元测试类.C(t, func(t *单元测试类.T) {
		var (
			err     error
			array   = 数组类.X创建(true)
			dirPath = 文件类.X取临时目录(时间类.X取文本时间戳纳秒())
		)
		err = 文件类.X创建目录(dirPath)
		t.AssertNil(err)

		_, err = 文件监控类.Add(dirPath, func(event *文件监控类.Event) {
			// 打印输出 event.String() 的结果
			array.Append别名(1)
		}, false)
		t.AssertNil(err)
		time.Sleep(time.Millisecond * 100)
		t.Assert(array.X取长度(), 0)

		f, err := 文件类.X创建文件与目录(文件类.X路径生成(dirPath, "1"))
		t.AssertNil(err)
		t.AssertNil(f.Close())
		time.Sleep(time.Millisecond * 100)
		t.Assert(array.X取长度(), 1)
	})
}
