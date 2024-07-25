// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31

package gcmd

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

// Run 调用与该命令绑定的自定义函数，根据os.Args执行。
// 如果发生任何错误，它将使进程退出并返回退出代码1。 md5:f6512536eb3555fe
func (c *Command) Run(ctx context.Context) {
	_ = c.RunWithValue(ctx)
}

// RunWithValue 调用与该命令绑定的 os.Args 中的自定义函数，传入值作为输出。如果发生任何错误，它将退出进程并返回退出码 1。 md5:4d204c2503673c10
func (c *Command) RunWithValue(ctx context.Context) (value interface{}) {
	value, err := c.RunWithValueError(ctx)
	if err != nil {
		var (
			code   = gerror.Code(err)
			detail = code.Detail()
			buffer = bytes.NewBuffer(nil)
		)
		if code.Code() == gcode.CodeNotFound.Code() {
			buffer.WriteString(fmt.Sprintf("ERROR: %s\n", gstr.Trim(err.Error())))
			if lastCmd, ok := detail.(*Command); ok {
				lastCmd.PrintTo(buffer)
			} else {
				c.PrintTo(buffer)
			}
		} else {
			buffer.WriteString(fmt.Sprintf("%+v\n", err))
		}
		if gtrace.GetTraceID(ctx) == "" {
			fmt.Println(buffer.String())
			os.Exit(1)
		}
		glog.Stack(false).Fatal(ctx, buffer.String())
	}
	return value
}

// RunWithError 调用与该命令关联的 os.Args 中的自定义函数，同时输出错误信息。 md5:59f4632a1aab9342
func (c *Command) RunWithError(ctx context.Context) (err error) {
	_, err = c.RunWithValueError(ctx)
	return
}

// RunWithValueError 使用os.Args中的值调用与此命令关联的自定义函数，并带有值和错误输出。 md5:007ad372fee78f96
func (c *Command) RunWithValueError(ctx context.Context) (value interface{}, err error) {
	return c.RunWithSpecificArgs(ctx, os.Args)
}

// RunWithSpecificArgs 使用绑定到该命令的特定参数调用自定义函数，并将值和错误输出传递给它。 md5:48c98cbef4733851
func (c *Command) RunWithSpecificArgs(ctx context.Context, args []string) (value interface{}, err error) {
	if len(args) == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "args can not be empty!")
	}
	parser, err := ParseArgs(args, nil)
	if err != nil {
		return nil, err
	}
	parsedArgs := parser.GetArgAll()
	if len(parsedArgs) == 1 {
		return c.doRun(ctx, args, parser)
	}

	// 排除根二进制文件名。 md5:74be69a4a70a25b8
	parsedArgs = parsedArgs[1:]

	// 找到匹配的命令并运行它。 md5:7cabb95b952de688
	lastCmd, foundCmd, newCtx := c.searchCommand(ctx, parsedArgs, 0)
	if foundCmd != nil {
		return foundCmd.doRun(newCtx, args, parser)
	}

	// 如果未找到命令，则打印错误和帮助信息。 md5:e8829411cb2fb3df
	err = gerror.NewCodef(
		gcode.WithCode(gcode.CodeNotFound, lastCmd),
		`command "%s" not found for command "%s", command line: %s`,
		gstr.Join(parsedArgs, " "),
		c.Name,
		gstr.Join(args, " "),
	)
	return
}

func (c *Command) doRun(ctx context.Context, args []string, parser *Parser) (value interface{}, err error) {
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && gerror.HasStack(v) {
				err = v
			} else {
				err = gerror.NewCodef(gcode.CodeInternalPanic, "exception recovered: %+v", exception)
			}
		}
	}()

	ctx = context.WithValue(ctx, CtxKeyCommand, c)
	// 检查内置的帮助命令。 md5:80aa5adefafed66d
	if parser.GetOpt(helpOptionName) != nil || parser.GetOpt(helpOptionNameShort) != nil {
		if c.HelpFunc != nil {
			return nil, c.HelpFunc(ctx, parser)
		}
		return nil, c.defaultHelpFunc(ctx, parser)
	}
	// 为命令提供OpenTelemetry。 md5:46407dd5b38f692f
	var (
		span trace.Span
		tr   = otel.GetTracerProvider().Tracer(
			tracingInstrumentName,
			trace.WithInstrumentationVersion(gf.VERSION),
		)
	)
	ctx, span = tr.Start(
		otel.GetTextMapPropagator().Extract(
			ctx,
			propagation.MapCarrier(genv.Map()),
		),
		gstr.Join(os.Args, " "),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	span.SetAttributes(gtrace.CommonLabels()...)
	// 为当前命令配置重新解析原始参数。 md5:6dfe6c6434a27ec5
	parser, err = c.reParse(ctx, args, parser)
	if err != nil {
		return nil, err
	}
	// 注册的命令函数调用。 md5:9e5739c9c6b28f0f
	if c.Func != nil {
		return nil, c.Func(ctx, parser)
	}
	if c.FuncWithValue != nil {
		return c.FuncWithValue(ctx, parser)
	}
	// 如果当前命令中没有定义函数，那么它会打印帮助信息。 md5:35f280d9901715f5
	if c.HelpFunc != nil {
		return nil, c.HelpFunc(ctx, parser)
	}
	return nil, c.defaultHelpFunc(ctx, parser)
}

// reParse使用当前命令的选项配置重新解析原始参数。 md5:c23561243bbefff3
func (c *Command) reParse(ctx context.Context, args []string, parser *Parser) (*Parser, error) {
	if len(c.Arguments) == 0 {
		return parser, nil
	}
	var (
		optionKey        string
		supportedOptions = make(map[string]bool)
	)
	for _, arg := range c.Arguments {
		if arg.IsArg {
			continue
		}
		if arg.Short != "" {
			optionKey = fmt.Sprintf(`%s,%s`, arg.Name, arg.Short)
		} else {
			optionKey = arg.Name
		}
		supportedOptions[optionKey] = !arg.Orphan
	}
	parser, err := ParseArgs(args, supportedOptions, ParserOption{
		CaseSensitive: c.CaseSensitive,
		Strict:        c.Strict,
	})
	if err != nil {
		return nil, err
	}
	// 如果config组件有"config"标签，从其中获取选项值。 md5:25fb126ffe7890dc
	if c.Config != "" && gcfg.Instance().Available(ctx) {
		value, err := gcfg.Instance().Get(ctx, c.Config)
		if err != nil {
			return nil, err
		}
		configMap := value.Map()
		for optionName := range parser.supportedOptions {
			// 命令行具有较高优先级。 md5:8326234bd7de1eaa
			if parser.GetOpt(optionName) != nil {
				continue
			}
			// 将配置值合并到解析器中。 md5:82c508be2619b799
			foundKey, foundValue := gutil.MapPossibleItemByKey(configMap, optionName)
			if foundKey != "" {
				parser.parsedOptions[optionName] = gconv.String(foundValue)
			}
		}
	}
	return parser, nil
}

// searchCommand 递归地根据给定的参数搜索命令。 md5:5a28ecf7bd849fd7
func (c *Command) searchCommand(
	ctx context.Context, args []string, fromArgIndex int,
) (lastCmd, foundCmd *Command, newCtx context.Context) {
	if len(args) == 0 {
		return c, nil, ctx
	}
	for _, cmd := range c.commands {
		// 递归搜索命令。
		// 字符串比较区分大小写。 md5:801cc6b5c74b2a82
		if cmd.Name == args[0] {
			leftArgs := args[1:]
			// 如果此命令需要参数，
			// 则使用参数索引标记将其剩余的所有参数传递给它。
			//
			// 注意，这里使用的args（采用默认的参数解析方式）可能与在命令中解析到的args有所不同。 md5:6f65480aaaabf1f3
			if cmd.hasArgumentFromIndex() || len(leftArgs) == 0 {
				ctx = context.WithValue(ctx, CtxKeyArgumentsIndex, fromArgIndex+1)
				return c, cmd, ctx
			}
			return cmd.searchCommand(ctx, leftArgs, fromArgIndex+1)
		}
	}
	return c, nil, ctx
}

func (c *Command) hasArgumentFromIndex() bool {
	for _, arg := range c.Arguments {
		if arg.IsArg {
			return true
		}
	}
	return false
}

func (c *Command) hasArgumentFromOption() bool {
	for _, arg := range c.Arguments {
		if !arg.IsArg {
			return true
		}
	}
	return false
}
