// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package gfsnotify

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/internal/intlog"
)

// watchLoop 启动循环以从底层inotify监控器监听事件。 md5:a057c294cb3f7186
func (w *Watcher) watchLoop() {
	go func() {
		for {
			select {
			// Close event.
			case <-w.closeChan:
				return

			// Event listening.
			case ev, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				// 过滤自定义持续时间内的重复事件。 md5:f7b5d987e84f8092
				_, err := w.cache.SetIfNotExist(
					context.Background(),
					ev.String(),
					func(ctx context.Context) (value interface{}, err error) {
						w.events.Push(&Event{
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

// eventLoop是核心事件处理器。 md5:7adb3bf9c821349b
func (w *Watcher) eventLoop() {
	go func() {
		for {
			if v := w.events.Pop(); v != nil {
				event := v.(*Event)
				// 如果该路径没有任何回调，就从监控中移除它。 md5:1d18925e16d1ccb5
				callbacks := w.getCallbacks(event.Path)
				if len(callbacks) == 0 {
					_ = w.watcher.Remove(event.Path)
					continue
				}
				switch {
				case event.IsRemove():
					// 它应该再次检查路径的存在。
					// 如果该路径仍然存在，它会将其重新添加到监视器中。 md5:216ebbce200ac7a4
					if fileExists(event.Path) {
						// 将路径重新添加到监控列表。
						// 我们不需要担心重复添加的问题。 md5:4487198f5d35bb60
						if err := w.watcher.Add(event.Path); err != nil {
							intlog.Errorf(context.TODO(), `%+v`, err)
						} else {
							intlog.Printf(context.TODO(), "fake remove event, watcher re-adds monitor for: %s", event.Path)
						}
						// 将事件更改为 RENAME，这意味着它将自己重命名为原始名称。 md5:7e6fbf14f9528be7
						event.Op = RENAME
					}

				case event.IsRename():
					// 它应该再次检查路径的存在。
					// 如果该路径仍然存在，它会将其重新添加到监视器中。 md5:216ebbce200ac7a4
					// Especially Some editors might do RENAME and then CHMOD when it's editing file.
					if fileExists(event.Path) {
						// 可能会丢失对路径的监控，因此我们需将路径重新添加到监控中。
						// 我们无需担心重复添加的问题。 md5:d6dd87eba165d9e7
						if err := w.watcher.Add(event.Path); err != nil {
							intlog.Errorf(context.TODO(), `%+v`, err)
						} else {
							intlog.Printf(context.TODO(), "fake rename event, watcher re-adds monitor for: %s", event.Path)
						}
						// 将事件更改为CHMOD。 md5:84c563944c4dfa07
						event.Op = CHMOD
					}

				case event.IsCreate():
					// =========================================
					// 注意，这里只是添加了要监控的路径，而不需要注册回调，
					// 因为它的父级已经具有了回调。
					// ========================================= md5:2b5f1f3849c5ccff
					if fileIsDir(event.Path) {
						// 如果这是一个文件夹，它会递归地添加以进行监控。 md5:3b1a61cf45e4cf3a
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
						// 如果它是一个文件，就直接将其添加到监控中。 md5:18b66bfd2946b42e
						if err := w.watcher.Add(event.Path); err != nil {
							intlog.Errorf(context.TODO(), `%+v`, err)
						} else {
							intlog.Printf(context.TODO(), "file creation event, watcher adds monitor for: %s", event.Path)
						}
					}
				}
				// 按顺序调用回调函数。 md5:426b787bf42f20fa
				for _, callback := range callbacks {
					go func(callback *Callback) {
						defer func() {
							if err := recover(); err != nil {
								switch err {
								case callbackExitEventPanicStr:
									w.RemoveCallback(callback.Id)
								default:
									if e, ok := err.(error); ok {
										panic(gerror.WrapCode(gcode.CodeInternalPanic, e))
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

// getCallbacks 搜索并返回所有具有给定 `path` 的回调。如果它们是递归的，还会在其父级中搜索回调。 md5:abe3c32241868912
func (w *Watcher) getCallbacks(path string) (callbacks []*Callback) {
	// 首先，添加自身的回调。 md5:474fe6dbf371de56
	if v := w.callbacks.Get(path); v != nil {
		for _, v := range v.(*glist.List).FrontAll() {
			callback := v.(*Callback)
			callbacks = append(callbacks, callback)
		}
	}
	// 其次，在其直接父级中搜索回调。
	// 这里有特殊的处理逻辑，区别于“递归”与“非递归”的处理方式，
	// 特指针对来源于`path`的直接父级目录的事件。 md5:4e4cd99683eb9f66
	dirPath := fileDir(path)
	if v := w.callbacks.Get(dirPath); v != nil {
		for _, v := range v.(*glist.List).FrontAll() {
			callback := v.(*Callback)
			callbacks = append(callbacks, callback)
		}
	}
	// 最后，递归地搜索`path`目录的所有父级以查找回调函数。 md5:24dea4c80a5e5c6d
	for {
		parentDirPath := fileDir(dirPath)
		if parentDirPath == dirPath {
			break
		}
		if v := w.callbacks.Get(parentDirPath); v != nil {
			for _, v := range v.(*glist.List).FrontAll() {
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
