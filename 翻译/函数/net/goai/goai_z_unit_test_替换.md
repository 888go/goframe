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

[t.Assert(oai.Components.Schemas.Get(`github.com.gogf.gf.v2.net.goai_test.CreateResourceReq`).Value.Type, goai.TypeObject)]
th=github.com.888go.goframe.net.goai_test.CreateResourceReq
cz=github.com.gogf.gf.v2.net.goai_test.CreateResourceReq

[t.Assert(len(oai.Components.Schemas.Get(`github.com.gogf.gf.v2.net.goai_test.SetSpecInfo`).Value.Properties.Map()), 3)]
th=github.com.888go.goframe.net.goai_test.SetSpecInfo
cz=github.com.gogf.gf.v2.net.goai_test.SetSpecInfo

[t.Assert(oai.Components.Schemas.Get(`github.com.gogf.gf.v2.net.goai_test.CategoryTreeItem`).Value.Type, goai.TypeObject)]
th=github.com.888go.goframe.net.goai_test.CategoryTreeItem
cz=github.com.gogf.gf.v2.net.goai_test.CategoryTreeItem

[t.Assert(b, `{"openapi":"3.0.0","components":{"schemas":{"github.com.gogf.gf.v2.net.goai_test.CreateReq":{"properties":{"nick_name":{"format":"string","properties":{},"type":"string"}},"type":"object"}}},"info":{"title":"","version":""},"paths":null}`)]
th=github.com.888go.goframe.net.goai_test.CreateReq
cz=github.com.gogf.gf.v2.net.goai_test.CreateReq

["github.com/gogf/gf/v2/net/goai_test.Status": #左中括号##右中括号#interface{}{StatusA, StatusB},]
th=github.com/888go/goframe/net/goai_test.Status
cz=github.com/gogf/gf/v2/net/goai_test.Status

[var reqKey = "github.com.gogf.gf.v2.net.goai_test.Req"]
th=github.com.888go.goframe.net.goai_test.Req
cz=github.com.gogf.gf.v2.net.goai_test.Req

[t.Assert(oai.String(), `{"openapi":"3.0.0","components":{"schemas":{"github.com.gogf.gf.v2.net.goai_test.GetListReq":{"properties":{"Page":{"default":1,"description":"Page number","format":"int","properties":{},"type":"integer","x-sort":"1"},"Size":{"default":10,"description":"Size for per page.","format":"int","properties":{},"type":"integer","x-sort":"2"}},"type":"object","x-group":"User/Info"}}},"info":{"title":"","version":""},"paths":null}`)]
th=github.com.888go.goframe.net.goai_test.GetListReq
cz=github.com.gogf.gf.v2.net.goai_test.GetListReq
