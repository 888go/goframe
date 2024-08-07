// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package http类

import (
	"fmt"
	"net/http"

	gjson "github.com/888go/goframe/encoding/gjson"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	gconv "github.com/888go/goframe/util/gconv"
)

// X写响应缓冲区 将 `content` 写入响应缓冲区。 md5:2f656734fbf8eab6
func (r *Response) X写响应缓冲区(内容 ...interface{}) {
	if r.IsHijacked() || len(内容) == 0 {
		return
	}
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	for _, v := range 内容 {
		switch value := v.(type) {
		case []byte:
			_, _ = r.BufferWriter.Write(value)
		case string:
			_, _ = r.BufferWriter.WriteString(value)
		default:
			_, _ = r.BufferWriter.WriteString(gconv.String(v))
		}
	}
}

// X写响应缓冲区并退出 将`content`写入响应缓冲区并退出当前处理器的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:afcb2dda1beb9358
func (r *Response) X写响应缓冲区并退出(内容 ...interface{}) {
	r.X写响应缓冲区(内容...)
	r.Request.X退出当前()
}

// X写覆盖响应缓冲区 将响应缓冲区用`content`重写。 md5:d68f13dc57329ab0
func (r *Response) X写覆盖响应缓冲区(内容 ...interface{}) {
	r.ClearBuffer()
	r.X写响应缓冲区(内容...)
}

// X写覆盖响应缓冲区并退出 将响应缓冲区用 `content` 替换，然后退出当前处理程序的执行。"Exit" 功能通常用于方便地替换处理程序中的返回语句。
// md5:968d387aea44eeab
func (r *Response) X写覆盖响应缓冲区并退出(内容 ...interface{}) {
	r.X写覆盖响应缓冲区(内容...)
	r.Request.X退出当前()
}

// X写响应缓冲区并格式化 使用 fmt.Sprintf 格式化字符串并写出响应。 md5:15163b759bd146b8
func (r *Response) X写响应缓冲区并格式化(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区(fmt.Sprintf(格式, 内容...))
}

// X写响应缓冲区并退出与格式化使用fmt.Sprintf格式化响应内容并退出当前处理器的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:01275db804fa4029
func (r *Response) X写响应缓冲区并退出与格式化(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区并格式化(格式, 内容...)
	r.Request.X退出当前()
}

// X写响应缓冲区并换行 使用`content`和换行符写入响应。 md5:574e18a271a92e20
func (r *Response) X写响应缓冲区并换行(内容 ...interface{}) {
	if len(内容) == 0 {
		r.X写响应缓冲区("\n")
		return
	}
	r.X写响应缓冲区(append(内容, "\n")...)
}

// X写响应缓冲区并退出与换行 写入包含`content`和换行符的响应，然后退出当前处理器的执行。"Exit"特性通常用于方便地替换处理器中的返回语句。
// md5:bb5f123bedaec380
func (r *Response) X写响应缓冲区并退出与换行(内容 ...interface{}) {
	r.X写响应缓冲区并换行(内容...)
	r.Request.X退出当前()
}

// X写响应缓冲区并格式化与换行 使用 fmt.Sprintf 格式化输出并将结果作为响应写入，同时添加换行符。 md5:154e6f5f52878f00
func (r *Response) X写响应缓冲区并格式化与换行(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区并换行(fmt.Sprintf(格式, 内容...))
}

// X写响应缓冲区并退出与格式化换行使用fmt.Sprintf格式化响应并添加换行符，然后退出当前处理程序的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:ee5288e61cdea4b2
func (r *Response) X写响应缓冲区并退出与格式化换行(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区并格式化与换行(格式, 内容...)
	r.Request.X退出当前()
}

// X写响应缓冲区JSON 将`content`以JSON格式写入到响应中。 md5:0ca8d5da1805456f
func (r *Response) X写响应缓冲区JSON(内容 interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
		// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:e14783864a1068a9
	switch 内容.(type) {
	case string, []byte:
		r.X写响应缓冲区(gconv.String(内容))
		return
	}
		// 否则，使用json.Marshal函数对参数进行编码。 md5:b140f4be3fab1fa1
	if b, err := json.Marshal(内容); err != nil {
		panic(gerror.X多层错误(err, `WriteJson failed`))
	} else {
		r.X写响应缓冲区(b)
	}
}

// X写响应缓冲区JSON并退出 将`content`以JSON格式写入响应，并在成功时退出当前处理器的执行。"Exit"功能常用于替代处理器中的返回语句，以提供便利。
// md5:0714d99528fcb93e
func (r *Response) X写响应缓冲区JSON并退出(内容 interface{}) {
	r.X写响应缓冲区JSON(内容)
	r.Request.X退出当前()
}

// X写响应缓冲区JSONP 将`content`按照JSONP格式写入响应。
// 
// 注意，对于JSONP格式，请求中应该包含一个名为"callback"的参数。
// md5:32a3e4fa6b4e92b0
func (r *Response) X写响应缓冲区JSONP(内容 interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
		// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:ff82edfcddee9a78
	switch 内容.(type) {
	case string, []byte:
		r.X写响应缓冲区(gconv.String(内容))
		return
	}
		// 否则，使用json.Marshal函数对参数进行编码。 md5:b140f4be3fab1fa1
	if b, err := json.Marshal(内容); err != nil {
		panic(gerror.X多层错误(err, `WriteJsonP failed`))
	} else {
				// 设置HTTP响应头中的"Content-Type"为"application/json". md5:8e0be5eb3c232d44
		if callback := r.Request.Get别名("callback").String(); callback != "" {
			buffer := []byte(callback)
			buffer = append(buffer, byte('('))
			buffer = append(buffer, b...)
			buffer = append(buffer, byte(')'))
			r.X写响应缓冲区(buffer)
		} else {
			r.X写响应缓冲区(b)
		}
	}
}

// X写响应缓冲区JSONP并退出 将 `content` 以 JSONP 格式写入响应，并在成功时退出当前处理器的执行。"Exit" 功能常用于替代处理器中的返回语句，以提供便利。
//
// 请注意，为了使用 JSONP 格式，请求中应该包含一个 "callback" 参数。
// md5:6c959e76945e075a
func (r *Response) X写响应缓冲区JSONP并退出(内容 interface{}) {
	r.X写响应缓冲区JSONP(内容)
	r.Request.X退出当前()
}

// X写响应缓冲区XML 将`content`以XML格式写入响应。 md5:850a872cf25d6a70
func (r *Response) X写响应缓冲区XML(内容 interface{}, 根标记 ...string) {
	r.Header().Set("Content-Type", contentTypeXml)
		// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:4fc9d6cf062e5bf9
	switch 内容.(type) {
	case string, []byte:
		r.X写响应缓冲区(gconv.String(内容))
		return
	}
	if b, err := gjson.X创建(内容).X取xml字节集(根标记...); err != nil {
		panic(gerror.X多层错误(err, `WriteXml failed`))
	} else {
		r.X写响应缓冲区(b)
	}
}

// X写响应缓冲区XML并退出 将`content`以XML格式写入响应，并在成功时退出当前处理器的执行。
// “退出”功能常用于便捷地替代处理器中return语句的使用。
// md5:12a9a20328b00f55
func (r *Response) X写响应缓冲区XML并退出(内容 interface{}, 根标记 ...string) {
	r.X写响应缓冲区XML(内容, 根标记...)
	r.Request.X退出当前()
}

// X写响应缓冲区与HTTP状态码 将HTTP状态码`status`和内容`content`写入响应。
// 请注意，它不会在这里设置Content-Type头。
// md5:8b7195f02ad8ced0
func (r *Response) X写响应缓冲区与HTTP状态码(状态码 int, 内容 ...interface{}) {
	r.WriteHeader(状态码)
	if len(内容) > 0 {
		r.X写响应缓冲区(内容...)
	} else {
		r.X写响应缓冲区(http.StatusText(状态码))
	}
}

// X写响应缓冲区与HTTP状态码并退出 将HTTP状态码`status`和内容`content`写入响应，并在成功时退出当前处理程序的执行。"Exit"特性通常用于方便地替代处理程序中返回语句的使用。
// md5:2e5cbf96316a0c3c
func (r *Response) X写响应缓冲区与HTTP状态码并退出(状态码 int, 内容 ...interface{}) {
	r.X写响应缓冲区与HTTP状态码(状态码, 内容...)
	r.Request.X退出当前()
}
