// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package fileinfo provides virtual os.FileInfo for given information.
package fileinfo

import (
	"os"
	"time"
)

type Info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}


// ff:
// modTime:
// mode:
// size:
// name:
func New(name string, size int64, mode os.FileMode, modTime time.Time) *Info {
	return &Info{
		name:    name,
		size:    size,
		mode:    mode,
		modTime: modTime,
	}
}


// ff:
func (i *Info) Name() string {
	return i.name
}


// ff:
func (i *Info) Size() int64 {
	return i.size
}


// ff:
func (i *Info) IsDir() bool {
	return i.mode.IsDir()
}


// ff:
func (i *Info) Mode() os.FileMode {
	return i.mode
}


// ff:
func (i *Info) ModTime() time.Time {
	return i.modTime
}


// ff:
func (i *Info) Sys() interface{} {
	return nil
}
