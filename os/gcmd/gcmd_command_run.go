// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gcmd
import (
	"bytes"
	"context"
	"fmt"
	"os"
	
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	
	"github.com/888go/goframe"
	"github.com/888go/goframe/errors/gcode"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/net/gtrace"
	"github.com/888go/goframe/os/gcfg"
	"github.com/888go/goframe/os/genv"
	"github.com/888go/goframe/os/glog"
	"github.com/888go/goframe/text/gstr"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
// Run调用与此命令绑定的自定义函数。
// 如果出现任何错误，它将使用退出码1退出此进程。
func (c *Command) Run(ctx context.Context) {
	_ = c.RunWithValue(ctx)
}

// RunWithValue 调用与此命令绑定的自定义函数，并传入输出的值。
// 如果发生任何错误，该过程将以退出码 1 退出。
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

// RunWithError 调用与此命令绑定的自定义函数，并带有错误输出。
func (c *Command) RunWithError(ctx context.Context) (err error) {
	_, err = c.RunWithValueError(ctx)
	return
}

// RunWithValueError 调用与此命令绑定的自定义函数，并带有值和错误输出。
func (c *Command) RunWithValueError(ctx context.Context) (value interface{}, err error) {
	// 使用默认算法解析命令行参数和选项。
	parser, err := Parse(nil)
	if err != nil {
		return nil, err
	}
	args := parser.GetArgAll()
	if len(args) == 1 {
		return c.doRun(ctx, parser)
	}

	// 排除根二进制名称。
	args = args[1:]

	// 查找匹配的命令并执行它。
	lastCmd, foundCmd, newCtx := c.searchCommand(ctx, args)
	if foundCmd != nil {
		return foundCmd.doRun(newCtx, parser)
	}

	// 如果未找到命令，则打印错误信息和帮助命令。
	err = gerror.NewCodef(
		gcode.WithCode(gcode.CodeNotFound, lastCmd),
		`command "%s" not found for command "%s", command line: %s`,
		gstr.Join(args, " "),
		c.Name,
		gstr.Join(os.Args, " "),
	)
	return
}

func (c *Command) doRun(ctx context.Context, parser *Parser) (value interface{}, err error) {
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
	// 检查内置帮助命令。
	if parser.GetOpt(helpOptionName) != nil || parser.GetOpt(helpOptionNameShort) != nil {
		if c.HelpFunc != nil {
			return nil, c.HelpFunc(ctx, parser)
		}
		return nil, c.defaultHelpFunc(ctx, parser)
	}
	// OpenTelemetry 用于命令。
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
	// 根据当前命令配置重新解析参数。
	parser, err = c.reParse(ctx, parser)
	if err != nil {
		return nil, err
	}
	// 调用已注册的命令函数
	if c.Func != nil {
		return nil, c.Func(ctx, parser)
	}
	if c.FuncWithValue != nil {
		return c.FuncWithValue(ctx, parser)
	}
	// 如果当前命令未定义任何函数，则打印帮助信息。
	if c.HelpFunc != nil {
		return nil, c.HelpFunc(ctx, parser)
	}
	return nil, c.defaultHelpFunc(ctx, parser)
}

// reParse 根据当前命令的选项配置重新解析参数。
func (c *Command) reParse(ctx context.Context, parser *Parser) (*Parser, error) {
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
	parser, err := Parse(supportedOptions, ParserOption{
		CaseSensitive: c.CaseSensitive,
		Strict:        c.Strict,
	})
	if err != nil {
		return nil, err
	}
	// 如果配置组件带有"config"标签，则从该组件中获取选项值。
	if c.Config != "" && gcfg.Instance().Available(ctx) {
		value, err := gcfg.Instance().Get(ctx, c.Config)
		if err != nil {
			return nil, err
		}
		configMap := value.Map()
		for optionName := range parser.supportedOptions {
			// 命令行参数具有高优先级
			if parser.GetOpt(optionName) != nil {
				continue
			}
			// 将配置值合并到解析器中。
			foundKey, foundValue := gutil.MapPossibleItemByKey(configMap, optionName)
			if foundKey != "" {
				parser.parsedOptions[optionName] = gconv.String(foundValue)
			}
		}
	}
	return parser, nil
}

// searchCommand 根据给定的参数递归搜索命令。
func (c *Command) searchCommand(ctx context.Context, args []string) (lastCmd, foundCmd *Command, newCtx context.Context) {
	if len(args) == 0 {
		return c, nil, ctx
	}
	for _, cmd := range c.commands {
		// 递归搜索命令。
		if cmd.Name == args[0] {
			leftArgs := args[1:]
// 如果该命令需要参数，
// 则将它左侧的所有参数传递给它。
			if cmd.hasArgumentFromIndex() {
				ctx = context.WithValue(ctx, CtxKeyArguments, leftArgs)
				return c, cmd, ctx
			}
			// Recursively searching.
			if len(leftArgs) == 0 {
				return c, cmd, ctx
			}
			return cmd.searchCommand(ctx, leftArgs)
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
