
<原文开始>
// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
<原文结束>

# <翻译开始>
// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。 md5:a9832f33b234e3f3
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
// 尚不支持此类转换。
// 函数 Test_Params_Parse_Attr_SliceSlice(t *testing.T) {
// 类型 User 结构体 {
//     Id     int
//     Name   string
//     Scores [][]int
// }
// // s := g.Server(guid.S())
// s.BindHandler("/parse", func(r *ghttp.Request) {
//     如果 m := r.GetMap(); m 的长度 > 0 {
//         var user *User
//         如果 err := r.Parse(&user); err 不等于 nil {
//             r.Response.WriteExit(err)
//         }
//         r.Response.WriteExit(user.Scores)
//     }
// })
// // s.SetDumpRouterMap(false)
// s.Start()
// defer s.Shutdown()
//
// time.Sleep(100 * time.Millisecond)
// gtest.C(t, func(t *gtest.T) {
//     client := g.Client()
//     client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", s.GetListenedPort()))
//     t.Assert(client.PostContent(ctx, "/parse", `{"id":1,"name":"john","scores":[[1,2,3]]}`), `1100`)
// })
// } md5:65bedf537ac79a00
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1488
<原文结束>

# <翻译开始>
// 这段注释引用的是一个GitHub问题（Issue）的链接，来自 "gf"（Go Foundation）项目。gf是一个用Go语言编写的轻量级、高性能的Web框架。1488号问题可能是指该框架中某个特定的问题或者讨论点。要了解详细内容，需要访问该链接查看相关讨论。 md5:dd9dcdabfe17c38c
# <翻译结束>

