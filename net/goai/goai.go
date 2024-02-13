// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式受 MIT 许可协议条款约束。
// 如果随此文件未分发 MIT 许可协议副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package goai 实现并提供了针对 OpenApi 规范的文档生成功能。
//
// 参考链接：https://editor.swagger.io/
package goai

import (
	"context"
	"fmt"
	"reflect"
	
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gtag"
)

// OpenApiV3 是从以下网址定义的结构体：
// https://swagger.io/specification/
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md
// （译：OpenApiV3 结构体是根据以下链接中定义的 OpenAPI 3.0 规范实现的：）
// （https://swagger.io/specification/，以及 OpenAPI 3.0 的具体版本规范：）
// （https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md）
type OpenApiV3 struct {
	Config       Config                `json:"-"`
	OpenAPI      string                `json:"openapi"`
	Components   Components            `json:"components,omitempty"`
	Info         Info                  `json:"info"`
	Paths        Paths                 `json:"paths"`
	Security     *SecurityRequirements `json:"security,omitempty"`
	Servers      *Servers              `json:"servers,omitempty"`
	Tags         *Tags                 `json:"tags,omitempty"`
	ExternalDocs *ExternalDocs         `json:"externalDocs,omitempty"`
}

const (
	TypeInteger    = `integer`
	TypeNumber     = `number`
	TypeBoolean    = `boolean`
	TypeArray      = `array`
	TypeString     = `string`
	TypeFile       = `file`
	TypeObject     = `object`
	FormatInt32    = `int32`
	FormatInt64    = `int64`
	FormatDouble   = `double`
	FormatByte     = `byte`
	FormatBinary   = `binary`
	FormatDate     = `date`
	FormatDateTime = `date-time`
	FormatPassword = `password`
)

const (
	ParameterInHeader = `header`
	ParameterInPath   = `path`
	ParameterInQuery  = `query`
	ParameterInCookie = `cookie`
)

const (
	validationRuleKeyForRequired = `required`
	validationRuleKeyForIn       = `in:`
)

var (
	defaultReadContentTypes  = []string{`application/json`}
	defaultWriteContentTypes = []string{`application/json`}
	shortTypeMapForTag       = map[string]string{
		gtag.DefaultShort:      gtag.Default,
		gtag.SummaryShort:      gtag.Summary,
		gtag.SummaryShort2:     gtag.Summary,
		gtag.DescriptionShort:  gtag.Description,
		gtag.DescriptionShort2: gtag.Description,
		gtag.ExampleShort:      gtag.Example,
		gtag.ExamplesShort:     gtag.Examples,
		gtag.ExternalDocsShort: gtag.ExternalDocs,
	}
)

// New 创建并返回一个实现了 OpenApiV3 的对象。
func New() *OpenApiV3 {
	oai := &OpenApiV3{}
	oai.fillWithDefaultValue()
	return oai
}

// AddInput 是函数 OpenApiV3.Add 的结构化参数。
type AddInput struct {
	Path   string      // Path 指定自定义路径，如果在结构体标签的 Meta 中未配置此路径，则使用该指定路径。
	Prefix string      // Prefix 指定自定义路由路径前缀，它将与结构体标签中 Meta 的 path 标签相结合。
	Method string      // Method 指定自定义的 HTTP 方法，如果在结构体标签的 Meta 中未配置该方法时使用。
	Object interface{} // Object 可以是结构体实例或路由函数。
}

// Add 将一个结构体实例或路由函数添加到OpenApiV3定义实现中。
func (oai *OpenApiV3) Add(in AddInput) error {
	var (
		reflectValue = reflect.ValueOf(in.Object)
	)
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	switch reflectValue.Kind() {
	case reflect.Struct:
		return oai.addSchema(in.Object)

	case reflect.Func:
		return oai.addPath(addPathInput{
			Path:     in.Path,
			Prefix:   in.Prefix,
			Method:   in.Method,
			Function: in.Object,
		})

	default:
		return 错误类.X创建错误码并格式化(
			错误码类.CodeInvalidParameter,
			`unsupported parameter type "%s", only struct/function type is supported`,
			reflect.TypeOf(in.Object).String(),
		)
	}
}

func (oai OpenApiV3) String() string {
	b, err := json.Marshal(oai)
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	return string(b)
}

func (oai *OpenApiV3) golangTypeToOAIType(t reflect.Type) string {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.String:
		return TypeString

	case reflect.Struct:
		switch t.String() {
		case `time.Time`, `gtime.Time`:
			return TypeString
		case `ghttp.UploadFile`:
			return TypeFile
		}
		return TypeObject

	case reflect.Slice, reflect.Array:
		switch t.String() {
		case `[]uint8`:
			return TypeString
		}
		return TypeArray

	case reflect.Bool:
		return TypeBoolean

	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return TypeInteger

	case
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		return TypeNumber

	default:
		return TypeObject
	}
}

// golangTypeToOAIFormat 将给定的Golang类型`t`转换并返回OpenAPI参数格式。
// 注意，它不会返回标准的OpenAPI参数格式，而是返回Golang类型的自定义格式。
func (oai *OpenApiV3) golangTypeToOAIFormat(t reflect.Type) string {
	format := t.String()
	switch 文本类.X过滤首字符并含空白(format, "*") {
	case `[]uint8`:
		return FormatBinary

	default:
		if oai.isEmbeddedStructDefinition(t) {
			return `EmbeddedStructDefinition`
		}
		return format
	}
}

func (oai *OpenApiV3) golangTypeToSchemaName(t reflect.Type) string {
	var (
		pkgPath    string
		schemaName = 文本类.X过滤首字符并含空白(t.String(), "*")
	)
	// 指针类型没有 PkgPath。
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if pkgPath = t.PkgPath(); pkgPath != "" && pkgPath != "." {
		if !oai.Config.IgnorePkgPath {
			schemaName = 文本类.X替换(pkgPath, `/`, `.`) + 文本类.SubStrFrom别名(schemaName, ".")
		}
	}
	schemaName = 文本类.Map替换(schemaName, map[string]string{
		` `: ``,
		`{`: ``,
		`}`: ``,
	})
	return schemaName
}

func (oai *OpenApiV3) fillMapWithShortTags(m map[string]string) map[string]string {
	for k, v := range shortTypeMapForTag {
		if m[v] == "" && m[k] != "" {
			m[v] = m[k]
		}
	}
	return m
}

func formatRefToBytes(ref string) []byte {
	return []byte(fmt.Sprintf(`{"$ref":"#/components/schemas/%s"}`, ref))
}

func isValidParameterName(key string) bool {
	return key != "-"
}
