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
// RuleRequiredWithoutAll 实现了 `required-without-all` 规则：
// 当所有给定的字段都为空时，该字段为必填。
//
// 格式： required-without-all:field1,field2,...
// 示例： required-without-all:id,name
type RuleRequiredWithoutAll struct{}

func init() {
	Register(RuleRequiredWithoutAll{})
}

func (r RuleRequiredWithoutAll) Name() string {
	return "required-without-all"
}

func (r RuleRequiredWithoutAll) Message() string {
	return "The {field} field is required"
}

func (r RuleRequiredWithoutAll) Run(in RunInput) error {
	var (
		required   = true
		array      = strings.Split(in.RulePattern, ",")
		foundValue interface{}
	)
	for i := 0; i < len(array); i++ {
		_, foundValue = gutil.MapPossibleItemByKey(in.Data.Map(), array[i])
		if !empty.IsEmpty(foundValue) {
			required = false
			break
		}
	}

	if required && isRequiredEmpty(in.Value.Val()) {
		return errors.New(in.Message)
	}
	return nil
}
