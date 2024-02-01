// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gcmd
import (
	"context"
	"fmt"
	"reflect"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/encoding/gjson"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
	"github.com/888go/goframe/util/gutil"
	"github.com/888go/goframe/util/gvalid"
	)
var (
	// defaultValueTags 是用于存储默认值的结构体标签名称。
	defaultValueTags = []string{"d", "default"}
)

// NewFromObject 通过给定的对象创建并返回一个根命令对象。
func NewFromObject(object interface{}) (rootCmd *Command, err error) {
	switch c := object.(type) {
	case Command:
		return &c, nil
	case *Command:
		return c, nil
	}

	originValueAndKind := reflection.OriginValueAndKind(object)
	if originValueAndKind.OriginKind != reflect.Struct {
		err = gerror.Newf(
			`input object should be type of struct, but got "%s"`,
			originValueAndKind.InputValue.Type().String(),
		)
		return
	}
	var reflectValue = originValueAndKind.InputValue
// 如果给定的`object`不是指针，它会创建一个临时指针，
// 其指向值为`reflectValue`。
// 然后可以获取结构体（包括结构体指针）的所有方法。
// 这段代码注释的翻译如下：
// ```go
// 如果传入的`object`不是一个指针类型，
// 则会创建一个临时指针变量，该指针指向`reflectValue`。
// 这样就可以获取到结构体及其指针类型的全部方法。
	if reflectValue.Kind() == reflect.Struct {
		newValue := reflect.New(reflectValue.Type())
		newValue.Elem().Set(reflectValue)
		reflectValue = newValue
	}

	// 创建根命令
	rootCmd, err = newCommandFromObjectMeta(object, "")
	if err != nil {
		return
	}
	// 子命令创建
	var (
		nameSet         = gset.NewStrSet()
		rootCommandName = gmeta.Get(object, gtag.Root).String()
		subCommands     []*Command
	)
	if rootCommandName == "" {
		rootCommandName = rootCmd.Name
	}
	for i := 0; i < reflectValue.NumMethod(); i++ {
		var (
			method      = reflectValue.Type().Method(i)
			methodValue = reflectValue.Method(i)
			methodType  = methodValue.Type()
			methodCmd   *Command
		)
		methodCmd, err = newCommandFromMethod(object, method, methodValue, methodType)
		if err != nil {
			return
		}
		if nameSet.Contains(methodCmd.Name) {
			err = gerror.Newf(
				`command name should be unique, found duplicated command name in method "%s"`,
				methodType.String(),
			)
			return
		}
		if rootCommandName == methodCmd.Name {
			methodToRootCmdWhenNameEqual(rootCmd, methodCmd)
		} else {
			subCommands = append(subCommands, methodCmd)
		}
	}
	if len(subCommands) > 0 {
		err = rootCmd.AddCommand(subCommands...)
	}
	return
}

func methodToRootCmdWhenNameEqual(rootCmd *Command, methodCmd *Command) {
	if rootCmd.Usage == "" {
		rootCmd.Usage = methodCmd.Usage
	}
	if rootCmd.Brief == "" {
		rootCmd.Brief = methodCmd.Brief
	}
	if rootCmd.Description == "" {
		rootCmd.Description = methodCmd.Description
	}
	if rootCmd.Examples == "" {
		rootCmd.Examples = methodCmd.Examples
	}
	if rootCmd.Func == nil {
		rootCmd.Func = methodCmd.Func
	}
	if rootCmd.FuncWithValue == nil {
		rootCmd.FuncWithValue = methodCmd.FuncWithValue
	}
	if rootCmd.HelpFunc == nil {
		rootCmd.HelpFunc = methodCmd.HelpFunc
	}
	if len(rootCmd.Arguments) == 0 {
		rootCmd.Arguments = methodCmd.Arguments
	}
	if !rootCmd.Strict {
		rootCmd.Strict = methodCmd.Strict
	}
	if rootCmd.Config == "" {
		rootCmd.Config = methodCmd.Config
	}
}

// `object` 是业务对象中的 Meta 属性，而 `name` 是命令名称，
// 通常来源于方法名，在 Meta 中未定义 name 标签时使用。
func newCommandFromObjectMeta(object interface{}, name string) (command *Command, err error) {
	var metaData = gmeta.Data(object)
	if err = gconv.Scan(metaData, &command); err != nil {
		return
	}
	// Name字段是必需的。
	if command.Name == "" {
		if name == "" {
			err = gerror.Newf(
				`command name cannot be empty, "name" tag not found in meta of struct "%s"`,
				reflect.TypeOf(object).String(),
			)
			return
		}
		command.Name = name
	}
	if command.Brief == "" {
		for _, tag := range []string{gtag.Summary, gtag.SummaryShort, gtag.SummaryShort2} {
			command.Brief = metaData[tag]
			if command.Brief != "" {
				break
			}
		}
	}
	if command.Description == "" {
		command.Description = metaData[gtag.DescriptionShort]
	}
	if command.Brief == "" && command.Description != "" {
		command.Brief = command.Description
		command.Description = ""
	}
	if command.Examples == "" {
		command.Examples = metaData[gtag.ExampleShort]
	}
	if command.Additional == "" {
		command.Additional = metaData[gtag.AdditionalShort]
	}
	return
}

func newCommandFromMethod(
	object interface{}, method reflect.Method, methodValue reflect.Value, methodType reflect.Type,
) (command *Command, err error) {
	// 对输入/输出参数和命名进行必要的验证。
	if methodType.NumIn() != 2 || methodType.NumOut() != 2 {
		if methodType.PkgPath() != "" {
			err = gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid command: %s.%s.%s defined as "%s", but "func(context.Context, Input)(Output, error)" is required`,
				methodType.PkgPath(), reflect.TypeOf(object).Name(), method.Name, methodType.String(),
			)
		} else {
			err = gerror.NewCodef(
				gcode.CodeInvalidParameter,
				`invalid command: %s.%s defined as "%s", but "func(context.Context, Input)(Output, error)" is required`,
				reflect.TypeOf(object).Name(), method.Name, methodType.String(),
			)
		}
		return
	}
	if !methodType.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid command: %s.%s defined as "%s", but the first input parameter should be type of "context.Context"`,
			reflect.TypeOf(object).Name(), method.Name, methodType.String(),
		)
		return
	}
	if !methodType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid command: %s.%s defined as "%s", but the last output parameter should be type of "error"`,
			reflect.TypeOf(object).Name(), method.Name, methodType.String(),
		)
		return
	}
	// 输入结构体应命名为 `xxxInput`。
	if !gstr.HasSuffix(methodType.In(1).String(), `Input`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for input: defined as "%s", but it should be named with "Input" suffix like "xxxInput"`,
			methodType.In(1).String(),
		)
		return
	}
	// 输出结构体的名称应命名为`xxxOutput`。
	if !gstr.HasSuffix(methodType.Out(0).String(), `Output`) {
		err = gerror.NewCodef(
			gcode.CodeInvalidParameter,
			`invalid struct naming for output: defined as "%s", but it should be named with "Output" suffix like "xxxOutput"`,
			methodType.Out(0).String(),
		)
		return
	}

	var inputObject reflect.Value
	if methodType.In(1).Kind() == reflect.Ptr {
		inputObject = reflect.New(methodType.In(1).Elem()).Elem()
	} else {
		inputObject = reflect.New(methodType.In(1)).Elem()
	}

	// Command creating.
	if command, err = newCommandFromObjectMeta(inputObject.Interface(), method.Name); err != nil {
		return
	}

	// Options creating.
	if command.Arguments, err = newArgumentsFromInput(inputObject.Interface()); err != nil {
		return
	}

// =============================================================================================
// 创建一个具有返回值的函数。
// =============================================================================================
	command.FuncWithValue = func(ctx context.Context, parser *Parser) (out interface{}, err error) {
		ctx = context.WithValue(ctx, CtxKeyParser, parser)
		var (
			data        = gconv.Map(parser.GetOptAll())
			argIndex    = 0
			arguments   = gconv.Strings(ctx.Value(CtxKeyArguments))
			inputValues = []reflect.Value{reflect.ValueOf(ctx)}
		)
		if data == nil {
			data = map[string]interface{}{}
		}
		// 处理孤立选项。
		for _, arg := range command.Arguments {
			if arg.IsArg {
				// 从命令行参数索引读取参数。
				if argIndex < len(arguments) {
					data[arg.Name] = arguments[argIndex]
					argIndex++
				}
			} else {
				// 从命令行选项名称读取参数。
				if arg.Orphan {
					if orphanValue := parser.GetOpt(arg.Name); orphanValue != nil {
						if orphanValue.String() == "" {
							// Eg: gf -f
							data[arg.Name] = "true"
						} else {
// 用户习惯适配器，包含通用用户习惯设置。
// 例如：
// `gf -f=0`：参数`f`将被解析为false
// `gf -f=1`：参数`f`将被解析为true
							data[arg.Name] = orphanValue.Bool()
						}
					}
				}
			}
		}
		// 结构体标签中的默认值。
		if err = mergeDefaultStructValue(data, inputObject.Interface()); err != nil {
			return nil, err
		}
		// 构造输入参数。
		if len(data) > 0 {
			intlog.PrintFunc(ctx, func() string {
				return fmt.Sprintf(`input command data map: %s`, gjson.MustEncode(data))
			})
			if inputObject.Kind() == reflect.Ptr {
				err = gconv.Scan(data, inputObject.Interface())
			} else {
				err = gconv.Struct(data, inputObject.Addr().Interface())
			}
			intlog.PrintFunc(ctx, func() string {
				return fmt.Sprintf(`input object assigned data: %s`, gjson.MustEncode(inputObject.Interface()))
			})
			if err != nil {
				return
			}
		}

		// 参数验证。
		if err = gvalid.New().Bail().Data(inputObject.Interface()).Assoc(data).Run(ctx); err != nil {
			err = gerror.Wrapf(gerror.Current(err), `arguments validation failed for command "%s"`, command.Name)
			return
		}
		inputValues = append(inputValues, inputObject)

		// 使用动态创建的参数值调用处理器。
		results := methodValue.Call(inputValues)
		out = results[0].Interface()
		if !results[1].IsNil() {
			if v, ok := results[1].Interface().(error); ok {
				err = v
			}
		}
		return
	}
	return
}

func newArgumentsFromInput(object interface{}) (args []Argument, err error) {
	var (
		fields   []gstructs.Field
		nameSet  = gset.NewStrSet()
		shortSet = gset.NewStrSet()
	)
	fields, err = gstructs.Fields(gstructs.FieldsInput{
		Pointer:         object,
		RecursiveOption: gstructs.RecursiveOptionEmbeddedNoTag,
	})
	for _, field := range fields {
		var (
			arg      = Argument{}
			metaData = field.TagMap()
		)
		if err = gconv.Scan(metaData, &arg); err != nil {
			return nil, err
		}
		if arg.Name == "" {
			arg.Name = field.Name()
		}
		if arg.Name == helpOptionName {
			return nil, gerror.Newf(
				`argument name "%s" defined in "%s.%s" is already token by built-in arguments`,
				arg.Name, reflect.TypeOf(object).String(), field.Name(),
			)
		}
		if arg.Short == helpOptionNameShort {
			return nil, gerror.Newf(
				`short argument name "%s" defined in "%s.%s" is already token by built-in arguments`,
				arg.Short, reflect.TypeOf(object).String(), field.Name(),
			)
		}
		if arg.Brief == "" {
			arg.Brief = field.TagDescription()
		}
		if v, ok := metaData[gtag.Arg]; ok {
			arg.IsArg = gconv.Bool(v)
		}
		if nameSet.Contains(arg.Name) {
			return nil, gerror.Newf(
				`argument name "%s" defined in "%s.%s" is already token by other argument`,
				arg.Name, reflect.TypeOf(object).String(), field.Name(),
			)
		}
		nameSet.Add(arg.Name)

		if arg.Short != "" {
			if shortSet.Contains(arg.Short) {
				return nil, gerror.Newf(
					`short argument name "%s" defined in "%s.%s" is already token by other argument`,
					arg.Short, reflect.TypeOf(object).String(), field.Name(),
				)
			}
			shortSet.Add(arg.Short)
		}

		args = append(args, arg)
	}

	return
}

// mergeDefaultStructValue 将请求参数与来自结构体标签定义的默认值进行合并。
func mergeDefaultStructValue(data map[string]interface{}, pointer interface{}) error {
	tagFields, err := gstructs.TagFields(pointer, defaultValueTags)
	if err != nil {
		return err
	}
	if len(tagFields) > 0 {
		var (
			foundKey   string
			foundValue interface{}
		)
		for _, field := range tagFields {
			var (
				nameValue  = field.Tag(tagNameName)
				shortValue = field.Tag(tagNameShort)
			)
			// 如果它已经有了值，则会忽略默认值。
			if value, ok := data[nameValue]; ok {
				data[field.Name()] = value
				continue
			}
			if value, ok := data[shortValue]; ok {
				data[field.Name()] = value
				continue
			}
			foundKey, foundValue = gutil.MapPossibleItemByKey(data, field.Name())
			if foundKey == "" {
				data[field.Name()] = field.TagValue
			} else {
				if utils.IsEmpty(foundValue) {
					data[foundKey] = field.TagValue
				}
			}
		}
	}
	return nil
}
