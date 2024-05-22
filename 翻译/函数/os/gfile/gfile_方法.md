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
#    re.F.String()
# }
# //zj:
# 备注结束

[func Mkdir(path string) (err error) {]
ff=创建目录
err=错误
path=目录

[func Create(path string) (*os.File, error) {]
ff=创建文件与目录
path=文件路径

[func Open(path string) (*os.File, error) {]
ff=打开并按只读模式
path=路径

[func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {]
ff=打开
perm=权限模式
flag=读写模式
path=路径

[func OpenWithFlag(path string, flag int) (*os.File, error) {]
ff=打开并按默认权限
flag=读写模式
path=路径

[func OpenWithFlagPerm(path string, flag int, perm os.FileMode) (*os.File, error) {]
ff=OpenWithFlagPerm别名

[func Join(paths ...string) string {]
ff=路径生成
paths=路径s

[func Exists(path string) bool {]
ff=是否存在
path=路径

[func IsDir(path string) bool {]
ff=是否存在目录
path=路径

[func Pwd() string {]
ff=取当前工作目录

[func Chdir(dir string) (err error) {]
ff=设置当前工作目录
err=错误
dir=目录

[func IsFile(path string) bool {]
ff=是否为文件
path=路径

[func Stat(path string) (os.FileInfo, error) {]
ff=取详情
path=路径

[func Move(src string, dst string) (err error) {]
ff=移动
err=错误
dst=新路径
src=路径

[func Rename(src string, dst string) error {]
ff=Rename别名

[func DirNames(path string) (#左中括号##右中括号#string, error) {]
ff=取文件列表
path=路径

[func Glob(pattern string, onlyNames ...bool) (#左中括号##右中括号#string, error) {]
ff=模糊查找
onlyNames=返回绝对路径
pattern=路径

[func Remove(path string) (err error) {]
ff=删除
err=错误
path=路径或文件夹

[func IsReadable(path string) bool {]
ff=是否可读
path=路径

[func IsWritable(path string) bool {]
ff=是否可写
path=路径

[func Chmod(path string, mode os.FileMode) (err error) {]
ff=更改权限
err=错误
mode=权限模式
path=路径

[func Abs(path string) string {]
ff=取绝对路径
path=路径

[func RealPath(path string) string {]
ff=取绝对路径且效验
path=路径

[func SelfPath() string {]
ff=取当前进程路径

[func SelfName() string {]
ff=取当前进程名

[func SelfDir() string {]
ff=取当前进程目录

[func Basename(path string) string {]
ff=路径取文件名
path=路径

[func Name(path string) string {]
ff=路径取文件名且不含扩展名
path=路径

[func Dir(path string) string {]
ff=路径取父目录
path=路径

[func IsEmpty(path string) bool {]
ff=是否为空
path=路径

[func Ext(path string) string {]
ff=路径取扩展名
path=路径

[func ExtName(path string) string {]
ff=路径取扩展名且不含点号
path=路径

[func Temp(names ...string) string {]
ff=取临时目录
names=可选路径
