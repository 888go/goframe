// 版权声明：GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码受MIT许可证条款约束。如果此文件未附带MIT许可证的副本，
// 您可以从 https://github.com/gogf/gf 获取。
// md5:12b80d680e9de440

package gfsnotify//bm:文件监控类

// String 返回当前事件的字符串表示。 md5:a1d293a81ad7d90a
// ff:
// e:
func (e *Event) String() string {
	return e.event.String()
}

// IsCreate 检查当前事件是否包含文件/目录创建事件。 md5:b4246a419de617f8
// ff:
// e:
func (e *Event) IsCreate() bool {
	return e.Op == 1 || e.Op&CREATE == CREATE
}

// IsWrite检查当前事件是否包含文件/文件夹写入事件。 md5:0f9ece45dfe47971
// ff:
// e:
func (e *Event) IsWrite() bool {
	return e.Op&WRITE == WRITE
}

// IsRemove 检查当前事件是否包含文件/文件夹删除事件。 md5:bb4f4468acdccf13
// ff:
// e:
func (e *Event) IsRemove() bool {
	return e.Op&REMOVE == REMOVE
}

// IsRename 检查当前事件是否包含文件/文件夹重命名事件。 md5:98b87df81be40ac3
// ff:
// e:
func (e *Event) IsRename() bool {
	return e.Op&RENAME == RENAME
}

// IsChmod 检查当前事件是否包含文件/文件夹权限变更事件。 md5:71f642290a74e6ad
// ff:
// e:
func (e *Event) IsChmod() bool {
	return e.Op&CHMOD == CHMOD
}
