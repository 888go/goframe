// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package builtin

import (
	"errors"
	"strings"
	
	"github.com/888go/goframe/internal/empty"
	"github.com/888go/goframe/util/gutil"
)

// RuleRequiredWithAll 实现了 `required-with-all` 规则：
// 当所有给定的字段都不为空时，该字段是必填的。
//
// 格式： required-with-all:field1,field2,...
// 示例： required-with-all:id,name
// 这段代码注释是Go语言的，翻译成中文后的含义为：
// RuleRequiredWithAll 实现了“required-with-all”规则：
// 如果所有给出的字段都不是空的，则该字段是必需的。
//
// 格式： required-with-all:字段1,字段2,...
// 示例： required-with-all:id,name
type RuleRequiredWithAll struct{}

func init() {
	Register(RuleRequiredWithAll{})
}

func (r RuleRequiredWithAll) Name() string {
	return "required-with-all"
}

func (r RuleRequiredWithAll) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWithAll) Run(in RunInput) error {
	var (
		required   = true
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = 工具类.MapPossibleItemByKey(in.Data.X取Map(), array[i])
		if empty.IsEmpty(foundValue) {
			required = false
			break
		}
	}

	if required && isRequiredEmpty(in.Value.X取值()) {
		return errors.New(in.Message)
	}
	return nil
}
