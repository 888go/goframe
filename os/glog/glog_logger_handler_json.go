// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package glog

import (
	"context"
	
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// HandlerOutputJson 是一个结构体，用于将日志内容以单个 JSON 的形式输出。
type HandlerOutputJson struct {
	Time       string `json:""`           // 格式化的时间字符串，如 "2016-01-09 12:00:00"。
	TraceId    string `json:",omitempty"` // 跟踪ID，仅在启用跟踪时可用。
	CtxStr     string `json:",omitempty"` // 从context中获取的字符串类型的上下文值，但只有在配置了Config.CtxKeys时才可用。
	Level      string `json:""`           // 格式化的级别字符串，如 "DEBU", "ERRO" 等。例如：ERRO
	CallerPath string `json:",omitempty"` // 调用日志记录的源文件路径及其行号，仅在设置了 F_FILE_SHORT 或 F_FILE_LONG 时可用。
	CallerFunc string `json:",omitempty"` // 如果设置了F_CALLER_FN，该变量记录调用日志函数的源函数名。
	Prefix     string `json:",omitempty"` // 自定义日志内容前缀字符串。
	Content    string `json:""`           // Content 是主要的日志内容，包含由 logger 生成的错误堆栈字符串。
	Stack      string `json:",omitempty"` // Stack 字符串由 logger 生成，仅在配置了 Config.StStatus 时可用。
}

// HandlerJson 是一个处理器，用于将输出的日志内容作为单个 JSON 字符串进行处理。
func HandlerJson(ctx context.Context, in *HandlerInput) {
	output := HandlerOutputJson{
		Time:       in.TimeFormat,
		TraceId:    in.TraceId,
		CtxStr:     in.CtxStr,
		Level:      in.LevelFormat,
		CallerFunc: in.CallerFunc,
		CallerPath: in.CallerPath,
		Prefix:     in.Prefix,
		Content:    in.Content,
		Stack:      in.Stack,
	}
	// 将values字符串内容进行转换
	var valueContent string
	for _, v := range in.Values {
		valueContent = gconv.String(v)
		if len(valueContent) == 0 {
			continue
		}
		if len(output.Content) > 0 {
			if output.Content[len(output.Content)-1] == '\n' {
				// 删除一个空行（\n\n）
				if valueContent[0] == '\n' {
					valueContent = valueContent[1:]
				}
				output.Content += valueContent
			} else {
				output.Content += " " + valueContent
			}
		} else {
			output.Content += valueContent
		}
	}
	// 输出json内容。
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	in.Buffer.Write(jsonBytes)
	in.Buffer.Write([]byte("\n"))
	in.Next(ctx)
}
