// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 日志类

import (
	"context"

	"github.com/888go/goframe/internal/json"
)

// HandlerOutputJson 是将日志内容作为单一JSON输出的结构体。
// 备注: 此配置结构不做名称翻译, 输出的日志会直接用结构字段名称作为日志内容的键名,导致TestLogger_SetHandlers_HandlerJson单元测试不通过.  (2024-07-21)
// HandlerOutputJson 是将日志内容作为单一JSON输出的结构体。
// 备注: 此配置结构不做名称翻译, 输出的日志会直接用结构字段名称作为日志内容的键名,导致TestLogger_SetHandlers_HandlerJson单元测试不通过.  (2024-07-21)
// HandlerOutputJson 是将日志内容作为单一JSON输出的结构体。
// 备注: 此配置结构不做名称翻译, 输出的日志会直接用结构字段名称作为日志内容的键名,导致TestLogger_SetHandlers_HandlerJson单元测试不通过.  (2024-07-21)
// HandlerOutputJson 是将日志内容作为单一JSON输出的结构体。
// 备注: 此配置结构不做名称翻译, 输出的日志会直接用结构字段名称作为日志内容的键名,导致TestLogger_SetHandlers_HandlerJson单元测试不通过.  (2024-07-21)
// md5:d9846a62089232e7
type HandlerOutputJson struct {
	Time       string `json:""`           // 格式化的日期时间字符串，例如 "2016-01-09 12:00:00"。 md5:4e5e7f760859f24e
	TraceId    string `json:",omitempty"` // 跟踪ID，仅在启用跟踪时可用。 md5:3b404f27a111b55f
	CtxStr     string `json:",omitempty"` // 从上下文中检索到的值字符串，仅在配置了Config.CtxKeys时可用。 md5:b666fc055085d9a9
	Level      string `json:""`           // 格式化的级别字符串，如"DEBU"、"ERRO"等。例如：ERRO. md5:b12f3fdf5d67e119
	CallerPath string `json:",omitempty"` // 调用日志的源文件路径及其行号，只有在设置F_FILE_SHORT或F_FILE_LONG时才可用。 md5:b3812538f3d66adf
	CallerFunc string `json:",omitempty"` // 调用日志的源函数名称，仅在设置F_CALLER_FN时可用。 md5:2bfd8148853e8e4c
	Prefix     string `json:",omitempty"` // 自定义日志内容的前缀字符串。 md5:c5186327c46919a1
	Content    string `json:""`           // Content 是日志的主要内容，包含由记录器生成的错误堆栈字符串。 md5:1939e6f99648f6e4
	Stack      string `json:",omitempty"` // 由logger生成的堆栈跟踪字符串，只有在配置了Config.StStatus时才可用。 md5:5951e7cd2f97d44d
}

// HandlerJson 是一个处理器，用于将输出日志内容作为单个 JSON 字符串。 md5:5f3ff01b64c4588b
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
	if len(in.Values) > 0 {
		if output.Content != "" {
			output.Content += " "
		}
		output.Content += in.ValuesContent()
	}
	// Output json content.
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	in.Buffer.Write(jsonBytes)
	in.Buffer.Write([]byte("\n"))
	in.Next(ctx)
}
