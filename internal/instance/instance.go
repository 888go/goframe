// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package instance provides instances management.
//
// Note that this package is not used for cache, as it has no cache expiration.
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

// Get returns the instance by given name.

// ff:
// name:
func Get(name string) interface{} {
	return getGroup(name).Get(name)
}

// Set sets an instance to the instance manager with given name.

// ff:设置值
// instance:
// name:
func Set(name string, instance interface{}) {
	getGroup(name).Set(name, instance)
}

// GetOrSet returns the instance by name,
// or set instance to the instance manager if it does not exist and returns this instance.

// ff:
// instance:
// name:
func GetOrSet(name string, instance interface{}) interface{} {
	return getGroup(name).GetOrSet(name, instance)
}

// GetOrSetFunc returns the instance by name,
// or sets instance with returned value of callback function `f` if it does not exist
// and then returns this instance.

// ff:
// f:
// name:
func GetOrSetFunc(name string, f func() interface{}) interface{} {
	return getGroup(name).GetOrSetFunc(name, f)
}

// GetOrSetFuncLock returns the instance by name,
// or sets instance with returned value of callback function `f` if it does not exist
// and then returns this instance.
//
// GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f`
// with mutex.Lock of the hash map.

// ff:
// f:
// name:
func GetOrSetFuncLock(name string, f func() interface{}) interface{} {
	return getGroup(name).GetOrSetFuncLock(name, f)
}

// SetIfNotExist sets `instance` to the map if the `name` does not exist, then returns true.
// It returns false if `name` exists, and `instance` would be ignored.

// ff:
// instance:
// name:
func SetIfNotExist(name string, instance interface{}) bool {
	return getGroup(name).SetIfNotExist(name, instance)
}

// Clear deletes all instances stored.

// ff:
func Clear() {
	for i := 0; i < groupNumber; i++ {
		groups[i].Clear()
	}
}
