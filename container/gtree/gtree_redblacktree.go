// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package gtree
import (
	"bytes"
	"fmt"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/util/gconv"
	"github.com/888go/goframe/util/gutil"
	)
type color bool

const (
	black, red color = true, false
)

// RedBlackTree 用于存储红黑树的元素。
type RedBlackTree struct {
	mu         rwmutex.RWMutex
	root       *RedBlackTreeNode
	size       int
	comparator func(v1, v2 interface{}) int
}

// RedBlackTreeNode 是树中的单个元素。
type RedBlackTreeNode struct {
	Key    interface{}
	Value  interface{}
	color  color
	left   *RedBlackTreeNode
	right  *RedBlackTreeNode
	parent *RedBlackTreeNode
}

// NewRedBlackTree 创建一个带有自定义键比较器的红黑树。
// 参数`safe`用于指定是否在并发安全的情况下使用该树，默认为false。
func NewRedBlackTree(comparator func(v1, v2 interface{}) int, safe ...bool) *RedBlackTree {
	return &RedBlackTree{
		mu:         rwmutex.Create(safe...),
		comparator: comparator,
	}
}

// NewRedBlackTreeFrom 通过自定义键比较器和 `data` 映射创建一个红黑树实例。
// 参数 `safe` 用于指定是否在并发安全的情况下使用该树，默认为 false。
func NewRedBlackTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *RedBlackTree {
	tree := NewRedBlackTree(comparator, safe...)
	for k, v := range data {
		tree.doSet(k, v)
	}
	return tree
}

// SetComparator 设置/更改用于排序的比较器。
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
		// 如果比较器发生改变，则重新对树进行排序。
		tree.root = nil
		tree.size = 0
		for k, v := range data {
			tree.doSet(k, v)
		}
	}
}

// Clone 返回一个新的树，其中包含当前树的副本。
func (tree *RedBlackTree) Clone() *RedBlackTree {
	newTree := NewRedBlackTree(tree.comparator, tree.mu.IsSafe())
	newTree.Sets(tree.Map())
	return newTree
}

// Set 将键值对项插入到树中。
func (tree *RedBlackTree) Set(key interface{}, value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.doSet(key, value)
}

// Sets批量设置键值对到树中。
func (tree *RedBlackTree) Sets(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// doSet 在没有互斥锁的情况下将键值对项插入到树中。
func (tree *RedBlackTree) doSet(key interface{}, value interface{}) {
	insertedNode := (*RedBlackTreeNode)(nil)
	if tree.root == nil {
		// 确保键是初始树中比较器的类型
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

// Get通过`key`在树中搜索节点，并返回其对应的值，如果在树中未找到该键，则返回nil。
func (tree *RedBlackTree) Get(key interface{}) (value interface{}) {
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

// GetOrSet 函数通过 key 返回对应的 value，
// 若该 key 不存在，则使用给定的 `value` 设置并返回这个设置后的值。
func (tree *RedBlackTree) GetOrSet(key interface{}, value interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键返回值，如果该键不存在，
// 则使用回调函数 `f` 返回的值进行设置，并随后返回这个设置后的值。
func (tree *RedBlackTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键返回值，如果该键不存在，则使用回调函数 `f` 返回的值设置并返回这个新值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
func (tree *RedBlackTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 通过给定的 `key` 返回一个包含其值的 gvar.Var。
// 返回的 gvar.Var 对象不支持并发安全。
func (tree *RedBlackTree) GetVar(key interface{}) *gvar.Var {
	return gvar.New(tree.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取结果的 gvar.Var。
// 返回的 gvar.Var 不是线程安全的。
func (tree *RedBlackTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
func (tree *RedBlackTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个 gvar.Var，其结果来自 GetOrSetFuncLock。
// 返回的 gvar.Var 并未实现并发安全。
func (tree *RedBlackTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果`key`不存在，则将`value`设置到map中，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (tree *RedBlackTree) SetIfNotExist(key interface{}, value interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置键值，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (tree *RedBlackTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
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
func (tree *RedBlackTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Contains 检查 `key` 是否存在于树中。
func (tree *RedBlackTree) Contains(key interface{}) bool {
	_, ok := tree.Search(key)
	return ok
}

// doRemove 在没有互斥锁的情况下，通过 `key` 从树中移除节点。
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

// Remove 通过 `key` 从树中移除节点。
func (tree *RedBlackTree) Remove(key interface{}) (value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	return tree.doRemove(key)
}

// 删除树中通过`keys`指定的一批值。
func (tree *RedBlackTree) Removes(keys []interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for _, key := range keys {
		tree.doRemove(key)
	}
}

// IsEmpty 返回 true 如果树中不包含任何节点。
func (tree *RedBlackTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size 返回树中节点的数量。
func (tree *RedBlackTree) Size() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.size
}

// Keys 返回所有按升序排列的键。
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

// Values 返回所有基于键升序排列的值。
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

// Map 返回所有键值对项作为映射（map）。
func (tree *RedBlackTree) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[key] = value
		return true
	})
	return m
}

// MapStrAny 返回所有键值对项作为 map[string]interface{} 类型。
func (tree *RedBlackTree) MapStrAny() map[string]interface{} {
	m := make(map[string]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[gconv.String(key)] = value
		return true
	})
	return m
}

// Left 返回树中最左边（最小值）的节点，如果树为空则返回nil。
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

// Right 返回最右侧（最大）的节点，如果树为空则返回nil。
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

// leftNode 返回最左边（最小值）的节点，如果树为空则返回 nil。
func (tree *RedBlackTree) leftNode() *RedBlackTreeNode {
	p := (*RedBlackTreeNode)(nil)
	n := tree.root
	for n != nil {
		p = n
		n = n.left
	}
	return p
}

// rightNode 返回最右侧（最大值）的节点，如果树为空则返回 nil。
func (tree *RedBlackTree) rightNode() *RedBlackTreeNode {
	p := (*RedBlackTreeNode)(nil)
	n := tree.root
	for n != nil {
		p = n
		n = n.right
	}
	return p
}

// Floor 查找输入键的下界节点，返回下界节点或在未找到时返回 nil。
// 第二个返回参数为布尔值，若找到下界节点则为 true，否则为 false。
//
// 下界节点定义为键小于等于给定 `key` 的最大节点。
// 可能无法找到下界节点，原因可能是树为空，或者树中所有节点都大于给定节点。
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

// Ceiling 函数查找大于或等于输入键的最小节点（即上限节点），并返回该上限节点；若未找到上限节点，则返回 nil。
// 第二个返回参数为布尔值，若找到上限节点则为 true，否则为 false。
//
// 上限节点定义为：其键大于或等于给定 `key` 的最小节点。
// 可能找不到上限节点，原因可能是树为空，或者树中所有节点都小于给定节点。
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

// Iterator 是 IteratorAsc 的别名。
func (tree *RedBlackTree) Iterator(f func(key, value interface{}) bool) {
	tree.IteratorAsc(f)
}

// IteratorFrom 是 IteratorAscFrom 的别名。
func (tree *RedBlackTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.IteratorAscFrom(key, match, f)
}

// IteratorAsc 以升序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *RedBlackTree) IteratorAsc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorAsc(tree.leftNode(), f)
}

// IteratorAscFrom 以升序遍历（只读）给定回调函数 `f` 的树。
// 参数 `key` 指定了遍历的起始项。`match` 指定了当 `key` 完全匹配时是否开始遍历，
// 否则使用索引搜索方式进行遍历。
// 若 `f` 返回 true，则继续遍历；若返回 false，则停止遍历。
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

// IteratorDesc 以降序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *RedBlackTree) IteratorDesc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorDesc(tree.rightNode(), f)
}

// IteratorDescFrom 从指定的键(key)开始以降序方式遍历树（只读模式），并使用给定的回调函数`f`。
// 参数`key`指定了遍历的起始条目。`match`参数指定了如果`key`完全匹配时是否开始遍历，
// 否则将采用索引搜索方式进行遍历。
// 若`f`返回值为true，则继续遍历；若返回false，则停止遍历。
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

// Clear 清除树中的所有节点。
func (tree *RedBlackTree) Clear() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
}

// 用给定的`data`替换树的数据。
func (tree *RedBlackTree) Replace(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
	for k, v := range data {
		tree.doSet(k, v)
	}
}

// String 返回 container 的字符串表示形式。
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

// Print 将树打印到标准输出（stdout）。
func (tree *RedBlackTree) Print() {
	fmt.Println(tree.String())
}

// Search 使用给定的`key`搜索树。
// 第二个返回参数`found`如果找到key则为真（true），否则为假（false）。
func (tree *RedBlackTree) Search(key interface{}) (value interface{}, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node, found := tree.doSearch(key)
	if found {
		return node.Value, true
	}
	return nil, false
}

// Flip 将树中的键值对进行交换，即将键和值互换。
// 注意，你需要确保值的类型与键相同，
// 否则比较器将会触发 panic 异常。
//
// 如果值的类型与键不同，你需要传入新的 `comparator`（比较器）。
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

// doSearch 在没有互斥锁的情况下，使用给定的 `key` 搜索树。
// 如果找到，则返回节点，否则返回 nil。
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

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
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

// UnmarshalJSON 实现了 json.Unmarshal 接口的 UnmarshalJSON 方法。
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

// UnmarshalValue 是一个接口实现，用于为 map 设置任意类型的值。
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

// getComparator 返回之前设置的比较器，如果之前未设置，则会引发panic。
func (tree *RedBlackTree) getComparator() func(a, b interface{}) int {
	if tree.comparator == nil {
		panic("comparator is missing for tree")
	}
	return tree.comparator
}
