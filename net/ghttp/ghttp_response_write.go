// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package ghttp
import (
	"fmt"
	"net/http"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
	)
// Write将`content`写入响应缓冲区。
func (r *Response) Write(content ...interface{}) {
	if r.writer.IsHijacked() || len(content) == 0 {
		return
	}
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	for _, v := range content {
		switch value := v.(type) {
		case []byte:
			r.buffer.Write(value)
		case string:
			r.buffer.WriteString(value)
		default:
			r.buffer.WriteString(gconv.String(v))
		}
	}
}

// WriteExit 将`content`写入响应缓冲区并退出当前处理器的执行。
// “Exit”特性通常用于为了方便起见，在处理器中替换 return 语句的使用。
func (r *Response) WriteExit(content ...interface{}) {
	r.Write(content...)
	r.Request.Exit()
}

// WriteOver 将`content`覆盖写入响应缓冲区。
func (r *Response) WriteOver(content ...interface{}) {
	r.ClearBuffer()
	r.Write(content...)
}

// WriteOverExit 用 `content` 覆盖响应缓冲区并退出当前处理器的执行。  
// "Exit" 特性通常用于为了方便起见，在处理器中替换 return 语句的使用。
func (r *Response) WriteOverExit(content ...interface{}) {
	r.WriteOver(content...)
	r.Request.Exit()
}

// Writef 使用 fmt.Sprintf 方法写入响应内容。
func (r *Response) Writef(format string, params ...interface{}) {
	r.Write(fmt.Sprintf(format, params...))
}

// WritefExit 通过 fmt.Sprintf 写入响应，并退出当前处理器的执行。
// "Exit" 特性通常用于为了方便，在处理器中替代 return 语句的使用。
func (r *Response) WritefExit(format string, params ...interface{}) {
	r.Writef(format, params...)
	r.Request.Exit()
}

// Writeln将`content`内容和换行符一起写入响应。
func (r *Response) Writeln(content ...interface{}) {
	if len(content) == 0 {
		r.Write("\n")
		return
	}
	r.Write(append(content, "\n")...)
}

// WritelnExit 将`content`内容及换行符写入响应，并终止当前处理器的执行。
// "Exit"特性通常用于为了方便起见，替代处理器中return语句的使用。
func (r *Response) WritelnExit(content ...interface{}) {
	r.Writeln(content...)
	r.Request.Exit()
}

// Writefln 使用 fmt.Sprintf 格式化输出并将内容与换行符写入响应。
func (r *Response) Writefln(format string, params ...interface{}) {
	r.Writeln(fmt.Sprintf(format, params...))
}

// WriteflnExit 通过 fmt.Sprintf 和换行符写出响应，并退出当前处理器的执行。
// "Exit" 特性通常用于为了方便，而替换处理器中 return 语句的使用。
func (r *Response) WriteflnExit(format string, params ...interface{}) {
	r.Writefln(format, params...)
	r.Request.Exit()
}

// WriteJson 将`content`以JSON格式写入响应中。
func (r *Response) WriteJson(content interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
	// 如果给定的是字符串或[]byte，直接将响应发送回客户端。
	switch content.(type) {
	case string, []byte:
		r.Write(gconv.String(content))
		return
	}
	// 否则使用json.Marshal函数对参数进行编码。
	if b, err := json.Marshal(content); err != nil {
		panic(gerror.Wrap(err, `WriteJson failed`))
	} else {
		r.Write(b)
	}
}

// WriteJsonExit将`content`以JSON格式写入响应，并在成功时退出当前处理器的执行。
// “Exit”特性通常用于为了方便起见，替换处理器中return语句的使用。
func (r *Response) WriteJsonExit(content interface{}) {
	r.WriteJson(content)
	r.Request.Exit()
}

// WriteJsonP 将`content`以JSONP格式写入响应中。
//
// 注意：对于JSONP格式，请求中应包含一个"callback"参数。
func (r *Response) WriteJsonP(content interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
	// 如果给定的是字符串或[]byte，直接将响应发送给客户端。
	switch content.(type) {
	case string, []byte:
		r.Write(gconv.String(content))
		return
	}
	// 否则使用json.Marshal函数对参数进行编码。
	if b, err := json.Marshal(content); err != nil {
		panic(gerror.Wrap(err, `WriteJsonP failed`))
	} else {
		// 设置HTTP响应头的"Content-Type"字段为"application/json"
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

// WriteJsonPExit 将`content`以JSONP格式写入响应，并在成功时退出当前处理器的执行。
// “Exit”特性通常用于替换处理器中return语句的使用，以便于简化代码。
//
// 注意，请求中应包含一个“callback”参数以适应JSONP格式。
func (r *Response) WriteJsonPExit(content interface{}) {
	r.WriteJsonP(content)
	r.Request.Exit()
}

// WriteXml 将`content`以XML格式写入响应。
func (r *Response) WriteXml(content interface{}, rootTag ...string) {
	r.Header().Set("Content-Type", contentTypeXml)
	// 如果给定的是字符串或[]byte，直接将其响应给客户端。
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

// WriteXmlExit将`content`以XML格式写入响应，并在成功时退出当前处理器的执行。
// "Exit"特性通常用于为了方便起见，在处理器中替换return语句的使用。
func (r *Response) WriteXmlExit(content interface{}, rootTag ...string) {
	r.WriteXml(content, rootTag...)
	r.Request.Exit()
}

// WriteStatus将HTTP状态码`status`和内容`content`写入响应中。
// 注意，这里没有设置Content-Type头信息。
func (r *Response) WriteStatus(status int, content ...interface{}) {
	r.WriteHeader(status)
	if len(content) > 0 {
		r.Write(content...)
	} else {
		r.Write(http.StatusText(status))
	}
}

// WriteStatusExit 将HTTP状态码`status`和内容`content`写入响应，并在成功时退出当前处理器的执行。
// "Exit"特性通常用于为了方便，替换处理器中return语句的使用。
func (r *Response) WriteStatusExit(status int, content ...interface{}) {
	r.WriteStatus(status, content...)
	r.Request.Exit()
}
