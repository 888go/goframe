
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
// It does not support this kind of converting yet.
// func Test_Params_Parse_Attr_SliceSlice(t *testing.T) {
//	type User struct {
//		Id     int
//		Name   string
//		Scores [][]int
//	}
//	//	s := g.Server(guid.S())
//	s.BindHandler("/parse", func(r *ghttp.Request) {
//		if m := r.GetMap(); len(m) > 0 {
//			var user *User
//			if err := r.Parse(&user); err != nil {
//				r.Response.WriteExit(err)
//			}
//			r.Response.WriteExit(user.Scores)
//		}
//	})
//	//	s.SetDumpRouterMap(false)
//	s.Start()
//	defer s.Shutdown()
//
//	time.Sleep(100 * time.Millisecond)
//	gtest.C(t, func(t *gtest.T) {
//		client := g.Client()
//		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
//		t.Assert(client.PostContent(ctx, "/parse", `{"id":1,"name":"john","scores":[[1,2,3]]}`), `1100`)
//	})
// }
<原文结束>

# <翻译开始>
// 目前还不支持这种类型的转换。
// func Test_Params_Parse_Attr_SliceSlice(t *testing.T) {
// 定义一个User结构体，包含Id（整型）、Name（字符串）和Scores（二维整数数组）属性
// }
//// 创建gServer实例（注释已被省略）
// s := g.Server(guid.S())
// 为"/parse"路径绑定处理器函数
// s.BindHandler("/parse", func(r *ghttp.Request) {
// 获取请求参数映射
// 如果参数映射不为空
// 定义一个指向User类型的指针变量user
// 尝试解析请求参数到user变量中，如果解析出错
// 则向响应写入错误信息并结束处理
// 解析成功则向响应写入user的Scores属性值并结束处理
// })
//// 设置不输出路由映射信息（注释已被省略）
// s.SetDumpRouterMap(false)
// 启动服务器
// 在程序结束时关闭服务器
// s.Start()
// defer s.Shutdown()
// 暂停100毫秒
// time.Sleep(100 * time.Millisecond)
// 使用gtest进行测试
// gtest.C(t, func(t *gtest.T) {
// 创建gClient实例
// 设置客户端请求前缀为服务器地址与端口
// 发送POST请求到"/parse"，内容为指定的JSON格式数据
// 断言返回的内容应为"1100"
// })
// }
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1488
<原文结束>

# <翻译开始>
// 这是GitHub上gogf/gf仓库中第1488号问题的链接
# <翻译结束>

