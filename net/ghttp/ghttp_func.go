// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package ghttp

import (
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/httputil"
	"github.com/888go/goframe/text/gstr"
)

// SupportedMethods 返回所有支持的HTTP方法。
func SupportedMethods() []string {
	return gstr.SplitAndTrim(supportedHttpMethods, ",")
}

// BuildParams 为 http 客户端构建请求字符串。`params` 参数可以是以下类型：
// string/[]byte/map/struct/*struct。
//
// 可选参数 `noUrlEncode` 指定是否忽略对数据的 URL 编码。
func BuildParams(params interface{}, noUrlEncode ...bool) (encodedParamStr string) {
	return httputil.BuildParams(params, noUrlEncode...)
}

// niceCallFunc 调用函数 `f` 并实现异常捕获逻辑。
func niceCallFunc(f func()) {
	defer func() {
		if exception := recover(); exception != nil {
			switch exception {
			case exceptionExit, exceptionExitAll:
				return

			default:
				if v, ok := exception.(error); ok && gerror.HasStack(v) {
					// 这已经是一个带有堆栈信息的错误。
					panic(v)
				}
// 创建一个包含堆栈信息的新错误。
// 注意，这里有一个skip参数用于指向实际错误点的堆栈跟踪起始位置。
				if v, ok := exception.(error); ok {
					if gerror.Code(v) != gcode.CodeNil {
						panic(v)
					} else {
						panic(gerror.WrapCodeSkip(
							gcode.CodeInternalPanic, 1, v, "exception recovered",
						))
					}
				} else {
					panic(gerror.NewCodeSkipf(
						gcode.CodeInternalPanic, 1, "exception recovered: %+v", exception,
					))
				}
			}
		}
	}()
	f()
}
