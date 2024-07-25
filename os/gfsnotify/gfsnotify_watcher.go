// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gfsnotify

import (
	"context"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
)

// 将监控器添加到观察者，监控的路径为`path`，回调函数为`callbackFunc`。
// 可选参数`recursive`指定是否递归监控`path`，默认值为真。
// md5:35e0c4a9c0901ef8
func (w *Watcher) Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	return w.AddOnce("", path, callbackFunc, recursive...)
}

// AddOnce 使用唯一的名称 `name` 监控 `path`，并调用回调函数 `callbackFunc` 只一次。如果多次使用相同的 `name` 参数调用 AddOnce，`path` 将只被添加一次监控。
// 
// 如果两次使用相同的 `name` 调用，将返回错误。
// 
// 可选参数 `recursive` 指定是否递归地监控 `path`，默认为 true。
// md5:6ead1d3d4bff4432
func (w *Watcher) AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
	w.nameSet.AddIfNotExistFuncLock(name, func() bool {
				// 首先，添加路径到观察者。 md5:8830f7aece4dab2e
		callback, err = w.addWithCallbackFunc(name, path, callbackFunc, recursive...)
		if err != nil {
			return false
		}
		// 如果是递归添加，那么它会将所有子文件夹都添加到监视中。
		// 注意：
		// 1. 它仅递归地向监视器添加**文件夹**，而不添加文件，
		//    因为如果文件夹被监视，其下属文件自然也会被监视。
		// 2. 它不对这些文件夹绑定回调函数，因为如果有任何事件产生，
		//    它会从父级开始递归地查找对应的回调函数。
		// md5:ebaf807e7d2c1bb2
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

// addWithCallbackFunc 将路径添加到底层监控器中，创建并返回一个回调对象。
// 请注意，如果多次调用该函数并传入相同的`path`，最新的调用将覆盖之前的设置。
// md5:bec1b4834bd3126d
func (w *Watcher) addWithCallbackFunc(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error) {
		// 检查并转换给定的路径为绝对路径。 md5:a7b26b31d4dc4d54
	if t := fileRealPath(path); t == "" {
		return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `"%s" does not exist`, path)
	} else {
		path = t
	}
		// 创建回调对象。 md5:35c1374926d9f0ab
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
		// 向监视器注册回调函数。 md5:803a60e7f5c04013
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
		// 将路径添加到基础监控器中。 md5:3136f0f0e2cd9407
	if err = w.watcher.Add(path); err != nil {
		err = gerror.Wrapf(err, `add watch failed for path "%s"`, path)
	} else {
		intlog.Printf(context.TODO(), "watcher adds monitor for: %s", path)
	}
		// 将回调添加到全局回调映射中。 md5:32fee24607f18b97
	callbackIdMap.Set(callback.Id, callback)
	return
}

// Close 关闭监听器。 md5:c20cd2708e199b34
func (w *Watcher) Close() {
	close(w.closeChan)
	if err := w.watcher.Close(); err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	w.events.Close()
}

// Remove 递归地移除与"path"关联的监视器和所有回调。 md5:e48d059cb96966c1
func (w *Watcher) Remove(path string) error {
		// 首先移除路径上的回调函数。 md5:15ba778318ad7bb9
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
		// 其次，移除所有没有回调函数的子文件的监控。 md5:477e906da20585f9
	if subPaths, err := fileScanDir(path, "*", true); err == nil && len(subPaths) > 0 {
		for _, subPath := range subPaths {
			if w.checkPathCanBeRemoved(subPath) {
				if internalErr := w.watcher.Remove(subPath); internalErr != nil {
					intlog.Errorf(context.TODO(), `%+v`, internalErr)
				}
			}
		}
	}
		// 最后，从底层监视器中删除路径的监视器。 md5:d85b24985f7de449
	err := w.watcher.Remove(path)
	if err != nil {
		err = gerror.Wrapf(err, `remove watch failed for path "%s"`, path)
	}
	return err
}

// checkPathCanBeRemoved 检查给定路径是否绑定了没有回调。 md5:affaad498733b441
func (w *Watcher) checkPathCanBeRemoved(path string) bool {
		// 首先直接检查监视器中的回调函数。 md5:7167bd75166cef8a
	if v := w.callbacks.Get(path); v != nil {
		return false
	}
		// 其次，检查其父级是否具有回调函数。 md5:e913cd965e3ce822
	dirPath := fileDir(path)
	if v := w.callbacks.Get(dirPath); v != nil {
		for _, c := range v.(*glist.List).FrontAll() {
			if c.(*Callback).recursive {
				return false
			}
		}
		return false
	}
		// 递归检查其父节点。 md5:4557bd1d1f1bec8a
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

// RemoveCallback 从观察者中移除具有给定回调ID的回调。 md5:78b678cca3a84b90
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
