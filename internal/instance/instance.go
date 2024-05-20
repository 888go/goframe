// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// Package instance 提供实例管理功能。
// 
// 注意，此包不用于缓存，因为它没有缓存过期。
// md5:9cde92d483190e72
package instance

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/ghash"
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

// Get 根据给定的名称返回实例。. md5:a44f9ed4c07f4bd7
func Get(name string) interface{} {
	return getGroup(name).Get(name)
}

// Set 将给定名称的实例设置到实例管理器中。. md5:b2ea0ff086c307ba
func Set(name string, instance interface{}) {
	getGroup(name).Set(name, instance)
}

// GetOrSet 通过名称获取实例，
// 如果不存在，则将其设置到实例管理器中并返回该实例。
// md5:6e30e1788811bdcf
func GetOrSet(name string, instance interface{}) interface{} {
	return getGroup(name).GetOrSet(name, instance)
}

// GetOrSetFunc 通过名称获取实例，
// 如果不存在，它将使用回调函数 `f` 返回的值设置实例，
// 然后返回这个实例。
// md5:3e2dff7c2a8267b6
func GetOrSetFunc(name string, f func() interface{}) interface{} {
	return getGroup(name).GetOrSetFunc(name, f)
}

// GetOrSetFuncLock 通过名称获取实例，
// 如果该实例不存在，则使用回调函数 `f` 的返回值设置实例，
// 并随后返回这个实例。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，
// 它在执行函数 `f` 时会对哈希映射加锁（mutex.Lock）。
// md5:d7adba14d37045fa
func GetOrSetFuncLock(name string, f func() interface{}) interface{} {
	return getGroup(name).GetOrSetFuncLock(name, f)
}

// SetIfNotExist 如果`name`不存在，则将`instance`设置到地图中，然后返回true。
// 如果`name`已经存在，则忽略`instance`并返回false。
// md5:0eb14110f7286ae3
func SetIfNotExist(name string, instance interface{}) bool {
	return getGroup(name).SetIfNotExist(name, instance)
}

// Clear 删除所有存储的实例。. md5:19c1efdd76e32ce6
func Clear() {
	for i := 0; i < groupNumber; i++ {
		groups[i].Clear()
	}
}
