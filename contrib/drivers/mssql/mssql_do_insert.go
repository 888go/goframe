// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	gset "github.com/888go/goframe/container/gset"
	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gstr "github.com/888go/goframe/text/gstr"
)

// X底层插入 为给定的表插入或更新数据。 md5:2a62d01f344269b8
func (d *Driver) X底层插入(ctx context.Context, link gdb.Link, table string, list gdb.Map切片, option gdb.DoInsertOption) (result sql.Result, err error) {
	switch option.InsertOption {
	case gdb.InsertOptionSave:
		return d.doSave(ctx, link, table, list, option)

	case gdb.InsertOptionReplace:
		return nil, gerror.X创建错误码(
			gcode.CodeNotSupported,
			`Replace operation is not supported by mssql driver`,
		)

	default:
		return d.Core.X底层插入(ctx, link, table, list, option)
	}
}

// doSave 支持SQL服务器的插入或更新操作. md5:8d7646245f001919
func (d *Driver) doSave(ctx context.Context,
	link gdb.Link, table string, list gdb.Map切片, option gdb.DoInsertOption,
) (result sql.Result, err error) {
	if len(option.OnConflict) == 0 {
		return nil, gerror.X创建错误码(
			gcode.CodeMissingParameter, `Please specify conflict columns`,
		)
	}

	if len(list) == 0 {
		return nil, gerror.X创建错误码(
			gcode.CodeInvalidRequest, `Save operation list is empty by mssql driver`,
		)
	}

	var (
		one          = list[0]
		oneLen       = len(one)
		charL, charR = d.X底层取数据库安全字符()

		conflictKeys   = option.OnConflict
		conflictKeySet = gset.X创建(false)

		// queryHolders：处理需要插入或更新的Holder数据
		// queryValues：处理需要插入或更新的值
		// insertKeys：处理需要插入的有效键
		// insertValues：处理需要插入的值
		// updateValues：处理需要更新的值
		// md5:7779ec7103105a5e
		queryHolders = make([]string, oneLen)
		queryValues  = make([]interface{}, oneLen)
		insertKeys   = make([]string, oneLen)
		insertValues = make([]string, oneLen)
		updateValues []string
	)

		// 将conflictKeys切片类型转换为集合（set）类型. md5:bec4a3b4ed209948
	for _, conflictKey := range conflictKeys {
		conflictKeySet.X加入(gstr.X到大写(conflictKey))
	}

	index := 0
	for key, value := range one {
		queryHolders[index] = "?"
		queryValues[index] = value
		insertKeys[index] = charL + key + charR
		insertValues[index] = "T2." + charL + key + charR

		// 过滤掉更新值中的冲突键。
		// 并且该键不是软创建字段。
		// md5:7882adbf4107a87d
		if !(conflictKeySet.X是否存在(key) || d.Core.IsSoftCreatedFieldName(key)) {
			updateValues = append(
				updateValues,
				fmt.Sprintf(`T1.%s = T2.%s`, charL+key+charR, charL+key+charR),
			)
		}
		index++
	}

	batchResult := new(gdb.SqlResult)
	sqlStr := parseSqlForUpsert(table, queryHolders, insertKeys, insertValues, updateValues, conflictKeys)
	r, err := d.X底层原生SQL执行(ctx, link, sqlStr, queryValues...)
	if err != nil {
		return r, err
	}
	if n, err := r.RowsAffected(); err != nil {
		return r, err
	} else {
		batchResult.X原生sql行记录 = r
		batchResult.X影响行数 += n
	}
	return batchResult, nil
}

// parseSqlForUpsert
// 合并到 {{table}} 为 T1
// 使用 ( VALUES( {{queryHolders}}) 为 T2 ({{insertKeyStr}})
// 当 T1.{{duplicateKey}} 等于 T2.{{duplicateKey}} 和...
// 如果未找到匹配项 THEN
// 插入 {{insertKeys}} 的值为 {{insertValues}}
// 当找到匹配项 THEN
// 更新 SET {{updateValues}}
// md5:f73865e975016dbf
func parseSqlForUpsert(table string,
	queryHolders, insertKeys, insertValues, updateValues, duplicateKey []string,
) (sqlStr string) {
	var (
		queryHolderStr  = strings.Join(queryHolders, ",")
		insertKeyStr    = strings.Join(insertKeys, ",")
		insertValueStr  = strings.Join(insertValues, ",")
		updateValueStr  = strings.Join(updateValues, ",")
		duplicateKeyStr string
		pattern         = gstr.X过滤首尾符并含空白(`MERGE INTO %s T1 USING (VALUES(%s)) T2 (%s) ON (%s) WHEN NOT MATCHED THEN INSERT(%s) VALUES (%s) WHEN MATCHED THEN UPDATE SET %s;`)
	)

	for index, keys := range duplicateKey {
		if index != 0 {
			duplicateKeyStr += " AND "
		}
		duplicateTmp := fmt.Sprintf("T1.%s = T2.%s", keys, keys)
		duplicateKeyStr += duplicateTmp
	}

	return fmt.Sprintf(pattern,
		table,
		queryHolderStr,
		insertKeyStr,
		duplicateKeyStr,
		insertKeyStr,
		insertValueStr,
		updateValueStr,
	)
}
