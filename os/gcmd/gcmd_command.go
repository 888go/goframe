// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package cmd类

import (
	"context"
	
	"github.com/888go/goframe/container/gset"
	"github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/text/gstr"
)

// Command 结构体保存了关于一个可以处理自定义逻辑的参数的信息。
type Command struct {
	Name          string        // 命令名称（区分大小写）。
	Usage         string        // 一行简短的描述，关于其使用方式，例如：gf build main.go [选项]
	Brief         string        // 这个命令将要执行的简短描述
	Description   string        // 一段详细描述
	Arguments     []Argument    // 参数数组，用于配置此命令的行为。
	Func          Function      // Custom function.
	FuncWithValue FuncWithValue // 自定义函数，带有输出参数，可以与命令调用者进行交互。
	HelpFunc      Function      // 自定义帮助函数
	Examples      string        // Usage examples.
	Additional    string        // 这个命令的附加信息，将会被添加到帮助信息的末尾。
	Strict        bool          // 严格解析选项，这意味着如果给出无效选项将返回错误。
	CaseSensitive bool          // CaseSensitive 解析选项，表示以区分大小写的方式解析输入选项。
	Config        string        // 配置节点名称，该名称同时从配置组件和命令行中获取值。
	parent        *Command      // 用于内部使用的父命令。
	commands      []*Command    // 此命令的子命令。
}

// Function 是一个自定义命令回调函数，它绑定到某个特定参数。
type Function func(ctx context.Context, parser *Parser) (err error)

// FuncWithValue 类似于 Func，但是带有输出参数，可以与命令调用者进行交互。
type FuncWithValue func(ctx context.Context, parser *Parser) (out interface{}, err error)

// Argument 是某些命令所使用的命令值。
type Argument struct {
	Name   string // Option name.
	Short  string // Option short.
	Brief  string // 该Option的简要信息，用于帮助信息中。
	IsArg  bool   // IsArg 标记这个参数从命令行参数而非选项中获取值。
	Orphan bool   // 是否此Option已绑定或未绑定值。
}

var (
	// defaultHelpOption 是默认的帮助选项，它将自动添加到每个命令中。
	defaultHelpOption = Argument{
		Name:   `help`,
		Short:  `h`,
		Brief:  `more information about this command`,
		Orphan: true,
	}
)

// CommandFromCtx 从上下文中检索并返回 Command。
func CommandFromCtx(ctx context.Context) *Command {
	if v := ctx.Value(CtxKeyCommand); v != nil {
		if p, ok := v.(*Command); ok {
			return p
		}
	}
	return nil
}

// AddCommand 向当前命令添加一个或多个子命令。
func (c *Command) AddCommand(commands ...*Command) error {
	for _, cmd := range commands {
		if err := c.doAddCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

// doAddCommand 向当前命令添加一个子命令。
func (c *Command) doAddCommand(command *Command) error {
	command.Name = 文本类.X过滤首尾符并含空白(command.Name)
	if command.Name == "" {
		return 错误类.X创建("command name should not be empty")
	}
	// Repeated check.
	var (
		commandNameSet = 集合类.X创建文本()
	)
	for _, cmd := range c.commands {
		commandNameSet.X加入(cmd.Name)
	}
	if commandNameSet.X是否存在(command.Name) {
		return 错误类.X创建并格式化(`command "%s" is already added to command "%s"`, command.Name, c.Name)
	}
	// 将给定的命令添加到其子命令数组中。
	command.parent = c
	c.commands = append(c.commands, command)
	return nil
}

// AddObject 通过结构体对象向当前命令添加一个或多个子命令。
func (c *Command) AddObject(objects ...interface{}) error {
	var commands []*Command
	for _, object := range objects {
		rootCommand, err := NewFromObject(object)
		if err != nil {
			return err
		}
		commands = append(commands, rootCommand)
	}
	return c.AddCommand(commands...)
}
