# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
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
# re.F.String()
# }
# //zj:
# 备注结束

[func (m *Model) Batch(batch int) *Model {]
ff=设置批量操作行数
batch=数量

[func (m *Model) Data(data ...interface{}) *Model {]
ff=设置数据
data=值

[func (m *Model) OnDuplicate(onDuplicate ...interface{}) *Model {]
ff=设置插入冲突更新字段
onDuplicate=字段名称

[func (m *Model) OnDuplicateEx(onDuplicateEx ...interface{}) *Model {]
ff=设置插入冲突不更新字段
onDuplicateEx=字段名称

[func (m *Model) Insert(data ...interface{}) (result sql.Result, err error) {]
ff=插入
err=错误
result=结果
data=值

[func (m *Model) InsertAndGetId(data ...interface{}) (lastInsertId int64, err error) {]
ff=插入并取ID
err=错误
lastInsertId=最后插入ID
data=值

[func (m *Model) InsertIgnore(data ...interface{}) (result sql.Result, err error) {]
ff=插入并跳过已存在
err=错误
result=结果
data=值

[func (m *Model) Replace(data ...interface{}) (result sql.Result, err error) {]
ff=插入并替换已存在
err=错误
result=结果
data=值

[func (m *Model) Save(data ...interface{}) (result sql.Result, err error) {]
ff=插入并更新已存在
err=错误
result=结果
data=值
