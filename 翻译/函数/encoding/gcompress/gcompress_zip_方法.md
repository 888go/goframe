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

[func ZipPath(fileOrFolderPaths, dstFilePath string, prefix ...string) error {]
ff=Zip压缩文件
prefix=可选路径前缀
dstFilePath=压缩文件路径
fileOrFolderPaths=目录或文件

[func ZipPathWriter(fileOrFolderPaths string, writer io.Writer, prefix ...string) error {]
ff=Zip压缩文件到Writer
prefix=可选路径前缀
fileOrFolderPaths=目录或文件

[func ZipPathContent(fileOrFolderPaths string, prefix ...string) (#左中括号##右中括号#byte, error) {]
ff=Zip压缩文件到字节集
prefix=可选路径前缀
fileOrFolderPaths=目录或文件

[func UnZipFile(zippedFilePath, dstFolderPath string, zippedPrefix ...string) error {]
ff=Zip解压文件
zippedPrefix=可选路径前缀
dstFolderPath=解压目录
zippedFilePath=压缩包路径

[func UnZipContent(zippedContent #左中括号##右中括号#byte, dstFolderPath string, zippedPrefix ...string) error {]
ff=Zip解压字节集
zippedPrefix=可选路径前缀
dstFolderPath=解压目录
zippedContent=zip字节集
