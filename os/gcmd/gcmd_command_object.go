// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package cmd类

import (
	"context"
	"fmt"
	"reflect"

	gset "github.com/888go/goframe/container/gset"
	gjson "github.com/888go/goframe/encoding/gjson"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/reflection"
	"github.com/888go/goframe/internal/utils"
	"github.com/888go/goframe/os/gstructs"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	gmeta "github.com/888go/goframe/util/gmeta"
	"github.com/888go/goframe/util/gtag"
	gutil "github.com/888go/goframe/util/gutil"
	gvalid "github.com/888go/goframe/util/gvalid"
)

var (
		// defaultValueTags 是用于存储默认值的结构体标签名称。 md5:5ce6fb87a220db49
	defaultValueTags = []string{"d", "default"}
)

// NewFromObject 使用给定的对象创建并返回一个根命令对象。 md5:3bdd362e3ec9f337
func NewFromObject(object interface{}) (rootCmd *Command, err error) {
	switch c := object.(type) {
	case Command:
		return &c, nil
	case *Command:
		return c, nil
	}

	originValueAndKind := reflection.OriginValueAndKind(object)
	if originValueAndKind.OriginKind != reflect.Struct {
		err = gerror.X创建并格式化(
			`input object should be type of struct, but got "%s"`,
			originValueAndKind.InputValue.Type().String(),
		)
		return
	}
	var reflectValue = originValueAndKind.InputValue
	// 如果给定的`object`不是指针，那么它会创建一个临时的，其值为`reflectValue`。
	// 然后它可以获取结构体/`*struct`的所有方法。
	// md5:1e216cd9c7839ef2
	if reflectValue.Kind() == reflect.Struct {
		newValue := reflect.New(reflectValue.Type())
		newValue.Elem().Set(reflectValue)
		reflectValue = newValue
	}

	// Root command creating.
	rootCmd, err = newCommandFromObjectMeta(object, "")
	if err != nil {
		return
	}
	// Sub command creating.
	var (
		nameSet         = gset.X创建文本()
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
		if nameSet.X是否存在(methodCmd.Name) {
			err = gerror.X创建并格式化(
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

// `object`是业务对象的元属性，`name`是命令名称，通常来自方法名。当Meta中没有定义名称标签时使用这个。
// md5:044898119694fdf5
func newCommandFromObjectMeta(object interface{}, name string) (command *Command, err error) {
	var metaData = gmeta.Data(object)
	if err = gconv.Scan(metaData, &command); err != nil {
		return
	}
		// 名称字段是必需的。 md5:be70066859cce69e
	if command.Name == "" {
		if name == "" {
			err = gerror.X创建并格式化(
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
		// 对输入/输出参数及命名进行必要的验证。 md5:9e72ac9f4181fad5
	if methodType.NumIn() != 2 || methodType.NumOut() != 2 {
		if methodType.PkgPath() != "" {
			err = gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				`invalid command: %s.%s.%s defined as "%s", but "func(context.Context, Input)(Output, error)" is required`,
				methodType.PkgPath(), reflect.TypeOf(object).Name(), method.Name, methodType.String(),
			)
		} else {
			err = gerror.X创建错误码并格式化(
				gcode.CodeInvalidParameter,
				`invalid command: %s.%s defined as "%s", but "func(context.Context, Input)(Output, error)" is required`,
				reflect.TypeOf(object).Name(), method.Name, methodType.String(),
			)
		}
		return
	}
	if !methodType.In(0).Implements(reflect.TypeOf((*context.Context)(nil)).Elem()) {
		err = gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`invalid command: %s.%s defined as "%s", but the first input parameter should be type of "context.Context"`,
			reflect.TypeOf(object).Name(), method.Name, methodType.String(),
		)
		return
	}
	if !methodType.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		err = gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`invalid command: %s.%s defined as "%s", but the last output parameter should be type of "error"`,
			reflect.TypeOf(object).Name(), method.Name, methodType.String(),
		)
		return
	}
		// 输入结构体应该命名为`xxxInput`。 md5:98fe3954e690f01f
	if !gstr.X末尾判断(methodType.In(1).String(), `Input`) {
		err = gerror.X创建错误码并格式化(
			gcode.CodeInvalidParameter,
			`invalid struct naming for input: defined as "%s", but it should be named with "Input" suffix like "xxxInput"`,
			methodType.In(1).String(),
		)
		return
	}
		// 输出结构体应该命名为`xxxOutput`。 md5:3c8e65a804dbe66c
	if !gstr.X末尾判断(methodType.Out(0).String(), `Output`) {
		err = gerror.X创建错误码并格式化(
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

		// 利用优先级标签进行输入结构体转换。 md5:501b3e1e29551f82
	var priorityTag = gstr.X连接([]string{tagNameName, tagNameShort}, ",")

	// =============================================================================================
	// 创建一个有返回值的函数。
	// =============================================================================================
	// md5:665eb5f5657321cc
	command.FuncWithValue = func(ctx context.Context, parser *Parser) (out interface{}, err error) {
		ctx = context.WithValue(ctx, CtxKeyParser, parser)
		var (
			data        = gconv.X取Map(parser.GetOptAll())
			argIndex    = 0
			arguments   = parser.GetArgAll()
			inputValues = []reflect.Value{reflect.ValueOf(ctx)}
		)
		if value := ctx.Value(CtxKeyArgumentsIndex); value != nil {
			argIndex = value.(int)
						// 使用左面的参数来给输入结构体对象赋值。 md5:7a575ddae03dd0d6
			if argIndex < len(arguments) {
				arguments = arguments[argIndex:]
			}
		}
		if data == nil {
			data = map[string]interface{}{}
		}
		// Handle orphan options.
		for _, arg := range command.Arguments {
			if arg.IsArg {
								// 从命令行索引读取参数。 md5:aa066778e00736af
				if argIndex < len(arguments) {
					data[arg.Name] = arguments[argIndex]
					argIndex++
				}
			} else {
								// 从命令行选项名称中读取参数。 md5:9fcf493c28108e47
				if arg.Orphan {
					if orphanValue := parser.GetOpt(arg.Name); orphanValue != nil {
						if orphanValue.String() == "" {
							// Example: gf -f
							data[arg.Name] = "true"
							if arg.Short != "" {
								data[arg.Short] = "true"
							}
						} else {
							// 适配常见的用户习惯。
							// 例如：
							// `gf -f=0`：参数 `f` 被解析为 false
							// `gf -f=1`：参数 `f` 被解析为 true
							// md5:72432f87d0fc818b
							data[arg.Name] = orphanValue.X取布尔()
						}
					}
				}
			}
		}
				// 来自结构体标签的默认值。 md5:13e4a73100683597
		if err = mergeDefaultStructValue(data, inputObject.Interface()); err != nil {
			return nil, err
		}
				// 构建输入参数。 md5:c74f2d54c503f98e
		if len(data) > 0 {
			intlog.PrintFunc(ctx, func() string {
				return fmt.Sprintf(`input command data map: %s`, gjson.X变量到json字节集PANI(data))
			})
			if inputObject.Kind() == reflect.Ptr {
				err = gconv.StructTag(data, inputObject.Interface(), priorityTag)
			} else {
				err = gconv.StructTag(data, inputObject.Addr().Interface(), priorityTag)
			}
			intlog.PrintFunc(ctx, func() string {
				return fmt.Sprintf(`input object assigned data: %s`, gjson.X变量到json字节集PANI(inputObject.Interface()))
			})
			if err != nil {
				return
			}
		}

		// Parameters validation.
		if err = gvalid.New().Bail().Data(inputObject.Interface()).Assoc(data).Run(ctx); err != nil {
			err = gerror.X多层错误并格式化(gerror.X取当前错误(err), `arguments validation failed for command "%s"`, command.Name)
			return
		}
		inputValues = append(inputValues, inputObject)

				// 使用动态创建的参数值调用处理器。 md5:991efec71cdcc95a
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
		nameSet  = gset.X创建文本()
		shortSet = gset.X创建文本()
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
			return nil, gerror.X创建并格式化(
				`argument name "%s" defined in "%s.%s" is already token by built-in arguments`,
				arg.Name, reflect.TypeOf(object).String(), field.Name(),
			)
		}
		if arg.Short == helpOptionNameShort {
			return nil, gerror.X创建并格式化(
				`short argument name "%s" defined in "%s.%s" is already token by built-in arguments`,
				arg.Short, reflect.TypeOf(object).String(), field.Name(),
			)
		}
		if arg.Brief == "" {
			arg.Brief = field.TagDescription()
		}
		if v, ok := metaData[gtag.Arg]; ok {
			arg.IsArg = gconv.X取布尔(v)
		}
		if nameSet.X是否存在(arg.Name) {
			return nil, gerror.X创建并格式化(
				`argument name "%s" defined in "%s.%s" is already token by other argument`,
				arg.Name, reflect.TypeOf(object).String(), field.Name(),
			)
		}
		nameSet.X加入(arg.Name)

		if arg.Short != "" {
			if shortSet.X是否存在(arg.Short) {
				return nil, gerror.X创建并格式化(
					`short argument name "%s" defined in "%s.%s" is already token by other argument`,
					arg.Short, reflect.TypeOf(object).String(), field.Name(),
				)
			}
			shortSet.X加入(arg.Short)
		}

		args = append(args, arg)
	}

	return
}

// mergeDefaultStructValue 将请求参数与结构体标签定义中的默认值合并。 md5:0a73ebb7f647201a
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
						// 如果已经有值，那么它将忽略默认值。 md5:e95a88514a952418
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
