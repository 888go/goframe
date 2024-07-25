// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。 md5:a114f4bdd106ab31

package gcmd

import (
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
)

// Command 包含有关可以处理自定义逻辑的参数的信息。 md5:b0e0f23cc6e868c5
type Command struct {
	Name          string        // 命令名称（大小写敏感）。 md5:44e7c13c9c0eced2
	Usage         string        // 关于其用途的简短描述，例如：gf build main.go [选项]. md5:e2660484a0edfee8
	Brief         string        // 一个简短的描述，说明此命令将执行的操作。 md5:4a0304d2ac452238
	Description   string        // 一个详细的描述。 md5:b83b3d2318b54bce
	Arguments     []Argument    // 参数数组，配置此命令的行为。 md5:9c82b0f6e377e648
	Func          Function      // Custom function.
	FuncWithValue FuncWithValue // 自定义函数，带有输出参数，能够与命令调用者进行交互。 md5:586037addaa736f6
	HelpFunc      Function      // Custom help function
	Examples      string        // Usage examples.
	Additional    string        // 关于此命令的附加信息，将追加到帮助信息的末尾。 md5:328b9830bf970895
	Strict        bool          // 严格的解析选项，这意味着如果提供无效的选项，它会返回错误。 md5:5e8eec207aef7c7a
	CaseSensitive bool          // 区分大小写的解析选项，这意味着它以区分大小写的方式解析输入选项。 md5:b18eddb5f60c7176
	Config        string        // 配置节点名称，它也会从配置组件中获取值，以及从命令行中获取。 md5:0f67ea7288e8e541
	parent        *Command      // 用于内部使用的父命令。 md5:6572369b29bb2e3e
	commands      []*Command    // 该命令的子命令。 md5:579fb8699f4ff8e3
}

// Function 是一个绑定到特定参数的自定义命令回调函数。 md5:74f820cae660a1b5
type Function func(ctx context.Context, parser *Parser) (err error)

// FuncWithValue 类似于 Func，但它带有输出参数，这些参数可以与命令调用者进行交互。 md5:e8459756fad8cbb9
type FuncWithValue func(ctx context.Context, parser *Parser) (out interface{}, err error)

// Argument 是某些命令使用的命令值。 md5:e5c110dcf519025a
type Argument struct {
	Name   string // Option name.
	Short  string // Option short.
	Brief  string // 这个选项的简要信息，用于帮助信息中。 md5:b913553040a0d889
	IsArg  bool   // IsArg 标记这个参数从命令行参数而不是选项中获取值。 md5:24e6cc6cb658557a
	Orphan bool   // 此选项是否有值与之绑定。 md5:bc1b6ee078e2683c
}

var (
	// defaultHelpOption 是默认的帮助选项，将会自动添加到每个命令中。 md5:3593428e8c7dfe0a
	defaultHelpOption = Argument{
		Name:   `help`,
		Short:  `h`,
		Brief:  `more information about this command`,
		Orphan: true,
	}
)

// CommandFromCtx从上下文检索并返回Command。 md5:81a6b36fc029401b
func CommandFromCtx(ctx context.Context) *Command {
	if v := ctx.Value(CtxKeyCommand); v != nil {
		if p, ok := v.(*Command); ok {
			return p
		}
	}
	return nil
}

// AddCommand向当前命令添加一个或多个子命令。 md5:f1582e4eafa78dd7
func (c *Command) AddCommand(commands ...*Command) error {
	for _, cmd := range commands {
		if err := c.doAddCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

// doAddCommand 向当前命令添加一个子命令。 md5:bd1d8d447805aafd
func (c *Command) doAddCommand(command *Command) error {
	command.Name = gstr.Trim(command.Name)
	if command.Name == "" {
		return gerror.New("command name should not be empty")
	}
	// Repeated check.
	var (
		commandNameSet = gset.NewStrSet()
	)
	for _, cmd := range c.commands {
		commandNameSet.Add(cmd.Name)
	}
	if commandNameSet.Contains(command.Name) {
		return gerror.Newf(`command "%s" is already added to command "%s"`, command.Name, c.Name)
	}
	// 将给定的命令添加到其子命令数组中。 md5:ddd450893c5e1fcc
	command.parent = c
	c.commands = append(c.commands, command)
	return nil
}

// AddObject 通过struct对象向当前命令添加一个或多个子命令。 md5:8de76f64f667f83d
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
