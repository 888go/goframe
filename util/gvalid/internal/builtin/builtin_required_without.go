// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"strings"

	"github.com/gogf/gf/v2/internal/empty"
	"github.com/gogf/gf/v2/util/gutil"
)

// RuleRequiredWithout 实现了 `required-without` 规则：
// 如果给定的任何字段为空，那么这个是必需的。
//
// 格式：  required-without:field1,field2,...
// 示例： required-without:id,name md5:2e4a4e507f20f3a1
type RuleRequiredWithout struct{}

func init() {
	Register(RuleRequiredWithout{})
}

func (r RuleRequiredWithout) Name() string {
	return "required-without"
}

func (r RuleRequiredWithout) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWithout) Run(in RunInput) error {
	var (
		required   = false
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
		dataMap    = in.Data.Map()
	)

	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(dataMap, array[i])
		if empty.IsEmpty(foundValue) {
			required = true
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
