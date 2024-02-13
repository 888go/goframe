// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 文件监控类

import (
	"context"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/internal/intlog"
)

// watchLoop 启动一个循环，用于从底层 inotify 监视器监听事件。
func (w *Watcher) watchLoop() {
	go func() {
		for {
			select {
			// Close event.
			case <-w.closeChan:
				return

			// Event listening.
			case ev := <-w.watcher.Events:
				// 在自定义时间段内过滤重复事件。
				_, err := w.cache.X设置值并跳过已存在(
					context.Background(),
					ev.String(),
					func(ctx context.Context) (value interface{}, err error) {
						w.events.X入栈(&Event{
							event:   ev,
							Path:    ev.Name,
							Op:      Op(ev.Op),
							Watcher: w,
						})
						return struct{}{}, nil
					}, repeatEventFilterDuration,
				)
				if err != nil {
					intlog.Errorf(context.TODO(), `%+v`, err)
				}

			case err := <-w.watcher.Errors:
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}
	}()
}

// eventLoop 是核心事件处理器。
func (w *Watcher) eventLoop() {
	go func() {
		for {
			if v := w.events.X出栈(); v != nil {
				event := v.(*Event)
				// 如果此路径没有任何回调函数，它将从监视器中移除该路径。
				callbacks := w.getCallbacks(event.Path)
				if len(callbacks) == 0 {
					_ = w.watcher.Remove(event.Path)
					continue
				}
				switch {
				case event.IsRemove():
// 应该再次检查路径是否存在。
// 如果路径仍然存在，则将其重新添加到监视器中。
					if fileExists(event.Path) {
// 它将路径添加回监控中。
// 我们无需担心重复添加的问题。
						if err := w.watcher.Add(event.Path); err != nil {
							intlog.Errorf(context.TODO(), `%+v`, err)
						} else {
							intlog.Printf(context.TODO(), "fake remove event, watcher re-adds monitor for: %s", event.Path)
						}
						// 将事件更改为 RENAME，这意味着它将自身重命名为原始名称。
						event.Op = RENAME
					}

				case event.IsRename():
// 应该再次检查路径是否存在。
// 如果路径仍然存在，则将其重新添加到监视器中。
					// Especially Some editors might do RENAME and then CHMOD when it's editing file.
					if fileExists(event.Path) {
// 可能会丢失对路径的监控，所以我们将路径重新添加回监控中。
// 我们无需担心重复添加的问题。
						if err := w.watcher.Add(event.Path); err != nil {
							intlog.Errorf(context.TODO(), `%+v`, err)
						} else {
							intlog.Printf(context.TODO(), "fake rename event, watcher re-adds monitor for: %s", event.Path)
						}
						// 将事件更改为 CHMOD。
						event.Op = CHMOD
					}

				case event.IsCreate():
// =========================================
// 注意，这里仅添加要监视的路径而无需注册任何回调函数，
// 因为其父级已经拥有这些回调函数。
// =========================================
					if fileIsDir(event.Path) {
						// 如果这是一个文件夹，那么它会递归地添加到监控中。
						for _, subPath := range fileAllDirs(event.Path) {
							if fileIsDir(subPath) {
								if err := w.watcher.Add(subPath); err != nil {
									intlog.Errorf(context.TODO(), `%+v`, err)
								} else {
									intlog.Printf(context.TODO(), "folder creation event, watcher adds monitor for: %s", subPath)
								}
							}
						}
					} else {
						// 如果它是一个文件，就直接将其添加到监控中。
						if err := w.watcher.Add(event.Path); err != nil {
							intlog.Errorf(context.TODO(), `%+v`, err)
						} else {
							intlog.Printf(context.TODO(), "file creation event, watcher adds monitor for: %s", event.Path)
						}
					}
				}
				// 按照顺序调用回调函数。
				for _, callback := range callbacks {
					go func(callback *Callback) {
						defer func() {
							if err := recover(); err != nil {
								switch err {
								case callbackExitEventPanicStr:
									w.RemoveCallback(callback.Id)
								default:
									if e, ok := err.(error); ok {
										panic(错误类.X多层错误码(错误码类.CodeInternalPanic, e))
									}
									panic(err)
								}
							}
						}()
						callback.Func(event)
					}(callback)
				}
			} else {
				break
			}
		}
	}()
}

// getCallbacks 搜索并返回给定 `path` 的所有回调函数。
// 如果它们是递归的，还会在其父级中搜索回调函数。
func (w *Watcher) getCallbacks(path string) (callbacks []*Callback) {
	// 首先添加自身的回调函数。
	if v := w.callbacks.X取值(path); v != nil {
		for _, v := range v.(*链表类.List).FrontAll() {
			callback := v.(*Callback)
			callbacks = append(callbacks, callback)
		}
	}
// 第二步，在其直接父级中搜索回调函数。
// 这里有特殊处理，这是`递归`和`非递归`逻辑之间的区别，
// 特指从`path`的直接父目录传递过来的事件。
	dirPath := fileDir(path)
	if v := w.callbacks.X取值(dirPath); v != nil {
		for _, v := range v.(*链表类.List).FrontAll() {
			callback := v.(*Callback)
			callbacks = append(callbacks, callback)
		}
	}
	// 最后，递归地在`path`目录的所有父级目录中搜索回调函数。
	for {
		parentDirPath := fileDir(dirPath)
		if parentDirPath == dirPath {
			break
		}
		if v := w.callbacks.X取值(parentDirPath); v != nil {
			for _, v := range v.(*链表类.List).FrontAll() {
				callback := v.(*Callback)
				if callback.recursive {
					callbacks = append(callbacks, callback)
				}
			}
		}
		dirPath = parentDirPath
	}
	return
}
