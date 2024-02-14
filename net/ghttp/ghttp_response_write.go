// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package http类

import (
	"fmt"
	"net/http"
	
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/util/gconv"
)

// Write将`content`写入响应缓冲区。
func (r *X响应) X写响应缓冲区(内容 ...interface{}) {
	if r.writer.IsHijacked() || len(内容) == 0 {
		return
	}
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	for _, v := range 内容 {
		switch value := v.(type) {
		case []byte:
			r.buffer.Write(value)
		case string:
			r.buffer.WriteString(value)
		default:
			r.buffer.WriteString(转换类.String(v))
		}
	}
}

// WriteExit 将`content`写入响应缓冲区并退出当前处理器的执行。
// “Exit”特性通常用于为了方便起见，在处理器中替换 return 语句的使用。
func (r *X响应) X写响应缓冲区并退出(内容 ...interface{}) {
	r.X写响应缓冲区(内容...)
	r.Request.X退出当前()
}

// WriteOver 将`content`覆盖写入响应缓冲区。
func (r *X响应) X写覆盖响应缓冲区(内容 ...interface{}) {
	r.X清空缓冲区()
	r.X写响应缓冲区(内容...)
}

// WriteOverExit 用 `content` 覆盖响应缓冲区并退出当前处理器的执行。  
// "Exit" 特性通常用于为了方便起见，在处理器中替换 return 语句的使用。
func (r *X响应) X写覆盖响应缓冲区并退出(内容 ...interface{}) {
	r.X写覆盖响应缓冲区(内容...)
	r.Request.X退出当前()
}

// Writef 使用 fmt.Sprintf 方法写入响应内容。
func (r *X响应) X写响应缓冲区并格式化(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区(fmt.Sprintf(格式, 内容...))
}

// WritefExit 通过 fmt.Sprintf 写入响应，并退出当前处理器的执行。
// "Exit" 特性通常用于为了方便，在处理器中替代 return 语句的使用。
func (r *X响应) X写响应缓冲区并退出与格式化(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区并格式化(格式, 内容...)
	r.Request.X退出当前()
}

// Writeln将`content`内容和换行符一起写入响应。
func (r *X响应) X写响应缓冲区并换行(内容 ...interface{}) {
	if len(内容) == 0 {
		r.X写响应缓冲区("\n")
		return
	}
	r.X写响应缓冲区(append(内容, "\n")...)
}

// WritelnExit 将`content`内容及换行符写入响应，并终止当前处理器的执行。
// "Exit"特性通常用于为了方便起见，替代处理器中return语句的使用。
func (r *X响应) X写响应缓冲区并退出与换行(内容 ...interface{}) {
	r.X写响应缓冲区并换行(内容...)
	r.Request.X退出当前()
}

// Writefln 使用 fmt.Sprintf 格式化输出并将内容与换行符写入响应。
func (r *X响应) X写响应缓冲区并格式化与换行(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区并换行(fmt.Sprintf(格式, 内容...))
}

// WriteflnExit 通过 fmt.Sprintf 和换行符写出响应，并退出当前处理器的执行。
// "Exit" 特性通常用于为了方便，而替换处理器中 return 语句的使用。
func (r *X响应) X写响应缓冲区并退出与格式化换行(格式 string, 内容 ...interface{}) {
	r.X写响应缓冲区并格式化与换行(格式, 内容...)
	r.Request.X退出当前()
}

// WriteJson 将`content`以JSON格式写入响应中。
func (r *X响应) X写响应缓冲区JSON(内容 interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
	// 如果给定的是字符串或[]byte，直接将响应发送回客户端。
	switch 内容.(type) {
	case string, []byte:
		r.X写响应缓冲区(转换类.String(内容))
		return
	}
	// 否则使用json.Marshal函数对参数进行编码。
	if b, err := json.Marshal(内容); err != nil {
		panic(错误类.X多层错误(err, `WriteJson failed`))
	} else {
		r.X写响应缓冲区(b)
	}
}

// WriteJsonExit将`content`以JSON格式写入响应，并在成功时退出当前处理器的执行。
// “Exit”特性通常用于为了方便起见，替换处理器中return语句的使用。
func (r *X响应) X写响应缓冲区JSON并退出(内容 interface{}) {
	r.X写响应缓冲区JSON(内容)
	r.Request.X退出当前()
}

// WriteJsonP 将`content`以JSONP格式写入响应中。
//
// 注意：对于JSONP格式，请求中应包含一个"callback"参数。
func (r *X响应) X写响应缓冲区JSONP(内容 interface{}) {
	r.Header().Set("Content-Type", contentTypeJson)
	// 如果给定的是字符串或[]byte，直接将响应发送给客户端。
	switch 内容.(type) {
	case string, []byte:
		r.X写响应缓冲区(转换类.String(内容))
		return
	}
	// 否则使用json.Marshal函数对参数进行编码。
	if b, err := json.Marshal(内容); err != nil {
		panic(错误类.X多层错误(err, `WriteJsonP failed`))
	} else {
		// 设置HTTP响应头的"Content-Type"字段为"application/json"
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

// WriteJsonPExit 将`content`以JSONP格式写入响应，并在成功时退出当前处理器的执行。
// “Exit”特性通常用于替换处理器中return语句的使用，以便于简化代码。
//
// 注意，请求中应包含一个“callback”参数以适应JSONP格式。
func (r *X响应) X写响应缓冲区JSONP并退出(内容 interface{}) {
	r.X写响应缓冲区JSONP(内容)
	r.Request.X退出当前()
}

// WriteXml 将`content`以XML格式写入响应。
func (r *X响应) X写响应缓冲区XML(内容 interface{}, 根标记 ...string) {
	r.Header().Set("Content-Type", contentTypeXml)
	// 如果给定的是字符串或[]byte，直接将其响应给客户端。
	switch 内容.(type) {
	case string, []byte:
		r.X写响应缓冲区(转换类.String(内容))
		return
	}
	if b, err := json类.X创建(内容).X取xml字节集(根标记...); err != nil {
		panic(错误类.X多层错误(err, `WriteXml failed`))
	} else {
		r.X写响应缓冲区(b)
	}
}

// WriteXmlExit将`content`以XML格式写入响应，并在成功时退出当前处理器的执行。
// "Exit"特性通常用于为了方便起见，在处理器中替换return语句的使用。
func (r *X响应) X写响应缓冲区XML并退出(内容 interface{}, 根标记 ...string) {
	r.X写响应缓冲区XML(内容, 根标记...)
	r.Request.X退出当前()
}

// WriteStatus将HTTP状态码`status`和内容`content`写入响应中。
// 注意，这里没有设置Content-Type头信息。
func (r *X响应) X写响应缓冲区与HTTP状态码(状态码 int, 内容 ...interface{}) {
	r.WriteHeader(状态码)
	if len(内容) > 0 {
		r.X写响应缓冲区(内容...)
	} else {
		r.X写响应缓冲区(http.StatusText(状态码))
	}
}

// WriteStatusExit 将HTTP状态码`status`和内容`content`写入响应，并在成功时退出当前处理器的执行。
// "Exit"特性通常用于为了方便，替换处理器中return语句的使用。
func (r *X响应) X写响应缓冲区与HTTP状态码并退出(状态码 int, 内容 ...interface{}) {
	r.X写响应缓冲区与HTTP状态码(状态码, 内容...)
	r.Request.X退出当前()
}
