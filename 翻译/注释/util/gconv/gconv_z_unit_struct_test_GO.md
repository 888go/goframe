
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
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3
# <翻译结束>


<原文开始>
// It does not support this kind of converting yet.
//func Test_Struct_Attr_Slice2(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		type User struct {
//			Scores [][]int
//		}
//		scores := []interface{}{[]interface{}{99, 100, 60, 140}}
//		user := new(User)
//		if err := gconv.Struct(g.Map{"Scores": scores}, user); err != nil {
//			t.Error(err)
//		} else {
//			t.Assert(user, &User{
//				Scores: [][]int{{99, 100, 60, 140}},
//			})
//		}
//	})
//}
<原文结束>

# <翻译开始>
// 目前还不支持这种转换。
//func Test_Struct_Attr_Slice2(t *testing.T) {
//	gtest.C(t, func(t *gtest.T) {
//		// 定义一个User结构体
//		type User struct {
//			Scores [][]int
//		}
//		// 准备一个scores切片，其中包含一个元素为接口切片，该接口切片内部是整数
//		scores := []interface{}{[]interface{}{99, 100, 60, 140}}
//		// 创建一个新的User实例
//		user := new(User)
//		// 尝试使用gconv.Struct将包含scores的映射转换为User实例
//		if err := gconv.Struct(g.Map{"Scores": scores}, user); err != nil {
//			// 如果转换出错，则在测试中报错
//			t.Error(err)
//		} else {
//			// 检查转换后的user实例是否与预期相符
//			t.Assert(user, &User{
//				Scores: [][]int{{99, 100, 60, 140}},
//			})
//		}
//	})
//}
// md5:530558b91cc069c7
# <翻译结束>


<原文开始>
// From: k8s.io/apimachinery@v0.22.0/pkg/apis/meta/v1/duration.go
<原文结束>

# <翻译开始>
// 从：k8s.io/apimachinery@v0.22.0/pkg/apis/meta/v1/duration.go. md5:08be22594408ed2b
# <翻译结束>


<原文开始>
// UnmarshalJSON implements the json.Unmarshaller interface.
<原文结束>

# <翻译开始>
// UnmarshalJSON 实现了 json.Unmarshaler 接口。. md5:7a13933e6d93ee98
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/775
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/775
// 
// 这个注释链接指向的是GitHub上的一个 issues（问题或讨论），gf（GoGF）是一个用Go语言编写的Web框架。775号issue可能是指该框架中的某个特定问题或者讨论的编号，具体内容需要查看链接才能得知。. md5:aad5ea6b7b59206e
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1387
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/1387
// 
// 这段注释链接指向的是GitHub上的一个Issue，GF（Go Foundation）是一个Go语言的库或框架。"1387"可能是Issue的编号。具体的内容需要查看该链接才能得知，大致意思是关于GF项目在1387号问题上的讨论、报告了一个错误或者提出了一个特性请求。. md5:7c877c3e7a856cb1
# <翻译结束>


<原文开始>
// Auto create struct when given pointer.
<原文结束>

# <翻译开始>
// 当给定指针时，自动创建结构体。. md5:950a2ff29f9f6798
# <翻译结束>


<原文开始>
// Implemented interface attribute.
<原文结束>

# <翻译开始>
// 实现的接口属性。. md5:cc0d5c4c14ca721e
# <翻译结束>


<原文开始>
// No implemented interface attribute.
<原文结束>

# <翻译开始>
// 没有实现的接口属性。. md5:441b8a44b6a99ee2
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1563
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/1563
// 
// 这段注释链接指向的是GitHub上的一个 issues 页面，其中包含关于 "gf"（Go语言框架）项目的问题1563。具体的内容需要点击链接查看，可能是一个bug报告、功能请求或者讨论点。. md5:40653716adf5b048
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/1597
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/1597 问题讨论. md5:934fd7bf7bd1b6b7
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/3449
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/3449
// 这段代码是链接到一个GitHub问题的注释，问题编号为3449。在Go语言中，这种注释用于提供有关代码的额外信息，比如引用外部资源、说明问题或者提供待解决的任务。由于内容是URL，它可能指向一个关于代码库gf的讨论、错误报告或改进请求。要查看具体的内容，需要访问该链接。. md5:641e41ce7485cc00
# <翻译结束>


<原文开始>
// https://github.com/gogf/gf/issues/2980
<原文结束>

# <翻译开始>
// https://github.com/gogf/gf/issues/2980
// 
// 这段注释是引用了GitHub上gf框架的一个问题链接，编号为2980。在Go语言中，这种注释通常用于指向相关问题、讨论或资源的链接，以便读者可以获取更多信息。. md5:5fa723daec963e70
# <翻译结束>

