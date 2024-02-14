// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package http类

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/encoding/gurl"
	"github.com/888go/goframe/encoding/gxml"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/text/gregex"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gvalid"
)

const (
	parseTypeRequest = 0
	parseTypeQuery   = 1
	parseTypeForm    = 2
)

var (
	// xmlHeaderBytes 是最常见的 XML 格式头部。
	xmlHeaderBytes = []byte("<?xml")
)

// Parse 是最常用的函数，用于将请求参数转换为结构体或结构体切片。
// 同时，它也会根据结构体上的验证标签自动校验结构体或结构体切片中的每个元素。
//
// 参数 `pointer` 可以是以下类型：*struct/**struct/*[]struct/*[]*struct。
//
// 它支持单个和多个结构体的转换：
// 1. 单个结构体，POST 内容如：{"id":1, "name":"john"} 或 ?id=1&name=john
// 2. 多个结构体，POST 内容如：[{"id":1, "name":"john"}, {"id":, "name":"smith"}]
//
// TODO: 通过减少在不同包中对同一变量重复使用 reflect 来提高性能。
func (r *X请求) X解析参数到结构(结构指针 interface{}) error {
	return r.doParse(结构指针, parseTypeRequest)
}

// ParseQuery 类似于函数 Parse，但它只解析查询参数。
func (r *X请求) X解析URL到结构(结构指针 interface{}) error {
	return r.doParse(结构指针, parseTypeQuery)
}

// ParseForm执行类似于函数Parse的功能，但只解析表单参数或主体内容。
func (r *X请求) X解析表单到结构(结构指针 interface{}) error {
	return r.doParse(结构指针, parseTypeForm)
}

// doParse 根据请求类型将请求数据解析到结构体/结构体数组中。
func (r *X请求) doParse(pointer interface{}, requestType int) error {
	var (
		reflectVal1  = reflect.ValueOf(pointer)
		reflectKind1 = reflectVal1.Kind()
	)
	if reflectKind1 != reflect.Ptr {
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`invalid parameter type "%v", of which kind should be of *struct/**struct/*[]struct/*[]*struct, but got: "%v"`,
			reflectVal1.Type(),
			reflectKind1,
		)
	}
	var (
		reflectVal2  = reflectVal1.Elem()
		reflectKind2 = reflectVal2.Kind()
	)
	switch reflectKind2 {
// Single 结构体，用于处理如下的POST内容：
// 1. {"id":1, "name":"john"} 
// 2. ?id=1&name=john
	case reflect.Ptr, reflect.Struct:
		var (
			err  error
			data map[string]interface{}
		)
		// Converting.
		switch requestType {
		case parseTypeQuery:
			if data, err = r.doGetQueryStruct(pointer); err != nil {
				return err
			}
		case parseTypeForm:
			if data, err = r.doGetFormStruct(pointer); err != nil {
				return err
			}
		default:
			if data, err = r.doGetRequestStruct(pointer); err != nil {
				return err
			}
		}
// TODO: 参考 https://github.com/gogf/gf/pull/2450
// 验证功能。
		if err = 效验类.New().
			Bail().
			Data(pointer).
			Assoc(data).
			Run(r.Context别名()); err != nil {
			return err
		}

// 多条数据结构，它仅支持类似以下JSON格式的POST内容：
// [{"id":1, "name":"john"}, {"id":, "name":"smith"}]
	case reflect.Array, reflect.Slice:
// 如果是结构体切片转换，它可能会发布 JSON/XML/... 格式的内容，
// 因此这里使用 `gjson` 进行转换。
		j, err := json类.X加载并自动识别格式(r.X取请求体字节集())
		if err != nil {
			return err
		}
		if err = j.X取泛型类().X取结构体指针(pointer); err != nil {
			return err
		}
		for i := 0; i < reflectVal2.Len(); i++ {
			if err = 效验类.New().
				Bail().
				Data(reflectVal2.Index(i)).
				Assoc(j.X取值(转换类.String(i)).X取Map()).
				Run(r.Context别名()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Get 是 GetRequest 的别名，它是用于获取参数的最常用函数之一。
// 请参阅 r.GetRequest。
func (r *X请求) Get别名(名称 string, 默认值 ...interface{}) *泛型类.Var {
	return r.X取参数(名称, 默认值...)
}

// GetBody 用于检索并返回请求正文内容作为字节。
// 它可以被多次调用，获取相同正文内容。
func (r *X请求) X取请求体字节集() []byte {
	if r.bodyContent == nil {
		r.bodyContent = r.MakeBodyRepeatableRead(true)
	}
	return r.bodyContent
}

// MakeBodyRepeatableRead 标记请求体是否可以被重复读取。
// 同时，它还会返回当前请求体的内容。
func (r *X请求) MakeBodyRepeatableRead(repeatableRead bool) []byte {
	if r.bodyContent == nil {
		var err error
		if r.bodyContent, err = io.ReadAll(r.Body); err != nil {
			errMsg := `Read from request Body failed`
			if 错误类.X是否包含(err, io.EOF) {
				errMsg += `, the Body might be closed or read manually from middleware/hook/other package previously`
			}
			panic(错误类.X多层错误码(错误码类.CodeInternalError, err, errMsg))
		}
	}
	r.Body = utils.NewReadCloser(r.bodyContent, repeatableRead)
	return r.bodyContent
}

// GetBodyString 用于获取并返回请求体内容作为字符串。
// 它可以被多次调用，以获取相同的请求体内容。
func (r *X请求) X取请求体文本() string {
	return string(r.X取请求体字节集())
}

// GetJson 将当前请求内容解析为JSON格式，并返回JSON对象。
// 注意：请求内容是从request BODY中读取，而不是从FORM的任何字段中读取。
func (r *X请求) X取请求体到json类() (*json类.Json, error) {
	return json类.X加载并按选项(r.X取请求体字节集(), json类.Options{
		Type:      json类.ContentTypeJson,
		StrNumber: true,
	})
}

// GetMap 是一个别名，也是一个方便获取请求映射的函数。
// 请参阅 GetRequestMap。
func (r *X请求) GetMap别名(默认值 ...map[string]interface{}) map[string]interface{} {
	return r.X取参数到Map(默认值...)
}

// GetMapStrStr 是 GetRequestMapStrStr 的别名和便捷函数。
// 请参阅 GetRequestMapStrStr。
func (r *X请求) GetMapStrStr别名(默认值 ...map[string]interface{}) map[string]string {
	return r.X取参数到MapStrStr(默认值...)
}

// GetStruct 是 GetRequestStruct 的别名和便捷函数。
// 请参阅 GetRequestStruct。
func (r *X请求) GetStruct别名(结构指针 interface{}, mapping ...map[string]string) error {
	return r.X取参数到结构体(结构指针, mapping...)
}

// parseQuery 将查询字符串解析到 r.queryMap 中。
func (r *X请求) parseQuery() {
	if r.parsedQuery {
		return
	}
	r.parsedQuery = true
	if r.URL.RawQuery != "" {
		var err error
		r.queryMap, err = 文本类.X参数解析(r.URL.RawQuery)
		if err != nil {
			panic(错误类.X多层错误码(错误码类.CodeInvalidParameter, err, "Parse Query failed"))
		}
	}
}

// parseBody 将请求原始数据解析到 r.rawMap 中。
// 注意，它还支持从客户端请求的 JSON 数据。
func (r *X请求) parseBody() {
	if r.parsedBody {
		return
	}
	r.parsedBody = true
	// 没有提交任何数据。
	if r.ContentLength == 0 {
		return
	}
	if body := r.X取请求体字节集(); len(body) > 0 {
		// 去除空格/换行符。
		body = bytes.TrimSpace(body)
		// JSON格式检查。
		if body[0] == '{' && body[len(body)-1] == '}' {
			_ = json.UnmarshalUseNumber(body, &r.bodyMap)
		}
		// XML格式检查。
		if len(body) > 5 && bytes.EqualFold(body[:5], xmlHeaderBytes) {
			r.bodyMap, _ = xml类.DecodeWithoutRoot(body)
		}
		if body[0] == '<' && body[len(body)-1] == '>' {
			r.bodyMap, _ = xml类.DecodeWithoutRoot(body)
		}
		// 默认参数解码
		if contentType := r.Header.Get("Content-Type"); (contentType == "" || !文本类.X是否包含(contentType, "multipart/")) && r.bodyMap == nil {
			r.bodyMap, _ = 文本类.X参数解析(r.X取请求体文本())
		}
	}
}

// parseForm 用于解析HTTP方法PUT、POST、PATCH的请求表单。
// 表单数据将被解析并存储到r.formMap中。
//
// 注意：如果表单首先被解析，请求体将会被清空，变成空内容。
func (r *X请求) parseForm() {
	if r.parsedForm {
		return
	}
	r.parsedForm = true
	// 没有提交任何数据。
	if r.ContentLength == 0 {
		return
	}
	if contentType := r.Header.Get("Content-Type"); contentType != "" {
		var err error
		if 文本类.X是否包含(contentType, "multipart/") {
			// multipart/form-data：这是一种HTTP内容类型，用于编码同一条请求中包含多种不同类型数据（如文本、文件等）的表单数据。在上传文件时尤为常见。
// multipart/mixed：这也是一种HTTP内容类型，用于表示消息体包含多个独立的部分，各个部分可以是不同的数据类型，且每个部分都有自己的Content-Type和边界标识符。它通常用于混合多部分消息，比如在一个HTTP请求中同时发送文本信息和附件。
			if err = r.ParseMultipartForm(r.X服务.config.X表单解析最大缓冲区长度); err != nil {
				panic(错误类.X多层错误码(错误码类.CodeInvalidRequest, err, "r.ParseMultipartForm failed"))
			}
		} else if 文本类.X是否包含(contentType, "form") {
			// application/x-www-form-urlencoded 是一种常见的HTTP内容类型，用于表示URL编码的表单数据。在Go语言中，通常在网络请求或表单提交时使用这种格式对键值对进行编码。
// ```go
// 这是HTTP请求Header中Content-Type的一种常见取值
// 表示请求体中的数据采用了"application/x-www-form-urlencoded"编码格式
// 该格式将表单字段名和值连接成键值对，并且对特殊字符进行URL编码
			if err = r.Request.ParseForm(); err != nil {
				panic(错误类.X多层错误码(错误码类.CodeInvalidRequest, err, "r.Request.ParseForm failed"))
			}
		}
		if len(r.PostForm) > 0 {
			// 使用统一解析方式解析表单数据。
			params := ""
			for name, values := range r.PostForm {
// 无效的参数名称。
// 只允许包含字符：'\w'（代表字母、数字或下划线）、'['、']' 和 '-'。
				if !正则类.X是否匹配文本(`^[\w\-\[\]]+$`, name) && len(r.PostForm) == 1 {
					// 这可能是一个JSON/XML内容。
					if s := 文本类.X过滤首尾符并含空白(name + strings.Join(values, " ")); len(s) > 0 {
						if s[0] == '{' && s[len(s)-1] == '}' || s[0] == '<' && s[len(s)-1] == '>' {
							r.bodyContent = []byte(s)
							params = ""
							break
						}
					}
				}
				if len(values) == 1 {
					if len(params) > 0 {
						params += "&"
					}
					params += name + "=" + url类.X编码(values[0])
				} else {
					if len(name) > 2 && name[len(name)-2:] == "[]" {
						name = name[:len(name)-2]
						for _, v := range values {
							if len(params) > 0 {
								params += "&"
							}
							params += name + "[]=" + url类.X编码(v)
						}
					} else {
						if len(params) > 0 {
							params += "&"
						}
						params += name + "=" + url类.X编码(values[len(values)-1])
					}
				}
			}
			if params != "" {
				if r.formMap, err = 文本类.X参数解析(params); err != nil {
					panic(错误类.X多层错误码(错误码类.CodeInvalidParameter, err, "Parse request parameters failed"))
				}
			}
		}
	}
	// 它解析请求体，但不检查 Content-Type。
	if r.formMap == nil {
		if r.Method != http.MethodGet {
			r.parseBody()
		}
		if len(r.bodyMap) > 0 {
			r.formMap = r.bodyMap
		}
	}
}

// GetMultipartForm 解析并返回表单为多部分表单形式。
func (r *X请求) X取multipart表单对象() *multipart.Form {
	r.parseForm()
	return r.MultipartForm
}

// GetMultipartFiles 解析并返回 POST 请求中的文件数组。
// 注意，请求表单的类型应当为 multipart。
func (r *X请求) X取multipart表单文件数组对象(名称 string) []*multipart.FileHeader {
	form := r.X取multipart表单对象()
	if form == nil {
		return nil
	}
	if v := form.File[名称]; len(v) > 0 {
		return v
	}
	// 支持 "name[]" 作为数组参数。
	if v := form.File[名称+"[]"]; len(v) > 0 {
		return v
	}
	// 支持 "name[0]"、"name[1]"、"name[2]" 等形式的数组参数。
	var (
		key   string
		files = make([]*multipart.FileHeader, 0)
	)
	for i := 0; ; i++ {
		key = fmt.Sprintf(`%s[%d]`, 名称, i)
		if v := form.File[key]; len(v) > 0 {
			files = append(files, v[0])
		} else {
			break
		}
	}
	if len(files) > 0 {
		return files
	}
	return nil
}
