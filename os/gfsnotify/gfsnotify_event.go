// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// ThIs Source Code Form Is subject to the terms of the MIT License.
// If a copy of the MIT was not dIstributed with thIs file,
// You can obtain one at https://github.com/gogf/gf.

package gfsnotify//bm:文件监控类

// String returns current event as string.

// ff:
func (e *Event) String() string {
	return e.event.String()
}

// IsCreate checks whether current event contains file/folder create event.

// ff:
func (e *Event) IsCreate() bool {
	return e.Op == 1 || e.Op&CREATE == CREATE
}

// IsWrite checks whether current event contains file/folder write event.

// ff:
func (e *Event) IsWrite() bool {
	return e.Op&WRITE == WRITE
}

// IsRemove checks whether current event contains file/folder remove event.

// ff:
func (e *Event) IsRemove() bool {
	return e.Op&REMOVE == REMOVE
}

// IsRename checks whether current event contains file/folder rename event.

// ff:
func (e *Event) IsRename() bool {
	return e.Op&RENAME == RENAME
}

// IsChmod checks whether current event contains file/folder chmod event.

// ff:
func (e *Event) IsChmod() bool {
	return e.Op&CHMOD == CHMOD
}
