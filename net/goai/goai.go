// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包 goai 实现并提供针对 OpenApi 规范的文档生成。
//
// https://editor.swagger.io/
// md5:cb37ff4d3e18479e
package goai

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gtag"
)

// OpenApiV3 是根据以下规范定义的结构体：
// https://swagger.io/specification/
// https://github.com/OAI/OpenAPI-Specification/blob/main/versions/3.0.0.md
// 
// 这个注释说明OpenApiV3这个结构体是依据Swagger/OpenAPI规范定义的，具体参照了给定的两个链接，分别指向Swagger官方网站的规范描述和OpenAPI Specification在GitHub上的3.0.0版本文档。
// md5:0153ee143f23e076
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

// New 创建并返回一个实现了OpenApiV3接口的对象。. md5:ccf57e0cf557df8a
func New() *OpenApiV3 {
	oai := &OpenApiV3{}
	oai.fillWithDefaultValue()
	return oai
}

// AddInput 是 OpenApiV3.Add 函数的结构化参数。. md5:0f162b41efe0b3d5
type AddInput struct {
	Path   string      // Path 如果在结构体标签的Meta中没有配置自定义路径，那么它会指定该路径。. md5:8948516e136d8d65
	Prefix string      // Prefix 指定自定义的路由路径前缀，它将与结构标签中 Meta 的 path 标签结合使用。. md5:e1653a4036580c9c
	Method string      // Method 指定了自定义的 HTTP 方法，如果这个没有在结构体标签的 Meta 中配置。. md5:2d3c92a67e5a1f3a
	Object interface{} // "Object"可以是一个结构体实例或路由函数。. md5:f828ebcbf7a4f386
}

// Add 将结构体实例或路由函数添加到 OpenApiV3 定义中实现。. md5:b29b3c78eb104250
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
		return gerror.NewCodef(
			gcode.CodeInvalidParameter,
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

// golangTypeToOAIFormat 将给定的 Go 语言类型 `t` 转换并返回为 OpenAPI 参数格式。
// 注意，它返回的不是标准的 OpenAPI 参数格式，而是 Go 语言类型中的自定义格式。
// md5:9fcc3831b2b211c9
func (oai *OpenApiV3) golangTypeToOAIFormat(t reflect.Type) string {
	format := t.String()
	switch gstr.TrimLeft(format, "*") {
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
		schemaName = gstr.TrimLeft(t.String(), "*")
	)
	// 指针类型没有PkgPath。. md5:38ccb85365da232e
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if pkgPath = t.PkgPath(); pkgPath != "" && pkgPath != "." {
		if !oai.Config.IgnorePkgPath {
			schemaName = gstr.Replace(pkgPath, `/`, `.`) + gstr.SubStrFrom(schemaName, ".")
		}
	}
	schemaName = gstr.ReplaceByMap(schemaName, map[string]string{
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
