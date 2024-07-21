// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gerror

import (
	"bytes"
	"container/list"
	"fmt"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/internal/consts"
	"github.com/gogf/gf/v2/internal/errors"
)

// stackInfo 管理特定错误的堆栈信息。 md5:79468b5dd3f5e0ad
type stackInfo struct {
	Index   int        // Index是整个错误堆栈中当前错误的索引。 md5:87ea995a56dc4559
	Message string     // 错误信息字符串。 md5:56ef906abe7cf21e
	Lines   *list.List // Lines contains all error stack lines of current error stack in sequence.
}

// stackLine 用于管理堆栈信息中的每一行。 md5:1a2dfe6d80506d43
type stackLine struct {
	Function string // 函数名称，包括其完整的包路径。 md5:7169ae6e7a85cc99
	FileLine string // FileLine 是 Function 的源文件名及其行号。 md5:e072d8f974569fd3
}

// Stack 返回错误堆栈信息作为字符串。 md5:467a08a63c68ff5e
// ff:
// err:
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

// filterLinesOfStackInfos 从顶级错误中移除后续堆栈中出现的重复行。 md5:da3a6efe58d9b63f
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

// formatStackInfos 格式化并返回错误堆栈信息作为字符串。 md5:edb267da17f1d379
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

// formatStackLines 将错误堆栈行格式化并返回为字符串。 md5:824c7dc18e84d9a7
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

// loopLinesOfStackInfo 遍历堆栈信息行，并生成堆栈行信息。 md5:104e6007b0f2fc07
func loopLinesOfStackInfo(st stack, info *stackInfo, isStackModeBrief bool) {
	if st == nil {
		return
	}
	for _, p := range st {
		if fn := runtime.FuncForPC(p - 1); fn != nil {
			file, line := fn.FileLine(p - 1)
			if isStackModeBrief {
				// 过滤整个GoFrame包堆栈路径。 md5:097ea4b9718ad7a8
				if strings.Contains(file, consts.StackFilterKeyForGoFrame) {
					continue
				}
			} else {
				// 包路径栈过滤。 md5:272fc7db0b653585
				if strings.Contains(file, stackFilterKeyLocal) {
					continue
				}
			}
			// 避免使用像"`自动生成`"这样的栈字符串. md5:d965e072120580f2
			if strings.Contains(file, "<") {
				continue
			}
			// Ignore GO ROOT paths.
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
