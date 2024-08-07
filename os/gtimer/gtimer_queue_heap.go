// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 定时类

// Len 用于实现 sort.Interface 接口。 md5:1e9f986dadbd5118
func (h *priorityQueueHeap) Len() int {
	return len(h.array)
}

// Less 用于实现 sort.Interface 接口。
// 较小的元素将被放在堆的顶部。
// md5:ad6834ec849095c3
func (h *priorityQueueHeap) Less(i, j int) bool {
	return h.array[i].priority < h.array[j].priority
}

// Swap 用于实现 sort.Interface 接口。 md5:ae0eab83cf38e0cb
func (h *priorityQueueHeap) Swap(i, j int) {
	if len(h.array) == 0 {
		return
	}
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// Push 将一个项目推送到堆栈。 md5:546121d20863ca8c
func (h *priorityQueueHeap) Push(x interface{}) {
	h.array = append(h.array, x.(priorityQueueItem))
}

// Pop 从堆中移除并返回最高优先级的项。 md5:5f06d92870330e44
func (h *priorityQueueHeap) Pop() interface{} {
	length := len(h.array)
	if length == 0 {
		return nil
	}
	item := h.array[length-1]
	h.array = h.array[0 : length-1]
	return item
}
