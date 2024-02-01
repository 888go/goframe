
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。
# <翻译结束>


<原文开始>
// Domain is used for route register for domains.
<原文结束>

# <翻译开始>
// Domain 用于为域名进行路由注册。
# <翻译结束>







<原文开始>
// Domain creates and returns a domain object for management for one or more domains.
<原文结束>

# <翻译开始>
// Domain 创建并返回一个域名管理对象，用于管理一个或多个域名。
# <翻译结束>


<原文开始>
// BindHandler binds the handler for the specified pattern.
<原文结束>

# <翻译开始>
// BindHandler 为指定模式绑定处理器。
# <翻译结束>


<原文开始>
// BindObject binds the object for the specified pattern.
<原文结束>

# <翻译开始>
// BindObject 为指定的模式绑定对象。
//
// BindObjectMethod和BindObject的区别：
// BindObjectMethod将对象中的指定方法与指定路由规则进行绑定，第三个method参数只能指定一个方法名称；
// BindObject注册时，所有的路由都是对象方法名称按照规则生成的，第三个methods参数可以指定多个注册的方法名称。
# <翻译结束>


<原文开始>
// BindObjectMethod binds the method for the specified pattern.
<原文结束>

# <翻译开始>
// BindObjectMethod 将指定模式的方法绑定。
// 
// BindObjectMethod和BindObject的区别：
// BindObjectMethod将对象中的指定方法与指定路由规则进行绑定，第三个method参数只能指定一个方法名称；
// BindObject注册时，所有的路由都是对象方法名称按照规则生成的，第三个methods参数可以指定多个注册的方法名称。
# <翻译结束>


<原文开始>
// BindObjectRest binds the RESTful API for the specified pattern.
<原文结束>

# <翻译开始>
// BindObjectRest 为指定模式绑定RESTful API。
// RESTful设计方式的控制器，通常用于API服务。
// 在这种模式下，HTTP的Method将会映射到控制器对应的方法名称，
// 例如：POST方式将会映射到控制器的Post方法中(公开方法，首字母大写)，DELETE方式将会映射到控制器的Delete方法中，以此类推。
// 其他非HTTP Method命名的方法，即使是定义的包公开方法，将不会自动注册，对于应用端不可见。
// 当然，如果控制器并未定义对应HTTP Method的方法，该Method请求下将会返回 HTTP Status 404。
# <翻译结束>


<原文开始>
// BindHookHandler binds the hook handler for the specified pattern.
<原文结束>

# <翻译开始>
// BindHookHandler 为指定模式绑定钩子处理器。
# <翻译结束>


<原文开始>
// BindHookHandlerByMap binds the hook handler for the specified pattern.
<原文结束>

# <翻译开始>
// BindHookHandlerByMap 通过映射为特定模式绑定钩子处理器。
# <翻译结束>


<原文开始>
// BindStatusHandler binds the status handler for the specified pattern.
<原文结束>

# <翻译开始>
// BindStatusHandler 为指定模式绑定状态处理器。
# <翻译结束>


<原文开始>
// BindStatusHandlerByMap binds the status handler for the specified pattern.
<原文结束>

# <翻译开始>
// BindStatusHandlerByMap 通过给定的模式绑定状态处理器。
# <翻译结束>


<原文开始>
// BindMiddleware binds the middleware for the specified pattern.
<原文结束>

# <翻译开始>
// BindMiddleware 为指定模式绑定中间件。
# <翻译结束>


<原文开始>
// BindMiddlewareDefault binds the default middleware for the specified pattern.
<原文结束>

# <翻译开始>
// BindMiddlewareDefault 为指定模式绑定默认中间件。
# <翻译结束>


<原文开始>
// Use adds middleware to the domain.
<原文结束>

# <翻译开始>
// Use 向域名添加中间件。
// Use 是 BindMiddlewareDefault 的别名。
// 请参阅 BindMiddlewareDefault。
# <翻译结束>







<原文开始>
// Support multiple domains.
<原文结束>

# <翻译开始>
// 支持多个域名。
# <翻译结束>

