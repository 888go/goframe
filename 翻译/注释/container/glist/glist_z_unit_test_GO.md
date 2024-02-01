
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
// Clear all elements by iterating
<原文结束>

# <翻译开始>
// 通过遍历清空所有元素
# <翻译结束>







<原文开始>
		// e := l.Front()
		// l.Remove(e)
		// checkListPointers(t, l, []*Element{e2})
		// l.Remove(e)
		// checkListPointers(t, l, []*Element{e2})
<原文结束>

# <翻译开始>
// 获取链表l的头部元素，并赋值给e
// e := l.Front()
// 从链表l中移除元素e
// l.Remove(e)
// 调用checkListPointers函数，检查链表l的指针情况是否符合预期，此时链表中应包含元素e2
// checkListPointers(t, l, []*Element{e2})
// 再次尝试从链表l中移除元素e（由于之前已经移除过，这里实际操作无意义）
// l.Remove(e)
// 再次调用checkListPointers函数，检查链表l的指针情况是否符合预期，此时链表中仍应只包含元素e2
// checkListPointers(t, l, []*Element{e2})
# <翻译结束>


<原文开始>
// l2 should not change because e is not an element of l2
<原文结束>

# <翻译开始>
// l2不应发生改变，因为e不是l2中的元素
# <翻译结束>


<原文开始>
	// if e.Next() != nil {
	//    t.Errorf("e.Next() != nil")
	// }
	// if e.Prev() != nil {
	//    t.Errorf("e.Prev() != nil")
	// }
<原文结束>

# <翻译开始>
// 如果e.Next()不为nil（非空）{
//    t.Errorf("e.Next() != nil") // 输出错误信息："e.Next() 不应为nil"
// }
// 如果e.Prev()不为nil（非空）{
//    t.Errorf("e.Prev() != nil") // 输出错误信息："e.Prev() 不应为nil"
// }
// 这段Go代码的注释翻译成中文后，含义如下：
// 在测试中，如果`e.Next()`方法返回的结果不是nil，则通过`t.Errorf`输出错误信息，表示期望`e.Next()`的结果应为nil；同理，如果`e.Prev()`方法返回的结果不是nil，也会输出相应的错误信息，表示期望`e.Prev()`的结果应为nil。
# <翻译结束>


<原文开始>
// Test PushFront, PushBack, PushFrontList, PushBackList with uninitialized List
<原文结束>

# <翻译开始>
// 测试未初始化的List使用PushFront、PushBack、PushFrontList和PushBackList函数
# <翻译结束>


<原文开始>
// Test that a list l is not modified when calling InsertBefore with a mark that is not an element of l.
<原文结束>

# <翻译开始>
// 测试当使用一个非列表l元素的标记调用InsertBefore时，列表l不会被修改。
# <翻译结束>


<原文开始>
// Test that a list l is not modified when calling InsertAfter with a mark that is not an element of l.
<原文结束>

# <翻译开始>
// 测试当使用一个不在列表l中的标记调用InsertAfter时，列表l不会被修改。
# <翻译结束>


<原文开始>
// Test that a list l is not modified when calling MoveAfter or MoveBefore with a mark that is not an element of l.
<原文结束>

# <翻译开始>
// 测试当在列表l上调用MoveAfter或MoveBefore方法时，如果移动标记不是l的元素，则列表l不会被修改。
# <翻译结束>

















<原文开始>
	// for e := l.Front(); e != nil; e = e.Next() {
	//	le := e.Value.(int)
	//	if le != es[i] {
	//		t.Errorf("elt[%d].Value() = %v, want %v", i, le, es[i])
	//	}
	//	i++
	// }
<原文结束>

# <翻译开始>
// 遍历链表l，从头部开始直到链表结束（e不为空）
// for e := l.Front(); e != nil; e = e.Next() {
// 获取当前元素e的值，并将其转换为int类型，赋值给le
// le := e.Value.(int)
// 如果le与期望的es[i]不相等，则输出错误信息
// if le != es[i] {
//		t.Errorf("elt[%d].Value() = %v, want %v", i, le, es[i])
// }
// 将索引i加1，继续遍历下一个元素
// i++
// }
# <翻译结束>


<原文开始>
// Single element list
<原文结束>

# <翻译开始>
// 单元素列表
# <翻译结束>


<原文开始>
// insert before front
<原文结束>

# <翻译开始>
// 在前面插入
# <翻译结束>


<原文开始>
// insert before middle
<原文结束>

# <翻译开始>
// 在中间之前插入
# <翻译结束>


<原文开始>
// insert before back
<原文结束>

# <翻译开始>
// 在结尾处插入
# <翻译结束>


<原文开始>
// insert after front
<原文结束>

# <翻译开始>
// 在前面之后插入
# <翻译结束>


<原文开始>
// insert after middle
<原文结束>

# <翻译开始>
// 在中间位置后插入
# <翻译结束>


<原文开始>
// Check standard iteration.
<原文结束>

# <翻译开始>
// 检查标准迭代。
# <翻译结束>

