// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gerror
import (
	"bytes"
	"container/list"
	"fmt"
	"runtime"
	"strings"
	
	"github.com/888go/goframe/internal/consts"
	"github.com/888go/goframe/internal/errors"
	)
// stackInfo 管理特定错误的堆栈信息。
type stackInfo struct {
	Index   int        // Index 是当前错误在完整错误堆栈中的索引。
	Message string     // 错误信息字符串。
	Lines   *list.List // Lines 包含当前错误堆栈中所有错误堆栈行，按顺序排列。
}

// stackLine 管理堆栈中每一行的信息。
type stackLine struct {
	Function string // 函数名称，包含其完整的包路径。
	FileLine string // FileLine 是 Function 的源文件名及其行号。
}

// Stack 返回错误堆栈信息作为字符串。
func (err *Error) Stack() string {
	if err == nil {
		return ""
	}
	var (
		loop             = err
		index            = 1
		infos            []*stackInfo
		isStackModeBrief = errors.IsStackModeBrief()
	)
	for loop != nil {
		info := &stackInfo{
			Index:   index,
			Message: fmt.Sprintf("%-v", loop),
		}
		index++
		infos = append(infos, info)
		loopLinesOfStackInfo(loop.stack, info, isStackModeBrief)
		if loop.error != nil {
			if e, ok := loop.error.(*Error); ok {
				loop = e
			} else {
				infos = append(infos, &stackInfo{
					Index:   index,
					Message: loop.error.Error(),
				})
				index++
				break
			}
		} else {
			break
		}
	}
	filterLinesOfStackInfos(infos)
	return formatStackInfos(infos)
}

// filterLinesOfStackInfos 从顶部错误中移除后续堆栈中存在的重复行。
func filterLinesOfStackInfos(infos []*stackInfo) {
	var (
		ok      bool
		set     = make(map[string]struct{})
		info    *stackInfo
		line    *stackLine
		removes []*list.Element
	)
	for i := len(infos) - 1; i >= 0; i-- {
		info = infos[i]
		if info.Lines == nil {
			continue
		}
		for n, e := 0, info.Lines.Front(); n < info.Lines.Len(); n, e = n+1, e.Next() {
			line = e.Value.(*stackLine)
			if _, ok = set[line.FileLine]; ok {
				removes = append(removes, e)
			} else {
				set[line.FileLine] = struct{}{}
			}
		}
		if len(removes) > 0 {
			for _, e := range removes {
				info.Lines.Remove(e)
			}
		}
		removes = removes[:0]
	}
}

// formatStackInfos 格式化并返回错误堆栈信息作为字符串。
func formatStackInfos(infos []*stackInfo) string {
	var buffer = bytes.NewBuffer(nil)
	for i, info := range infos {
		buffer.WriteString(fmt.Sprintf("%d. %s\n", i+1, info.Message))
		if info.Lines != nil && info.Lines.Len() > 0 {
			formatStackLines(buffer, info.Lines)
		}
	}
	return buffer.String()
}

// formatStackLines 格式化并返回错误堆栈的行信息，以字符串形式。
func formatStackLines(buffer *bytes.Buffer, lines *list.List) string {
	var (
		line   *stackLine
		space  = "  "
		length = lines.Len()
	)
	for i, e := 0, lines.Front(); i < length; i, e = i+1, e.Next() {
		line = e.Value.(*stackLine)
		// Graceful indent.
		if i >= 9 {
			space = " "
		}
		buffer.WriteString(fmt.Sprintf(
			"   %d).%s%s\n        %s\n",
			i+1, space, line.Function, line.FileLine,
		))
	}
	return buffer.String()
}

// loopLinesOfStackInfo 遍历堆栈信息行，并生成堆栈行信息。
func loopLinesOfStackInfo(st stack, info *stackInfo, isStackModeBrief bool) {
	if st == nil {
		return
	}
	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)
			if isStackModeBrief {
				// 过滤整个GoFrame包堆栈路径。
				if strings.Contains(file, consts.StackFilterKeyForGoFrame) {
					continue
				}
			} else {
				// 包路径栈过滤功能
				if strings.Contains(file, stackFilterKeyLocal) {
					continue
				}
			}
			// 避免使用类似"`自动生成`"的堆栈字符串
			if strings.Contains(file, "<") {
				continue
			}
			// 忽略 GO ROOT 路径。
			if goRootForFilter != "" &&
				len(file) >= len(goRootForFilter) &&
				file[0:len(goRootForFilter)] == goRootForFilter {
				continue
			}
			if info.Lines == nil {
				info.Lines = list.New()
			}
			info.Lines.PushBack(&stackLine{
				Function: fn.Name(),
				FileLine: fmt.Sprintf(`%s:%d`, file, line),
			})
		}
	}
}
