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
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
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
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func New(data interface{}, safe ...bool) *Json {]
ff=创建
safe=并发安全
data=值

[func NewWithTag(data interface{}, tags string, safe ...bool) *Json {]
ff=创建并按类型标签
safe=并发安全
tags=类型标签
data=值

[func NewWithOptions(data interface{}, options Options) *Json {]
ff=创建并按选项
options=选项
data=值

[func Load(path string, safe ...bool) (*Json, error) {]
ff=加载文件
safe=并发安全
path=路径

[func LoadWithOptions(data interface{}, options Options) (*Json, error) {]
ff=加载并按选项
options=选项
data=值

[func LoadJson(data interface{}, safe ...bool) (*Json, error) {]
ff=加载json
safe=并发安全
data=值

[func LoadXml(data interface{}, safe ...bool) (*Json, error) {]
ff=加载xml
safe=并发安全
data=值

[func LoadIni(data interface{}, safe ...bool) (*Json, error) {]
ff=加载ini
safe=并发安全
data=值

[func LoadYaml(data interface{}, safe ...bool) (*Json, error) {]
ff=加载Yaml
safe=并发安全
data=值

[func LoadToml(data interface{}, safe ...bool) (*Json, error) {]
ff=加载Toml
safe=并发安全
data=值

[func LoadProperties(data interface{}, safe ...bool) (*Json, error) {]
ff=加载Properties
safe=并发安全
data=值

[func LoadContent(data interface{}, safe ...bool) (*Json, error) {]
ff=加载并自动识别格式
safe=并发安全
data=值

[func LoadContentType(dataType ContentType, data interface{}, safe ...bool) (*Json, error) {]
ff=加载并按格式
safe=并发安全
data=值
dataType=类型标签

[func IsValidDataType(dataType ContentType) bool {]
ff=检查类型
dataType=待判断值
