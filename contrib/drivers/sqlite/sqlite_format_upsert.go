// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package sqlite

import (
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// FormatUpsert 返回用于SQLite的UPSERT类型的SQL子句。
// 例如：ON CONFLICT (id) DO UPDATE SET ...
// md5:49f955d5c160f808
func (d *Driver) FormatUpsert(columns []string, list gdb.List, option gdb.DoInsertOption) (string, error) {
	if len(option.OnConflict) == 0 {
		return "", gerror.NewCode(
			gcode.CodeMissingParameter, `Please specify conflict columns`,
		)
	}

	var onDuplicateStr string
	if option.OnDuplicateStr != "" {
		onDuplicateStr = option.OnDuplicateStr
	} else if len(option.OnDuplicateMap) > 0 {
		for k, v := range option.OnDuplicateMap {
			if len(onDuplicateStr) > 0 {
				onDuplicateStr += ","
			}
			switch v.(type) {
			case gdb.Raw, *gdb.Raw:
				onDuplicateStr += fmt.Sprintf(
					"%s=%s",
					d.Core.QuoteWord(k),
					v,
				)
			default:
				onDuplicateStr += fmt.Sprintf(
					"%s=EXCLUDED.%s",
					d.Core.QuoteWord(k),
					d.Core.QuoteWord(gconv.String(v)),
				)
			}
		}
	} else {
		for _, column := range columns {
			// 如果是SAVE操作，不要自动更新创建时间。. md5:409c9c162d30afae
			if d.Core.IsSoftCreatedFieldName(column) {
				continue
			}
			if len(onDuplicateStr) > 0 {
				onDuplicateStr += ","
			}
			onDuplicateStr += fmt.Sprintf(
				"%s=EXCLUDED.%s",
				d.Core.QuoteWord(column),
				d.Core.QuoteWord(column),
			)
		}
	}

	conflictKeys := gstr.Join(option.OnConflict, ",")

	return fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET ", conflictKeys) + onDuplicateStr, nil
}
