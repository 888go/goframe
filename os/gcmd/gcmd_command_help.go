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
	"io"
	"os"
	
	"github.com/888go/goframe/text/gstr"
	)
// Print 将当前命令的帮助信息打印到标准输出（stdout）中。
func (c *Command) Print() {
	c.PrintTo(os.Stdout)
}

// PrintTo 将帮助信息打印到自定义的 io.Writer。
func (c *Command) PrintTo(writer io.Writer) {
	var (
		prefix    = gstr.Repeat(" ", 4)
		buffer    = bytes.NewBuffer(nil)
		arguments = make([]Argument, len(c.Arguments))
	)
	// 复制打印选项。
	copy(arguments, c.Arguments)
	// 添加内置帮助选项，仅用于提供信息。
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
				wordwrapPrefix = gstr.Repeat(" ", len(prefix+cmd.Name)+spaceLength+4)
			)
			c.printLineBrief(printLineBriefInput{
				Buffer:         buffer,
				Name:           cmd.Name,
				Prefix:         prefix,
				Brief:          gstr.Trim(cmd.Brief),
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
				wordwrapPrefix = gstr.Repeat(" ", len(prefix+arg.Name)+spaceLength+4)
			)
			c.printLineBrief(printLineBriefInput{
				Buffer:         buffer,
				Name:           arg.Name,
				Prefix:         prefix,
				Brief:          gstr.Trim(arg.Brief),
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
				brief          = gstr.Trim(arg.Brief)
				spaceLength    = maxSpaceLength - len(nameStr)
				wordwrapPrefix = gstr.Repeat(" ", len(prefix+nameStr)+spaceLength+4)
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
		for _, line := range gstr.SplitAndTrim(gstr.Trim(c.Examples), "\n") {
			buffer.WriteString(prefix)
			buffer.WriteString(gstr.WordWrap(gstr.Trim(line), maxLineChars, "\n"+prefix))
			buffer.WriteString("\n")
		}
		buffer.WriteString("\n")
	}

	// Description.
	if c.Description != "" {
		buffer.WriteString("DESCRIPTION\n")
		for _, line := range gstr.SplitAndTrim(gstr.Trim(c.Description), "\n") {
			buffer.WriteString(prefix)
			buffer.WriteString(gstr.WordWrap(gstr.Trim(line), maxLineChars, "\n"+prefix))
			buffer.WriteString("\n")
		}
		buffer.WriteString("\n")
	}

	// Additional.
	if c.Additional != "" {
		lineStr := gstr.WordWrap(gstr.Trim(c.Additional), maxLineChars, "\n")
		buffer.WriteString(lineStr)
		buffer.WriteString("\n")
	}
	content := buffer.String()
	content = gstr.Replace(content, "\t", "    ")
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
	briefArray := gstr.SplitAndTrim(in.Brief, "\n")
	if len(briefArray) == 0 {
		// 如果命令简介为空，则仅打印其命令名称。
		briefArray = []string{""}
	}
	for i, line := range briefArray {
		var lineStr string
		if i == 0 {
			lineStr = fmt.Sprintf(
				"%s%s%s%s\n",
				in.Prefix, in.Name, gstr.Repeat(" ", in.SpaceLength+4), line,
			)
		} else {
			lineStr = fmt.Sprintf(
				"%s%s%s%s\n",
				in.Prefix, gstr.Repeat(" ", len(in.Name)), gstr.Repeat(" ", in.SpaceLength+4), line,
			)
		}
		lineStr = gstr.WordWrap(lineStr, maxLineChars, "\n"+in.WordwrapPrefix)
		in.Buffer.WriteString(lineStr)
	}
}

func (c *Command) defaultHelpFunc(ctx context.Context, parser *Parser) error {
	// 向标准输出(stdout)打印命令帮助信息。
	c.Print()
	return nil
}
