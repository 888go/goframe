// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

// 包 builtin 实现了内置的验证规则。
//
// 参考了 Laravel 验证规则：
// https://laravel.com/docs/master/validation#available-validation-rules
// md5:855d785bc4982d21
package builtin

import (
	"reflect"

	gvar "github.com/888go/goframe/container/gvar"
)

type Rule interface {
	// Name 返回该规则的内置名称。 md5:a71cebab229252cc
	Name() string

	// Message返回规则的默认错误消息。 md5:0b3669ebda3f3237
	Message() string

	// Run开始运行规则，如果成功则返回nil，否则返回一个错误。 md5:1e3c94009a35745d
	Run(in RunInput) error
}

type RunInput struct {
	RuleKey     string       // RuleKey 类似于规则 "max: 6" 中的 "max". md5:b823133b6b2adfb7
	RulePattern string       // RulePattern 类似于规则 "max:6" 中的 "6". md5:1ddcd9889a484fa2
	Field       string       // Value字段的名称。 md5:25ed2e11e3b4820f
	ValueType   reflect.Type // ValueType 指定了值的类型，可能为 nil。 md5:b1ad5cfd9a152a1d
	Value       *gvar.Var    // Value 指定此规则用于验证的值。 md5:29bdb57107181fe6
	Data        *gvar.Var    // Data 指定传递给Validator的数据。 md5:d01d5d82737671aa
	Message     string       // Message 指定此规则的自定义错误消息或配置的 i18n 消息。 md5:407649d2c7943432
	Option      RunOption    // Option 为验证规则提供额外的配置选项。 md5:37c94f95ca3b8bbc
}

type RunOption struct {
	CaseInsensitive bool // CaseInsensitive 表示在字符串比较时忽略大小写。 md5:46dfd02eaccd7444
}

var (
	// ruleMap 存储所有预置的验证规则。 md5:9a86d7fe68fc3d28
	ruleMap = map[string]Rule{}
)

// Register 将内置规则注册到管理器中。 md5:491be1bd9cfa9577
func Register(rule Rule) {
	ruleMap[rule.Name()] = rule
}

// GetRule 通过 `name` 获取并返回规则。 md5:a84a676e5ebc496b
func GetRule(name string) Rule {
	return ruleMap[name]
}
