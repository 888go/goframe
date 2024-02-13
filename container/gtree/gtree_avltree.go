// 版权所有 GoFrame 作者（https://goframe.org）。保留所有权利。
//
// 本源代码形式遵循 MIT 许可协议条款。如果随此文件未分发 MIT 许可副本，
// 您可以在 https://github.com/gogf/gf 获取一份。

package 树形类

import (
	"bytes"
	"fmt"
	
	"github.com/888go/goframe/container/gvar"
	"github.com/888go/goframe/internal/json"
	"github.com/888go/goframe/internal/rwmutex"
	"github.com/888go/goframe/util/gconv"
)

// AVLTree 用于存储 AVL 树中的元素。
type AVLTree struct {
	mu         rwmutex.RWMutex
	root       *AVLTreeNode
	comparator func(v1, v2 interface{}) int
	size       int
}

// AVLTreeNode 是树中的一个单元素。
type AVLTreeNode struct {
	Key      interface{}
	Value    interface{}
	parent   *AVLTreeNode
	children [2]*AVLTreeNode
	b        int8
}

// NewAVLTree 通过自定义键比较器实例化一个 AVL 树。
// 参数 `safe` 用于指定是否在并发安全模式下使用该树，默认为 false。
func NewAVLTree(comparator func(v1, v2 interface{}) int, safe ...bool) *AVLTree {
	return &AVLTree{
		mu:         rwmutex.Create(safe...),
		comparator: comparator,
	}
}

// NewAVLTreeFrom 通过自定义键比较器和数据映射创建一个AVL树。
// 参数 `safe` 用于指定是否在并发安全环境下使用该树，默认为 false。
func NewAVLTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *AVLTree {
	tree := NewAVLTree(comparator, safe...)
	for k, v := range data {
		tree.put(k, v, nil, &tree.root)
	}
	return tree
}

// Clone 返回一个新的树，其中包含当前树的副本。
func (tree *AVLTree) Clone() *AVLTree {
	newTree := NewAVLTree(tree.comparator, tree.mu.IsSafe())
	newTree.Sets(tree.Map())
	return newTree
}

// Set 将节点插入到树中。
func (tree *AVLTree) X设置值(key interface{}, value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.put(key, value, nil, &tree.root)
}

// Sets批量设置键值对到树中。
func (tree *AVLTree) Sets(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for key, value := range data {
		tree.put(key, value, nil, &tree.root)
	}
}

// Search 使用给定的`key`搜索树。
// 第二个返回参数`found`如果找到key则为真（true），否则为假（false）。
func (tree *AVLTree) Search(key interface{}) (value interface{}, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	if node, found := tree.doSearch(key); found {
		return node.Value, true
	}
	return nil, false
}

// doSearch 用给定的 `key` 搜索树。
// 第二个返回参数 `found` 如果找到了 key，则为 true，否则为 false。
func (tree *AVLTree) doSearch(key interface{}) (node *AVLTreeNode, found bool) {
	node = tree.root
	for node != nil {
		cmp := tree.getComparator()(key, node.Key)
		switch {
		case cmp == 0:
			return node, true
		case cmp < 0:
			node = node.children[0]
		case cmp > 0:
			node = node.children[1]
		}
	}
	return nil, false
}

// Get通过`key`在树中搜索节点，并返回其对应的值，如果在树中未找到该键，则返回nil。
func (tree *AVLTree) Get(key interface{}) (value interface{}) {
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
func (tree *AVLTree) doSetWithLockCheck(key interface{}, value interface{}) interface{} {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	if node, found := tree.doSearch(key); found {
		return node.Value
	}
	if f, ok := value.(func() interface{}); ok {
		value = f()
	}
	if value != nil {
		tree.put(key, value, nil, &tree.root)
	}
	return value
}

// GetOrSet 函数通过 key 返回对应的 value，
// 若该 key 不存在，则使用给定的 `value` 设置并返回这个设置后的值。
func (tree *AVLTree) GetOrSet(key interface{}, value interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFunc 通过键返回值，如果该键不存在，
// 则使用回调函数 `f` 返回的值进行设置，并随后返回这个设置后的值。
func (tree *AVLTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f())
	} else {
		return v
	}
}

// GetOrSetFuncLock 通过键返回值，如果该键不存在，则使用回调函数 `f` 返回的值设置并返回这个新值。
//
// GetOrSetFuncLock 与 GetOrSetFunc 函数的不同之处在于，它在哈希映射的 mutex.Lock 保护下执行函数 `f`。
func (tree *AVLTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 通过给定的 `key` 返回一个包含其值的 gvar.Var。
// 返回的 gvar.Var 对象不支持并发安全。
func (tree *AVLTree) GetVar(key interface{}) *泛型类.Var {
	return 泛型类.X创建(tree.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取结果的 gvar.Var。
// 返回的 gvar.Var 不是线程安全的。
func (tree *AVLTree) GetVarOrSet(key interface{}, value interface{}) *泛型类.Var {
	return 泛型类.X创建(tree.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
func (tree *AVLTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *泛型类.Var {
	return 泛型类.X创建(tree.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个 gvar.Var，其结果来自 GetOrSetFuncLock。
// 返回的 gvar.Var 并未实现并发安全。
func (tree *AVLTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *泛型类.Var {
	return 泛型类.X创建(tree.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果`key`不存在，则将`value`设置到map中，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (tree *AVLTree) SetIfNotExist(key interface{}, value interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置键值，并返回true。
// 若`key`已存在，则返回false，同时`value`将被忽略。
func (tree *AVLTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
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
func (tree *AVLTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Contains 检查 `key` 是否存在于树中。
func (tree *AVLTree) Contains(key interface{}) bool {
	_, ok := tree.Search(key)
	return ok
}

// Remove 通过键从树中移除节点。
// 键应遵循比较器的类型断言，否则该方法会引发恐慌。
func (tree *AVLTree) Remove(key interface{}) (value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	value, _ = tree.remove(key, &tree.root)
	return
}

// 删除树中通过`keys`指定的一批值。
func (tree *AVLTree) Removes(keys []interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for _, key := range keys {
		tree.remove(key, &tree.root)
	}
}

// IsEmpty 返回 true 如果树中不包含任何节点。
func (tree *AVLTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size 返回树中节点的数量。
func (tree *AVLTree) Size() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.size
}

// Keys 返回所有按升序排列的键。
func (tree *AVLTree) Keys() []interface{} {
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
func (tree *AVLTree) Values() []interface{} {
	values := make([]interface{}, tree.Size())
	index := 0
	tree.IteratorAsc(func(key, value interface{}) bool {
		values[index] = value
		index++
		return true
	})
	return values
}

// Left 返回 AVL 树中的最小元素，
// 若树为空，则返回 nil。
func (tree *AVLTree) Left() *AVLTreeNode {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.bottom(0)
	if tree.mu.IsSafe() {
		return &AVLTreeNode{
			Key:   node.Key,
			Value: node.Value,
		}
	}
	return node
}

// Right 返回 AVL 树中最大的元素，
// 如果树为空，则返回 nil。
func (tree *AVLTree) Right() *AVLTreeNode {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	node := tree.bottom(1)
	if tree.mu.IsSafe() {
		return &AVLTreeNode{
			Key:   node.Key,
			Value: node.Value,
		}
	}
	return node
}

// Floor 查找输入键的下界节点，返回下界节点或若未找到则返回 nil。
// 第二个返回参数为布尔值，表示是否找到了下界节点（找到为 true，否则为 false）。
//
// 下界节点定义为小于等于给定节点的最大节点。
// 可能无法找到下界节点，原因可能是树为空，或者树中所有节点都大于给定节点。
//
// 键应遵循比较器的类型断言，否则该方法将引发 panic。
func (tree *AVLTree) Floor(key interface{}) (floor *AVLTreeNode, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	n := tree.root
	for n != nil {
		c := tree.getComparator()(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.children[0]
		case c > 0:
			floor, found = n, true
			n = n.children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Ceiling 查找大于或等于输入键的最小节点，返回该天花板节点（ceiling node）；若未找到满足条件的节点，则返回 nil。
// 第二个返回参数为布尔值，表示是否找到了天花板节点，找到则为 true，否则为 false。
//
// “天花板节点”定义为大于或等于给定节点的最小节点。
// 可能找不到天花板节点，原因可能包括：树为空，或者树中所有节点都小于给定节点。
//
// 键需符合比较器的类型断言，否则该方法将引发 panic。
func (tree *AVLTree) Ceiling(key interface{}) (ceiling *AVLTreeNode, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	n := tree.root
	for n != nil {
		c := tree.getComparator()(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c > 0:
			n = n.children[1]
		case c < 0:
			ceiling, found = n, true
			n = n.children[0]
		}
	}
	if found {
		return
	}
	return nil, false
}

// Clear 清除树中的所有节点。
func (tree *AVLTree) Clear() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
}

// 用给定的`data`替换树的数据。
func (tree *AVLTree) Replace(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
	for key, value := range data {
		tree.put(key, value, nil, &tree.root)
	}
}

// String 返回容器的字符串表示形式
func (tree *AVLTree) String() string {
	if tree == nil {
		return ""
	}
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	str := ""
	if tree.size != 0 {
		output(tree.root, "", true, &str)
	}
	return str
}

// Print 将树打印到标准输出（stdout）。
func (tree *AVLTree) Print() {
	fmt.Println(tree.String())
}

// Map 返回所有键值对项作为映射（map）。
func (tree *AVLTree) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[key] = value
		return true
	})
	return m
}

// MapStrAny 返回所有键值对项作为 map[string]interface{} 类型。
func (tree *AVLTree) X取MapStrAny() map[string]interface{} {
	m := make(map[string]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[转换类.String(key)] = value
		return true
	})
	return m
}

// Flip 将树中的键值对进行交换，即将键和值互换。
// 注意，你需要确保值的类型与键相同，
// 否则比较器将会触发 panic 异常。
//
// 如果值的类型与键不同，你需要传入新的 `comparator`（比较器）。
func (tree *AVLTree) Flip(comparator ...func(v1, v2 interface{}) int) {
	t := (*AVLTree)(nil)
	if len(comparator) > 0 {
		t = NewAVLTree(comparator[0], tree.mu.IsSafe())
	} else {
		t = NewAVLTree(tree.comparator, tree.mu.IsSafe())
	}
	tree.IteratorAsc(func(key, value interface{}) bool {
		t.put(value, key, nil, &t.root)
		return true
	})
	tree.mu.Lock()
	tree.root = t.root
	tree.size = t.size
	tree.mu.Unlock()
}

// Iterator 是 IteratorAsc 的别名。
func (tree *AVLTree) X遍历(f func(key, value interface{}) bool) {
	tree.IteratorAsc(f)
}

// IteratorFrom 是 IteratorAscFrom 的别名。
func (tree *AVLTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.IteratorAscFrom(key, match, f)
}

// IteratorAsc 以升序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *AVLTree) IteratorAsc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorAsc(tree.bottom(0), f)
}

// IteratorAscFrom 以升序遍历（只读）给定回调函数 `f` 的树。
// 参数 `key` 指定了遍历的起始项。`match` 指定了当 `key` 完全匹配时是否开始遍历，
// 否则使用索引搜索方式进行遍历。
// 若 `f` 返回 true，则继续遍历；若返回 false，则停止遍历。
func (tree *AVLTree) IteratorAscFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
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

func (tree *AVLTree) doIteratorAsc(node *AVLTreeNode, f func(key, value interface{}) bool) {
	for node != nil {
		if !f(node.Key, node.Value) {
			return
		}
		node = node.Next()
	}
}

// IteratorDesc 以降序遍历给定回调函数 `f` 的只读树。
// 如果 `f` 返回 true，则继续迭代；如果返回 false，则停止遍历。
func (tree *AVLTree) IteratorDesc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorDesc(tree.bottom(1), f)
}

// IteratorDescFrom 从指定的键(key)开始以降序方式遍历树（只读模式），并使用给定的回调函数`f`。
// 参数`key`指定了遍历的起始条目。`match`参数指定了如果`key`完全匹配时是否开始遍历，
// 否则将采用索引搜索方式进行遍历。
// 若`f`返回值为true，则继续遍历；若返回false，则停止遍历。
func (tree *AVLTree) IteratorDescFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
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

func (tree *AVLTree) doIteratorDesc(node *AVLTreeNode, f func(key, value interface{}) bool) {
	for node != nil {
		if !f(node.Key, node.Value) {
			return
		}
		node = node.Prev()
	}
}

func (tree *AVLTree) put(key interface{}, value interface{}, p *AVLTreeNode, qp **AVLTreeNode) bool {
	q := *qp
	if q == nil {
		tree.size++
		*qp = &AVLTreeNode{Key: key, Value: value, parent: p}
		return true
	}

	c := tree.getComparator()(key, q.Key)
	if c == 0 {
		q.Key = key
		q.Value = value
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	if tree.put(key, value, q, &q.children[a]) {
		return putFix(int8(c), qp)
	}
	return false
}

func (tree *AVLTree) remove(key interface{}, qp **AVLTreeNode) (value interface{}, fix bool) {
	q := *qp
	if q == nil {
		return nil, false
	}

	c := tree.getComparator()(key, q.Key)
	if c == 0 {
		tree.size--
		value = q.Value
		fix = true
		if q.children[1] == nil {
			if q.children[0] != nil {
				q.children[0].parent = q.parent
			}
			*qp = q.children[0]
			return
		}
		if removeMin(&q.children[1], &q.Key, &q.Value) {
			return value, removeFix(-1, qp)
		}
		return
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	value, fix = tree.remove(key, &q.children[a])
	if fix {
		return value, removeFix(int8(-c), qp)
	}
	return value, false
}

func removeMin(qp **AVLTreeNode, minKey *interface{}, minVal *interface{}) bool {
	q := *qp
	if q.children[0] == nil {
		*minKey = q.Key
		*minVal = q.Value
		if q.children[1] != nil {
			q.children[1].parent = q.parent
		}
		*qp = q.children[1]
		return true
	}
	fix := removeMin(&q.children[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

func putFix(c int8, t **AVLTreeNode) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.children[(c+1)/2].b == c {
		s = singleRotate(c, s)
	} else {
		s = doubleRotate(c, s)
	}
	*t = s
	return false
}

func removeFix(c int8, t **AVLTreeNode) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return false
	}

	if s.b == -c {
		s.b = 0
		return true
	}

	a := (c + 1) / 2
	if s.children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}

	if s.children[a].b == c {
		s = singleRotate(c, s)
	} else {
		s = doubleRotate(c, s)
	}
	*t = s
	return true
}

func singleRotate(c int8, s *AVLTreeNode) *AVLTreeNode {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doubleRotate(c int8, s *AVLTreeNode) *AVLTreeNode {
	a := (c + 1) / 2
	r := s.children[a]
	s.children[a] = rotate(-c, s.children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0
	return p
}

func rotate(c int8, s *AVLTreeNode) *AVLTreeNode {
	a := (c + 1) / 2
	r := s.children[a]
	s.children[a] = r.children[a^1]
	if s.children[a] != nil {
		s.children[a].parent = s
	}
	r.children[a^1] = s
	r.parent = s.parent
	s.parent = r
	return r
}

func (tree *AVLTree) bottom(d int) *AVLTreeNode {
	n := tree.root
	if n == nil {
		return nil
	}

	for c := n.children[d]; c != nil; c = n.children[d] {
		n = c
	}
	return n
}

// Prev 返回在AVL树中 inorder 遍历的前一个元素。
func (node *AVLTreeNode) Prev() *AVLTreeNode {
	return node.walk1(0)
}

// Next 返回AVL树中 inorder 遍历的下一个元素。
func (node *AVLTreeNode) Next() *AVLTreeNode {
	return node.walk1(1)
}

func (node *AVLTreeNode) walk1(a int) *AVLTreeNode {
	if node == nil {
		return nil
	}
	n := node
	if n.children[a] != nil {
		n = n.children[a]
		for n.children[a^1] != nil {
			n = n.children[a^1]
		}
		return n
	}

	p := n.parent
	for p != nil && p.children[a] == n {
		n = p
		p = p.parent
	}
	return p
}

func output(node *AVLTreeNode, prefix string, isTail bool, str *string) {
	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += fmt.Sprintf("%v\n", node.Key)
	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.children[0], newPrefix, true, str)
	}
}

// MarshalJSON 实现了 json.Marshal 接口所需的 MarshalJSON 方法。
func (tree AVLTree) MarshalJSON() (jsonBytes []byte, err error) {
	if tree.root == nil {
		return []byte("null"), nil
	}
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('{')
	tree.X遍历(func(key, value interface{}) bool {
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
func (tree *AVLTree) getComparator() func(a, b interface{}) int {
	if tree.comparator == nil {
		panic("comparator is missing for tree")
	}
	return tree.comparator
}
