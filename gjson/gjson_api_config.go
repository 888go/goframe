// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gjson

// SetSplitChar 设置分隔符字符，用于层级数据访问。
func (j *Json) SetSplitChar(char byte) {
	j.mu.Lock()
	j.c = char
	j.mu.Unlock()
}

// SetViolenceCheck 设置是否对层级数据访问进行暴力检查。
func (j *Json) SetViolenceCheck(enabled bool) {
	j.mu.Lock()
	j.vc = enabled
	j.mu.Unlock()
}
