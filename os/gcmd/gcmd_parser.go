// 版权所有 GoFrame 作者(https://goframe.org)。保留所有权利。
//
// 此源代码形式受 MIT 许可证的条款约束。
// 如果未随此文件一起分发 MIT 许可证的副本，
// 您可以在 https://github.com/gogf/gf 获取一个。
// md5:a114f4bdd106ab31

package cmd类

import (
	"context"
	"os"
	"strings"

	gvar "github.com/888go/goframe/container/gvar"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/command"
	"github.com/888go/goframe/internal/json"
	gregex "github.com/888go/goframe/text/gregex"
	gstr "github.com/888go/goframe/text/gstr"
)

// ParserOption负责管理解析选项。 md5:6294496b49d5c3bb
type ParserOption struct {
	CaseSensitive bool // 以区分大小写的方式标记选项解析。 md5:8d9524f23421dc60
	Strict        bool // 如果传入了无效的选项，是否停止解析并返回错误。 md5:2564adc332d2fd51
}

// Parser for arguments.
type Parser struct {
	option           ParserOption      // Parse option.
	parsedArgs       []string          // As name described.
	parsedOptions    map[string]string // As name described.
	passedOptions    map[string]bool   // 用户传递的受支持选项，如：map[string]bool{"name,n":true}. md5:ae5a3d920682c314
	supportedOptions map[string]bool   // 选项 [OptionName:是否需要参数]，例如：map[string]bool{"name":true, "n":true}. md5:d57dd0851c5ab783
	commandFuncMap   map[string]func() // 函数处理程序的命令函数映射。 md5:0061f09955d9b987
}

// ParserFromCtx 从上下文中检索并返回解析器。 md5:260bf6b7d06ebc7c
func ParserFromCtx(ctx context.Context) *Parser {
	if v := ctx.Value(CtxKeyParser); v != nil {
		if p, ok := v.(*Parser); ok {
			return p
		}
	}
	return nil
}

// Parse 创建并返回一个新的Parser，使用os.Args和受支持的选项。
//
// 请注意，参数`supportedOptions`是[key: need argument]形式，其中
// `supportedOptions`的值项表示相应的选项名是否需要参数。
//
// 可选参数`strict`指定如果遇到无效选项时，是否停止解析并返回错误。
// md5:136e728aecd2a3b5
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

// ParseArgs 创建并返回一个新的Parser，具有给定的参数和支持的选项。
// 
// 注意，参数`supportedOptions`是一个[选项名称: 需要参数]的映射，这意味着`supportedOptions`的值项表示对应选项名称是否需要参数。
// 
// 可选参数`strict`指定是否在遇到无效选项时停止解析并返回错误。
// md5:5c367c6c4d6d78be
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
			array, _ := gregex.X匹配文本(`^(.+?)=(.+)$`, option)
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
						return nil, gerror.X创建错误码并格式化(gcode.CodeInvalidParameter, `invalid option '%s'`, args[i])
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

// parseMultiOption 解析多个有效选项，如：--dav。如果给定的选项不是多选项，它将返回nil。
// md5:d70d0f096bf48cc4
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
	array, _ := gregex.X匹配文本(`^\-{1,2}(.+)$`, argument)
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
		if gstr.X相等比较并忽略大小写(optionName, name) {
			return true
		}
	}
	return false
}

func (p *Parser) isOptionNeedArgument(name string) bool {
	return p.supportedOptions[name]
}

// setOptionValue 为名称和相应的别名设置选项值。 md5:9b55fb71d527f4c6
func (p *Parser) setOptionValue(name, value string) {
		// 准确的选项名称匹配。 md5:92eb07ef58b2270c
	for optionName := range p.passedOptions {
		optionNameAndShort := gstr.X分割并忽略空值(optionName, ",")
		for _, optionNameItem := range optionNameAndShort {
			if optionNameItem == name {
				for _, v := range optionNameAndShort {
					p.parsedOptions[v] = value
				}
				return
			}
		}
	}
		// 模糊选项名称匹配。 md5:84dde7ce64941c27
	for optionName := range p.passedOptions {
		optionNameAndShort := gstr.X分割并忽略空值(optionName, ",")
		for _, optionNameItem := range optionNameAndShort {
			if strings.EqualFold(optionNameItem, name) {
				for _, v := range optionNameAndShort {
					p.parsedOptions[v] = value
				}
				return
			}
		}
	}
}

// GetOpt 作为gvar.Var返回名为`name`的选项值。 md5:1859b868ee779be0
func (p *Parser) GetOpt(name string, def ...interface{}) *gvar.Var {
	if p == nil {
		return nil
	}
	if v, ok := p.parsedOptions[name]; ok {
		return gvar.X创建(v)
	}
	if len(def) > 0 {
		return gvar.X创建(def[0])
	}
	return nil
}

// GetOptAll 返回所有已解析的选项。 md5:6de4d266d8991786
func (p *Parser) GetOptAll() map[string]string {
	if p == nil {
		return nil
	}
	return p.parsedOptions
}

// GetArg 作为gvar.Var返回索引为`index`的参数。 md5:12ea2f8d74c6370d
func (p *Parser) GetArg(index int, def ...string) *gvar.Var {
	if p == nil {
		return nil
	}
	if index >= 0 && index < len(p.parsedArgs) {
		return gvar.X创建(p.parsedArgs[index])
	}
	if len(def) > 0 {
		return gvar.X创建(def[0])
	}
	return nil
}

// GetArgAll 返回所有解析的参数。 md5:85cc0fd5995d4878
func (p *Parser) GetArgAll() []string {
	if p == nil {
		return nil
	}
	return p.parsedArgs
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
func (p Parser) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"parsedArgs":       p.parsedArgs,
		"parsedOptions":    p.parsedOptions,
		"passedOptions":    p.passedOptions,
		"supportedOptions": p.supportedOptions,
	})
}
