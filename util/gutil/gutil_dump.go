// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 工具类

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
)

// iString 用于类型断言API，用于String()。 md5:8ec0af717c4f530e
type iString interface {
	String() string
}

// iError用于类型断言错误信息。 md5:ca9885066be22039
type iError interface {
	Error() string
}

// iMarshalJSON 是自定义 JSON 序列化接口。 md5:8f96cb97a90bdc48
type iMarshalJSON interface {
	MarshalJSON() ([]byte, error)
}

// DumpOption 指定了 Export 函数的行为。 md5:2a73bcd0ce073910
type DumpOption struct {
	WithType     bool // WithType 指定以包含类型信息的方式转储内容。 md5:f0b7a9863381d552
	ExportedOnly bool // 只导出结构体的Exported字段。 md5:b19bd21abecb4c21
}

// X调试输出 将变量 `values` 打印到标准输出，以更人工可读的方式。 md5:05206ddf9d48510d
func X调试输出(值s ...interface{}) {
	for _, value := range 值s {
		X调试输出并带选项(value, DumpOption{
			WithType:     false,
			ExportedOnly: false,
		})
	}
}

// X调试输出并带类型 类似于 Dump，但带有类型信息。同时参阅 Dump。
// md5:faabab79589d38a3
func X调试输出并带类型(值s ...interface{}) {
	for _, value := range 值s {
		X调试输出并带选项(value, DumpOption{
			WithType:     true,
			ExportedOnly: false,
		})
	}
}

// X调试输出并带选项 函数将变量 `values` 以更易于人工阅读的字符串形式返回。 md5:99fec3f0f209dcf7
func X调试输出并带选项(值 interface{}, 选项 DumpOption) {
	buffer := bytes.NewBuffer(nil)
	X调试输出到Writer(buffer, 值, DumpOption{
		WithType:     选项.WithType,
		ExportedOnly: 选项.ExportedOnly,
	})
	fmt.Println(buffer.String())
}

// X调试输出到Writer 将变量 `values` 作为字符串写入到 `writer` 中，提供更易人工阅读的格式. md5:68fd8fc9ea0dfc4b
func X调试输出到Writer(writer io.Writer, 值 interface{}, 选项 DumpOption) {
	buffer := bytes.NewBuffer(nil)
	doDump(值, "", buffer, doDumpOption{
		WithType:     选项.WithType,
		ExportedOnly: 选项.ExportedOnly,
	})
	_, _ = writer.Write(buffer.Bytes())
}

type doDumpOption struct {
	WithType         bool
	ExportedOnly     bool
	DumpedPointerSet map[string]struct{}
}

func doDump(value interface{}, indent string, buffer *bytes.Buffer, option doDumpOption) {
	if option.DumpedPointerSet == nil {
		option.DumpedPointerSet = map[string]struct{}{}
	}

	if value == nil {
		buffer.WriteString(`<nil>`)
		return
	}
	var reflectValue reflect.Value
	if v, ok := value.(reflect.Value); ok {
		reflectValue = v
		if v.IsValid() && v.CanInterface() {
			value = v.Interface()
		} else {
			if convertedValue, ok := reflection.ValueToInterface(v); ok {
				value = convertedValue
			}
		}
	} else {
		reflectValue = reflect.ValueOf(value)
	}
	var reflectKind = reflectValue.Kind()
		// 二次确认空值。 md5:7122d7415991f3ef
	if value == nil || reflectKind == reflect.Invalid {
		buffer.WriteString(`<nil>`)
		return
	}
	var (
		reflectTypeName = reflectValue.Type().String()
		ptrAddress      string
		newIndent       = indent + dumpIndent
	)
	reflectTypeName = strings.ReplaceAll(reflectTypeName, `[]uint8`, `[]byte`)
	for reflectKind == reflect.Ptr {
		if ptrAddress == "" {
			ptrAddress = fmt.Sprintf(`0x%x`, reflectValue.Pointer())
		}
		reflectValue = reflectValue.Elem()
		reflectKind = reflectValue.Kind()
	}
	var (
		exportInternalInput = doDumpInternalInput{
			Value:            value,
			Indent:           indent,
			NewIndent:        newIndent,
			Buffer:           buffer,
			Option:           option,
			PtrAddress:       ptrAddress,
			ReflectValue:     reflectValue,
			ReflectTypeName:  reflectTypeName,
			ExportedOnly:     option.ExportedOnly,
			DumpedPointerSet: option.DumpedPointerSet,
		}
	)
	switch reflectKind {
	case reflect.Slice, reflect.Array:
		doDumpSlice(exportInternalInput)

	case reflect.Map:
		doDumpMap(exportInternalInput)

	case reflect.Struct:
		doDumpStruct(exportInternalInput)

	case reflect.String:
		doDumpString(exportInternalInput)

	case reflect.Bool:
		doDumpBool(exportInternalInput)

	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128:
		doDumpNumber(exportInternalInput)

	case reflect.Chan:
		buffer.WriteString(fmt.Sprintf(`<%s>`, reflectValue.Type().String()))

	case reflect.Func:
		if reflectValue.IsNil() || !reflectValue.IsValid() {
			buffer.WriteString(`<nil>`)
		} else {
			buffer.WriteString(fmt.Sprintf(`<%s>`, reflectValue.Type().String()))
		}

	case reflect.Interface:
		doDump(exportInternalInput.ReflectValue.Elem(), indent, buffer, option)

	default:
		doDumpDefault(exportInternalInput)
	}
}

type doDumpInternalInput struct {
	Value            interface{}
	Indent           string
	NewIndent        string
	Buffer           *bytes.Buffer
	Option           doDumpOption
	ReflectValue     reflect.Value
	ReflectTypeName  string
	PtrAddress       string
	ExportedOnly     bool
	DumpedPointerSet map[string]struct{}
}

func doDumpSlice(in doDumpInternalInput) {
	if b, ok := in.Value.([]byte); ok {
		if !in.Option.WithType {
			in.Buffer.WriteString(fmt.Sprintf(`"%s"`, addSlashesForString(string(b))))
		} else {
			in.Buffer.WriteString(fmt.Sprintf(
				`%s(%d) "%s"`,
				in.ReflectTypeName,
				len(string(b)),
				string(b),
			))
		}
		return
	}
	if in.ReflectValue.Len() == 0 {
		if !in.Option.WithType {
			in.Buffer.WriteString("[]")
		} else {
			in.Buffer.WriteString(fmt.Sprintf("%s(0) []", in.ReflectTypeName))
		}
		return
	}
	if !in.Option.WithType {
		in.Buffer.WriteString("[\n")
	} else {
		in.Buffer.WriteString(fmt.Sprintf("%s(%d) [\n", in.ReflectTypeName, in.ReflectValue.Len()))
	}
	for i := 0; i < in.ReflectValue.Len(); i++ {
		in.Buffer.WriteString(in.NewIndent)
		doDump(in.ReflectValue.Index(i), in.NewIndent, in.Buffer, in.Option)
		in.Buffer.WriteString(",\n")
	}
	in.Buffer.WriteString(fmt.Sprintf("%s]", in.Indent))
}

func doDumpMap(in doDumpInternalInput) {
	var mapKeys = make([]reflect.Value, 0)
	for _, key := range in.ReflectValue.MapKeys() {
		if !key.CanInterface() {
			continue
		}
		mapKey := key
		mapKeys = append(mapKeys, mapKey)
	}
	if len(mapKeys) == 0 {
		if !in.Option.WithType {
			in.Buffer.WriteString("{}")
		} else {
			in.Buffer.WriteString(fmt.Sprintf("%s(0) {}", in.ReflectTypeName))
		}
		return
	}
	var (
		maxSpaceNum = 0
		tmpSpaceNum = 0
		mapKeyStr   = ""
	)
	for _, key := range mapKeys {
		tmpSpaceNum = len(fmt.Sprintf(`%v`, key.Interface()))
		if tmpSpaceNum > maxSpaceNum {
			maxSpaceNum = tmpSpaceNum
		}
	}
	if !in.Option.WithType {
		in.Buffer.WriteString("{\n")
	} else {
		in.Buffer.WriteString(fmt.Sprintf("%s(%d) {\n", in.ReflectTypeName, len(mapKeys)))
	}
	for _, mapKey := range mapKeys {
		tmpSpaceNum = len(fmt.Sprintf(`%v`, mapKey.Interface()))
		if mapKey.Kind() == reflect.String {
			mapKeyStr = fmt.Sprintf(`"%v"`, mapKey.Interface())
		} else {
			mapKeyStr = fmt.Sprintf(`%v`, mapKey.Interface())
		}
				// 映射键和缩进字符串的转储。 md5:2c8156f9f5e204bd
		if !in.Option.WithType {
			in.Buffer.WriteString(fmt.Sprintf(
				"%s%v:%s",
				in.NewIndent,
				mapKeyStr,
				strings.Repeat(" ", maxSpaceNum-tmpSpaceNum+1),
			))
		} else {
			in.Buffer.WriteString(fmt.Sprintf(
				"%s%s(%v):%s",
				in.NewIndent,
				mapKey.Type().String(),
				mapKeyStr,
				strings.Repeat(" ", maxSpaceNum-tmpSpaceNum+1),
			))
		}
		// Map value dump.
		doDump(in.ReflectValue.MapIndex(mapKey), in.NewIndent, in.Buffer, in.Option)
		in.Buffer.WriteString(",\n")
	}
	in.Buffer.WriteString(fmt.Sprintf("%s}", in.Indent))
}

func doDumpStruct(in doDumpInternalInput) {
	if in.PtrAddress != "" {
		if _, ok := in.DumpedPointerSet[in.PtrAddress]; ok {
			in.Buffer.WriteString(fmt.Sprintf(`<cycle dump %s>`, in.PtrAddress))
			return
		}
	}
	in.DumpedPointerSet[in.PtrAddress] = struct{}{}

	structFields, _ := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         in.Value,
		RecursiveOption: gstructs.RecursiveOptionEmbedded,
	})
	var (
		hasNoExportedFields = true
		_, isReflectValue   = in.Value.(reflect.Value)
	)
	for _, field := range structFields {
		if field.IsExported() {
			hasNoExportedFields = false
			break
		}
	}
	if !isReflectValue && (len(structFields) == 0 || hasNoExportedFields) {
		var (
			structContentStr  = ""
			attributeCountStr = "0"
		)
		if v, ok := in.Value.(iString); ok {
			structContentStr = v.String()
		} else if v, ok := in.Value.(iError); ok {
			structContentStr = v.Error()
		} else if v, ok := in.Value.(iMarshalJSON); ok {
			b, _ := v.MarshalJSON()
			structContentStr = string(b)
		} else {
						// 没有实现任何公共接口。 md5:e916242484fe8e89
			if len(structFields) != 0 {
				goto dumpStructFields
			}
		}
		if structContentStr == "" {
			structContentStr = "{}"
		} else {
			structContentStr = fmt.Sprintf(`"%s"`, addSlashesForString(structContentStr))
			attributeCountStr = fmt.Sprintf(`%d`, len(structContentStr)-2)
		}
		if !in.Option.WithType {
			in.Buffer.WriteString(structContentStr)
		} else {
			in.Buffer.WriteString(fmt.Sprintf(
				"%s(%s) %s",
				in.ReflectTypeName,
				attributeCountStr,
				structContentStr,
			))
		}
		return
	}

dumpStructFields:
	var (
		maxSpaceNum = 0
		tmpSpaceNum = 0
	)
	for _, field := range structFields {
		if in.ExportedOnly && !field.IsExported() {
			continue
		}
		tmpSpaceNum = len(field.Name())
		if tmpSpaceNum > maxSpaceNum {
			maxSpaceNum = tmpSpaceNum
		}
	}
	if !in.Option.WithType {
		in.Buffer.WriteString("{\n")
	} else {
		in.Buffer.WriteString(fmt.Sprintf("%s(%d) {\n", in.ReflectTypeName, len(structFields)))
	}
	for _, field := range structFields {
		if in.ExportedOnly && !field.IsExported() {
			continue
		}
		tmpSpaceNum = len(fmt.Sprintf(`%v`, field.Name()))
		in.Buffer.WriteString(fmt.Sprintf(
			"%s%s:%s",
			in.NewIndent,
			field.Name(),
			strings.Repeat(" ", maxSpaceNum-tmpSpaceNum+1),
		))
		doDump(field.Value, in.NewIndent, in.Buffer, in.Option)
		in.Buffer.WriteString(",\n")
	}
	in.Buffer.WriteString(fmt.Sprintf("%s}", in.Indent))
}

func doDumpNumber(in doDumpInternalInput) {
	if v, ok := in.Value.(iString); ok {
		s := v.String()
		if !in.Option.WithType {
			in.Buffer.WriteString(fmt.Sprintf(`"%v"`, addSlashesForString(s)))
		} else {
			in.Buffer.WriteString(fmt.Sprintf(
				`%s(%d) "%v"`,
				in.ReflectTypeName,
				len(s),
				addSlashesForString(s),
			))
		}
	} else {
		doDumpDefault(in)
	}
}

func doDumpString(in doDumpInternalInput) {
	s := in.ReflectValue.String()
	if !in.Option.WithType {
		in.Buffer.WriteString(fmt.Sprintf(`"%v"`, addSlashesForString(s)))
	} else {
		in.Buffer.WriteString(fmt.Sprintf(
			`%s(%d) "%v"`,
			in.ReflectTypeName,
			len(s),
			addSlashesForString(s),
		))
	}
}

func doDumpBool(in doDumpInternalInput) {
	var s string
	if in.ReflectValue.Bool() {
		s = `true`
	} else {
		s = `false`
	}
	if in.Option.WithType {
		s = fmt.Sprintf(`bool(%s)`, s)
	}
	in.Buffer.WriteString(s)
}

func doDumpDefault(in doDumpInternalInput) {
	var s string
	if in.ReflectValue.IsValid() && in.ReflectValue.CanInterface() {
		s = fmt.Sprintf("%v", in.ReflectValue.Interface())
	}
	if s == "" {
		s = fmt.Sprintf("%v", in.Value)
	}
	s = gstr.X过滤首尾符并含空白(s, `<>`)
	if !in.Option.WithType {
		in.Buffer.WriteString(s)
	} else {
		in.Buffer.WriteString(fmt.Sprintf("%s(%s)", in.ReflectTypeName, s))
	}
}

func addSlashesForString(s string) string {
	return gstr.Map替换(s, map[string]string{
		`"`:  `\"`,
		"\r": `\r`,
		"\t": `\t`,
		"\n": `\n`,
	})
}

// X调试输出json 将 JSON 内容以美化的方式输出到标准输出。 md5:9f4c95e099395360
func X调试输出json(value any) {
	switch result := value.(type) {
	case []byte:
		doDumpJson(result)
	case string:
		doDumpJson([]byte(result))
	default:
		jsonContent, err := json.Marshal(value)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		doDumpJson(jsonContent)
	}
}

func doDumpJson(jsonContent []byte) {
	var (
		buffer    = bytes.NewBuffer(nil)
		jsonBytes = jsonContent
	)
	if err := json.Indent(buffer, jsonBytes, "", "    "); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(buffer.String())
}
