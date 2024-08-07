// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/httputil"
	gstr "github.com/888go/goframe/text/gstr"
)

// X取所支持的HTTP方法 返回所有支持的HTTP方法。 md5:1fc8637928ce346f
func X取所支持的HTTP方法() []string {
	return gstr.X分割并忽略空值(supportedHttpMethods, ",")
}

// X生成请求参数 为 httpClient 构建请求字符串。`params` 可以是以下类型之一：
// string/[]byte/map/struct/*struct。
//
// 可选参数 `noUrlEncode` 指定是否忽略数据的 URL 编码。
// md5:6816cb48e0c8752b
func X生成请求参数(参数 interface{}, 不进行URL编码 ...bool) (请求参数 string) {
	return httputil.BuildParams(参数, 不进行URL编码...)
}

// niceCallFunc 使用异常捕获逻辑调用函数 `f`。 md5:cd4c356f9e76fa6e
func niceCallFunc(f func()) {
	defer func() {
		if exception := recover(); exception != nil {
			switch exception {
			case exceptionExit, exceptionExitAll:
				return

			default:
				if v, ok := exception.(error); ok && gerror.X判断是否带堆栈(v) {
										// 它已经是一个带有堆栈信息的错误。 md5:ec045ebe21bca18d
					panic(v)
				}
				// 创建一个带有堆栈信息的新错误。
				// 注意，skip 参数指定了从哪个调用栈开始追踪真正的错误点。
				// md5:e23da1f0a4a0c90f
				if v, ok := exception.(error); ok {
					if gerror.X取错误码(v) != gcode.CodeNil {
						panic(v)
					} else {
						panic(gerror.X多层错误码并跳过堆栈(
							gcode.CodeInternalPanic, 1, v, "exception recovered",
						))
					}
				} else {
					panic(gerror.X创建错误码并跳过堆栈与格式化(
						gcode.CodeInternalPanic, 1, "exception recovered: %+v", exception,
					))
				}
			}
		}
	}()
	f()
}
