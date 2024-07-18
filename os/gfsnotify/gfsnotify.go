// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包gfsnotify提供了一个与平台无关的接口，用于文件系统通知。 md5:85a6a9e7b52e09e5
package gfsnotify

import (
	"context"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/os/gcache"
)

// Watcher是文件更改的监视器。 md5:82c8c6038aefc543
type Watcher struct {
	watcher   *fsnotify.Watcher // 基础的fsnotify对象。 md5:33936d13041bf0ad
	events    *gqueue.Queue     // 用于内部事件管理。 md5:3b850c6c87220b73
	cache     *gcache.Cache     // 用于重复事件过滤。 md5:bc4edc625e5962e4
	nameSet   *gset.StrSet      // 用于添加一次功能。 md5:840c98179b4051aa
	callbacks *gmap.StrAnyMap   // 回调函数的路径（文件/文件夹）映射。 md5:ffdb0824b6a6c3e3
	closeChan chan struct{}     // 用于监视器关闭的通知。 md5:bd08433f373aee64
}

// Callback是Watcher的回调函数。 md5:fa67ea2e048fd039
type Callback struct {
	Id        int                // 回调对象的唯一标识符。 md5:5e0b273b79bf867f
	Func      func(event *Event) // Callback function.
	Path      string             // 限定的文件路径（绝对）。 md5:300376a734dfb30f
	name      string             // AddOnce的注册名称。 md5:eb134382190fde28
	elem      *glist.Element     // 观察者回调中的元素。 md5:abe4631c343e0396
	recursive bool               // 是否递归地绑定到路径。 md5:b5e23c5fc9be130a
}

// Event 是底层 fsnotify 产生的事件。 md5:cf38e0981dbdfa84
type Event struct {
	event   fsnotify.Event // Underlying event.
	Path    string         // Absolute file path.
	Op      Op             // File operation.
	Watcher *Watcher       // Parent watcher.
}

// Op 是文件操作的位标志联合。 md5:b9c63ca71272a7d3
type Op uint32

// internalPanic 是用于内部使用的自定义恐慌。 md5:287806e552654f1d
type internalPanic string

const (
	CREATE Op = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
)

const (
	repeatEventFilterDuration               = time.Millisecond // 重复事件过滤器的持续时间。 md5:2f6fec9f3496777a
	callbackExitEventPanicStr internalPanic = "exit"           // 用于内部使用的自定义退出事件。 md5:7c86231249e45c6a
)

var (
	mu                  sync.Mutex                // 用于并发安全的defaultWatcher的Mutex。 md5:0b64dbe940db78a8
	defaultWatcher      *Watcher                  // Default watcher.
	callbackIdMap       = gmap.NewIntAnyMap(true) // ID到回调函数的映射。 md5:641a269397a47639
	callbackIdGenerator = gtype.NewInt()          // 用于回调的原子性ID生成器。 md5:2caf00d0d805af7b
)

// New creates and returns a new watcher.
// Note that the watcher number is limited by the file handle setting of the system.
// ff:
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

// 使用默认的观察者(`watcher`)监控路径`path`，并调用回调函数`callbackFunc`。
// 可选参数`recursive`指定是否递归地监控路径`path`，默认为true。
// md5:e660326b83136bd1
// ff:
// path:
// callbackFunc:
// event:
// recursive:
// callback:
// err:
func Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	w, err := getDefaultWatcher()
	if err != nil {
		return nil, err
	}
	return w.Add(path, callbackFunc, recursive...)
}

// AddOnce 使用唯一名称 `name` 及回调函数 `callbackFunc`，仅使用默认监视器监控 `path` 一次。
// 如果多次调用 AddOnce 并传入相同的 `name` 参数，`path` 仅会被添加监控一次。如果使用相同的 `name` 调用两次，它将返回错误。
//
// 可选参数 `recursive` 指定是否递归监控 `path`，默认为 true。
// md5:c28c83d5a2230d07
// ff:
// name:
// path:
// callbackFunc:
// event:
// recursive:
// callback:
// err:
func AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	w, err := getDefaultWatcher()
	if err != nil {
		return nil, err
	}
	return w.AddOnce(name, path, callbackFunc, recursive...)
}

// Remove 递归地从监视器中删除给定`path`的所有监控回调。 md5:63888786f53ffca5
// ff:
// path:
func Remove(path string) error {
	w, err := getDefaultWatcher()
	if err != nil {
		return err
	}
	return w.Remove(path)
}

// RemoveCallback 从观察者中移除具有给定ID的指定回调。 md5:af906f3547f93046
// ff:
// callbackId:
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

// Exit 只在回调函数中使用，可以用于从观察者中移除当前的回调。
// md5:697f4cd00adc082e
// ff:
func Exit() {
	panic(callbackExitEventPanicStr)
}

// getDefaultWatcher 创建并返回默认的监视器。
// 这用于惰性初始化的目的。
// md5:c1a7b4f4102130c0
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
