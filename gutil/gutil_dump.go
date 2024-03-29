// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 工具类

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	
	"github.com/888go/goframe/gutil/internal/reflection"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/text/gstr"
)

// iString 用于在进行类型断言时，配合 String() 方法使用。
type iString interface {
	String() string
}

// iError 用于对 Error() 方法进行类型断言。
type iError interface {
	Error() string
}

// iMarshalJSON 是用于自定义 JSON 序列化的接口。
type iMarshalJSON interface {
	MarshalJSON() ([]byte, error)
}

// DumpOption 定义了 Export 函数的行为。
type DumpOption struct {
	WithType     bool // WithType 指定在导出内容时包含类型信息。
	ExportedOnly bool // 只导出结构体的公开字段。
}

// Dump 将变量 `values` 以更易于人工阅读的方式打印到标准输出（stdout）中。
func X调试输出(值s ...interface{}) {
	for _, value := range 值s {
		X调试输出并带选项(value, DumpOption{
			WithType:     false,
			ExportedOnly: false,
		})
	}
}

// DumpWithType 的行为类似于 Dump，但会包含类型信息。
// 也可参考 Dump。
func X调试输出并带类型(值s ...interface{}) {
	for _, value := range 值s {
		X调试输出并带选项(value, DumpOption{
			WithType:     true,
			ExportedOnly: false,
		})
	}
}

// DumpWithOption 使用自定义选项返回变量 `values`，将其格式化为更易读的字符串形式。
func X调试输出并带选项(值 interface{}, 选项 DumpOption) {
	buffer := bytes.NewBuffer(nil)
	X调试输出到Writer(buffer, 值, DumpOption{
		WithType:     选项.WithType,
		ExportedOnly: 选项.ExportedOnly,
	})
	fmt.Println(buffer.String())
}

// DumpTo 将变量 `values` 转换为字符串并写入到 `writer` 中，以更易于人工阅读的方式
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
	// 双重检查空值。
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
		// 对映射键和缩进字符串进行转储。
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
			// 没有实现共同接口。
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
	s = gstr.Trim(s, `<>`)
	if !in.Option.WithType {
		in.Buffer.WriteString(s)
	} else {
		in.Buffer.WriteString(fmt.Sprintf("%s(%s)", in.ReflectTypeName, s))
	}
}

func addSlashesForString(s string) string {
	return gstr.ReplaceByMap(s, map[string]string{
		`"`:  `\"`,
		"\r": `\r`,
		"\t": `\t`,
		"\n": `\n`,
	})
}

// DumpJson 将 JSON 内容格式化输出到标准输出（stdout）。
func X调试输出json(json值 string) {
	var (
		buffer    = bytes.NewBuffer(nil)
		jsonBytes = []byte(json值)
	)
	if err := json.Indent(buffer, jsonBytes, "", "\t"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(buffer.String())
}
