// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package builtin

import (
	"errors"
	"fmt"
	"reflect"

	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	"github.com/888go/goframe/internal/json"
	gstr "github.com/888go/goframe/text/gstr"
	gconv "github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gtag"
)

// RuleEnums 实现了 `enums` 规则：
// 值应在其常量类型的枚举值中。
//
// 格式：enums
// md5:e45f74add2129f75
type RuleEnums struct{}

func init() {
	Register(RuleEnums{})
}

func (r RuleEnums) Name() string {
	return "enums"
}

func (r RuleEnums) Message() string {
	return "The {field} value `{value}` should be in enums of: {enums}"
}

func (r RuleEnums) Run(in RunInput) error {
	if in.ValueType == nil {
		return gerror.NewCode(
			gcode.CodeInvalidParameter,
			`value type cannot be empty to use validation rule "enums"`,
		)
	}
	var (
		pkgPath  = in.ValueType.PkgPath()
		typeName = in.ValueType.Name()
	)
	if in.ValueType.Kind() == reflect.Slice {
		pkgPath = in.ValueType.Elem().PkgPath()
		typeName = in.ValueType.Elem().Name()
	}
	if pkgPath == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidOperation,
			`no pkg path found for type "%s"`,
			in.ValueType.String(),
		)
	}
	var (
		typeId   = fmt.Sprintf(`%s.%s`, pkgPath, typeName)
		tagEnums = gtag.GetEnumsByType(typeId)
	)
	if tagEnums == "" {
		return gerror.NewCodef(
			gcode.CodeInvalidOperation,
			`no enums found for type "%s", missing using command "gf gen enums"?`,
			typeId,
		)
	}
	var enumsValues = make([]interface{}, 0)
	if err := json.Unmarshal([]byte(tagEnums), &enumsValues); err != nil {
		return err
	}
	if !gstr.InArray(gconv.Strings(enumsValues), in.Value.String()) {
		return errors.New(gstr.Replace(
			in.Message, `{enums}`, tagEnums,
		))
	}
	return nil
}
