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

[func Intn(max int) int {]
max=最大值

[func B(n int) #左中括号##右中括号#byte {]
ff=字节集
n=长度

[func N(min, max int) int {]
ff=区间整数
max=最大值
min=最小值

[func S(n int, symbols ...bool) string {]
ff=文本
symbols=包含特殊字符
n=长度

[func D(min, max time.Duration) time.Duration {]
ff=时长
max=最大值
min=最小值

[func Str(s string, n int) string {]
ff=从文本生成文本
n=长度
s=给定文本

[func Digits(n int) string {]
ff=数字文本
n=长度

[func Letters(n int) string {]
ff=字母文本
n=长度

[func Symbols(n int) string {]
ff=特殊字符文本
n=长度

[func Perm(n int) #左中括号##右中括号#int {]
ff=整数数组
n=长度
