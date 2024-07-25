// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31

package ghttp

import (
	"fmt"
	"net/http"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/util/gconv"
)

// Write 将 `content` 写入响应缓冲区。 md5:2f656734fbf8eab6
func (r *Response) Write(content ...interface{}) {
	if r.IsHijacked() || len(content) == 0 {
		return
	}
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	for _, v := range content {
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

// WriteExit 将`content`写入响应缓冲区并退出当前处理器的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。 md5:afcb2dda1beb9358
func (r *Response) WriteExit(content ...interface{}) {
	r.Write(content...)
	r.Request.Exit()
}

// WriteOver 将响应缓冲区用`content`重写。 md5:d68f13dc57329ab0
func (r *Response) WriteOver(content ...interface{}) {
	r.ClearBuffer()
	r.Write(content...)
}

// WriteOverExit 将响应缓冲区用 `content` 替换，然后退出当前处理程序的执行。"Exit" 功能通常用于方便地替换处理程序中的返回语句。 md5:968d387aea44eeab
func (r *Response) WriteOverExit(content ...interface{}) {
	r.WriteOver(content...)
	r.Request.Exit()
}

// Writef 使用 fmt.Sprintf 格式化字符串并写出响应。 md5:15163b759bd146b8
func (r *Response) Writef(format string, params ...interface{}) {
	r.Write(fmt.Sprintf(format, params...))
}

// WritefExit使用fmt.Sprintf格式化响应内容并退出当前处理器的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。 md5:01275db804fa4029
func (r *Response) WritefExit(format string, params ...interface{}) {
	r.Writef(format, params...)
	r.Request.Exit()
}

// Writeln 使用`content`和换行符写入响应。 md5:574e18a271a92e20
func (r *Response) Writeln(content ...interface{}) {
	if len(content) == 0 {
		r.Write("\n")
		return
	}
	r.Write(append(content, "\n")...)
}

// WritelnExit 写入包含`content`和换行符的响应，然后退出当前处理器的执行。"Exit"特性通常用于方便地替换处理器中的返回语句。 md5:bb5f123bedaec380
func (r *Response) WritelnExit(content ...interface{}) {
	r.Writeln(content...)
	r.Request.Exit()
}

// Writefln 使用 fmt.Sprintf 格式化输出并将结果作为响应写入，同时添加换行符。 md5:154e6f5f52878f00
func (r *Response) Writefln(format string, params ...interface{}) {
	r.Writeln(fmt.Sprintf(format, params...))
}

// WriteflnExit使用fmt.Sprintf格式化响应并添加换行符，然后退出当前处理程序的执行。
// "Exit"功能常用于替代处理器中的返回语句，以提供便利。 md5:ee5288e61cdea4b2
func (r *Response) WriteflnExit(format string, params ...interface{}) {
	r.Writefln(format, params...)
	r.Request.Exit()
}

// WriteJson 将`content`以JSON格式写入到响应中。 md5:0ca8d5da1805456f
func (r *Response) WriteJson(content interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
	// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:e14783864a1068a9
	switch content.(type) {
	case string, []byte:
		r.Write(gconv.String(content))
		return
	}
	// 否则，使用json.Marshal函数对参数进行编码。 md5:b140f4be3fab1fa1
	if b, err := json.Marshal(content); err != nil {
		panic(gerror.Wrap(err, `WriteJson failed`))
	} else {
		r.Write(b)
	}
}

// WriteJsonExit 将`content`以JSON格式写入响应，并在成功时退出当前处理器的执行。"Exit"功能常用于替代处理器中的返回语句，以提供便利。 md5:0714d99528fcb93e
func (r *Response) WriteJsonExit(content interface{}) {
	r.WriteJson(content)
	r.Request.Exit()
}

// WriteJsonP 将`content`按照JSONP格式写入响应。
//
// 注意，对于JSONP格式，请求中应该包含一个名为"callback"的参数。 md5:32a3e4fa6b4e92b0
func (r *Response) WriteJsonP(content interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
	// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:ff82edfcddee9a78
	switch content.(type) {
	case string, []byte:
		r.Write(gconv.String(content))
		return
	}
	// 否则，使用json.Marshal函数对参数进行编码。 md5:b140f4be3fab1fa1
	if b, err := json.Marshal(content); err != nil {
		panic(gerror.Wrap(err, `WriteJsonP failed`))
	} else {
		// 设置HTTP响应头中的"Content-Type"为"application/json". md5:8e0be5eb3c232d44
		if callback := r.Request.Get("callback").String(); callback != "" {
			buffer := []byte(callback)
			buffer = append(buffer, byte('('))
			buffer = append(buffer, b...)
			buffer = append(buffer, byte(')'))
			r.Write(buffer)
		} else {
			r.Write(b)
		}
	}
}

// WriteJsonPExit 将 `content` 以 JSONP 格式写入响应，并在成功时退出当前处理器的执行。"Exit" 功能常用于替代处理器中的返回语句，以提供便利。
//
// 请注意，为了使用 JSONP 格式，请求中应该包含一个 "callback" 参数。 md5:6c959e76945e075a
func (r *Response) WriteJsonPExit(content interface{}) {
	r.WriteJsonP(content)
	r.Request.Exit()
}

// WriteXml 将`content`以XML格式写入响应。 md5:850a872cf25d6a70
func (r *Response) WriteXml(content interface{}, rootTag ...string) {
	r.Header().Set("Content-Type", contentTypeXml)
	// 如果给定字符串/字节切片，直接将其返回给客户端。 md5:4fc9d6cf062e5bf9
	switch content.(type) {
	case string, []byte:
		r.Write(gconv.String(content))
		return
	}
	if b, err := gjson.New(content).ToXml(rootTag...); err != nil {
		panic(gerror.Wrap(err, `WriteXml failed`))
	} else {
		r.Write(b)
	}
}

// WriteXmlExit 将`content`以XML格式写入响应，并在成功时退出当前处理器的执行。
// “退出”功能常用于便捷地替代处理器中return语句的使用。 md5:12a9a20328b00f55
func (r *Response) WriteXmlExit(content interface{}, rootTag ...string) {
	r.WriteXml(content, rootTag...)
	r.Request.Exit()
}

// WriteStatus 将HTTP状态码`status`和内容`content`写入响应。
// 请注意，它不会在这里设置Content-Type头。 md5:8b7195f02ad8ced0
func (r *Response) WriteStatus(status int, content ...interface{}) {
	r.WriteHeader(status)
	if len(content) > 0 {
		r.Write(content...)
	} else {
		r.Write(http.StatusText(status))
	}
}

// WriteStatusExit 将HTTP状态码`status`和内容`content`写入响应，并在成功时退出当前处理程序的执行。"Exit"特性通常用于方便地替代处理程序中返回语句的使用。 md5:2e5cbf96316a0c3c
func (r *Response) WriteStatusExit(status int, content ...interface{}) {
	r.WriteStatus(status, content...)
	r.Request.Exit()
}
