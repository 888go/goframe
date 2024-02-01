// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package gfsnotify 提供了一个跨平台的文件系统通知接口。
package gfsnotify
import (
	"context"
	"sync"
	"time"
	
	"github.com/fsnotify/fsnotify"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/container/gqueue"
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/os/gcache"
	)
// Watcher 是用于文件变化监测的监视器。
type Watcher struct {
	watcher   *fsnotify.Watcher // 底层 fsnotify 对象。
	events    *gqueue.Queue     // 用于内部事件管理。
	cache     *gcache.Cache     // 用于重复事件过滤。
	nameSet   *gset.StrSet      // 用于AddOnce功能。
	callbacks *gmap.StrAnyMap   // Path(文件/文件夹)到回调函数的映射。
	closeChan chan struct{}     // 用于通知观察者关闭。
}

// Callback是Watcher的回调函数。
type Callback struct {
	Id        int                // 唯一标识回调对象的ID。
	Func      func(event *Event) // 回调函数。
	Path      string             // 绑定文件路径（绝对路径）。
	name      string             // AddOnce的注册名称。
	elem      *glist.Element     // watcher回调中的元素。
	recursive bool               // 是否递归绑定到路径
}

// Event 是由底层 fsnotify 产生的事件。
type Event struct {
	event   fsnotify.Event // Underlying event.
	Path    string         // 绝对文件路径。
	Op      Op             // File operation.
	Watcher *Watcher       // Parent watcher.
}

// Op 是用于文件操作的位集合（联合体）。
type Op uint32

// internalPanic 是用于内部使用的自定义恐慌函数。
type internalPanic string

const (
	CREATE Op = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
)

const (
	repeatEventFilterDuration               = time.Millisecond // 重复事件过滤的持续时间。
	callbackExitEventPanicStr internalPanic = "exit"           // 自定义退出事件，用于内部使用。
)

var (
	mu                  sync.Mutex                // 用于保护defaultWatcher在并发环境下的安全性的互斥锁。
	defaultWatcher      *Watcher                  // Default watcher.
	callbackIdMap       = gmap.NewIntAnyMap(true) // Id 到回调函数的映射。
	callbackIdGenerator = gtype.NewInt()          // 原子式ID生成器，用于回调。
)

// New 创建并返回一个新的监视器。
// 注意，监视器的数量受系统文件句柄设置的限制。
// 例如：在Linux系统中，fs.inotify.max_user_instances系统变量。
func New() (*Watcher, error) {
	w := &Watcher{
		cache:     gcache.New(),
		events:    gqueue.New(),
		nameSet:   gset.NewStrSet(true),
		closeChan: make(chan struct{}),
		callbacks: gmap.NewStrAnyMap(true),
	}
	if watcher, err := fsnotify.NewWatcher(); err == nil {
		w.watcher = watcher
	} else {
		intlog.Printf(context.TODO(), "New watcher failed: %v", err)
		return nil, err
	}
	w.watchLoop()
	w.eventLoop()
	return w, nil
}

// Add 使用默认观察器监控`path`，并使用回调函数`callbackFunc`。
// 可选参数`recursive`指定了是否递归地监控`path`，默认为true。
func Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	w, err := getDefaultWatcher()
	if err != nil {
		return nil, err
	}
	return w.Add(path, callbackFunc, recursive...)
}

// AddOnce 使用默认的监视器，通过回调函数 `callbackFunc` 仅对 `path` 进行一次性监控，并使用唯一的名称 `name` 标识。
// 如果多次调用 AddOnce 并传入相同的 `name` 参数，`path` 只会被添加监控一次。如果两次调用时使用了相同的 `name`，将会返回错误。
//
// 可选参数 `recursive` 指定是否递归地监控 `path`，默认情况下其值为 true。
func AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	w, err := getDefaultWatcher()
	if err != nil {
		return nil, err
	}
	return w.AddOnce(name, path, callbackFunc, recursive...)
}

// Remove 递归地从 watcher 中移除指定 `path` 的所有监控回调函数。
func Remove(path string) error {
	w, err := getDefaultWatcher()
	if err != nil {
		return err
	}
	return w.Remove(path)
}

// RemoveCallback 从观察者中移除具有给定id的指定回调函数。
func RemoveCallback(callbackId int) error {
	w, err := getDefaultWatcher()
	if err != nil {
		return err
	}
	callback := (*Callback)(nil)
	if r := callbackIdMap.Get(callbackId); r != nil {
		callback = r.(*Callback)
	}
	if callback == nil {
		return gerror.NewCodef(gcode.CodeInvalidParameter, `callback for id %d not found`, callbackId)
	}
	w.RemoveCallback(callbackId)
	return nil
}

// Exit 仅在回调函数中使用，可用于从监视器中移除自身的当前回调。
func Exit() {
	panic(callbackExitEventPanicStr)
}

// getDefaultWatcher 创建并返回默认的监视器。
// 这是为了实现延迟初始化的目的。
func getDefaultWatcher() (*Watcher, error) {
	mu.Lock()
	defer mu.Unlock()
	if defaultWatcher != nil {
		return defaultWatcher, nil
	}
	var err error
	defaultWatcher, err = New()
	return defaultWatcher, err
}
