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

[func Encrypt(data interface{}) (encrypt string, err error) {]
err=错误
encrypt=md5值
data=值

[func MustEncrypt(data interface{}) string {]
ff=加密PANI
data=值

[func EncryptBytes(data #左中括号##右中括号#byte) (encrypt string, err error) {]
ff=加密字节集
err=错误
encrypt=md5值
data=字节集

[func MustEncryptBytes(data #左中括号##右中括号#byte) string {]
ff=加密字节集PANI
data=字节集

[func EncryptString(data string) (encrypt string, err error) {]
ff=加密文本
err=错误
encrypt=md5值
data=值

[func MustEncryptString(data string) string {]
ff=加密文本PANI
data=值

[func EncryptFile(path string) (encrypt string, err error) {]
ff=加密文件
err=错误
encrypt=md5值
path=路径

[func MustEncryptFile(path string) string {]
ff=加密文件PANI
path=路径
