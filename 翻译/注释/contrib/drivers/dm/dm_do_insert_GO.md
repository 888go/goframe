
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// DoInsert inserts or updates data for given table.
<原文结束>

# <翻译开始>
// DoInsert 为给定的表插入或更新数据。 md5:2a62d01f344269b8
# <翻译结束>


<原文开始>
// TODO:: Should be Supported
<原文结束>

# <翻译开始>
// 待办:: 应当予以支持. md5:e307b954f3ef009f
# <翻译结束>


<原文开始>
// doSave support upsert for dm
<原文结束>

# <翻译开始>
// doSave 支持dm的更新插入操作. md5:590008c012e35ec7
# <翻译结束>


<原文开始>
		// queryHolders:	Handle data with Holder that need to be upsert
		// queryValues:		Handle data that need to be upsert
		// insertKeys:		Handle valid keys that need to be inserted
		// insertValues:	Handle values that need to be inserted
		// updateValues:	Handle values that need to be updated
<原文结束>

# <翻译开始>
		// queryHolders：处理需要插入或更新的Holder数据
		// queryValues：处理需要插入或更新的值
		// insertKeys：处理需要插入的有效键
		// insertValues：处理需要插入的值
		// updateValues：处理需要更新的值
		// md5:7779ec7103105a5e
# <翻译结束>


<原文开始>
// conflictKeys slice type conv to set type
<原文结束>

# <翻译开始>
// 将conflictKeys切片类型转换为集合（set）类型. md5:bec4a3b4ed209948
# <翻译结束>


<原文开始>
		// filter conflict keys in updateValues.
		// And the key is not a soft created field.
<原文结束>

# <翻译开始>
		// 过滤掉更新值中的冲突键。
		// 并且该键不是软创建字段。
		// md5:7882adbf4107a87d
# <翻译结束>


<原文开始>
// parseSqlForUpsert
// MERGE INTO {{table}} T1
// USING ( SELECT {{queryHolders}} FROM DUAL T2
// ON (T1.{{duplicateKey}} = T2.{{duplicateKey}} AND ...)
// WHEN NOT MATCHED THEN
// INSERT {{insertKeys}} VALUES {{insertValues}}
// WHEN MATCHED THEN
// UPDATE SET {{updateValues}}
<原文结束>

# <翻译开始>
// parseSqlForUpsert
// MERGE INTO {{table}} T1
// 使用 ( SELECT {{queryHolders}} FROM DUAL T2
// ON (T1.{{duplicateKey}} = T2.{{duplicateKey}} AND ...)
// 当未找到匹配时
// INSERT {{insertKeys}} VALUES {{insertValues}}
// 当找到匹配时
// UPDATE SET {{updateValues}}
// md5:7a233cb2881f0359
# <翻译结束>

