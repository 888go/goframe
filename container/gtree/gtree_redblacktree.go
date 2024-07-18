// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gtree

import (
	"bytes"
	"fmt"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/internal/json"
	"github.com/gogf/gf/v2/internal/rwmutex"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gutil"
)

type color bool

const (
	black, red color = true, false
)

// RedBlackTree 保存红黑树中的元素。 md5:5b20879b021304e2
type RedBlackTree struct {
	mu         rwmutex.RWMutex
	root       *RedBlackTreeNode
	size       int
	comparator func(v1, v2 interface{}) int
}

// RedBlackTreeNode 是树中的一个元素。 md5:782ce0fe3b02c5f4
type RedBlackTreeNode struct {
	Key    interface{}
	Value  interface{}
	color  color
	left   *RedBlackTreeNode
	right  *RedBlackTreeNode
	parent *RedBlackTreeNode
}

// NewRedBlackTree 使用自定义键比较器创建一个红黑树。
// 参数 `safe` 用于指定是否在并发安全模式下使用树，默认为 false。
// md5:ee97b0369d4155af
// ff:
// comparator:
// v1:
// v2:
// safe:
func NewRedBlackTree(comparator func(v1, v2 interface{}) int, safe ...bool) *RedBlackTree {
	return &RedBlackTree{
		mu:         rwmutex.Create(safe...),
		comparator: comparator,
	}
}

// NewRedBlackTreeFrom 根据自定义的键比较器和`data`映射实例化一个红黑树。
// 参数 `safe` 用于指定是否需要在并发安全的环境下使用该树，
// 默认情况下为false。
// md5:d8480b30bce45a04
// ff:
// comparator:
// v1:
// v2:
// data:
// safe:
func NewRedBlackTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *RedBlackTree {
	tree := NewRedBlackTree(comparator, safe...)
	for k, v := range data {
		tree.doSet(k, v)
	}
	return tree
}

// SetComparator 设置或更改排序的比较器。 md5:5108f29cfa1a4664
// ff:
// tree:
// comparator:
// a:
// b:
func (tree *RedBlackTree) SetComparator(comparator func(a, b interface{}) int) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.comparator = comparator
	if tree.size > 0 {
		data := make(map[interface{}]interface{}, tree.size)
		tree.doIteratorAsc(tree.leftNode(), func(key, value interface{}) bool {
			data[key] = value
			return true
		})
		// 如果比较器改变，对树进行排序。 md5:478d516e0a469cc0
		tree.root = nil
		tree.size = 0
		for k, v := range data {
			tree.doSet(k, v)
		}
	}
}

// Clone 返回一个新的树，其中包含当前树的副本。 md5:256477216ae712b7
// ff:
// tree:
func (tree *RedBlackTree) Clone() *RedBlackTree {
	newTree := NewRedBlackTree(tree.comparator, tree.mu.IsSafe())
	newTree.Sets(tree.Map())
	return newTree
}

// Set 将键值对插入到树中。 md5:af4d398e6bf21959
// yx:true
// ff:设置值
// tree:
// key:
// value:
func (tree *RedBlackTree) Set(key interface{}, value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.doSet(key, value)
}

// 设置批处理将键值对添加到树中。 md5:70c6ec85c8b7476c
// ff:
// tree:
// data:
func (tree *RedBlackTree) Sets(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// doSet 无需互斥地将键值对插入到树中。 md5:ecf43365b126f78e
func (tree *RedBlackTree) doSet(key interface{}, value interface{}) {
	insertedNode := (*RedBlackTreeNode)(nil)
	if tree.root == nil {
		// 断言键为比较器类型的初始树. md5:59bc0d818f986858
		tree.getComparator()(key, key)
		tree.root = &RedBlackTreeNode{Key: key, Value: value, color: red}
		insertedNode = tree.root
	} else {
		node := tree.root
		loop := true
		for loop {
			compare := tree.getComparator()(key, node.Key)
			switch {
			case compare == 0:
				// node.Key   = key
				node.Value = value
				return
			case compare < 0:
				if node.left == nil {
					node.left = &RedBlackTreeNode{Key: key, Value: value, color: red}
					insertedNode = node.left
					loop = false
				} else {
					node = node.left
				}
			case compare > 0:
				if node.right == nil {
					node.right = &RedBlackTreeNode{Key: key, Value: value, color: red}
					insertedNode = node.right
					loop = false
				} else {
					node = node.right
				}
			}
		}
		insertedNode.parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}

// Get 通过`key`在树中搜索节点，并返回其值，如果`key`在树中未找到，则返回nil。 md5:2e2483db20a69167
// ff:
// tree:
// key:
// value:
func (tree *RedBlackTree) Get(key interface{}) (value interface{}) {
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
func (tree *RedBlackTree) doSetWithLockCheck(key interface{}, value interface{}) interface{} {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	if node, found := tree.doSearch(key); found {
		return node.Value
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
// ff:
// tree:
// key:
// value:
func (tree *RedBlackTree) GetOrSet(key interface{}, value interface{}) interface{} {
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
// ff:
// tree:
// key:
// f:
func (tree *RedBlackTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
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
// ff:
// tree:
// key:
// f:
func (tree *RedBlackTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 函数通过给定的 `key` 返回一个 gvar.Var，其值为对应的变量。
// 返回的 gvar.Var 不是并发安全的。
// md5:a04747902e4bf242
// ff:
// tree:
// key:
func (tree *RedBlackTree) GetVar(key interface{}) *gvar.Var {
	return gvar.New(tree.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取的结果的 gvar.Var。返回的 gvar.Var 不是线程安全的。
// md5:089beb08264e18cf
// ff:
// tree:
// key:
// value:
func (tree *RedBlackTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
// md5:8c97b145faade5ae
// ff:
// tree:
// key:
// f:
func (tree *RedBlackTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个gvar.Var，其结果来自GetOrSetFuncLock。
// 返回的gvar.Var是非并发安全的。
// md5:90c22300c2187ce4
// ff:
// tree:
// key:
// f:
func (tree *RedBlackTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
// ff:
// tree:
// key:
// value:
func (tree *RedBlackTree) SetIfNotExist(key interface{}, value interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
// ff:
// tree:
// key:
// f:
func (tree *RedBlackTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
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
// ff:
// tree:
// key:
// f:
func (tree *RedBlackTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Contains 检查键 `key` 是否存在于树中。 md5:77fd85af8e586867
// ff:
// tree:
// key:
func (tree *RedBlackTree) Contains(key interface{}) bool {
	_, ok := tree.Search(key)
	return ok
}

// doRemove 函数在无锁状态下根据键(key)从树中移除节点。 md5:457e87555a234351
func (tree *RedBlackTree) doRemove(key interface{}) (value interface{}) {
	child := (*RedBlackTreeNode)(nil)
	node, found := tree.doSearch(key)
	if !found {
		return
	}
	value = node.Value
	if node.left != nil && node.right != nil {
		p := node.left.maximumNode()
		node.Key = p.Key
		node.Value = p.Value
		node = p
	}
	if node.left == nil || node.right == nil {
		if node.right == nil {
			child = node.left
		} else {
			child = node.right
		}
		if node.color == black {
			node.color = tree.nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size--
	return
}

// Remove 通过 `key` 从树中移除节点。 md5:42fcfa1d28b3945f
// ff:
// tree:
// key:
// value:
func (tree *RedBlackTree) Remove(key interface{}) (value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	return tree.doRemove(key)
}

// 通过`keys`移除树中的批量删除值。 md5:4620c81ac88b2936
// ff:
// tree:
// keys:
func (tree *RedBlackTree) Removes(keys []interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for _, key := range keys {
		tree.doRemove(key)
	}
}

// IsEmpty 如果树中不包含任何节点，则返回true。 md5:8f7ae813360d880b
// ff:
// tree:
func (tree *RedBlackTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size 返回树中的节点数量。 md5:d437d5852f80de5c
// ff:
// tree:
func (tree *RedBlackTree) Size() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.size
}

// Keys 返回所有键，按升序排列。 md5:c2a692ea3491e160
// ff:
// tree:
func (tree *RedBlackTree) Keys() []interface{} {
	var (
		keys  = make([]interface{}, tree.Size())
		index = 0
	)
	tree.IteratorAsc(func(key, value interface{}) bool {
		keys[index] = key
		index++
		return true
	})
	return keys
}

// Values返回根据键值升序排列的所有值。 md5:6268d60d7aa20c91
// ff:
// tree:
func (tree *RedBlackTree) Values() []interface{} {
	var (
		values = make([]interface{}, tree.Size())
		index  = 0
	)
	tree.IteratorAsc(func(key, value interface{}) bool {
		values[index] = value
		index++
		return true
	})
	return values
}

// Map 返回所有键值对项作为一个映射。 md5:c12ca822a6c71dc1
// ff:
// tree:
func (tree *RedBlackTree) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[key] = value
		return true
	})
	return m
}

// MapStrAny 将所有键值对作为 map[string]interface{} 返回。 md5:412456aafc43f7a8
// yx:true
// ff:取MapStrAny
// tree:
func (tree *RedBlackTree) MapStrAny() map[string]interface{} {
	m := make(map[string]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[gconv.String(key)] = value
		return true
	})
	return m
}

// Left 返回最左边（最小）的节点，如果树为空则返回nil。 md5:73ad6a74ff7ce0d2
// ff:
// tree:
func (tree *RedBlackTree) Left() *RedBlackTreeNode {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.leftNode()
	if tree.mu.IsSafe() {
		return &RedBlackTreeNode{
			Key:   node.Key,
			Value: node.Value,
		}
	}
	return node
}

// Right 返回最右边（最大）的节点，如果树为空则返回nil。 md5:92003588329d774d
// ff:
// tree:
func (tree *RedBlackTree) Right() *RedBlackTreeNode {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.rightNode()
	if tree.mu.IsSafe() {
		return &RedBlackTreeNode{
			Key:   node.Key,
			Value: node.Value,
		}
	}
	return node
}

// leftNode 返回最左侧（最小）的节点，如果树为空则返回 nil。 md5:7b369ec1ba2f4fd6
func (tree *RedBlackTree) leftNode() *RedBlackTreeNode {
	p := (*RedBlackTreeNode)(nil)
	n := tree.root
	for n != nil {
		p = n
		n = n.left
	}
	return p
}

// rightNode 返回右端（最大）节点，如果树为空则返回nil。 md5:38a67b5b179ccc5a
func (tree *RedBlackTree) rightNode() *RedBlackTreeNode {
	p := (*RedBlackTreeNode)(nil)
	n := tree.root
	for n != nil {
		p = n
		n = n.right
	}
	return p
}

// Floor 找到输入键的地板节点，如果没有找到地板节点，则返回nil。第二个返回参数表示是否找到了地板（true表示找到，false表示未找到）。
// 
// 地板节点定义为其键小于或等于给定`key`的最大节点。可能无法找到地板节点，因为树为空，或者树中的所有节点都大于给定节点。
// md5:e76dc868738ea5a0
// ff:
// tree:
// key:
// floor:
// found:
func (tree *RedBlackTree) Floor(key interface{}) (floor *RedBlackTreeNode, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	n := tree.root
	for n != nil {
		compare := tree.getComparator()(key, n.Key)
		switch {
		case compare == 0:
			return n, true
		case compare < 0:
			n = n.left
		case compare > 0:
			floor, found = n, true
			n = n.right
		}
	}
	if found {
		return
	}
	return nil, false
}

// Ceiling找到输入键的天花板节点，如果没有找到天花板节点则返回nil。第二个返回参数表示是否找到了天花板，否则为false。
// 
// 定义天花板节点为其键大于或等于给定`key`的最小节点。可能找不到天花板节点，原因可能是树为空，或者树中的所有节点都小于给定节点。
// md5:3d6d229626a6b5b2
// ff:
// tree:
// key:
// ceiling:
// found:
func (tree *RedBlackTree) Ceiling(key interface{}) (ceiling *RedBlackTreeNode, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	n := tree.root
	for n != nil {
		compare := tree.getComparator()(key, n.Key)
		switch {
		case compare == 0:
			return n, true
		case compare > 0:
			n = n.right
		case compare < 0:
			ceiling, found = n, true
			n = n.left
		}
	}
	if found {
		return
	}
	return nil, false
}

// Iterator 是 IteratorAsc 的别名。 md5:1bfdea306db62845
// yx:true
// ff:X遍历
// tree:
// f:
// key:
// value:
func (tree *RedBlackTree) Iterator(f func(key, value interface{}) bool) {
	tree.IteratorAsc(f)
}

// IteratorFrom是IteratorAscFrom的别名。 md5:6d3d506bcb5fe942
// ff:
// tree:
// key:
// match:
// f:
// key:
// value:
func (tree *RedBlackTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.IteratorAscFrom(key, match, f)
}

// IteratorAsc 使用给定的回调函数 `f` 以升序遍历树（只读）。如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c13b99ae40add3b0
// ff:
// tree:
// f:
// key:
// value:
func (tree *RedBlackTree) IteratorAsc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorAsc(tree.leftNode(), f)
}

// IteratorAscFrom 从给定的回调函数 `f` 以升序遍历树。
// 参数 `key` 指定了遍历的起始条目。`match` 参数指定如果 `key` 完全匹配时是否开始遍历，否则使用索引搜索进行遍历。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c04855bbd3989808
// ff:
// tree:
// key:
// match:
// f:
// key:
// value:
func (tree *RedBlackTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, found := tree.doSearch(key)
	if match {
		if found {
			tree.doIteratorAsc(node, f)
		}
	} else {
		tree.doIteratorAsc(node, f)
	}
}

func (tree *RedBlackTree) doIteratorAsc(node *RedBlackTreeNode, f func(key, value interface{}) bool) {
loop:
	if node == nil {
		return
	}
	if !f(node.Key, node.Value) {
		return
	}
	if node.right != nil {
		node = node.right
		for node.left != nil {
			node = node.left
		}
		goto loop
	}
	if node.parent != nil {
		old := node
		for node.parent != nil {
			node = node.parent
			if tree.getComparator()(old.Key, node.Key) <= 0 {
				goto loop
			}
		}
	}
}

// IteratorDesc 以降序方式遍历树，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:f6740ea55dafe4bb
// ff:
// tree:
// f:
// key:
// value:
func (tree *RedBlackTree) IteratorDesc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorDesc(tree.rightNode(), f)
}

// IteratorDescFrom 以降序方式遍历树，使用给定的回调函数 `f`。参数 `key` 指定开始遍历的条目。`match` 表示是否在 `key` 完全匹配时开始遍历，否则使用索引搜索遍历。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:e6bb2f7d12ab34f6
// ff:
// tree:
// key:
// match:
// f:
// key:
// value:
func (tree *RedBlackTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, found := tree.doSearch(key)
	if match {
		if found {
			tree.doIteratorDesc(node, f)
		}
	} else {
		tree.doIteratorDesc(node, f)
	}
}

func (tree *RedBlackTree) doIteratorDesc(node *RedBlackTreeNode, f func(key, value interface{}) bool) {
loop:
	if node == nil {
		return
	}
	if !f(node.Key, node.Value) {
		return
	}
	if node.left != nil {
		node = node.left
		for node.right != nil {
			node = node.right
		}
		goto loop
	}
	if node.parent != nil {
		old := node
		for node.parent != nil {
			node = node.parent
			if tree.getComparator()(old.Key, node.Key) >= 0 {
				goto loop
			}
		}
	}
}

// Clear 从树中移除所有节点。 md5:a7db742922264980
// ff:
// tree:
func (tree *RedBlackTree) Clear() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
}

// 使用给定的`data`替换树中的数据。 md5:ff636c579597f294
// ff:
// tree:
// data:
func (tree *RedBlackTree) Replace(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// String 返回容器的字符串表示形式。 md5:d27ed19a0d92f8aa
// ff:
// tree:
func (tree *RedBlackTree) String() string {
	if tree == nil {
		return ""
	}
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	str := ""
	if tree.size != 0 {
		tree.output(tree.root, "", true, &str)
	}
	return str
}

// Print 将树打印到标准输出。 md5:24fd6288549a501b
// ff:
// tree:
func (tree *RedBlackTree) Print() {
	fmt.Println(tree.String())
}

// Search 函数使用给定的 `key` 在树中进行查找。
// 第二个返回参数 `found` 为 true 表示找到了键，否则为 false。
// md5:d151c3783cadda2c
// ff:
// tree:
// key:
// value:
// found:
func (tree *RedBlackTree) Search(key interface{}) (value interface{}, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, found := tree.doSearch(key)
	if found {
		return node.Value, true
	}
	return nil, false
}

// Flip 将树中的键值对交换为值键。
// 请注意，你应该确保值的类型与键相同，否则比较器会panic。
//
// 如果值的类型与键不同，你需要传递新的 `comparator`。
// md5:e71ceac22aee55f1
// ff:
// tree:
// comparator:
// v1:
// v2:
func (tree *RedBlackTree) Flip(comparator ...func(v1, v2 interface{}) int) {
	t := (*RedBlackTree)(nil)
	if len(comparator) > 0 {
		t = NewRedBlackTree(comparator[0], tree.mu.IsSafe())
	} else {
		t = NewRedBlackTree(tree.comparator, tree.mu.IsSafe())
	}
	tree.IteratorAsc(func(key, value interface{}) bool {
		t.doSet(value, key)
		return true
	})
	tree.mu.Lock()
	tree.root = t.root
	tree.size = t.size
	tree.mu.Unlock()
}

func (tree *RedBlackTree) output(node *RedBlackTreeNode, prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		tree.output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += fmt.Sprintf("%v\n", node.Key)
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		tree.output(node.left, newPrefix, true, str)
	}
}

// doSearch 使用给定的`key`在树中进行搜索，但不使用互斥锁。
// 如果找到节点，则返回该节点；否则返回nil。
// md5:b8bbc8f49418f189
func (tree *RedBlackTree) doSearch(key interface{}) (node *RedBlackTreeNode, found bool) {
	node = tree.root
	for node != nil {
		compare := tree.getComparator()(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.left
		case compare > 0:
			node = node.right
		}
	}
	return node, false
}

func (node *RedBlackTreeNode) grandparent() *RedBlackTreeNode {
	if node != nil && node.parent != nil {
		return node.parent.parent
	}
	return nil
}

func (node *RedBlackTreeNode) uncle() *RedBlackTreeNode {
	if node == nil || node.parent == nil || node.parent.parent == nil {
		return nil
	}
	return node.parent.sibling()
}

func (node *RedBlackTreeNode) sibling() *RedBlackTreeNode {
	if node == nil || node.parent == nil {
		return nil
	}
	if node == node.parent.left {
		return node.parent.right
	}
	return node.parent.left
}

func (tree *RedBlackTree) rotateLeft(node *RedBlackTreeNode) {
	right := node.right
	tree.replaceNode(node, right)
	node.right = right.left
	if right.left != nil {
		right.left.parent = node
	}
	right.left = node
	node.parent = right
}

func (tree *RedBlackTree) rotateRight(node *RedBlackTreeNode) {
	left := node.left
	tree.replaceNode(node, left)
	node.left = left.right
	if left.right != nil {
		left.right.parent = node
	}
	left.right = node
	node.parent = left
}

func (tree *RedBlackTree) replaceNode(old *RedBlackTreeNode, new *RedBlackTreeNode) {
	if old.parent == nil {
		tree.root = new
	} else {
		if old == old.parent.left {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}
	if new != nil {
		new.parent = old.parent
	}
}

func (tree *RedBlackTree) insertCase1(node *RedBlackTreeNode) {
	if node.parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RedBlackTree) insertCase2(node *RedBlackTreeNode) {
	if tree.nodeColor(node.parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *RedBlackTree) insertCase3(node *RedBlackTreeNode) {
	uncle := node.uncle()
	if tree.nodeColor(uncle) == red {
		node.parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RedBlackTree) insertCase4(node *RedBlackTreeNode) {
	grandparent := node.grandparent()
	if node == node.parent.right && node.parent == grandparent.left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == grandparent.right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}

func (tree *RedBlackTree) insertCase5(node *RedBlackTreeNode) {
	node.parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.parent.left && node.parent == grandparent.left {
		tree.rotateRight(grandparent)
	} else if node == node.parent.right && node.parent == grandparent.right {
		tree.rotateLeft(grandparent)
	}
}

func (node *RedBlackTreeNode) maximumNode() *RedBlackTreeNode {
	if node == nil {
		return nil
	}
	for node.right != nil {
		return node.right
	}
	return node
}

func (tree *RedBlackTree) deleteCase1(node *RedBlackTreeNode) {
	if node.parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *RedBlackTree) deleteCase2(node *RedBlackTreeNode) {
	sibling := node.sibling()
	if tree.nodeColor(sibling) == red {
		node.parent.color = red
		sibling.color = black
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *RedBlackTree) deleteCase3(node *RedBlackTreeNode) {
	sibling := node.sibling()
	if tree.nodeColor(node.parent) == black &&
		tree.nodeColor(sibling) == black &&
		tree.nodeColor(sibling.left) == black &&
		tree.nodeColor(sibling.right) == black {
		sibling.color = red
		tree.deleteCase1(node.parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RedBlackTree) deleteCase4(node *RedBlackTreeNode) {
	sibling := node.sibling()
	if tree.nodeColor(node.parent) == red &&
		tree.nodeColor(sibling) == black &&
		tree.nodeColor(sibling.left) == black &&
		tree.nodeColor(sibling.right) == black {
		sibling.color = red
		node.parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RedBlackTree) deleteCase5(node *RedBlackTreeNode) {
	sibling := node.sibling()
	if node == node.parent.left &&
		tree.nodeColor(sibling) == black &&
		tree.nodeColor(sibling.left) == red &&
		tree.nodeColor(sibling.right) == black {
		sibling.color = red
		sibling.left.color = black
		tree.rotateRight(sibling)
	} else if node == node.parent.right &&
		tree.nodeColor(sibling) == black &&
		tree.nodeColor(sibling.right) == red &&
		tree.nodeColor(sibling.left) == black {
		sibling.color = red
		sibling.right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *RedBlackTree) deleteCase6(node *RedBlackTreeNode) {
	sibling := node.sibling()
	sibling.color = tree.nodeColor(node.parent)
	node.parent.color = black
	if node == node.parent.left && tree.nodeColor(sibling.right) == red {
		sibling.right.color = black
		tree.rotateLeft(node.parent)
	} else if tree.nodeColor(sibling.left) == red {
		sibling.left.color = black
		tree.rotateRight(node.parent)
	}
}

func (tree *RedBlackTree) nodeColor(node *RedBlackTreeNode) color {
	if node == nil {
		return black
	}
	return node.color
}

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。 md5:43c3b36e60a18f9a
// ff:
// tree:
// jsonBytes:
// err:
func (tree RedBlackTree) MarshalJSON() (jsonBytes []byte, err error) {
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

// UnmarshalJSON实现了json.Unmarshal接口的UnmarshalJSON方法。 md5:f6766b88cf3d63c2
// ff:
// tree:
// b:
func (tree *RedBlackTree) UnmarshalJSON(b []byte) error {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	if tree.comparator == nil {
		tree.comparator = gutil.ComparatorString
	}
	var data map[string]interface{}
	if err := json.UnmarshalUseNumber(b, &data); err != nil {
		return err
	}
	for k, v := range data {
		tree.doSet(k, v)
	}
	return nil
}

// UnmarshalValue 是一个接口实现，用于将任何类型的值设置到映射中。 md5:6f3087a6f7df5477
// ff:
// tree:
// value:
// err:
func (tree *RedBlackTree) UnmarshalValue(value interface{}) (err error) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	if tree.comparator == nil {
		tree.comparator = gutil.ComparatorString
	}
	for k, v := range gconv.Map(value) {
		tree.doSet(k, v)
	}
	return
}

// getComparator 如果之前已设置比较器，则返回该比较器，否则将引发恐慌。
// md5:03eac9fd6d838369
func (tree *RedBlackTree) getComparator() func(a, b interface{}) int {
	if tree.comparator == nil {
		panic("comparator is missing for tree")
	}
	return tree.comparator
}
