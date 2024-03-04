// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gcfg

import (
	"context"
	
	"github.com/888go/goframe/gcfg/internal/intlog"
)

// SetContent 为指定的`file`设置自定义配置内容。
// `file`参数不是必须的，默认值是DefaultConfigFile。
func (a *AdapterFile) SetContent(content string, file ...string) {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	// 清除缓存了`name`的实例的文件缓存。
	localInstances.LockFunc(func(m map[string]interface{}) {
		if customConfigContentMap.Contains(name) {
			for _, v := range m {
				if configInstance, ok := v.(*Config); ok {
					if fileConfig, ok := configInstance.GetAdapter().(*AdapterFile); ok {
						fileConfig.jsonMap.Remove(name)
					}
				}
			}
		}
		customConfigContentMap.Set(name, content)
	})
}

// GetContent 返回为指定 `file` 的自定义配置内容。
// 其中 `file` 参数非必要，默认为 DefaultConfigFile。
func (a *AdapterFile) GetContent(file ...string) string {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	return customConfigContentMap.Get(name)
}

// RemoveContent 删除具有指定 `file` 的全局配置。
// 如果未传递 `name`，则删除默认组名的配置。
func (a *AdapterFile) RemoveContent(file ...string) {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	// 清除缓存了`name`的实例的文件缓存。
	localInstances.LockFunc(func(m map[string]interface{}) {
		if customConfigContentMap.Contains(name) {
			for _, v := range m {
				if configInstance, ok := v.(*Config); ok {
					if fileConfig, ok := configInstance.GetAdapter().(*AdapterFile); ok {
						fileConfig.jsonMap.Remove(name)
					}
				}
			}
			customConfigContentMap.Remove(name)
		}
	})

	intlog.Printf(context.TODO(), `RemoveContent: %s`, name)
}

// ClearContent 清除所有全局配置内容。
func (a *AdapterFile) ClearContent() {
	customConfigContentMap.Clear()
	// 清除所有实例的缓存。
	localInstances.LockFunc(func(m map[string]interface{}) {
		for _, v := range m {
			if configInstance, ok := v.(*Config); ok {
				if fileConfig, ok := configInstance.GetAdapter().(*AdapterFile); ok {
					fileConfig.jsonMap.Clear()
				}
			}
		}
	})
	intlog.Print(context.TODO(), `RemoveConfig`)
}
