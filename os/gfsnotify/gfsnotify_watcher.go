// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gfsnotify

import (
	"context"
	
	"github.com/888go/goframe/container/glist"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
)

// 添加监控，将回调函数`callbackFunc`添加到watcher中，并监控路径`path`。
// 可选参数`recursive`指定了是否递归地监控路径`path`，默认为true。
func (w *Watcher) Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	return w.AddOnce("", path, callbackFunc, recursive...)
}

// AddOnce 通过唯一名称 `name` 使用回调函数 `callbackFunc` 仅对 `path` 进行一次性监控。
// 如果多次调用 AddOnce 并传入相同的 `name` 参数，`path` 只会被添加一次进行监控。
//
// 若同一 `name` 被调用两次，将会返回错误。
//
// 可选参数 `recursive` 指定是否递归地监控 `path`，默认情况下其值为 true。
func (w *Watcher) AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	w.nameSet.AddIfNotExistFuncLock(name, func() bool {
		// 首先将路径添加到监视器中。
		callback, err = w.addWithCallbackFunc(name, path, callbackFunc, recursive...)
		if err != nil {
			return false
		}
// 如果是递归添加，则将所有子文件夹添加到监视器中。
// 注意：
// 1. 它仅递归地向监视器添加**文件夹**，而不添加文件，
//    因为如果监视了文件夹及其子文件也会被监视。
// 2. 它没有给这些文件夹绑定任何回调函数，因为在产生任何事件时，它会从其父级开始递归地搜索回调函数。
		if fileIsDir(path) && (len(recursive) == 0 || recursive[0]) {
			for _, subPath := range fileAllDirs(path) {
				if fileIsDir(subPath) {
					if err = w.watcher.Add(subPath); err != nil {
						err = gerror.Wrapf(err, `add watch failed for path "%s"`, subPath)
					} else {
						intlog.Printf(context.TODO(), "watcher adds monitor for: %s", subPath)
					}
				}
			}
		}
		if name == "" {
			return false
		}
		return true
	})
	return
}

// addWithCallbackFunc 将路径添加到底层监视器中，创建并返回一个回调对象。
// 非常注意，如果对同一`path`调用多次，最新的一次将会覆盖之前的所有内容。
func (w *Watcher) addWithCallbackFunc(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	// 检查并把给定的路径转换为绝对路径。
	if t := fileRealPath(path); t == "" {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `"%s" does not exist`, path)
	} else {
		path = t
	}
	// 创建回调对象。
	callback = &Callback{
		Id:        callbackIdGenerator.Add(1),
		Func:      callbackFunc,
		Path:      path,
		name:      name,
		recursive: true,
	}
	if len(recursive) > 0 {
		callback.recursive = recursive[0]
	}
	// 将回调函数注册到监视器中。
	w.callbacks.LockFunc(func(m map[string]interface{}) {
		list := (*glist.List)(nil)
		if v, ok := m[path]; !ok {
			list = glist.New(true)
			m[path] = list
		} else {
			list = v.(*glist.List)
		}
		callback.elem = list.PushBack(callback)
	})
	// 将路径添加到基础监视器中。
	if err = w.watcher.Add(path); err != nil {
		err = gerror.Wrapf(err, `add watch failed for path "%s"`, path)
	} else {
		intlog.Printf(context.TODO(), "watcher adds monitor for: %s", path)
	}
	// 将回调函数添加到全局回调映射中。
	callbackIdMap.Set(callback.Id, callback)
	return
}

// Close 关闭监视器。
func (w *Watcher) Close() {
	w.events.Close()
	if err := w.watcher.Close(); err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	close(w.closeChan)
}

// Remove 递归地移除与`path`关联的监视器及其所有回调。
func (w *Watcher) Remove(path string) error {
	// 首先移除该路径的回调函数。
	if value := w.callbacks.Remove(path); value != nil {
		list := value.(*glist.List)
		for {
			if item := list.PopFront(); item != nil {
				callbackIdMap.Remove(item.(*Callback).Id)
			} else {
				break
			}
		}
	}
	// 其次，移除所有无回调函数的子文件的监控
	if subPaths, err := fileScanDir(path, "*", true); err == nil && len(subPaths) > 0 {
		for _, subPath := range subPaths {
			if w.checkPathCanBeRemoved(subPath) {
				if internalErr := w.watcher.Remove(subPath); internalErr != nil {
					intlog.Errorf(context.TODO(), `%+v`, internalErr)
				}
			}
		}
	}
	// 最后，从底层监视器中移除该路径的监控。
	err := w.watcher.Remove(path)
	if err != nil {
		err = gerror.Wrapf(err, `remove watch failed for path "%s"`, path)
	}
	return err
}

// checkPathCanBeRemoved 检查给定路径是否未绑定任何回调函数。
func (w *Watcher) checkPathCanBeRemoved(path string) bool {
	// 首先检查watcher中直接的回调函数。
	if v := w.callbacks.Get(path); v != nil {
		return false
	}
	// 第二步检查其父级是否具有回调函数。
	dirPath := fileDir(path)
	if v := w.callbacks.Get(dirPath); v != nil {
		for _, c := range v.(*glist.List).FrontAll() {
			if c.(*Callback).recursive {
				return false
			}
		}
		return false
	}
	// 递归检查其父级
	parentDirPath := ""
	for {
		parentDirPath = fileDir(dirPath)
		if parentDirPath == dirPath {
			break
		}
		if v := w.callbacks.Get(parentDirPath); v != nil {
			for _, c := range v.(*glist.List).FrontAll() {
				if c.(*Callback).recursive {
					return false
				}
			}
			return false
		}
		dirPath = parentDirPath
	}
	return true
}

// RemoveCallback 从 watcher 中移除具有给定回调 id 的回调函数。
func (w *Watcher) RemoveCallback(callbackId int) {
	callback := (*Callback)(nil)
	if r := callbackIdMap.Get(callbackId); r != nil {
		callback = r.(*Callback)
	}
	if callback != nil {
		if r := w.callbacks.Get(callback.Path); r != nil {
			r.(*glist.List).Remove(callback.elem)
		}
		callbackIdMap.Remove(callbackId)
		if callback.name != "" {
			w.nameSet.Remove(callback.name)
		}
	}
}
