# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如:
# //ff:取文本

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: 
# package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如:
# type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
#    re.F.String()
# }
# //zj:
# 备注结束

[func (m *Model) All(where ...interface{}) (Result, error) {]
ff=查询
where=查询条件

[func (m *Model) AllAndCount(useFieldForCount bool) (result Result, totalCount int, err error) {]
ff=查询与行数
err=错误
totalCount=行数
result=结果
useFieldForCount=是否用字段计数

[func (m *Model) Chunk(size int, handler ChunkHandler) {]
ff=分割
handler=处理函数
size=数量

[func (m *Model) One(where ...interface{}) (Record, error) {]
ff=查询一条
where=条件

[func (m *Model) Array(fieldsAndWhere ...interface{}) (#左中括号##右中括号#Value, error) {]
ff=查询数组
fieldsAndWhere=条件

[func (m *Model) Scan(pointer interface{}, where ...interface{}) error {]
ff=查询到结构体指针
where=条件
pointer=数据指针

[func (m *Model) ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount bool) (err error) {]
ff=查询与行数到指针
err=错误
useFieldForCount=是否用字段计数
totalCount=行数指针
pointer=数据指针

[func (m *Model) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error) {]
ff=查询到指针列表
err=错误
relationAttrNameAndFields=结构体属性关联
bindToAttrName=绑定到结构体属性名称
structSlicePointer=结构体切片指针

[func (m *Model) Value(fieldsAndWhere ...interface{}) (Value, error) {]
ff=查询一条值
fieldsAndWhere=字段和条件

[func (m *Model) Count(where ...interface{}) (int, error) {]
ff=查询行数
where=条件

[func (m *Model) CountColumn(column string) (int, error) {]
ff=查询字段行数
column=字段名称

[func (m *Model) Min(column string) (float64, error) {]
ff=查询最小值
column=字段名称

[func (m *Model) Max(column string) (float64, error) {]
ff=查询最大值
column=字段名称

[func (m *Model) Avg(column string) (float64, error) {]
ff=查询平均值
column=字段名称

[func (m *Model) Sum(column string) (float64, error) {]
ff=查询求和
column=字段名称

[func (m *Model) Union(unions ...*Model) *Model {]
ff=多表去重查询
unions=Model对象

[func (m *Model) UnionAll(unions ...*Model) *Model {]
ff=多表查询
unions=Model对象

[func (m *Model) Limit(limit ...int) *Model {]
ff=设置条数
limit=条数或两个数字

[func (m *Model) Distinct() *Model {]
ff=设置去重

[func (m *Model) Page(page, limit int) *Model {]
ff=设置分页
limit=条数
page=第几页

[func (m *Model) Having(having interface{}, args ...interface{}) *Model {]
ff=设置分组条件
args=参数
having=条件
