// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT协议条款。如果随此文件未分发MIT协议副本，
// 您可以在https://github.com/gogf/gf获取一份。

package gfsnotify

// String 返回当前事件的字符串表示。
func (e *Event) String() string {
	return e.event.String()
}

// IsCreate 检查当前事件是否包含文件/文件夹创建事件。
func (e *Event) IsCreate() bool {
	return e.Op == 1 || e.Op&CREATE == CREATE
}

// IsWrite 检查当前事件是否包含文件/文件夹写入事件。
func (e *Event) IsWrite() bool {
	return e.Op&WRITE == WRITE
}

// IsRemove 检查当前事件是否包含文件/文件夹移除事件。
func (e *Event) IsRemove() bool {
	return e.Op&REMOVE == REMOVE
}

// IsRename 检查当前事件是否包含文件/文件夹重命名事件。
func (e *Event) IsRename() bool {
	return e.Op&RENAME == RENAME
}

// IsChmod 检查当前事件是否包含文件/文件夹权限更改事件。
func (e *Event) IsChmod() bool {
	return e.Op&CHMOD == CHMOD
}
