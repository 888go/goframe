// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

// Package builtin 实现了内置的验证规则。
//
// 参考了 Laravel 验证：
// https://laravel.com/docs/master/validation#available-validation-rules
// （该链接为 Laravel 框架关于可用验证规则的文档）
package builtin
import (
	"reflect"
	
	"github.com/888go/goframe/container/gvar"
	)
type Rule interface {
	// Name 返回规则的内置名称。
	Name() string

	// Message 返回该规则的默认错误消息。
	Message() string

	// Run 开始运行规则，如果运行成功则返回 nil，否则返回错误。
	Run(in RunInput) error
}

type RunInput struct {
	RuleKey     string       // RuleKey 类似于规则中的 "max"，如 "max: 6" 中的 "max"
	RulePattern string       // RulePattern 类似于规则 "max:6" 中的 "6"
	Field       string       // Value的字段名称。
	ValueType   reflect.Type // ValueType 指定了值的类型，该值可能为 nil。
	Value       *gvar.Var    // Value 指定此规则验证的值。
	Data        *gvar.Var    // Data 指定传递给 Validator 的 `data`。
	Message     string       // Message 指定了该规则的自定义错误消息或配置好的国际化（i18n）消息。
	Option      RunOption    // Option 提供了验证规则的额外配置选项。
}

type RunOption struct {
	CaseInsensitive bool // CaseInsensitive 表示在进行字符串比较时采用不区分大小写的方式。
}

var (
	// ruleMap 存储所有内置验证规则。
	ruleMap = map[string]Rule{}
)

// Register 将内建规则注册到管理器中。
func Register(rule Rule) {
	ruleMap[rule.Name()] = rule
}

// GetRule 通过`name`检索并返回规则。
func GetRule(name string) Rule {
	return ruleMap[name]
}
