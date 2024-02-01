// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package instance 提供了实例管理功能。
//
// 注意，此包并不用于缓存，因为它没有缓存过期机制。
package instance
import (
	"github.com/888go/goframe/container/gmap"
	"github.com/888go/goframe/encoding/ghash"
	)
const (
	groupNumber = 64
)

var (
	groups = make([]*gmap.StrAnyMap, groupNumber)
)

func init() {
	for i := 0; i < groupNumber; i++ {
		groups[i] = gmap.NewStrAnyMap(true)
	}
}

func getGroup(key string) *gmap.StrAnyMap {
	return groups[int(ghash.DJB([]byte(key))%groupNumber)]
}

// Get通过给定的名称返回实例。
func Get(name string) interface{} {
	return getGroup(name).Get(name)
}

// Set 将具有给定名称的实例设置到实例管理器中。
func Set(name string, instance interface{}) {
	getGroup(name).Set(name, instance)
}

// GetOrSet 函数通过名称获取实例，
// 如果实例不存在，则将其设置到实例管理器中并返回该实例。
func GetOrSet(name string, instance interface{}) interface{} {
	return getGroup(name).GetOrSet(name, instance)
}

// GetOrSetFunc 函数通过名称返回实例，
// 如果实例不存在，则使用回调函数 `f` 返回的值设置该实例，
// 然后返回这个已设置的实例。
func GetOrSetFunc(name string, f func() interface{}) interface{} {
	return getGroup(name).GetOrSetFunc(name, f)
}

// GetOrSetFuncLock 通过名称返回实例，
// 如果实例不存在，则使用回调函数 `f` 返回的值设置该实例，
// 然后返回这个实例。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
func GetOrSetFuncLock(name string, f func() interface{}) interface{} {
	return getGroup(name).GetOrSetFuncLock(name, f)
}

// SetIfNotExist 如果`name`不存在，则将`instance`设置到map中，并返回true。
// 若`name`已存在，则返回false，同时`instance`将被忽略。
func SetIfNotExist(name string, instance interface{}) bool {
	return getGroup(name).SetIfNotExist(name, instance)
}

// Clear 删除所有已存储的实例。
func Clear() {
	for i := 0; i < groupNumber; i++ {
		groups[i].Clear()
	}
}
