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

[func New() *Locker {]
ff=创建

[func (l *Locker) Lock(key string) {]
ff=写锁定
key=名称

[func (l *Locker) TryLock(key string) bool {]
ff=非阻塞写锁定
key=名称

[func (l *Locker) Unlock(key string) {]
ff=退出写锁定
key=名称

[func (l *Locker) RLock(key string) {]
ff=读锁定
key=名称

[func (l *Locker) TryRLock(key string) bool {]
ff=非阻塞读锁定
key=名称

[func (l *Locker) RUnlock(key string) {]
ff=退出读锁定
key=名称

[func (l *Locker) LockFunc(key string, f func()) {]
ff=写锁定_函数
f=回调函数
key=名称

[func (l *Locker) RLockFunc(key string, f func()) {]
ff=读锁定_函数
f=回调函数
key=名称

[func (l *Locker) TryLockFunc(key string, f func()) bool {]
ff=非阻塞写锁定_函数
f=回调函数
key=名称

[func (l *Locker) TryRLockFunc(key string, f func()) bool {]
ff=非阻塞读锁定_函数
f=回调函数
key=名称

[func (l *Locker) Remove(key string) {]
ff=删除锁
key=名称

[func (l *Locker) Clear() {]
ff=移除所有锁
