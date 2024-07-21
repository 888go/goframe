
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
// SELECT * FROM `user` WHERE `name`='john'
<原文结束>

# <翻译开始>
// 从"user"表中选择所有列，其中"name"为'john'. md5:032af229cd8affac
# <翻译结束>


<原文开始>
// SELECT * FROM `user_detail` WHERE `uid`=1
<原文结束>

# <翻译开始>
// 从"user_detail"表中选择所有列，WHERE子句的条件是`uid`等于1. md5:d5e73807445a5607
# <翻译结束>


<原文开始>
// SELECT * FROM `user_scores` WHERE `uid`=1
<原文结束>

# <翻译开始>
// 从`user_scores`表中SELECT * WHERE `uid`=1. md5:d5e5d47d2cdd7d33
# <翻译结束>


<原文开始>
// Result ScanList with struct elements and pointer attributes.
<原文结束>

# <翻译开始>
// Result 使用具有结构体元素和指针属性的ScanList。 md5:b23d106d13859ad5
# <翻译结束>


<原文开始>
// Result ScanList with pointer elements and pointer attributes.
<原文结束>

# <翻译开始>
// 使用指针元素和指针属性的ScanList结果。 md5:137ae715e99be611
# <翻译结束>


<原文开始>
// Result ScanList with struct elements and struct attributes.
<原文结束>

# <翻译开始>
// 使用结构体元素和属性扫描Result。 md5:3af2572786856fc5
# <翻译结束>


<原文开始>
// Result ScanList with pointer elements and struct attributes.
<原文结束>

# <翻译开始>
// 结果 ScanList 包含指针元素和结构体属性。 md5:57e6be396681268f
# <翻译结束>


<原文开始>
// Model ScanList with pointer elements and pointer attributes.
<原文结束>

# <翻译开始>
// Model 是一个 ScanList 类型，其中的元素和属性为指针。 md5:815b904cdb4dea16
# <翻译结束>


<原文开始>
	//db.SetDebug(true)
	// Result ScanList with struct elements and pointer attributes.
<原文结束>

# <翻译开始>
	// 将db的调试模式设置为true
	// 使用结构体元素和指针属性扫描结果列表
	// md5:2e0de268c0f1e08f
# <翻译结束>


<原文开始>
			// Detail.
			// _, err = db.Insert(ctx, tableUserDetail, g.Map{
			//	"uid":     i,
			//	"address": fmt.Sprintf(`address_%d`, i),
			// })
			// t.AssertNil(err)
			// Scores.
			// for j := 1; j <= 5; j++ {
			//	_, err = db.Insert(ctx, tableUserScores, g.Map{
			//		"uid":   i,
			//		"score": j,
			//	})
			//	t.AssertNil(err)
			// }
<原文结束>

# <翻译开始>
			// 细节。
			// 插入用户详细信息到数据库，其中uid为i，address为格式化的`address_i`
			// t.AssertNil(err) 验证错误是否为nil
			// 分数。
			// 循环遍历1到5，插入用户分数到数据库，其中uid为i，score为j
			// t.AssertNil(err) 验证错误是否为nil
			// md5:1b3376e15e2dc36d
# <翻译结束>


<原文开始>
// SELECT * FROM `user_scores` WHERE `uid` IN(1,2,3,4,5)
<原文结束>

# <翻译开始>
// 从 `user_scores` 表中 SELECT * WHERE `uid` 在 (1,2,3,4,5) 中. md5:b0a4359d4663bf31
# <翻译结束>


<原文开始>
// SELECT * FROM `user_detail` WHERE `uid` IN(1,2,3,4,5)
<原文结束>

# <翻译开始>
// 从'user_detail'表中选择所有列，其中`uid`在(1,2,3,4,5)范围内. md5:fc3208d19b9f10f6
# <翻译结束>

