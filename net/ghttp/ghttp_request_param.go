// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	gvar "github.com/888go/goframe/container/gvar"
	gjson "github.com/888go/goframe/encoding/gjson"
	gurl "github.com/888go/goframe/encoding/gurl"
	gxml "github.com/888go/goframe/encoding/gxml"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/utils"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gvalid "github.com/888go/goframe/util/gvalid"
)

const (
	parseTypeRequest = iota
	parseTypeQuery
	parseTypeForm
)

var (
		// xmlHeaderBytes是最常见的XML格式头。 md5:a1aeea32c5f6c441
	xmlHeaderBytes = []byte("<?xml")
)

// X解析参数到结构 是最常用的函数，它将请求参数转换为结构体或结构体切片。同时，根据结构体上的验证标签，自动对结构体或结构体切片的每个元素进行验证。
//
// 参数 `pointer` 可以是以下类型之一：*struct/*struct/*[]struct/*[]*struct。
//
// 它支持单个和多个结构体的转换：
// 1. 单个结构体时，请求内容格式如：{"id":1, "name":"john"} 或 ?id=1&name=john
// 2. 多个结构体时，请求内容格式如：[{"id":1, "name":"john"}, {"id":, "name":"smith"}]
//
// 待办事项：通过减少跨包对同一变量的重复反射使用，来提升性能。
// md5:ad971f0fee54e93d
func (r *Request) X解析参数到结构(结构指针 interface{}) error {
	return r.doParse(结构指针, parseTypeRequest)
}

// X解析URL到结构 的行为类似于 Parse 函数，但只解析查询参数。 md5:4104abbe70053960
func (r *Request) X解析URL到结构(结构指针 interface{}) error {
	return r.doParse(结构指针, parseTypeQuery)
}

// X解析表单到结构 类似于 Parse 函数，但只解析表单参数或主体内容。 md5:c384eb18ba068958
func (r *Request) X解析表单到结构(结构指针 interface{}) error {
	return r.doParse(结构指针, parseTypeForm)
}

// doParse 根据请求类型解析请求数据到结构体/结构体中。 md5:82daab462d052004
func (r *Request) doParse(pointer interface{}, requestType int) error {
	var (
		reflectVal1  = reflect.ValueOf(pointer)
		reflectKind1 = reflectVal1.Kind()
	)
	if reflectKind1 != reflect.Ptr {
		return gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
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
	// 单个结构体，帖子内容格式如下：
	// 1. {"id":1, "name":"john"} 
	// 2. ?id=1&name=john
	// md5:968f64e28941480c
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
		// 待办事项: https:		//github.com/gogf/gf/pull/2450
		// 验证。
		// md5:ec24b1494dabb977
		if err = gvalid.New().
			Bail().
			Data(pointer).
			Assoc(data).
			Run(r.Context别名()); err != nil {
			return err
		}

	// 多个结构体，它只支持像这样的JSON类型POST内容：
	// [{"id":1, "name":"john"}, {"id":2, "name":"smith"}]
	// md5:b759870b71d2ffab
	case reflect.Array, reflect.Slice:
		// 如果是结构体切片转换，可能会包含JSON/XML等内容，因此它使用`gjson`进行转换。
		// md5:e60fd34347047253
		j, err := gjson.X加载并自动识别格式(r.X取请求体字节集())
		if err != nil {
			return err
		}
		if err = j.X取泛型类().X取结构体指针(pointer); err != nil {
			return err
		}
		for i := 0; i < reflectVal2.Len(); i++ {
			if err = gvalid.New().
				Bail().
				Data(reflectVal2.Index(i)).
				Assoc(j.X取值(gconv.String(i)).X取Map()).
				Run(r.Context别名()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Get别名 是 GetRequest 的别名，它是用于检索参数的最常用函数之一。
// 请参见 r.GetRequest。
// md5:80825e01a3c06041
func (r *Request) Get别名(名称 string, 默认值 ...interface{}) *gvar.Var {
	return r.X取参数(名称, 默认值...)
}

// X取请求体字节集 读取并返回请求体内容为字节。
// 可以多次调用，每次都返回相同的正文内容。
// md5:be66d2484fd786ca
func (r *Request) X取请求体字节集() []byte {
	if r.bodyContent == nil {
		r.bodyContent = r.MakeBodyRepeatableRead(true)
	}
	return r.bodyContent
}

// MakeBodyRepeatableRead 标记请求体是否可以重复读取。它还会返回当前请求体的内容。
// md5:3cda0a2da5c712d7
func (r *Request) MakeBodyRepeatableRead(repeatableRead bool) []byte {
	if r.bodyContent == nil {
		var err error
		if r.bodyContent, err = io.ReadAll(r.Body); err != nil {
			errMsg := `Read from request Body failed`
			if gerror.X是否包含(err, io.EOF) {
				errMsg += `, the Body might be closed or read manually from middleware/hook/other package previously`
			}
			panic(gerror.X多层错误码(gcode.CodeInternalError, err, errMsg))
		}
	}
	r.Body = utils.NewReadCloser(r.bodyContent, repeatableRead)
	return r.bodyContent
}

// X取请求体文本 用于检索并返回请求体内容作为字符串。可以多次调用以获取相同的内容。
// md5:503c28317dc909ca
func (r *Request) X取请求体文本() string {
	return string(r.X取请求体字节集())
}

// X取请求体到json类 将当前请求内容解析为JSON格式，并返回JSON对象。
// 注意，请求内容是从请求体(BODY)中读取的，而不是从表单的任何字段中读取。
// md5:166af4b89b6a5a68
func (r *Request) X取请求体到json类() (*gjson.Json, error) {
	return gjson.X加载并按选项(r.X取请求体字节集(), gjson.Options{
		Type:      gjson.ContentTypeJson,
		StrNumber: true,
	})
}

// GetMap别名 是 GetRequestMap 函数的别名，提供便利的使用方式。
// 参考 GetRequestMap。
// md5:395e8bbf3fea416a
func (r *Request) GetMap别名(默认值 ...map[string]interface{}) map[string]interface{} {
	return r.X取参数到Map(默认值...)
}

// GetMapStrStr别名是GetRequestMapStrStr的别名，提供便捷的功能。详情请参阅GetRequestMapStrStr。
// md5:1828f3886ccd906d
func (r *Request) GetMapStrStr别名(默认值 ...map[string]interface{}) map[string]string {
	return r.X取参数到MapStrStr(默认值...)
}

// GetStruct别名 是 GetRequestStruct 的别名和便捷函数。详情请参阅 GetRequestStruct。
// md5:c558debb875b77cd
func (r *Request) GetStruct别名(结构指针 interface{}, mapping ...map[string]string) error {
	return r.X取参数到结构体(结构指针, mapping...)
}

// parseQuery 将查询字符串解析到 r.queryMap 中。 md5:9a26b305dc518866
func (r *Request) parseQuery() {
	if r.parsedQuery {
		return
	}
	r.parsedQuery = true
	if r.URL.RawQuery != "" {
		var err error
		r.queryMap, err = gstr.X参数解析(r.URL.RawQuery)
		if err != nil {
			panic(gerror.X多层错误码(gcode.CodeInvalidParameter, err, "Parse Query failed"))
		}
	}
}

// parseBody 将请求的原始数据解析到 r.rawMap 中。
// 请注意，它还支持从客户端请求的 JSON 数据。
// md5:f8f001deccef59e6
func (r *Request) parseBody() {
	if r.parsedBody {
		return
	}
	r.parsedBody = true
		// 没有提交任何数据。 md5:cf70840053024c2b
	if r.ContentLength == 0 {
		return
	}
	if body := r.X取请求体字节集(); len(body) > 0 {
				// 去除空格和换行符。 md5:0cf77adc8fee1e9a
		body = bytes.TrimSpace(body)
		// JSON format checks.
		if body[0] == '{' && body[len(body)-1] == '}' {
			_ = json.UnmarshalUseNumber(body, &r.bodyMap)
		}
		// XML format checks.
		if len(body) > 5 && bytes.EqualFold(body[:5], xmlHeaderBytes) {
			r.bodyMap, _ = gxml.DecodeWithoutRoot(body)
		}
		if body[0] == '<' && body[len(body)-1] == '>' {
			r.bodyMap, _ = gxml.DecodeWithoutRoot(body)
		}
				// 默认参数解码。 md5:941d9de3ebb46554
		if contentType := r.Header.Get("Content-Type"); (contentType == "" || !gstr.X是否包含(contentType, "multipart/")) && r.bodyMap == nil {
			r.bodyMap, _ = gstr.X参数解析(r.X取请求体文本())
		}
	}
}

// parseForm 解析HTTP方法PUT，POST，PATCH的请求表单。
// 表单数据被解析到r.formMap中。
//
// 请注意，如果已经先解析了表单，那么请求体将会被清空。
// md5:97f04aa06758375b
func (r *Request) parseForm() {
	if r.parsedForm {
		return
	}
	r.parsedForm = true
		// 没有提交任何数据。 md5:cf70840053024c2b
	if r.ContentLength == 0 {
		return
	}
	if contentType := r.Header.Get("Content-Type"); contentType != "" {
		var (
			err            error
			repeatableRead = true
		)
		if gstr.X是否包含(contentType, "multipart/") {
			// 为了避免大量消耗内存。
			// `multipart/` 类型的表单始终包含二进制数据，没有必要读取两次。
			// md5:d95befcac4fa7fd0
			repeatableRead = false
						// 这两个注释是在描述MIME类型。`multipart/form-data`通常用于通过HTTP发送表单数据，如文件上传。`multipart/mixed`则用于包含多个部分的混合内容，每个部分可以是不同的MIME类型，常用于邮件或API请求中包含多种类型的附件或数据。 md5:5f5a1e86722f47ec
			if err = r.ParseMultipartForm(r.X服务.config.FormParsingMemory); err != nil {
				panic(gerror.X多层错误码(gcode.CodeInvalidRequest, err, "r.ParseMultipartForm failed"))
			}
		} else if gstr.X是否包含(contentType, "form") {
									// 应用程序/x-www-form-urlencoded. md5:6de553b2a7019beb
			if err = r.Request.ParseForm(); err != nil {
				panic(gerror.X多层错误码(gcode.CodeInvalidRequest, err, "r.Request.ParseForm failed"))
			}
		}
		if repeatableRead {
			r.MakeBodyRepeatableRead(true)
		}
		if len(r.PostForm) > 0 {
						// 使用统一的解析方式解析表单数据。 md5:21f3f94370e18b5d
			params := ""
			for name, values := range r.PostForm {
				// 非法的参数名称。
				// 只允许使用以下字符：`\w`，`[`，`]`，`-`。
				// md5:72a7ff7f2d38a973
				if !gregex.X是否匹配文本(`^[\w\-\[\]]+$`, name) && len(r.PostForm) == 1 {
										// 它可能是JSON或XML内容。 md5:105b844bbc2857c0
					if s := gstr.X过滤首尾符并含空白(name + strings.Join(values, " ")); len(s) > 0 {
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
					params += name + "=" + gurl.X编码(values[0])
				} else {
					if len(name) > 2 && name[len(name)-2:] == "[]" {
						name = name[:len(name)-2]
						for _, v := range values {
							if len(params) > 0 {
								params += "&"
							}
							params += name + "[]=" + gurl.X编码(v)
						}
					} else {
						if len(params) > 0 {
							params += "&"
						}
						params += name + "=" + gurl.X编码(values[len(values)-1])
					}
				}
			}
			if params != "" {
				if r.formMap, err = gstr.X参数解析(params); err != nil {
					panic(gerror.X多层错误码(gcode.CodeInvalidParameter, err, "Parse request parameters failed"))
				}
			}
		}
	}
		// 它解析请求体，而不检查Content-Type。 md5:89cfec67836d4575
	if r.formMap == nil {
		if r.Method != http.MethodGet {
			r.parseBody()
		}
		if len(r.bodyMap) > 0 {
			r.formMap = r.bodyMap
		}
	}
}

// X取multipart表单对象 解析并返回表单为多部分形式。 md5:c80c641ed3887bea
func (r *Request) X取multipart表单对象() *multipart.Form {
	r.parseForm()
	return r.MultipartForm
}

// X取multipart表单文件切片对象 解析并返回表单中的文件数组。
// 请注意，请求表单的类型应该是multipart。
// md5:33503fc76a60c149
func (r *Request) X取multipart表单文件切片对象(名称 string) []*multipart.FileHeader {
	form := r.X取multipart表单对象()
	if form == nil {
		return nil
	}
	if v := form.File[名称]; len(v) > 0 {
		return v
	}
		// 支持" name[]"作为数组参数。 md5:f1460d96fee37609
	if v := form.File[名称+"[]"]; len(v) > 0 {
		return v
	}
		// 支持将"name[0]","name[1]","name[2]"等作为数组参数使用。 md5:a9545b3b88169505
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
