// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtree

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/util/gconv"
)

// BTree 保存了 B-树 的元素。
type BTree struct {
	mu         rwmutex.RWMutex
	root       *BTreeNode
	comparator func(v1, v2 interface{}) int
	size       int // 树中键的总数
	m          int // order （最大子节点数）
}

// BTreeNode 是树中的单个元素。
type BTreeNode struct {
	Parent   *BTreeNode
	Entries  []*BTreeEntry // 节点中包含的键
	Children []*BTreeNode  // Children nodes
}

// BTreeEntry 代表节点中包含的键值对。
type BTreeEntry struct {
	Key   interface{}
	Value interface{}
}

// NewBTree 创建一个具有`m`（最大子节点数量）的B树，并使用自定义键比较器。
// 参数`safe`用于指定是否在并发安全环境下使用该树，默认为false。
// 注意，`m`必须大于或等于3，否则会引发panic。
func NewBTree(m int, comparator func(v1, v2 interface{}) int, safe ...bool) *BTree {
	if m < 3 {
		panic("Invalid order, should be at least 3")
	}
	return &BTree{
		comparator: comparator,
		mu:         rwmutex.Create(safe...),
		m:          m,
	}
}

// NewBTreeFrom 通过给定的最大子节点数 `m`，自定义键比较器和数据映射来实例化一个 B-树。
// 参数 `safe` 用于指定是否在并发安全的情况下使用该树，默认为 false。
func NewBTreeFrom(m int, comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *BTree {
	tree := NewBTree(m, comparator, safe...)
	for k, v := range data {
		tree.doSet(k, v)
	}
	return tree
}

// Clone 返回一个新的树，其中包含当前树的副本。
func (tree *BTree) Clone() *BTree {
	newTree := NewBTree(tree.m, tree.comparator, tree.mu.IsSafe())
	newTree.Sets(tree.Map())
	return newTree
}

// Set 将键值对项插入到树中。
func (tree *BTree) Set(key interface{}, value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.doSet(key, value)
}

// doSet 将键值对节点插入到树中。
// 如果键已存在，则用新值更新其原有值。
func (tree *BTree) doSet(key interface{}, value interface{}) {
	entry := &BTreeEntry{Key: key, Value: value}
	if tree.root == nil {
		tree.root = &BTreeNode{Entries: []*BTreeEntry{entry}, Children: []*BTreeNode{}}
		tree.size++
		return
	}

	if tree.insert(tree.root, entry) {
		tree.size++
	}
}

// Sets批量设置键值对到树中。
func (tree *BTree) Sets(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// Get通过`key`在树中搜索节点，并返回其对应的值，如果在树中未找到该键，则返回nil。
func (tree *BTree) Get(key interface{}) (value interface{}) {
	value, _ = tree.Search(key)
	return
}

// doSetWithLockCheck 检查在对 mutex 锁定后，给定 key 对应的值是否存在，
// 如果不存在，则使用给定的 `key` 将 value 设置到映射中；
// 否则，直接返回已存在的 value。
//
// 在设置值的过程中，如果 `value` 的类型为 <func() interface {}>，
// 会在哈希映射的 mutex 锁定下执行该函数，
// 并将函数的返回值以 `key` 为键设置到映射中。
//
// 最终返回给定 `key` 对应的值。
func (tree *BTree) doSetWithLockCheck(key interface{}, value interface{}) interface{} {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	if entry := tree.doSearch(key); entry != nil {
		return entry.Value
	}
	if f, ok := value.(func() interface{}); ok {
		value = f()
	}
	if value != nil {
		tree.doSet(key, value)
	}
	return value
}

// GetOrSet 函数通过 key 返回对应的 value，
// 若该 key 不存在，则使用给定的 `value` 设置并返回这个设置后的值。
func (tree *BTree) GetOrSet(key interface{}, value interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键返回值，如果该键不存在，
// 则使用回调函数 `f` 返回的值进行设置，并随后返回这个设置后的值。
func (tree *BTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键返回值，如果该键不存在，则使用回调函数 `f` 返回的值设置并返回这个新值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
func (tree *BTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 通过给定的 `key` 返回一个包含其值的 gvar.Var。
// 返回的 gvar.Var 对象不支持并发安全。
func (tree *BTree) GetVar(key interface{}) *gvar.Var {
	return gvar.New(tree.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取结果的 gvar.Var。
// 返回的 gvar.Var 不是线程安全的。
func (tree *BTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
func (tree *BTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个 gvar.Var，其结果来自 GetOrSetFuncLock。
// 返回的 gvar.Var 并未实现并发安全。
func (tree *BTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果`key`不存在，则将`value`设置到map中，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (tree *BTree) SetIfNotExist(key interface{}, value interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置键值，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (tree *BTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock 函数用于设置键值对，其值为回调函数 `f` 的返回值，并在设置成功时返回 true。
// 若 `key` 已存在，则返回 false，并且将忽略 `value` 参数。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在执行回调函数 `f` 时会锁定哈希表的 mutex 锁。
func (tree *BTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Contains 检查 `key` 是否存在于树中。
func (tree *BTree) Contains(key interface{}) bool {
	_, ok := tree.Search(key)
	return ok
}

// doRemove 通过键从树中移除节点。
// 键应遵循比较器的类型断言，否则方法将引发恐慌。
func (tree *BTree) doRemove(key interface{}) (value interface{}) {
	node, index, found := tree.searchRecursively(tree.root, key)
	if found {
		value = node.Entries[index].Value
		tree.delete(node, index)
		tree.size--
	}
	return
}

// Remove 通过 `key` 从树中移除节点。
func (tree *BTree) Remove(key interface{}) (value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	return tree.doRemove(key)
}

// 删除树中通过`keys`指定的一批值。
func (tree *BTree) Removes(keys []interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for _, key := range keys {
		tree.doRemove(key)
	}
}

// IsEmpty 返回 true 如果树中不包含任何节点
func (tree *BTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size 返回树中节点的数量。
func (tree *BTree) Size() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.size
}

// Keys 返回所有按升序排列的键。
func (tree *BTree) Keys() []interface{} {
	keys := make([]interface{}, tree.Size())
	index := 0
	tree.IteratorAsc(func(key, value interface{}) bool {
		keys[index] = key
		index++
		return true
	})
	return keys
}

// Values 返回所有基于键升序排列的值。
func (tree *BTree) Values() []interface{} {
	values := make([]interface{}, tree.Size())
	index := 0
	tree.IteratorAsc(func(key, value interface{}) bool {
		values[index] = value
		index++
		return true
	})
	return values
}

// Map 返回所有键值对项作为映射（map）。
func (tree *BTree) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[key] = value
		return true
	})
	return m
}

// MapStrAny 返回所有键值对项作为 map[string]interface{} 类型。
func (tree *BTree) MapStrAny() map[string]interface{} {
	m := make(map[string]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[gconv.String(key)] = value
		return true
	})
	return m
}

// Clear 清除树中的所有节点。
func (tree *BTree) Clear() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
}

// 用给定的`data`替换树的数据。
func (tree *BTree) Replace(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// Height 返回树的高度。
func (tree *BTree) Height() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.root.height()
}

// Left 返回最左边（最小）的元素，如果树为空则返回 nil。
func (tree *BTree) Left() *BTreeEntry {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.left(tree.root)
	if node != nil {
		return node.Entries[0]
	}
	return nil
}

// Right 返回最右侧（最大）的元素，如果树为空，则返回nil。
func (tree *BTree) Right() *BTreeEntry {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.right(tree.root)
	if node != nil {
		return node.Entries[len(node.Entries)-1]
	}
	return nil
}

// String 返回容器的字符串表示形式（用于调试目的）
func (tree *BTree) String() string {
	if tree == nil {
		return ""
	}
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	var buffer bytes.Buffer
	if tree.size != 0 {
		tree.output(&buffer, tree.root, 0, true)
	}
	return buffer.String()
}

// Search 使用给定的`key`搜索树。
// 第二个返回参数`found`如果找到key则为真（true），否则为假（false）。
func (tree *BTree) Search(key interface{}) (value interface{}, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, index, found := tree.searchRecursively(tree.root, key)
	if found {
		return node.Entries[index].Value, true
	}
	return nil, false
}

// Search 在没有使用互斥锁的情况下，通过给定的`key`搜索树。
// 如果找到对应的项则返回该条目，否则返回nil。
func (tree *BTree) doSearch(key interface{}) *BTreeEntry {
	node, index, found := tree.searchRecursively(tree.root, key)
	if found {
		return node.Entries[index]
	}
	return nil
}

// Print 将树打印到标准输出（stdout）。
func (tree *BTree) Print() {
	fmt.Println(tree.String())
}

// Iterator 是 IteratorAsc 的别名。
func (tree *BTree) Iterator(f func(key, value interface{}) bool) {
	tree.IteratorAsc(f)
}

// IteratorFrom 是 IteratorAscFrom 的别名。
func (tree *BTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.IteratorAscFrom(key, match, f)
}

// IteratorAsc 以升序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *BTree) IteratorAsc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.left(tree.root)
	if node == nil {
		return
	}
	tree.doIteratorAsc(node, node.Entries[0], 0, f)
}

// IteratorAscFrom 以升序遍历（只读）给定回调函数 `f` 的树。
// 参数 `key` 指定了遍历的起始项。`match` 指定了当 `key` 完全匹配时是否开始遍历，
// 否则使用索引搜索方式进行遍历。
// 若 `f` 返回 true，则继续遍历；若返回 false，则停止遍历。
func (tree *BTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, index, found := tree.searchRecursively(tree.root, key)
	if match {
		if found {
			tree.doIteratorAsc(node, node.Entries[index], index, f)
		}
	} else {
		if index >= 0 && index < len(node.Entries) {
			tree.doIteratorAsc(node, node.Entries[index], index, f)
		}
	}
}

func (tree *BTree) doIteratorAsc(node *BTreeNode, entry *BTreeEntry, index int, f func(key, value interface{}) bool) {
	first := true
loop:
	if entry == nil {
		return
	}
	if !f(entry.Key, entry.Value) {
		return
	}
	// 在当前节点中查找当前条目的位置
	if !first {
		index, _ = tree.search(node, entry.Key)
	} else {
		first = false
	}
	// 尝试转到当前条目右侧的子节点
	if index+1 < len(node.Children) {
		node = node.Children[index+1]
		// 尝试移动到当前节点的左孩子节点
		for len(node.Children) > 0 {
			node = node.Children[0]
		}
		// 返回最左边的条目
		entry = node.Entries[0]
		goto loop
	}
	// 上述代码确保我们已经到达一个叶节点，因此返回当前节点中的下一个条目（如果存在）
	if index+1 < len(node.Entries) {
		entry = node.Entries[index+1]
		goto loop
	}
	// 已到达叶节点，并且当前条目右侧没有其他条目，因此向上返回到父节点
	for node.Parent != nil {
		node = node.Parent
		// 在当前节点中查找下一个条目位置（注意：搜索返回第一个等于或大于给定条目的位置）
		index, _ = tree.search(node, entry.Key)
		// 检查当前节点中是否存在下一个条目位置
		if index < len(node.Entries) {
			entry = node.Entries[index]
			goto loop
		}
	}
}

// IteratorDesc 以降序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *BTree) IteratorDesc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.right(tree.root)
	if node == nil {
		return
	}
	index := len(node.Entries) - 1
	entry := node.Entries[index]
	tree.doIteratorDesc(node, entry, index, f)
}

// IteratorDescFrom 从指定的键(key)开始以降序方式遍历树（只读模式），并使用给定的回调函数`f`。
// 参数`key`指定了遍历的起始条目。`match`参数指定了如果`key`完全匹配时是否开始遍历，
// 否则将采用索引搜索方式进行遍历。
// 若`f`返回值为true，则继续遍历；若返回false，则停止遍历。
func (tree *BTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, index, found := tree.searchRecursively(tree.root, key)
	if match {
		if found {
			tree.doIteratorDesc(node, node.Entries[index], index, f)
		}
	} else {
		if index >= 0 && index < len(node.Entries) {
			tree.doIteratorDesc(node, node.Entries[index], index, f)
		}
	}
}

// IteratorDesc 以降序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *BTree) doIteratorDesc(node *BTreeNode, entry *BTreeEntry, index int, f func(key, value interface{}) bool) {
	first := true
loop:
	if entry == nil {
		return
	}
	if !f(entry.Key, entry.Value) {
		return
	}
	// 在当前节点中查找当前条目的位置
	if !first {
		index, _ = tree.search(node, entry.Key)
	} else {
		first = false
	}
	// 尝试转到当前条目左侧的子节点
	if index < len(node.Children) {
		node = node.Children[index]
		// 尝试移动到当前节点右侧的子节点
		for len(node.Children) > 0 {
			node = node.Children[len(node.Children)-1]
		}
		// 返回最右侧的条目
		entry = node.Entries[len(node.Entries)-1]
		goto loop
	}
	// 上述代码确保我们已经到达一个叶节点，因此返回当前节点中（如果有的话）的前一个条目
	if index-1 >= 0 {
		entry = node.Entries[index-1]
		goto loop
	}

	// 已到达叶节点，并且当前条目左侧没有条目，因此向上移动到父节点
	for node.Parent != nil {
		node = node.Parent
		// 在当前节点中查找前一个条目位置（注意：搜索返回第一个等于或大于给定条目的位置）
		index, _ = tree.search(node, entry.Key)
		// 检查当前节点中是否存在前一个条目位置
		if index-1 >= 0 {
			entry = node.Entries[index-1]
			goto loop
		}
	}
}

func (tree *BTree) output(buffer *bytes.Buffer, node *BTreeNode, level int, isTail bool) {
	for e := 0; e < len(node.Entries)+1; e++ {
		if e < len(node.Children) {
			tree.output(buffer, node.Children[e], level+1, true)
		}
		if e < len(node.Entries) {
			if _, err := buffer.WriteString(strings.Repeat("    ", level)); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
			if _, err := buffer.WriteString(fmt.Sprintf("%v", node.Entries[e].Key) + "\n"); err != nil {
				intlog.Errorf(context.TODO(), `%+v`, err)
			}
		}
	}
}

func (node *BTreeNode) height() int {
	h := 0
	n := node
	for ; n != nil; n = n.Children[0] {
		h++
		if len(n.Children) == 0 {
			break
		}
	}
	return h
}

func (tree *BTree) isLeaf(node *BTreeNode) bool {
	return len(node.Children) == 0
}

// 函数 (tree *BTree) isFull(node *BTreeNode) bool 的作用是：
// 判断给定的 BTreeNode 节点（node）是否已满。
// 当节点中的 Entries 数组长度等于 BTree 的最大条目数（通过 tree.maxEntries() 方法获取）时，返回 true，表示节点已满；否则返回 false。

func (tree *BTree) shouldSplit(node *BTreeNode) bool {
	return len(node.Entries) > tree.maxEntries()
}

func (tree *BTree) maxChildren() int {
	return tree.m
}

func (tree *BTree) minChildren() int {
	return (tree.m + 1) / 2 // ceil(m/2)
}

func (tree *BTree) maxEntries() int {
	return tree.maxChildren() - 1
}

func (tree *BTree) minEntries() int {
	return tree.minChildren() - 1
}

func (tree *BTree) middle() int {
	// 当分裂节点时，使用“-1”倾向于使右侧节点拥有更多的键
	return (tree.m - 1) / 2
}

// search 在单个节点的条目中执行搜索
func (tree *BTree) search(node *BTreeNode, key interface{}) (index int, found bool) {
	low, mid, high := 0, 0, len(node.Entries)-1
	for low <= high {
		mid = low + (high-low)/2
		compare := tree.getComparator()(key, node.Entries[mid].Key)
		switch {
		case compare > 0:
			low = mid + 1
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false
}

// searchRecursively 从起始节点startNode开始，递归地向下搜索整个树结构
func (tree *BTree) searchRecursively(startNode *BTreeNode, key interface{}) (node *BTreeNode, index int, found bool) {
	if tree.size == 0 {
		return nil, -1, false
	}
	node = startNode
	for {
		index, found = tree.search(node, key)
		if found {
			return node, index, true
		}
		if tree.isLeaf(node) {
			return node, index, false
		}
		node = node.Children[index]
	}
}

func (tree *BTree) insert(node *BTreeNode, entry *BTreeEntry) (inserted bool) {
	if tree.isLeaf(node) {
		return tree.insertIntoLeaf(node, entry)
	}
	return tree.insertIntoInternal(node, entry)
}

func (tree *BTree) insertIntoLeaf(node *BTreeNode, entry *BTreeEntry) (inserted bool) {
	insertPosition, found := tree.search(node, entry.Key)
	if found {
		node.Entries[insertPosition] = entry
		return false
	}
	// 在节点中间插入entry的键
	node.Entries = append(node.Entries, nil)
	copy(node.Entries[insertPosition+1:], node.Entries[insertPosition:])
	node.Entries[insertPosition] = entry
	tree.split(node)
	return true
}

func (tree *BTree) insertIntoInternal(node *BTreeNode, entry *BTreeEntry) (inserted bool) {
	insertPosition, found := tree.search(node, entry.Key)
	if found {
		node.Entries[insertPosition] = entry
		return false
	}
	return tree.insert(node.Children[insertPosition], entry)
}

func (tree *BTree) split(node *BTreeNode) {
	if !tree.shouldSplit(node) {
		return
	}

	if node == tree.root {
		tree.splitRoot()
		return
	}

	tree.splitNonRoot(node)
}

func (tree *BTree) splitNonRoot(node *BTreeNode) {
	middle := tree.middle()
	parent := node.Parent

	left := &BTreeNode{Entries: append([]*BTreeEntry(nil), node.Entries[:middle]...), Parent: parent}
	right := &BTreeNode{Entries: append([]*BTreeEntry(nil), node.Entries[middle+1:]...), Parent: parent}

	// 将待拆分节点的子节点移动到左右两个新节点中
	if !tree.isLeaf(node) {
		left.Children = append([]*BTreeNode(nil), node.Children[:middle+1]...)
		right.Children = append([]*BTreeNode(nil), node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := tree.search(parent, node.Entries[middle].Key)

	// 将中间键插入到父节点中
	parent.Entries = append(parent.Entries, nil)
	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
	parent.Entries[insertPosition] = node.Entries[middle]

	// 将父节点中插入键左侧的子节点设置为新创建的左节点
	parent.Children[insertPosition] = left

	// 在父节点中，将插入键的右子节点设置为新创建的右节点
	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	tree.split(parent)
}

func (tree *BTree) splitRoot() {
	middle := tree.middle()
	left := &BTreeNode{Entries: append([]*BTreeEntry(nil), tree.root.Entries[:middle]...)}
	right := &BTreeNode{Entries: append([]*BTreeEntry(nil), tree.root.Entries[middle+1:]...)}

	// 将待拆分节点的子节点移动到左右两个新节点中
	if !tree.isLeaf(tree.root) {
		left.Children = append([]*BTreeNode(nil), tree.root.Children[:middle+1]...)
		right.Children = append([]*BTreeNode(nil), tree.root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	// Root 是一个节点，包含一个键值对，并且拥有两个子节点（左、右子节点）
	newRoot := &BTreeNode{
		Entries:  []*BTreeEntry{tree.root.Entries[middle]},
		Children: []*BTreeNode{left, right},
	}

	left.Parent = newRoot
	right.Parent = newRoot
	tree.root = newRoot
}

func setParent(nodes []*BTreeNode, parent *BTreeNode) {
	for _, node := range nodes {
		node.Parent = parent
	}
}

func (tree *BTree) left(node *BTreeNode) *BTreeNode {
	if tree.size == 0 {
		return nil
	}
	current := node
	for {
		if tree.isLeaf(current) {
			return current
		}
		current = current.Children[0]
	}
}

func (tree *BTree) right(node *BTreeNode) *BTreeNode {
	if tree.size == 0 {
		return nil
	}
	current := node
	for {
		if tree.isLeaf(current) {
			return current
		}
		current = current.Children[len(current.Children)-1]
	}
}

// leftSibling 返回给定节点的左兄弟节点及其在父节点中的子索引（如果存在的话），否则返回 (nil, -1)
// key 是节点中任意一个键（甚至可能是已删除的键）。
// 这段 Go 语言代码的注释翻译成中文为：
// ```go
// leftSibling 函数会返回该节点的左侧兄弟节点及其在父节点中的子节点位置索引，如果存在这样的兄弟节点，则返回相应的信息；否则返回 (nil, -1)。
// 参数 key 是该节点中的任意一个键值（甚至可能是一个已被删除的键值）。
func (tree *BTree) leftSibling(node *BTreeNode, key interface{}) (*BTreeNode, int) {
	if node.Parent != nil {
		index, _ := tree.search(node.Parent, key)
		index--
		if index >= 0 && index < len(node.Parent.Children) {
			return node.Parent.Children[index], index
		}
	}
	return nil, -1
}

// rightSibling 返回给定节点的右兄弟节点及其在父节点中的子节点索引（如果存在的话），否则返回 (nil, -1)
// key 是节点中任意一个键（甚至可能是已被删除的键）。
func (tree *BTree) rightSibling(node *BTreeNode, key interface{}) (*BTreeNode, int) {
	if node.Parent != nil {
		index, _ := tree.search(node.Parent, key)
		index++
		if index < len(node.Parent.Children) {
			return node.Parent.Children[index], index
		}
	}
	return nil, -1
}

// delete 删除在节点中entries指定索引处的条目
// 参考文献: https://en.wikipedia.org/wiki/B-tree#删除
func (tree *BTree) delete(node *BTreeNode, index int) {
	// 从叶子节点进行删除
	if tree.isLeaf(node) {
		deletedKey := node.Entries[index].Key
		tree.deleteEntry(node, index)
		tree.reBalance(node, deletedKey)
		if len(tree.root.Entries) == 0 {
			tree.root = nil
		}
		return
	}

	// 从内部节点进行删除
	leftLargestNode := tree.right(node.Children[index]) // 左子树中最大的节点（假设存在）
	leftLargestEntryIndex := len(leftLargestNode.Entries) - 1
	node.Entries[index] = leftLargestNode.Entries[leftLargestEntryIndex]
	deletedKey := leftLargestNode.Entries[leftLargestEntryIndex].Key
	tree.deleteEntry(leftLargestNode, leftLargestEntryIndex)
	tree.reBalance(leftLargestNode, deletedKey)
}

// reBalance 在必要时在删除操作后重新平衡树并返回 true，否则返回 false。
// 注意，我们首先删除条目，然后调用 reBalance，因此将已删除键作为参考传入。
func (tree *BTree) reBalance(node *BTreeNode, deletedKey interface{}) {
	// 检查是否需要进行平衡调整
	if node == nil || len(node.Entries) >= tree.minEntries() {
		return
	}

	// 尝试从左兄弟节点借用
	leftSibling, leftSiblingIndex := tree.leftSibling(node, deletedKey)
	if leftSibling != nil && len(leftSibling.Entries) > tree.minEntries() {
		// rotate right
		node.Entries = append([]*BTreeEntry{node.Parent.Entries[leftSiblingIndex]}, node.Entries...) // 在节点的条目前添加父级的分隔符条目
		node.Parent.Entries[leftSiblingIndex] = leftSibling.Entries[len(leftSibling.Entries)-1]
		tree.deleteEntry(leftSibling, len(leftSibling.Entries)-1)
		if !tree.isLeaf(leftSibling) {
			leftSiblingRightMostChild := leftSibling.Children[len(leftSibling.Children)-1]
			leftSiblingRightMostChild.Parent = node
			node.Children = append([]*BTreeNode{leftSiblingRightMostChild}, node.Children...)
			tree.deleteChild(leftSibling, len(leftSibling.Children)-1)
		}
		return
	}

	// 尝试从右兄弟节点借用
	rightSibling, rightSiblingIndex := tree.rightSibling(node, deletedKey)
	if rightSibling != nil && len(rightSibling.Entries) > tree.minEntries() {
		// rotate left
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1]) // 将父节点的分隔符条目添加到当前节点的条目中
		node.Parent.Entries[rightSiblingIndex-1] = rightSibling.Entries[0]
		tree.deleteEntry(rightSibling, 0)
		if !tree.isLeaf(rightSibling) {
			rightSiblingLeftMostChild := rightSibling.Children[0]
			rightSiblingLeftMostChild.Parent = node
			node.Children = append(node.Children, rightSiblingLeftMostChild)
			tree.deleteChild(rightSibling, 0)
		}
		return
	}

	// 与兄弟节点合并
	if rightSibling != nil {
		// 与右侧兄弟节点合并
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1])
		node.Entries = append(node.Entries, rightSibling.Entries...)
		deletedKey = node.Parent.Entries[rightSiblingIndex-1].Key
		tree.deleteEntry(node.Parent, rightSiblingIndex-1)
		tree.appendChildren(node.Parent.Children[rightSiblingIndex], node)
		tree.deleteChild(node.Parent, rightSiblingIndex)
	} else if leftSibling != nil {
		// 与左兄弟节点合并
		entries := append([]*BTreeEntry(nil), leftSibling.Entries...)
		entries = append(entries, node.Parent.Entries[leftSiblingIndex])
		node.Entries = append(entries, node.Entries...)
		deletedKey = node.Parent.Entries[leftSiblingIndex].Key
		tree.deleteEntry(node.Parent, leftSiblingIndex)
		tree.prependChildren(node.Parent.Children[leftSiblingIndex], node)
		tree.deleteChild(node.Parent, leftSiblingIndex)
	}

	// 如果合并节点的父节点是根节点，并且根节点为空，则将合并节点设置为根节点
	if node.Parent == tree.root && len(tree.root.Entries) == 0 {
		tree.root = node
		node.Parent = nil
		return
	}

	// 父节点可能发生了下溢（即子节点数量不平衡），因此如果有必要，尝试进行重新平衡
	tree.reBalance(node.Parent, deletedKey)
}

func (tree *BTree) prependChildren(fromNode *BTreeNode, toNode *BTreeNode) {
	children := append([]*BTreeNode(nil), fromNode.Children...)
	toNode.Children = append(children, toNode.Children...)
	setParent(fromNode.Children, toNode)
}

func (tree *BTree) appendChildren(fromNode *BTreeNode, toNode *BTreeNode) {
	toNode.Children = append(toNode.Children, fromNode.Children...)
	setParent(fromNode.Children, toNode)
}

func (tree *BTree) deleteEntry(node *BTreeNode, index int) {
	copy(node.Entries[index:], node.Entries[index+1:])
	node.Entries[len(node.Entries)-1] = nil
	node.Entries = node.Entries[:len(node.Entries)-1]
}

func (tree *BTree) deleteChild(node *BTreeNode, index int) {
	if index >= len(node.Children) {
		return
	}
	copy(node.Children[index:], node.Children[index+1:])
	node.Children[len(node.Children)-1] = nil
	node.Children = node.Children[:len(node.Children)-1]
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (tree BTree) MarshalJSON() (jsonBytes []byte, err error) {
	if tree.root == nil {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('{')
	tree.Iterator(func(key, value interface{}) bool {
		valueBytes, valueJsonErr := json.Marshal(value)
		if valueJsonErr != nil {
			err = valueJsonErr
			return false
		}
		if buffer.Len() > 1 {
			buffer.WriteByte(',')
		}
		buffer.WriteString(fmt.Sprintf(`"%v":%s`, key, valueBytes))
		return true
	})
	buffer.WriteByte('}')
	return buffer.Bytes(), nil
}

// getComparator 返回之前设置的比较器，如果之前未设置，则会引发panic。
func (tree *BTree) getComparator() func(a, b interface{}) int {
	if tree.comparator == nil {
		panic("comparator is missing for tree")
	}
	return tree.comparator
}
