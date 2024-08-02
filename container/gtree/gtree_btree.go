// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package 树形类

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	gvar "github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/intlog"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	gconv "github.com/888go/goframe/util/gconv"
)

// BTree 存储 B 树的元素。 md5:191d2e09c9c918ab
type BTree struct {
	mu         rwmutex.RWMutex
	root       *BTreeNode
	comparator func(v1, v2 interface{}) int
	size       int //树中的总键数. md5:894ec399ab2f88ea
	m          int // 顺序（最大子节点数量）. md5:9c909788c23fe0a9
}

// BTreeNode 是树中的一个单个元素。 md5:f6e73ea6e5510845
type BTreeNode struct {
	Parent   *BTreeNode
	Entries  []*BTreeEntry // Contained keys in node
	Children []*BTreeNode  // Children nodes
}

// BTreeEntry表示节点中包含的键值对。 md5:215d17b4d8c2f829
type BTreeEntry struct {
	Key   interface{}
	Value interface{}
}

// NewBTree 创建一个具有 `m`（最大子节点数）和自定义键比较器的 B 树。参数 `safe` 用于指定是否在并发安全模式下使用树，其默认值为 false。
// 注意，`m` 必须大于或等于 3，否则将引发 panic。
// md5:63e15eb274ca4e1d
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

// NewBTreeFrom 根据给定的参数实例化一个 B-树，包括孩子节点的最大数量 `m`、自定义键比较器和数据映射。
// 参数 `safe` 用于指定是否需要并发安全，默认情况下为 false。
// md5:7a8fbca9b49feb70
func NewBTreeFrom(m int, comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *BTree {
	tree := NewBTree(m, comparator, safe...)
	for k, v := range data {
		tree.doSet(k, v)
	}
	return tree
}

// Clone 返回一个新的树，其中包含当前树的副本。 md5:256477216ae712b7
func (tree *BTree) Clone() *BTree {
	newTree := NewBTree(tree.m, tree.comparator, tree.mu.IsSafe())
	newTree.Sets(tree.Map())
	return newTree
}

// Set 将键值对插入到树中。 md5:af4d398e6bf21959
func (tree *BTree) Set(key interface{}, value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.doSet(key, value)
}

// doSet 将键值对节点插入到树中。
// 如果键已存在，则用新值更新其值。
// md5:dd34c6d624358b26
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

// 设置批处理将键值对添加到树中。 md5:70c6ec85c8b7476c
func (tree *BTree) Sets(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// Get 通过`key`在树中搜索节点，并返回其值，如果`key`在树中未找到，则返回nil。 md5:2e2483db20a69167
func (tree *BTree) Get(key interface{}) (value interface{}) {
	value, _ = tree.Search(key)
	return
}

// doSetWithLockCheck 使用互斥锁(mutex.Lock)检查键的值是否存在，
// 如果不存在，则将给定的`key`和`value`设置到映射中，
// 否则直接返回已存在的值。
//
// 在设置值时，如果`value`是<func() interface {}>类型，
// 它将在哈希映射的互斥锁(mutex.Lock)保护下执行，
// 并将其返回值设置到以`key`为键的映射中。
//
// 返回给定`key`对应的值。
// md5:1de9ffab89f3c38a
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

// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
func (tree *BTree) GetOrSet(key interface{}, value interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键获取值，
// 如果键不存在，则使用回调函数`f`的返回值设置值，
// 并返回这个设置的值。
// md5:f584dd7547dfbcc0
func (tree *BTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键获取值，
// 如果不存在，它将使用回调函数 `f` 的返回值设置该值，然后返回这个值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在执行函数 `f` 时会先锁定哈希映射的 mutex。
// md5:d32fdee586d84dde
func (tree *BTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 函数通过给定的 `key` 返回一个 gvar.Var，其值为对应的变量。
// 返回的 gvar.Var 不是并发安全的。
// md5:a04747902e4bf242
func (tree *BTree) GetVar(key interface{}) *gvar.Var {
	return gvar.New(tree.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取的结果的 gvar.Var。返回的 gvar.Var 不是线程安全的。
// md5:089beb08264e18cf
func (tree *BTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
// md5:8c97b145faade5ae
func (tree *BTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个gvar.Var，其结果来自GetOrSetFuncLock。
// 返回的gvar.Var是非并发安全的。
// md5:90c22300c2187ce4
func (tree *BTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
func (tree *BTree) SetIfNotExist(key interface{}, value interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
func (tree *BTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f())
		return true
	}
	return false
}

// SetIfNotExistFuncLock 使用回调函数 `f` 的返回值设置值，然后返回 true。
// 如果 `key` 已存在，则返回 false，`value` 将被忽略。
//
// SetIfNotExistFuncLock 与 SetIfNotExistFunc 函数的区别在于，
// 它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
// md5:a6ee84b157328f61
func (tree *BTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Contains 检查键 `key` 是否存在于树中。 md5:77fd85af8e586867
func (tree *BTree) Contains(key interface{}) bool {
	_, ok := tree.Search(key)
	return ok
}

// doRemove 通过键从树中删除节点。
// 键应符合比较器的类型断言，否则方法将 panic。
// md5:748baf2fba8b968d
func (tree *BTree) doRemove(key interface{}) (value interface{}) {
	node, index, found := tree.searchRecursively(tree.root, key)
	if found {
		value = node.Entries[index].Value
		tree.delete(node, index)
		tree.size--
	}
	return
}

// Remove 通过 `key` 从树中移除节点。 md5:42fcfa1d28b3945f
func (tree *BTree) Remove(key interface{}) (value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	return tree.doRemove(key)
}

// 通过`keys`移除树中的批量删除值。 md5:4620c81ac88b2936
func (tree *BTree) Removes(keys []interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for _, key := range keys {
		tree.doRemove(key)
	}
}

// IsEmpty 返回true当树中不包含任何节点时. md5:d43f280c082bb0fd
func (tree *BTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size 返回树中的节点数量。 md5:d437d5852f80de5c
func (tree *BTree) Size() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.size
}

// Keys 返回所有键，按升序排列。 md5:c2a692ea3491e160
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

// Values返回根据键值升序排列的所有值。 md5:6268d60d7aa20c91
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

// Map 返回所有键值对项作为一个映射。 md5:c12ca822a6c71dc1
func (tree *BTree) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[key] = value
		return true
	})
	return m
}

// MapStrAny 将所有键值对作为 map[string]interface{} 返回。 md5:412456aafc43f7a8
func (tree *BTree) MapStrAny() map[string]interface{} {
	m := make(map[string]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[gconv.String(key)] = value
		return true
	})
	return m
}

// Clear 从树中移除所有节点。 md5:a7db742922264980
func (tree *BTree) Clear() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
}

// 使用给定的`data`替换树中的数据。 md5:ff636c579597f294
func (tree *BTree) Replace(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// Height 返回树的高度。 md5:c3af563cbe50966a
func (tree *BTree) Height() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.root.height()
}

// Left 返回最左边（最小）的条目，如果树为空则返回 nil。 md5:57cf05edc8d10b88
func (tree *BTree) Left() *BTreeEntry {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.left(tree.root)
	if node != nil {
		return node.Entries[0]
	}
	return nil
}

// Right 返回最右边（最大）的条目，如果树为空则返回 nil。 md5:cd331b29b9cc98f8
func (tree *BTree) Right() *BTreeEntry {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.right(tree.root)
	if node != nil {
		return node.Entries[len(node.Entries)-1]
	}
	return nil
}

// String 返回一个表示容器的字符串（用于调试目的）. md5:2d28c3cbf692ce78
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

// Search 函数使用给定的 `key` 在树中进行查找。
// 第二个返回参数 `found` 为 true 表示找到了键，否则为 false。
// md5:d151c3783cadda2c
func (tree *BTree) Search(key interface{}) (value interface{}, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, index, found := tree.searchRecursively(tree.root, key)
	if found {
		return node.Entries[index].Value, true
	}
	return nil, false
}

// Search 使用给定的 `key` 在不加锁的情况下搜索树。如果找到相应的条目，则返回该条目，否则返回 nil。
// md5:2f4ee3482351a19d
func (tree *BTree) doSearch(key interface{}) *BTreeEntry {
	node, index, found := tree.searchRecursively(tree.root, key)
	if found {
		return node.Entries[index]
	}
	return nil
}

// Print 将树打印到标准输出。 md5:24fd6288549a501b
func (tree *BTree) Print() {
	fmt.Println(tree.String())
}

// Iterator 是 IteratorAsc 的别名。 md5:1bfdea306db62845
func (tree *BTree) Iterator(f func(key, value interface{}) bool) {
	tree.IteratorAsc(f)
}

// IteratorFrom是IteratorAscFrom的别名。 md5:6d3d506bcb5fe942
func (tree *BTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.IteratorAscFrom(key, match, f)
}

// IteratorAsc 使用给定的回调函数 `f` 以升序遍历树（只读）。如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c13b99ae40add3b0
func (tree *BTree) IteratorAsc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.left(tree.root)
	if node == nil {
		return
	}
	tree.doIteratorAsc(node, node.Entries[0], 0, f)
}

// IteratorAscFrom 从给定的回调函数 `f` 以升序遍历树。
// 参数 `key` 指定了遍历的起始条目。`match` 参数指定如果 `key` 完全匹配时是否开始遍历，否则使用索引搜索进行遍历。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c04855bbd3989808
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
		// 在当前节点中找到当前条目的位置. md5:0a7b8dbdf0511756
	if !first {
		index, _ = tree.search(node, entry.Key)
	} else {
		first = false
	}
		// 尝试进入当前条目右侧的子级. md5:76c7333c8aa6548c
	if index+1 < len(node.Children) {
		node = node.Children[index+1]
						// 尝试下降到当前节点的左子节点. md5:500eb3344ae2e9dc
		for len(node.Children) > 0 {
			node = node.Children[0]
		}
				// 返回最左边的条目. md5:e1e0f1ca5a78a394
		entry = node.Entries[0]
		goto loop
	}
		// 上面的代码确保我们已经到达了一个叶节点，所以返回当前节点（如果有）的下一个条目. md5:16db742c6c56694a
	if index+1 < len(node.Entries) {
		entry = node.Entries[index+1]
		goto loop
	}
		// 已到达叶子节点，并且当前条目右侧没有更多条目，因此返回到父节点. md5:2c6528856bee4df1
	for node.Parent != nil {
		node = node.Parent
						// 在当前节点中查找下一个条目位置（注意：搜索返回第一个等于或大于条目的位置）. md5:f184010d524512f9
		index, _ = tree.search(node, entry.Key)
				// 检查当前节点是否有下一个条目位置. md5:c0e01af4b4d09d1d
		if index < len(node.Entries) {
			entry = node.Entries[index]
			goto loop
		}
	}
}

// IteratorDesc 以降序方式遍历树，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:f6740ea55dafe4bb
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

// IteratorDescFrom 以降序方式遍历树，使用给定的回调函数 `f`。参数 `key` 指定开始遍历的条目。`match` 表示是否在 `key` 完全匹配时开始遍历，否则使用索引搜索遍历。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:e6bb2f7d12ab34f6
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

// IteratorDesc 以降序方式遍历树，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:f6740ea55dafe4bb
func (tree *BTree) doIteratorDesc(node *BTreeNode, entry *BTreeEntry, index int, f func(key, value interface{}) bool) {
	first := true
loop:
	if entry == nil {
		return
	}
	if !f(entry.Key, entry.Value) {
		return
	}
		// 在当前节点中找到当前条目的位置. md5:0a7b8dbdf0511756
	if !first {
		index, _ = tree.search(node, entry.Key)
	} else {
		first = false
	}
		// 尝试向下进入当前条目左侧的子项. md5:c5a2056515034dc8
	if index < len(node.Children) {
		node = node.Children[index]
				// 尝试前往当前节点右侧的子节点。 md5:dc7b90ce22f5e3db
		for len(node.Children) > 0 {
			node = node.Children[len(node.Children)-1]
		}
				// 返回最右边的条目. md5:d99f4b49cd9c2ea6
		entry = node.Entries[len(node.Entries)-1]
		goto loop
	}
		// 以上确保我们已经到达叶子节点，因此返回当前节点（如果有）的前一个条目. md5:8c1c1d33dbf6920a
	if index-1 >= 0 {
		entry = node.Entries[index-1]
		goto loop
	}

		// 到达叶子节点，且当前条目左侧没有更多项，因此向上移到父节点. md5:68f4a5ccc4125b55
	for node.Parent != nil {
		node = node.Parent
				// 在当前节点中查找前一个条目的位置（注意：搜索返回第一个等于或大于该条目的位置）. md5:cf405018e6f98ac2
		index, _ = tree.search(node, entry.Key)
						// 检查当前节点中是否存在前一个条目的位置. md5:38ebf51611534fa0
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

// 函数（tree *BTree）isFull（node *BTreeNode）bool：
// 返回当前节点（node）的条目数等于BTree类型的maxEntries方法的返回值，即满载状态。
// md5:3a4924d98a84d807

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
			// 当分割时，倾向于将更多的键分配给右侧节点，使用"-1". md5:589fa3b8c8c0ac7b
	return (tree.m - 1) / 2
}

// search 仅在单个节点的条目中进行搜索. md5:708796b6f6c04ad5
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

// searchRecursively 从startNode开始递归地在树中搜索. md5:6c5effca3e12cf15
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
			// 将条目的键插入节点的中间. md5:77832e5f62b079b2
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

			// 将节点要被分割的子节点移动到左右子节点中. md5:3e37e30e3dd2cb2c
	if !tree.isLeaf(node) {
		left.Children = append([]*BTreeNode(nil), node.Children[:middle+1]...)
		right.Children = append([]*BTreeNode(nil), node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := tree.search(parent, node.Entries[middle].Key)

		// 将中间键插入到父节点中. md5:90770d4875d60061
	parent.Entries = append(parent.Entries, nil)
	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
	parent.Entries[insertPosition] = node.Entries[middle]

			// 将插入键在父节点的左侧子节点设置为创建的左侧节点. md5:3ef858cf2ae0942a
	parent.Children[insertPosition] = left

			// 在父节点中将插入键的子节点设置为创建的右节点. md5:3ab277966ef065b7
	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	tree.split(parent)
}

func (tree *BTree) splitRoot() {
	middle := tree.middle()
	left := &BTreeNode{Entries: append([]*BTreeEntry(nil), tree.root.Entries[:middle]...)}
	right := &BTreeNode{Entries: append([]*BTreeEntry(nil), tree.root.Entries[middle+1:]...)}

			// 将节点要被分割的子节点移动到左右子节点中. md5:3e37e30e3dd2cb2c
	if !tree.isLeaf(tree.root) {
		left.Children = append([]*BTreeNode(nil), tree.root.Children[:middle+1]...)
		right.Children = append([]*BTreeNode(nil), tree.root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

		// Root 是一个具有一个入口和两个子节点（左和右）的节点. md5:5c7ab1f314ee5149
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

// leftSibling 函数返回节点的左兄弟节点以及该节点在父节点中的索引（如果存在的话），否则返回（nil, -1）。
// key 是节点中的任意一个键（即使是已被删除的键）。
// md5:5df6d39676db1b43
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

// rightSibling 返回节点的右兄弟节点及其在父节点中的子索引，如果存在的话，否则返回 (nil,-1)。
// key 可以是节点中的任意键（甚至可能是已删除的键）。
// md5:d987c8284e77dafa
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

// delete 删除node中entries索引处的条目
// 参考：https://en.wikipedia.org/wiki/B-tree#Deletion
// md5:b876a095ea679730
func (tree *BTree) delete(node *BTreeNode, index int) {
			// 从叶节点删除. md5:7876d56e8045e7f9
	if tree.isLeaf(node) {
		deletedKey := node.Entries[index].Key
		tree.deleteEntry(node, index)
		tree.reBalance(node, deletedKey)
		if len(tree.root.Entries) == 0 {
			tree.root = nil
		}
		return
	}

		// 从内部节点删除. md5:4bd2fbac4d732f59
	leftLargestNode := tree.right(node.Children[index]) // 左子树中最大的节点（假设存在）. md5:58dc1797a99c50fe
	leftLargestEntryIndex := len(leftLargestNode.Entries) - 1
	node.Entries[index] = leftLargestNode.Entries[leftLargestEntryIndex]
	deletedKey := leftLargestNode.Entries[leftLargestEntryIndex].Key
	tree.deleteEntry(leftLargestNode, leftLargestEntryIndex)
	tree.reBalance(leftLargestNode, deletedKey)
}

// reBalance 在必要时重新平衡树并返回true，否则返回false。
// 注意，我们首先删除条目，然后调用reBalance，因此将传递已删除的键作为引用。
// md5:c3feadb6a7f38094
func (tree *BTree) reBalance(node *BTreeNode, deletedKey interface{}) {
		// 检查是否需要重新平衡. md5:1a872a2636208ac3
	if node == nil || len(node.Entries) >= tree.minEntries() {
		return
	}

			// 尝试从左侧兄弟节点借用. md5:93535f4b1bfcf27f
	leftSibling, leftSiblingIndex := tree.leftSibling(node, deletedKey)
	if leftSibling != nil && len(leftSibling.Entries) > tree.minEntries() {
		// rotate right
		node.Entries = append([]*BTreeEntry{node.Parent.Entries[leftSiblingIndex]}, node.Entries...) // 将父节点的分隔符条目添加到节点的条目中. md5:aa1e7a85adad7bb6
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

			// 尝试从右侧兄弟节点借入. md5:e0e1cfcfc7caad95
	rightSibling, rightSiblingIndex := tree.rightSibling(node, deletedKey)
	if rightSibling != nil && len(rightSibling.Entries) > tree.minEntries() {
		// rotate left
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1]) // 将父节点的分隔符条目追加到节点的条目中. md5:00ee3de89c558897
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

	// merge with siblings
	if rightSibling != nil {
				// 与右侧兄弟节点合并. md5:cf809f1e8e2d90dc
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1])
		node.Entries = append(node.Entries, rightSibling.Entries...)
		deletedKey = node.Parent.Entries[rightSiblingIndex-1].Key
		tree.deleteEntry(node.Parent, rightSiblingIndex-1)
		tree.appendChildren(node.Parent.Children[rightSiblingIndex], node)
		tree.deleteChild(node.Parent, rightSiblingIndex)
	} else if leftSibling != nil {
						// 与左兄弟节点合并. md5:411c166f3c82c9dc
		entries := append([]*BTreeEntry(nil), leftSibling.Entries...)
		entries = append(entries, node.Parent.Entries[leftSiblingIndex])
		node.Entries = append(entries, node.Entries...)
		deletedKey = node.Parent.Entries[leftSiblingIndex].Key
		tree.deleteEntry(node.Parent, leftSiblingIndex)
		tree.prependChildren(node.Parent.Children[leftSiblingIndex], node)
		tree.deleteChild(node.Parent, leftSiblingIndex)
	}

			// 如果合并节点的父节点是根节点并且根节点为空，将合并节点设为新的根. md5:304d193d00ef0afc
	if node.Parent == tree.root && len(tree.root.Entries) == 0 {
		tree.root = node
		node.Parent = nil
		return
	}

		// 父元素可能会下溢，因此如果必要的话尝试重新平衡. md5:157406767c643099
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

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
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

// getComparator 如果之前已设置比较器，则返回该比较器，否则将引发恐慌。
// md5:03eac9fd6d838369
func (tree *BTree) getComparator() func(a, b interface{}) int {
	if tree.comparator == nil {
		panic("comparator is missing for tree")
	}
	return tree.comparator
}
