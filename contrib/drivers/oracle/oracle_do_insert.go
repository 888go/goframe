// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package oracle

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	gset "github.com/888go/goframe/container/gset"
	gstr "github.com/888go/goframe/text/gstr"

	gdb "github.com/888go/goframe/database/gdb"
	gcode "github.com/888go/goframe/errors/gcode"
	gerror "github.com/888go/goframe/errors/gerror"
	gconv "github.com/888go/goframe/util/gconv"
)

// X底层插入 为给定的表插入或更新数据。 md5:2a62d01f344269b8
func (d *Driver) X底层插入(
	ctx context.Context, link gdb.Link, table string, list gdb.Map切片, option gdb.DoInsertOption,
) (result sql.Result, err error) {
	switch option.InsertOption {
	case gdb.InsertOptionSave:
		return d.doSave(ctx, link, table, list, option)

	case gdb.InsertOptionReplace:
		return nil, gerror.X创建错误码(
			gcode.CodeNotSupported,
			`Replace operation is not supported by oracle driver`,
		)
	}
	var (
		keys   []string
		values []string
		params []interface{}
	)
		// 获取表字段和长度。 md5:d3d13ee5d6edab01
	var (
		listLength  = len(list)
		valueHolder = make([]string, 0)
	)
	for k := range list[0] {
		keys = append(keys, k)
		valueHolder = append(valueHolder, "?")
	}
	var (
		batchResult    = new(gdb.SqlResult)
		charL, charR   = d.X底层取数据库安全字符()
		keyStr         = charL + strings.Join(keys, charL+","+charR) + charR
		valueHolderStr = strings.Join(valueHolder, ",")
	)
		// 格式化 "INSERT...INTO..." 语句。 md5:bc835784d4de298b
	intoStrArray := make([]string, 0)
	for i := 0; i < len(list); i++ {
		for _, k := range keys {
			if s, ok := list[i][k].(gdb.Raw); ok {
				params = append(params, gconv.String(s))
			} else {
				params = append(params, list[i][k])
			}
		}
		values = append(values, valueHolderStr)
		intoStrArray = append(
			intoStrArray,
			fmt.Sprintf(
				"INTO %s(%s) VALUES(%s)",
				table, keyStr, valueHolderStr,
			),
		)
		if len(intoStrArray) == option.BatchCount || (i == listLength-1 && len(valueHolder) > 0) {
			r, err := d.X底层原生SQL执行(ctx, link, fmt.Sprintf(
				"INSERT ALL %s SELECT * FROM DUAL",
				strings.Join(intoStrArray, " "),
			), params...)
			if err != nil {
				return r, err
			}
			if n, err := r.RowsAffected(); err != nil {
				return r, err
			} else {
				batchResult.X原生sql行记录 = r
				batchResult.X影响行数 += n
			}
			params = params[:0]
			intoStrArray = intoStrArray[:0]
		}
	}
	return batchResult, nil
}

// doSave 支持Oracle的upsert操作. md5:29379eec8ad5635d
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
			gcode.CodeInvalidRequest, `Save operation list is empty by oracle driver`,
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
		keyWithChar := charL + key + charR
		queryHolders[index] = fmt.Sprintf("? AS %s", keyWithChar)
		queryValues[index] = value
		insertKeys[index] = keyWithChar
		insertValues[index] = fmt.Sprintf("T2.%s", keyWithChar)

		// 过滤掉更新值中的冲突键。
		// 并且该键不是软创建字段。
		// md5:7882adbf4107a87d
		if !(conflictKeySet.X是否存在(key) || d.Core.IsSoftCreatedFieldName(key)) {
			updateValues = append(
				updateValues,
				fmt.Sprintf(`T1.%s = T2.%s`, keyWithChar, keyWithChar),
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
// MERGE INTO {{table}} T1
// 使用 ( SELECT {{queryHolders}} FROM DUAL T2
// ON (T1.{{duplicateKey}} = T2.{{duplicateKey}} AND ...)
// 当未找到匹配时
// INSERT {{insertKeys}} VALUES {{insertValues}}
// 当找到匹配时
// UPDATE SET {{updateValues}}
// md5:7a233cb2881f0359
func parseSqlForUpsert(table string,
	queryHolders, insertKeys, insertValues, updateValues, duplicateKey []string,
) (sqlStr string) {
	var (
		queryHolderStr  = strings.Join(queryHolders, ",")
		insertKeyStr    = strings.Join(insertKeys, ",")
		insertValueStr  = strings.Join(insertValues, ",")
		updateValueStr  = strings.Join(updateValues, ",")
		duplicateKeyStr string
		pattern         = gstr.X过滤首尾符并含空白(`MERGE INTO %s T1 USING (SELECT %s FROM DUAL) T2 ON (%s) WHEN NOT MATCHED THEN INSERT(%s) VALUES (%s) WHEN MATCHED THEN UPDATE SET %s`)
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
		duplicateKeyStr,
		insertKeyStr,
		insertValueStr,
		updateValueStr,
	)
}
