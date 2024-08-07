// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package cmd类

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	gstr "github.com/888go/goframe/text/gstr"
)

// Print 将当前命令的帮助信息打印到标准输出（stdout）。 md5:f96bdfd3fe6f19a6
func (c *Command) Print() {
	c.PrintTo(os.Stdout)
}

// PrintTo 将帮助信息打印到自定义的io.Writer。 md5:bd4a5ae4f69d2d7d
func (c *Command) PrintTo(writer io.Writer) {
	var (
		prefix    = gstr.X生成重复文本(" ", 4)
		buffer    = bytes.NewBuffer(nil)
		arguments = make([]Argument, len(c.Arguments))
	)
		// 复制打印选项。 md5:e1c3b9f9f9585d88
	copy(arguments, c.Arguments)
		// 添加内置帮助选项，仅用于信息用途。 md5:7ddb319e073cbc2c
	arguments = append(arguments, defaultHelpOption)

	// Usage.
	if c.Usage != "" || c.Name != "" {
		buffer.WriteString("USAGE\n")
		buffer.WriteString(prefix)
		if c.Usage != "" {
			buffer.WriteString(c.Usage)
		} else {
			var (
				p    = c
				name = c.Name
			)
			for p.parent != nil {
				name = p.parent.Name + " " + name
				p = p.parent
			}
			buffer.WriteString(name)
			if len(c.commands) > 0 {
				buffer.WriteString(` COMMAND`)
			}
			if c.hasArgumentFromIndex() {
				buffer.WriteString(` ARGUMENT`)
			}
			buffer.WriteString(` [OPTION]`)
		}
		buffer.WriteString("\n\n")
	}
	// Command.
	if len(c.commands) > 0 {
		buffer.WriteString("COMMAND\n")
		var (
			maxSpaceLength = 0
		)
		for _, cmd := range c.commands {
			if len(cmd.Name) > maxSpaceLength {
				maxSpaceLength = len(cmd.Name)
			}
		}
		for _, cmd := range c.commands {
			var (
				spaceLength    = maxSpaceLength - len(cmd.Name)
				wordwrapPrefix = gstr.X生成重复文本(" ", len(prefix+cmd.Name)+spaceLength+4)
			)
			c.printLineBrief(printLineBriefInput{
				Buffer:         buffer,
				Name:           cmd.Name,
				Prefix:         prefix,
				Brief:          gstr.X过滤首尾符并含空白(cmd.Brief),
				WordwrapPrefix: wordwrapPrefix,
				SpaceLength:    spaceLength,
			})
		}
		buffer.WriteString("\n")
	}

	// Argument.
	if c.hasArgumentFromIndex() {
		buffer.WriteString("ARGUMENT\n")
		var (
			maxSpaceLength = 0
		)
		for _, arg := range arguments {
			if !arg.IsArg {
				continue
			}
			if len(arg.Name) > maxSpaceLength {
				maxSpaceLength = len(arg.Name)
			}
		}
		for _, arg := range arguments {
			if !arg.IsArg {
				continue
			}
			var (
				spaceLength    = maxSpaceLength - len(arg.Name)
				wordwrapPrefix = gstr.X生成重复文本(" ", len(prefix+arg.Name)+spaceLength+4)
			)
			c.printLineBrief(printLineBriefInput{
				Buffer:         buffer,
				Name:           arg.Name,
				Prefix:         prefix,
				Brief:          gstr.X过滤首尾符并含空白(arg.Brief),
				WordwrapPrefix: wordwrapPrefix,
				SpaceLength:    spaceLength,
			})
		}
		buffer.WriteString("\n")
	}

	// Option.
	if c.hasArgumentFromOption() {
		buffer.WriteString("OPTION\n")
		var (
			nameStr        string
			maxSpaceLength = 0
		)
		for _, arg := range arguments {
			if arg.IsArg {
				continue
			}
			if arg.Short != "" {
				nameStr = fmt.Sprintf("-%s,--%s", arg.Short, arg.Name)
			} else {
				nameStr = fmt.Sprintf("-/--%s", arg.Name)
			}
			if len(nameStr) > maxSpaceLength {
				maxSpaceLength = len(nameStr)
			}
		}
		for _, arg := range arguments {
			if arg.IsArg {
				continue
			}
			if arg.Short != "" {
				nameStr = fmt.Sprintf("-%s, --%s", arg.Short, arg.Name)
			} else {
				nameStr = fmt.Sprintf("-/--%s", arg.Name)
			}
			var (
				brief          = gstr.X过滤首尾符并含空白(arg.Brief)
				spaceLength    = maxSpaceLength - len(nameStr)
				wordwrapPrefix = gstr.X生成重复文本(" ", len(prefix+nameStr)+spaceLength+4)
			)
			c.printLineBrief(printLineBriefInput{
				Buffer:         buffer,
				Name:           nameStr,
				Prefix:         prefix,
				Brief:          brief,
				WordwrapPrefix: wordwrapPrefix,
				SpaceLength:    spaceLength,
			})
		}
		buffer.WriteString("\n")
	}

	// Example.
	if c.Examples != "" {
		buffer.WriteString("EXAMPLE\n")
		for _, line := range gstr.X分割并忽略空值(gstr.X过滤首尾符并含空白(c.Examples), "\n") {
			buffer.WriteString(prefix)
			buffer.WriteString(gstr.X按字符数量换行(gstr.X过滤首尾符并含空白(line), maxLineChars, "\n"+prefix))
			buffer.WriteString("\n")
		}
		buffer.WriteString("\n")
	}

	// Description.
	if c.Description != "" {
		buffer.WriteString("DESCRIPTION\n")
		for _, line := range gstr.X分割并忽略空值(gstr.X过滤首尾符并含空白(c.Description), "\n") {
			buffer.WriteString(prefix)
			buffer.WriteString(gstr.X按字符数量换行(gstr.X过滤首尾符并含空白(line), maxLineChars, "\n"+prefix))
			buffer.WriteString("\n")
		}
		buffer.WriteString("\n")
	}

	// Additional.
	if c.Additional != "" {
		lineStr := gstr.X按字符数量换行(gstr.X过滤首尾符并含空白(c.Additional), maxLineChars, "\n")
		buffer.WriteString(lineStr)
		buffer.WriteString("\n")
	}
	content := buffer.String()
	content = gstr.X替换(content, "\t", "    ")
	_, _ = writer.Write([]byte(content))
}

type printLineBriefInput struct {
	Buffer         *bytes.Buffer
	Name           string
	Prefix         string
	Brief          string
	WordwrapPrefix string
	SpaceLength    int
}

func (c *Command) printLineBrief(in printLineBriefInput) {
	briefArray := gstr.X分割并忽略空值(in.Brief, "\n")
	if len(briefArray) == 0 {
				// 如果命令摘要为空，它只打印其命令名。 md5:07ab5664d0594d07
		briefArray = []string{""}
	}
	for i, line := range briefArray {
		var lineStr string
		if i == 0 {
			lineStr = fmt.Sprintf(
				"%s%s%s%s\n",
				in.Prefix, in.Name, gstr.X生成重复文本(" ", in.SpaceLength+4), line,
			)
		} else {
			lineStr = fmt.Sprintf(
				"%s%s%s%s\n",
				in.Prefix, gstr.X生成重复文本(" ", len(in.Name)), gstr.X生成重复文本(" ", in.SpaceLength+4), line,
			)
		}
		lineStr = gstr.X按字符数量换行(lineStr, maxLineChars, "\n"+in.WordwrapPrefix)
		in.Buffer.WriteString(lineStr)
	}
}

func (c *Command) defaultHelpFunc(ctx context.Context, parser *Parser) error {
		// 将命令帮助信息打印到标准输出（stdout）。 md5:0971e4539e0f02f6
	c.Print()
	return nil
}
