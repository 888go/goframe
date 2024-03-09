// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 定时类

// Len 实现了 sort.Interface 接口中的方法。
// （译注：在 Go 语言中，sort.Interface 是一个接口类型，其中包含三个需要实现的方法：Len()、Swap(i, j int) 和 Less(i, j int) bool，用于对任意数据结构进行排序。此处的 Len 方法就是返回待排序序列的长度。）
func (h *priorityQueueHeap) Len() int {
	return len(h.array)
}

// Less 用于实现 sort.Interface 接口。
// 其中较小的元素会被放置在堆顶。
func (h *priorityQueueHeap) Less(i, j int) bool {
	return h.array[i].priority < h.array[j].priority
}

// Swap用于实现sort.Interface接口。
func (h *priorityQueueHeap) Swap(i, j int) {
	if len(h.array) == 0 {
		return
	}
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// Push 向堆中添加一个元素。
func (h *priorityQueueHeap) Push(x interface{}) {
	h.array = append(h.array, x.(priorityQueueItem))
}

// Pop从堆中获取、移除并返回优先级最高的元素。
func (h *priorityQueueHeap) Pop() interface{} {
	length := len(h.array)
	if length == 0 {
		return nil
	}
	item := h.array[length-1]
	h.array = h.array[0 : length-1]
	return item
}
