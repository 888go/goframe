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

[func CaseTypeMatch(caseStr string) CaseType {]
ff=命名方式判断
caseStr=待判断名称

[func CaseConvert(s string, caseType CaseType) string {]
ff=命名转换
caseType=类型
s=待转换文本

[func CaseCamel(s string) string {]
ff=命名转换到首字母大写驼峰
s=待转换文本

[func CaseCamelLower(s string) string {]
ff=命名转换到首字母小写驼峰
s=待转换文本

[func CaseSnake(s string) string {]
ff=命名转换到全小写蛇形
s=待转换文本

[func CaseSnakeScreaming(s string) string {]
ff=命名转换到大写蛇形
s=待转换文本

[func CaseSnakeFirstUpper(word string, underscore ...string) string {]
ff=命名转换到全小写蛇形2
underscore=可选连接符
word=待转换文本

[func CaseKebab(s string) string {]
ff=命名转换到小写短横线
s=待转换文本

[func CaseKebabScreaming(s string) string {]
ff=命名转换到大写驼峰短横线
s=待转换文本

[func CaseDelimited(s string, del byte) string {]
ff=命名转换按符号
del=连接符号
s=待转换文本

[func CaseDelimitedScreaming(s string, del uint8, screaming bool) string {]
ff=命名转换按符号与大小写
screaming=是否全大写
del=连接符号
s=待转换文本
