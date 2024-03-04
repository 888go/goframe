// 版权所有，GoFrame作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循MIT许可证条款。
// 如果随此文件未分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf 获取一份。
//

package gcmd

import (
	"context"
	"os"
	"strings"
	
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/888go/goframe/gcmd/internal/command"
	"github.com/888go/goframe/gcmd/internal/json"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// ParserOption 管理解析选项。
type ParserOption struct {
	CaseSensitive bool // 以区分大小写的方式标记选项解析
	Strict        bool // 如果传递了无效的选项，则停止解析并返回错误。
}

// 参数解析器
type Parser struct {
	option           ParserOption      // Parse option.
	parsedArgs       []string          // 如名称所述。
	parsedOptions    map[string]string // 如名称所述。
	passedOptions    map[string]bool   // 用户传入支持的选项，格式如：map[string]bool{"name,n":true}
// （该行代码注释表明，用户可以通过一个字符串到布尔值的映射表来传递支持的选项，其中键（key）为选项名，可能包含逗号分隔的别名，值（value）为true表示该选项被启用或选中。例如，在这个示例中，“name,n”表示选项名为"name"且其别名为 "n"，并且这个选项是被支持并启用的。）
	supportedOptions map[string]bool   // Option [选项名称:是否需要参数], 格式如：map[string]bool{"name":true, "n":true}
// 其中，键(如"name"和"n")代表选项名称，值(true或false)表示该选项在使用时是否需要携带参数。
	commandFuncMap   map[string]func() // Command 函数映射，用于函数处理器。
}

// ParserFromCtx 从上下文中检索并返回 Parser。
func ParserFromCtx(ctx context.Context) *Parser {
	if v := ctx.Value(CtxKeyParser); v != nil {
		if p, ok := v.(*Parser); ok {
			return p
		}
	}
	return nil
}

// Parse 创建并返回一个新的 Parser，其参数为 os.Args 以及支持的选项。
//
// 注意，参数 `supportedOptions` 形式为 [选项名: 是否需要参数]，这意味着
// `supportedOptions` 的值项表示对应的选项名是否需要参数。
//
// 可选参数 `strict` 指定在遇到无效选项时是否停止解析并返回错误。
func Parse(supportedOptions map[string]bool, option ...ParserOption) (*Parser, error) {
	if supportedOptions == nil {
		command.Init(os.Args...)
		return &Parser{
			parsedArgs:    GetArgAll(),
			parsedOptions: GetOptAll(),
		}, nil
	}
	return ParseArgs(os.Args, supportedOptions, option...)
}

// ParseArgs 创建并返回一个新的解析器，该解析器包含给定的参数及支持的选项。
//
// 注意，参数 `supportedOptions` 形式为 [选项名称: 是否需要参数]，这意味着
// `supportedOptions` 中的值项表示对应的选项名称是否需要参数。
//
// 可选参数 `strict` 指定了当遇到无效选项时，是否停止解析并返回错误。
func ParseArgs(args []string, supportedOptions map[string]bool, option ...ParserOption) (*Parser, error) {
	if supportedOptions == nil {
		command.Init(args...)
		return &Parser{
			parsedArgs:    GetArgAll(),
			parsedOptions: GetOptAll(),
		}, nil
	}
	var parserOption ParserOption
	if len(option) > 0 {
		parserOption = option[0]
	}
	parser := &Parser{
		option:           parserOption,
		parsedArgs:       make([]string, 0),
		parsedOptions:    make(map[string]string),
		passedOptions:    supportedOptions,
		supportedOptions: make(map[string]bool),
		commandFuncMap:   make(map[string]func()),
	}
	for name, needArgument := range supportedOptions {
		for _, v := range strings.Split(name, ",") {
			parser.supportedOptions[strings.TrimSpace(v)] = needArgument
		}
	}

	for i := 0; i < len(args); {
		if option := parser.parseOption(args[i]); option != "" {
			array, _ := gregex.MatchString(`^(.+?)=(.+)$`, option)
			if len(array) == 3 {
				if parser.isOptionValid(array[1]) {
					parser.setOptionValue(array[1], array[2])
				}
			} else {
				if parser.isOptionValid(option) {
					if parser.isOptionNeedArgument(option) {
						if i < len(args)-1 {
							parser.setOptionValue(option, args[i+1])
							i += 2
							continue
						}
					} else {
						parser.setOptionValue(option, "")
						i++
						continue
					}
				} else {
					// Multiple options?
					if array = parser.parseMultiOption(option); len(array) > 0 {
						for _, v := range array {
							parser.setOptionValue(v, "")
						}
						i++
						continue
					} else if parser.option.Strict {
						return nil, gerror.NewCodef(gcode.CodeInvalidParameter, `invalid option '%s'`, args[i])
					}
				}
			}
		} else {
			parser.parsedArgs = append(parser.parsedArgs, args[i])
		}
		i++
	}
	return parser, nil
}

// parseMultiOption解析选项为多个有效选项，如：--dav。
// 如果给定的选项不是多选项，则返回nil。
func (p *Parser) parseMultiOption(option string) []string {
	for i := 1; i <= len(option); i++ {
		s := option[:i]
		if p.isOptionValid(s) && !p.isOptionNeedArgument(s) {
			if i == len(option) {
				return []string{s}
			}
			array := p.parseMultiOption(option[i:])
			if len(array) == 0 {
				return nil
			}
			return append(array, s)
		}
	}
	return nil
}

func (p *Parser) parseOption(argument string) string {
	array, _ := gregex.MatchString(`^\-{1,2}(.+)$`, argument)
	if len(array) == 2 {
		return array[1]
	}
	return ""
}

func (p *Parser) isOptionValid(name string) bool {
	// Case-Sensitive.
	if p.option.CaseSensitive {
		_, ok := p.supportedOptions[name]
		return ok
	}
	// Case-InSensitive.
	for optionName := range p.supportedOptions {
		if gstr.Equal(optionName, name) {
			return true
		}
	}
	return false
}

func (p *Parser) isOptionNeedArgument(name string) bool {
	return p.supportedOptions[name]
}

// setOptionValue 为名称name及其别名设置选项值。
func (p *Parser) setOptionValue(name, value string) {
	for optionName := range p.passedOptions {
		array := gstr.SplitAndTrim(optionName, ",")
		for _, v := range array {
			if strings.EqualFold(v, name) {
				for _, v := range array {
					p.parsedOptions[v] = value
				}
				return
			}
		}
	}
}

// GetOpt 函数返回名为 `name` 的选项值，类型为 gvar.Var。
func (p *Parser) GetOpt(name string, def ...interface{}) *gvar.Var {
	if v, ok := p.parsedOptions[name]; ok {
		return gvar.New(v)
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// GetOptAll 返回所有已解析的选项。
func (p *Parser) GetOptAll() map[string]string {
	return p.parsedOptions
}

// GetArg 返回位于`index`处的参数作为gvar.Var类型。
func (p *Parser) GetArg(index int, def ...string) *gvar.Var {
	if index >= 0 && index < len(p.parsedArgs) {
		return gvar.New(p.parsedArgs[index])
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return nil
}

// GetArgAll 返回所有已解析的参数。
func (p *Parser) GetArgAll() []string {
	return p.parsedArgs
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (p Parser) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"parsedArgs":       p.parsedArgs,
		"parsedOptions":    p.parsedOptions,
		"passedOptions":    p.passedOptions,
		"supportedOptions": p.supportedOptions,
	})
}
