// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gcfg

import (
	"context"

	"github.com/gogf/gf/v2/internal/intlog"
)

// SetContent 为指定的`file`设置自定义配置内容。
// `file`是可选参数，默认值为DefaultConfigFile。
// md5:49ae38cf671e3b96
// ff:
// a:
// content:
// file:
func (a *AdapterFile) SetContent(content string, file ...string) {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	// 清除缓存了`name`的实例的文件缓存。 md5:214f88d3763fe8e1
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

// GetContent 为指定的 `file`(默认为DefaultConfigFile)返回自定义配置内容。
// `file` 参数是不必要的。
// md5:633823fe81267232
// ff:
// a:
// file:
func (a *AdapterFile) GetContent(file ...string) string {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	return customConfigContentMap.Get(name)
}

// RemoveContent 移除具有指定`file`的全局配置。如果未传递`name`，则删除默认组名的配置。
// md5:93cf717e8dc7670b
// ff:
// a:
// file:
func (a *AdapterFile) RemoveContent(file ...string) {
	name := DefaultConfigFileName
	if len(file) > 0 {
		name = file[0]
	}
	// 清除缓存了`name`的实例的文件缓存。 md5:214f88d3763fe8e1
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

// ClearContent 清除所有全局配置内容。 md5:89d91d7f2a05e7da
// ff:
// a:
func (a *AdapterFile) ClearContent() {
	customConfigContentMap.Clear()
	// 清除所有实例的缓存。 md5:9be897e5e4dd65a8
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
