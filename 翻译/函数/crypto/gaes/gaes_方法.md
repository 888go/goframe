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

[func Decrypt(cipherText #左中括号##右中括号#byte, key #左中括号##右中括号#byte, iv ...#左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=Decrypt别名

[func EncryptCBC(plainText #左中括号##右中括号#byte, key #左中括号##右中括号#byte, iv ...#左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=加密CBC
key=秘钥
plainText=待加密

[func DecryptCBC(cipherText #左中括号##右中括号#byte, key #左中括号##右中括号#byte, iv ...#左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=解密CBC
key=秘钥
cipherText=待解密

[func EncryptCFB(plainText #左中括号##右中括号#byte, key #左中括号##右中括号#byte, padding *int, iv ...#左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=加密CFB
key=秘钥
plainText=待加密

[func DecryptCFB(cipherText #左中括号##右中括号#byte, key #左中括号##右中括号#byte, unPadding int, iv ...#左中括号##右中括号#byte) (#左中括号##右中括号#byte, error) {]
ff=解密CFB
key=秘钥
cipherText=待解密
