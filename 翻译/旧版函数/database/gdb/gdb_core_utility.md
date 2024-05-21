
# <翻译开始>
) GetDB
X取DB对象
# <翻译结束>

# <翻译开始>
) GetLink(ctx context.Context, master
主节点
# <翻译结束>

# <翻译开始>
) GetLink(ctx
上下文
# <翻译结束>

# <翻译开始>
) GetLink
X取数据库链接对象
# <翻译结束>

# <翻译开始>
) MasterLink
X底层MasterLink
# <翻译结束>

# <翻译开始>
) SlaveLink
X底层SlaveLink
# <翻译结束>

# <翻译开始>
) QuoteWord
X底层QuoteWord
# <翻译结束>

# <翻译开始>
) QuoteString
X底层QuoteString
# <翻译结束>

# <翻译开始>
) QuotePrefixTableName(table
表名称
# <翻译结束>

# <翻译开始>
) QuotePrefixTableName
X底层添加前缀字符和引用字符
# <翻译结束>

# <翻译开始>
) GetChars() (charLeft string, charRight
右字符
# <翻译结束>

# <翻译开始>
) GetChars() (charLeft
左字符
# <翻译结束>

# <翻译开始>
) GetChars
X底层取数据库安全字符
# <翻译结束>

# <翻译开始>
) Tables(ctx context.Context, schema ...string) (tables []string, err
错误
# <翻译结束>

# <翻译开始>
) Tables(ctx context.Context, schema ...string) (tables
表名称数组
# <翻译结束>

# <翻译开始>
) Tables(ctx
上下文
# <翻译结束>

# <翻译开始>
) Tables
X取表名称数组
# <翻译结束>

# <翻译开始>
) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err
错误
# <翻译结束>

# <翻译开始>
) TableFields(ctx context.Context, table string, schema ...string) (fields
字段信息Map
# <翻译结束>

# <翻译开始>
) TableFields(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) TableFields(ctx
上下文
# <翻译结束>

# <翻译开始>
) TableFields
X取表字段信息Map
# <翻译结束>

# <翻译开始>
) ClearTableFields(ctx context.Context, table string, schema ...string) (err
错误
# <翻译结束>

# <翻译开始>
) ClearTableFields(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) ClearTableFields(ctx
上下文
# <翻译结束>

# <翻译开始>
) ClearTableFields
X删除表字段缓存
# <翻译结束>

# <翻译开始>
) ClearTableFieldsAll(ctx context.Context) (err
错误
# <翻译结束>

# <翻译开始>
) ClearTableFieldsAll(ctx
上下文
# <翻译结束>

# <翻译开始>
) ClearTableFieldsAll
X删除表字段所有缓存
# <翻译结束>

# <翻译开始>
) ClearCache(ctx context.Context, table string) (err
错误
# <翻译结束>

# <翻译开始>
) ClearCache(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) ClearCache(ctx
上下文
# <翻译结束>

# <翻译开始>
) ClearCache
X删除表查询缓存
# <翻译结束>

# <翻译开始>
) ClearCacheAll(ctx context.Context) (err
错误
# <翻译结束>

# <翻译开始>
) ClearCacheAll(ctx
上下文
# <翻译结束>

# <翻译开始>
) ClearCacheAll
X删除所有表查询缓存
# <翻译结束>

# <翻译开始>
) HasField(ctx context.Context, table, field
字段名称
# <翻译结束>

# <翻译开始>
) HasField(ctx context.Context, table
表名称
# <翻译结束>

# <翻译开始>
) HasField(ctx
上下文
# <翻译结束>

# <翻译开始>
) HasField
X是否存在字段
# <翻译结束>
