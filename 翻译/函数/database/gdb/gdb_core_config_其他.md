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

[Host                 string        `json:"host"`]
qm=地址
cz=Host                 string        `json:"host"`

[Port                 string        `json:"port"`]
qm=端口
cz=Port                 string        `json:"port"`

[User                 string        `json:"user"`]
qm=账号
cz=User                 string        `json:"user"

[Pass                 string        `json:"pass"`]
qm=密码
cz=Pass                 string        `json:"pass"

[Name                 string        `json:"name"`]
qm=名称
cz=Name                 string        `json:"name"

[Type                 string        `json:"type"`]
qm=类型
cz=Type                 string        `json:"type"`

[Link                 string        `json:"link"`]
qm=自定义链接信息
cz=Link                 string        `json:"link"`

[Extra                string        `json:"extra"`]
qm=额外
cz=Extra                string        `json:"extra"

[Role                 string        `json:"role"`]
qm=节点角色
cz=Role                 string        `json:"role"`

[Debug                bool          `json:"debug"`]
qm=调试模式
cz=Debug                bool          `json:"debug"

[Prefix               string        `json:"prefix"`]
qm=表前缀
cz=Prefix               string        `json:"prefix"`

[DryRun               bool          `json:"dryRun"`]
qm=空跑特性
cz=DryRun               bool          `json:"dryRun"`

[Weight               int           `json:"weight"`]
qm=负载均衡权重
cz=Weight               int           `json:"weight"

[Charset              string        `json:"charset"`]
qm=字符集
cz=Charset              string        `json:"charset"

[Protocol             string        `json:"protocol"`]
qm=协议
cz=Protocol             string        `json:"protocol"`

[Timezone             string        `json:"timezone"`]
qm=时区
cz=Timezone             string        `json:"timezone"

[Namespace            string        `json:"namespace"`]
qm=命名空间
cz=Namespace            string        `json:"namespace"

[MaxIdleConnCount     int           `json:"maxIdle"`]
qm=最大闲置连接数
cz=MaxIdleConnCount     int           `json:"maxIdle"`

[MaxOpenConnCount     int           `json:"maxOpen"`]
qm=最大打开连接数
cz=MaxOpenConnCount     int           `json:"maxOpen"`

[MaxConnLifeTime      time.Duration `json:"maxLifeTime"`]
qm=最大空闲时长
cz=MaxConnLifeTime      time.Duration `json:"maxLifeTime"`

[QueryTimeout         time.Duration `json:"queryTimeout"`]
qm=查询超时时长
cz=QueryTimeout         time.Duration `json:"queryTimeout"

[ExecTimeout          time.Duration `json:"execTimeout"`]
qm=执行超时时长
cz=ExecTimeout          time.Duration `json:"execTimeout"

[TranTimeout          time.Duration `json:"tranTimeout"`]
qm=事务超时时长
cz=TranTimeout          time.Duration `json:"tranTimeout"

[PrepareTimeout       time.Duration `json:"prepareTimeout"`]
qm=预准备SQL超时时长
cz=PrepareTimeout       time.Duration `json:"prepareTimeout"

[CreatedAt            string        `json:"createdAt"`]
qm=创建时间字段名
cz=CreatedAt            string        `json:"createdAt"

[UpdatedAt            string        `json:"updatedAt"`]
qm=更新时间字段名
cz=UpdatedAt            string        `json:"updatedAt

[DeletedAt            string        `json:"deletedAt"`]
qm=软删除时间字段名
cz=DeletedAt            string        `json:"deletedAt"`

[TimeMaintainDisabled bool          `json:"timeMaintainDisabled"`]
qm=禁用时间自动更新特性
cz=TimeMaintainDisabled bool          `json:"timeMaintainDisabled"
