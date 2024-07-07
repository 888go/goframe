
# <翻译开始>
) All(where
查询条件
# <翻译结束>

# <翻译开始>
) All
X查询
# <翻译结束>

# <翻译开始>
) AllAndCount(useFieldForCount bool) (result Result, totalCount int, err
错误
# <翻译结束>

# <翻译开始>
) AllAndCount(useFieldForCount bool) (result Result, totalCount
行数
# <翻译结束>

# <翻译开始>
) AllAndCount(useFieldForCount bool) (result
结果
# <翻译结束>

# <翻译开始>
) AllAndCount(useFieldForCount
是否用字段计数
# <翻译结束>

# <翻译开始>
) AllAndCount
X查询与行数
# <翻译结束>

# <翻译开始>
) Chunk(size int, handler
处理函数
# <翻译结束>

# <翻译开始>
) Chunk(size
数量
# <翻译结束>

# <翻译开始>
) Chunk
X分割
# <翻译结束>

# <翻译开始>
) One(where
条件
# <翻译结束>

# <翻译开始>
) One
X查询一条
# <翻译结束>

# <翻译开始>
) Array(fieldsAndWhere
条件
# <翻译结束>

# <翻译开始>
) Array
X查询切片
# <翻译结束>

# <翻译开始>
) Scan(pointer interface{}, where
条件
# <翻译结束>

# <翻译开始>
) Scan(pointer
数据指针
# <翻译结束>

# <翻译开始>
func (m *Model) Scan
X查询到结构体指针
# <翻译结束>

# <翻译开始>
ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount bool) (err
错误
# <翻译结束>

# <翻译开始>
) ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount
是否用字段计数
# <翻译结束>

# <翻译开始>
) ScanAndCount(pointer interface{}, totalCount
行数指针
# <翻译结束>

# <翻译开始>
) ScanAndCount(pointer
数据指针
# <翻译结束>

# <翻译开始>
) ScanAndCount
X查询与行数到指针
# <翻译结束>

# <翻译开始>
) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err
错误
# <翻译结束>

# <翻译开始>
) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields
结构体属性关联
# <翻译结束>

# <翻译开始>
) ScanList(structSlicePointer interface{}, bindToAttrName
绑定到结构体属性名称
# <翻译结束>

# <翻译开始>
) ScanList(structSlicePointer
结构体切片指针
# <翻译结束>

# <翻译开始>
) ScanList
X查询到指针列表
# <翻译结束>

# <翻译开始>
) Value(fieldsAndWhere
字段和条件
# <翻译结束>

# <翻译开始>
) Value
X查询一条值
# <翻译结束>

# <翻译开始>
) Count(where
条件
# <翻译结束>

# <翻译开始>
) Count
X查询行数
# <翻译结束>

# <翻译开始>
) CountColumn(column
字段名称
# <翻译结束>

# <翻译开始>
) CountColumn
X查询字段行数
# <翻译结束>

# <翻译开始>
) Min(column
字段名称
# <翻译结束>

# <翻译开始>
) Min
X查询最小值
# <翻译结束>

# <翻译开始>
) Max(column
字段名称
# <翻译结束>

# <翻译开始>
) Max
X查询最大值
# <翻译结束>

# <翻译开始>
) Avg(column
字段名称
# <翻译结束>

# <翻译开始>
) Avg
X查询平均值
# <翻译结束>

# <翻译开始>
) Sum(column
字段名称
# <翻译结束>

# <翻译开始>
) Sum
X查询求和
# <翻译结束>

# <翻译开始>
) Union(unions
Model对象
# <翻译结束>

# <翻译开始>
) Union
X多表去重查询
# <翻译结束>

# <翻译开始>
) UnionAll(unions
Model对象
# <翻译结束>

# <翻译开始>
) UnionAll
X多表查询
# <翻译结束>

# <翻译开始>
) Limit(limit
条数或两个数字
# <翻译结束>

# <翻译开始>
) Limit
X设置条数
# <翻译结束>

# <翻译开始>
) Distinct
X设置去重
# <翻译结束>

# <翻译开始>
) Page(page, limit
条数
# <翻译结束>

# <翻译开始>
) Page(page
第几页
# <翻译结束>

# <翻译开始>
) Page
X设置分页
# <翻译结束>

# <翻译开始>
) Having(having interface{}, args
参数
# <翻译结束>

# <翻译开始>
) Having(having
条件
# <翻译结束>

# <翻译开始>
) Having
X设置分组条件
# <翻译结束>
