
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
// SELECT * FROM `user` WHERE `name`='john'
<原文结束>

# <翻译开始>
// 从`user`表中选择所有列，条件为`name`字段等于'john'
// 即：查询用户表中名为'john'的所有记录
# <翻译结束>


<原文开始>
// SELECT * FROM `user_detail` WHERE `uid`=1
<原文结束>

# <翻译开始>
// 从`user_detail`表中选取所有列，其条件为`uid`等于1
# <翻译结束>


<原文开始>
// SELECT * FROM `user_scores` WHERE `uid`=1
<原文结束>

# <翻译开始>
// 从`user_scores`表中选择所有列，其条件是`uid`等于1
# <翻译结束>


<原文开始>
// Result ScanList with struct elements and pointer attributes.
<原文结束>

# <翻译开始>
// Result ScanList，用于包含结构体元素和指针属性。
# <翻译结束>


<原文开始>
// Result ScanList with pointer elements and pointer attributes.
<原文结束>

# <翻译开始>
// Result ScanList，具有指针元素和指针属性。
# <翻译结束>


<原文开始>
// Result ScanList with struct elements and struct attributes.
<原文结束>

# <翻译开始>
// Result ScanList 用于包含结构体元素及结构体属性的扫描列表。
# <翻译结束>


<原文开始>
// Result ScanList with pointer elements and struct attributes.
<原文结束>

# <翻译开始>
// Result 扫描并生成一个具有指针元素和结构体属性的列表。
# <翻译结束>


<原文开始>
// Model ScanList with pointer elements and pointer attributes.
<原文结束>

# <翻译开始>
// Model ScanList，其中包含指针元素和指针属性。
# <翻译结束>


<原文开始>
	//db.SetDebug(true)
// Result ScanList with struct elements and pointer attributes.
<原文结束>

# <翻译开始>
// 设置数据库调试模式为开启状态
// db.SetDebug(true)
// 使用具有结构体元素和指针属性的Result ScanList
// 这段代码注释的翻译为：
// ```go
// 设置数据库调试模式为true（即开启调试）
// Result ScanList 用于处理具有结构体元素和指针属性的结果集
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
// 用户详情.
// _, err = db.Insert(ctx, tableUserDetail, g.Map{
//	"uid":     i, // 用户ID
//	"address": fmt.Sprintf(`address_%d`, i), // 格式化后的用户地址信息，索引为i
// })
// t.AssertNil(err) // 断言错误是否为nil，即检查插入用户详情操作是否成功
// 用户得分.
// for j := 1; j <= 5; j++ { // 遍历1到5的得分值
//	_, err = db.Insert(ctx, tableUserScores, g.Map{
//		"uid":   i, // 用户ID
//		"score": j, // 用户对应得分值
//	})
//	t.AssertNil(err) // 断言错误是否为nil，即检查插入用户得分记录操作是否成功
// }
# <翻译结束>


<原文开始>
// SELECT * FROM `user_scores` WHERE `uid` IN(1,2,3,4,5)
<原文结束>

# <翻译开始>
// 从`user_scores`表中选取所有列，其条件是`uid`字段的值存在于列表(1,2,3,4,5)中
// 即：查询`user_scores`表中uid为1、2、3、4或5的所有记录
# <翻译结束>







<原文开始>
// SELECT * FROM `user_detail` WHERE `uid` IN(1,2,3,4,5)
<原文结束>

# <翻译开始>
// 从`user_detail`表中选择所有列，其条件是`uid`在(1,2,3,4,5)这个列表内
# <翻译结束>


<原文开始>
// Initialize the data.
<原文结束>

# <翻译开始>
// 初始化数据
# <翻译结束>

