
# <追加函数开始>
// 追加函数
func (j *Json) X取文本() string {
	return j.String()
}
# <追加函数结束>

# <追加函数开始>
// 追加函数
func (j *Json) X取any数组() []interface{} {
return j.Interfaces()
}
# <追加函数结束>

# <追加函数开始>
// 追加函数
func (j *Json) X取MapStrAny() map[string]interface{} {
return j.MapStrAny()
}
# <追加函数结束>
