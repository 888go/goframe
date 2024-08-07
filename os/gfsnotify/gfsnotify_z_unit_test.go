// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 文件监控类_test

import (
	"testing"
	"time"

	garray "github.com/888go/goframe/container/garray"
	gtype "github.com/888go/goframe/container/gtype"
	gfile "github.com/888go/goframe/os/gfile"
	gfsnotify "github.com/888go/goframe/os/gfsnotify"
	gtime "github.com/888go/goframe/os/gtime"
	gtest "github.com/888go/goframe/test/gtest"
	gconv "github.com/888go/goframe/util/gconv"
)

func TestWatcher_AddOnce(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		value := gtype.New()
		path := gfile.X取临时目录(gconv.String(gtime.X取时间戳纳秒()))
		err := gfile.X写入文本(path, "init")
		t.AssertNil(err)
		defer gfile.X删除(path)

		time.Sleep(100 * time.Millisecond)
		callback1, err := gfsnotify.AddOnce("mywatch", path, func(event *gfsnotify.Event) {
			value.X设置值(1)
		})
		t.AssertNil(err)
		callback2, err := gfsnotify.AddOnce("mywatch", path, func(event *gfsnotify.Event) {
			value.X设置值(2)
		})
		t.AssertNil(err)
		t.Assert(callback2, nil)

		err = gfile.X写入文本(path, "1")
		t.AssertNil(err)

		time.Sleep(100 * time.Millisecond)
		t.Assert(value, 1)

		err = gfsnotify.RemoveCallback(callback1.Id)
		t.AssertNil(err)

		err = gfile.X写入文本(path, "2")
		t.AssertNil(err)

		time.Sleep(100 * time.Millisecond)
		t.Assert(value, 1)
	})
}

func TestWatcher_AddRemove(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path1 := gfile.X取临时目录() + gfile.Separator + gconv.String(gtime.X取时间戳纳秒())
		path2 := gfile.X取临时目录() + gfile.Separator + gconv.String(gtime.X取时间戳纳秒()) + "2"
		gfile.X写入文本(path1, "1")
		defer func() {
			gfile.X删除(path1)
			gfile.X删除(path2)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.X设置值(2)
				return
			}
			if event.IsRename() {
				v.X设置值(3)
				gfsnotify.Exit()
				return
			}
		})
		t.AssertNil(err)
		t.AssertNE(callback, nil)

		gfile.X写入文本(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 2)

		gfile.Rename别名(path1, path2)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 3)
	})

	gtest.C(t, func(t *gtest.T) {
		path1 := gfile.X取临时目录() + gfile.Separator + gconv.String(gtime.X取时间戳纳秒())
		gfile.X写入文本(path1, "1")
		defer func() {
			gfile.X删除(path1)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
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

		gfile.X写入文本(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 2)

		gfile.X删除(path1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 4)

		gfile.X写入文本(path1, "1")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 4)
	})
}

func TestWatcher_Callback1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		path1 := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		gfile.X写入文本(path1, "1")
		defer func() {
			gfile.X删除(path1)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.X设置值(2)
				return
			}
		})
		t.AssertNil(err)
		t.AssertNE(callback, nil)

		gfile.X写入文本(path1, "2")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 2)

		v.X设置值(3)
		gfsnotify.RemoveCallback(callback.Id)
		gfile.X写入文本(path1, "3")
		time.Sleep(100 * time.Millisecond)
		t.Assert(v.X取值(), 3)
	})
}

func TestWatcher_Callback2(t *testing.T) {
	// multiple callbacks
	gtest.C(t, func(t *gtest.T) {
		path1 := gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		t.Assert(gfile.X写入文本(path1, "1"), nil)
		defer func() {
			gfile.X删除(path1)
		}()
		v1 := gtype.NewInt(1)
		v2 := gtype.NewInt(1)
		callback1, err1 := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v1.X设置值(2)
				return
			}
		})
		callback2, err2 := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v2.X设置值(2)
				return
			}
		})
		t.Assert(err1, nil)
		t.Assert(err2, nil)
		t.AssertNE(callback1, nil)
		t.AssertNE(callback2, nil)

		t.Assert(gfile.X写入文本(path1, "2"), nil)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v1.X取值(), 2)
		t.Assert(v2.X取值(), 2)

		v1.X设置值(3)
		v2.X设置值(3)
		gfsnotify.RemoveCallback(callback1.Id)
		t.Assert(gfile.X写入文本(path1, "3"), nil)
		time.Sleep(100 * time.Millisecond)
		t.Assert(v1.X取值(), 3)
		t.Assert(v2.X取值(), 2)
	})
}

func TestWatcher_WatchFolderWithoutRecursively(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			err     error
			array   = garray.X创建(true)
			dirPath = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
		)
		err = gfile.X创建目录(dirPath)
		t.AssertNil(err)

		_, err = gfsnotify.Add(dirPath, func(event *gfsnotify.Event) {
									// 打印事件的字符串表示. md5:03df73520ca66c12
			array.Append别名(1)
		}, false)
		t.AssertNil(err)
		time.Sleep(time.Millisecond * 100)
		t.Assert(array.X取长度(), 0)

		f, err := gfile.X创建文件与目录(gfile.X路径生成(dirPath, "1"))
		t.AssertNil(err)
		t.AssertNil(f.Close())
		time.Sleep(time.Millisecond * 100)
		t.Assert(array.X取长度(), 1)
	})
}

func TestWatcher_WatchClose(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			err     error
			dirPath = gfile.X取临时目录(gtime.X取文本时间戳纳秒())
			watcher *gfsnotify.Watcher
		)

		err = gfile.X创建目录(dirPath)
		t.AssertNil(err)

		watcher, err = gfsnotify.New()
		t.AssertNil(err)
		t.AssertNE(watcher, nil)

		time.Sleep(time.Millisecond * 100)
		watcher.Close()
	})
}
