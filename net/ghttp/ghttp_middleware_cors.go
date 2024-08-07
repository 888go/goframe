// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package http类

// X中间件跨域函数 是一个使用默认选项的CORS中间件处理器。 md5:522ca96c4772b84d
func X中间件跨域函数(r *Request) {
	r.X响应.X跨域请求全允许()
	r.X中间件管理器.Next()
}
