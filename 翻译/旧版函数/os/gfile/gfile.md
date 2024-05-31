
# <翻译开始>
func Mkdir(path string) (err
错误
# <翻译结束>

# <翻译开始>
func Mkdir(path
目录
# <翻译结束>

# <翻译开始>
func Mkdir
X创建目录
# <翻译结束>

# <翻译开始>
func Create(path
文件路径
# <翻译结束>

# <翻译开始>
func Create
X创建文件与目录
# <翻译结束>

# <翻译开始>
func Open(path
路径
# <翻译结束>

# <翻译开始>
func Open
X打开并按只读模式
# <翻译结束>

# <翻译开始>
func OpenFile(path string, flag int, perm
权限模式
# <翻译结束>

# <翻译开始>
func OpenFile(path string, flag
读写模式
# <翻译结束>

# <翻译开始>
func OpenFile(path
路径
# <翻译结束>

# <翻译开始>
func OpenFile
X打开
# <翻译结束>

# <翻译开始>
func OpenWithFlag(path string, flag
读写模式
# <翻译结束>

# <翻译开始>
func OpenWithFlag(path
路径
# <翻译结束>

# <翻译开始>
func OpenWithFlag
X打开并按默认权限
# <翻译结束>

# <翻译开始>
func OpenWithFlagPerm
OpenWithFlagPerm别名
# <翻译结束>

# <翻译开始>
func Join(paths
路径s
# <翻译结束>

# <翻译开始>
func Join
X路径生成
# <翻译结束>

# <翻译开始>
func Exists(path
路径
# <翻译结束>

# <翻译开始>
func Exists
X是否存在
# <翻译结束>

# <翻译开始>
func IsDir(path
路径
# <翻译结束>

# <翻译开始>
func IsDir
X是否存在目录
# <翻译结束>

# <翻译开始>
func Pwd
X取当前工作目录
# <翻译结束>

# <翻译开始>
func Chdir(dir string) (err
错误
# <翻译结束>

# <翻译开始>
func Chdir(dir
目录
# <翻译结束>

# <翻译开始>
func Chdir
X设置当前工作目录
# <翻译结束>

# <翻译开始>
func IsFile(path
路径
# <翻译结束>

# <翻译开始>
func IsFile
X是否为文件
# <翻译结束>

# <翻译开始>
func Stat(path
路径
# <翻译结束>

# <翻译开始>
func Stat
X取详情
# <翻译结束>

# <翻译开始>
func Move(src string, dst string) (err
错误
# <翻译结束>

# <翻译开始>
func Move(src string, dst
新路径
# <翻译结束>

# <翻译开始>
func Move(src
路径
# <翻译结束>

# <翻译开始>
func Move
X移动
# <翻译结束>

# <翻译开始>
func Rename
Rename别名
# <翻译结束>

# <翻译开始>
func DirNames(path
路径
# <翻译结束>

# <翻译开始>
func DirNames
X取文件列表
# <翻译结束>

# <翻译开始>
func Glob(pattern string, onlyNames
返回绝对路径
# <翻译结束>

# <翻译开始>
func Glob(pattern
路径
# <翻译结束>

# <翻译开始>
func Glob
X模糊查找
# <翻译结束>

# <翻译开始>
func Remove(path string) (err
错误
# <翻译结束>

# <翻译开始>
func Remove(path
路径或文件夹
# <翻译结束>

# <翻译开始>
func Remove
X删除
# <翻译结束>

# <翻译开始>
func IsReadable(path
路径
# <翻译结束>

# <翻译开始>
func IsReadable
X是否可读
# <翻译结束>

# <翻译开始>
func IsWritable(path
路径
# <翻译结束>

# <翻译开始>
func IsWritable
X是否可写
# <翻译结束>

# <翻译开始>
func Chmod(path string, mode os.FileMode) (err
错误
# <翻译结束>

# <翻译开始>
func Chmod(path string, mode
权限模式
# <翻译结束>

# <翻译开始>
func Chmod(path
路径
# <翻译结束>

# <翻译开始>
func Chmod
X更改权限
# <翻译结束>

# <翻译开始>
func Abs(path
路径
# <翻译结束>

# <翻译开始>
func Abs
X取绝对路径
# <翻译结束>

# <翻译开始>
func RealPath(path
路径
# <翻译结束>

# <翻译开始>
func RealPath
X取绝对路径且效验
# <翻译结束>

# <翻译开始>
func SelfPath
X取当前进程路径
# <翻译结束>

# <翻译开始>
func SelfName
X取当前进程名
# <翻译结束>

# <翻译开始>
func SelfDir
X取当前进程目录
# <翻译结束>

# <翻译开始>
func Basename(path
路径
# <翻译结束>

# <翻译开始>
func Basename
X路径取文件名
# <翻译结束>

# <翻译开始>
func Name(path
路径
# <翻译结束>

# <翻译开始>
func Name
X路径取文件名且不含扩展名
# <翻译结束>

# <翻译开始>
func Dir(path
路径
# <翻译结束>

# <翻译开始>
func Dir
X路径取父目录
# <翻译结束>

# <翻译开始>
func IsEmpty(path
路径
# <翻译结束>

# <翻译开始>
func IsEmpty
X是否为空
# <翻译结束>

# <翻译开始>
func Ext(path
路径
# <翻译结束>

# <翻译开始>
func Ext
X路径取扩展名
# <翻译结束>

# <翻译开始>
func ExtName(path
路径
# <翻译结束>

# <翻译开始>
func ExtName
X路径取扩展名且不含点号
# <翻译结束>

# <翻译开始>
func Temp(names
可选路径
# <翻译结束>

# <翻译开始>
func Temp
X取临时目录
# <翻译结束>
