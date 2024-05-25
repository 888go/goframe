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
)

// AVLTree 存储AVL树的元素。. md5:d108e2d6ca60747c
type AVLTree struct {
	mu         rwmutex.RWMutex
	root       *AVLTreeNode
	comparator func(v1, v2 interface{}) int
	size       int
}

// AVLTreeNode 是树中的一个元素。. md5:ae1bf04ae171ca4e
type AVLTreeNode struct {
	Key      interface{}
	Value    interface{}
	parent   *AVLTreeNode
	children [2]*AVLTreeNode
	b        int8
}

// NewAVLTree 创建一个带有自定义键比较器的AVL树。
// 参数 `safe` 用于指定是否需要并发安全，默认为false。
// md5:dfb13f71bc07620c
func NewAVLTree(comparator func(v1, v2 interface{}) int, safe ...bool) *AVLTree {
	return &AVLTree{
		mu:         rwmutex.Create(safe...),
		comparator: comparator,
	}
}

// NewAVLTreeFrom 使用自定义的键比较器和数据映射初始化一个AVL树。
// 参数`safe`用于指定是否在并发安全模式下使用树，其默认值为false。
// md5:856b75ecd9dc1540
func NewAVLTreeFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *AVLTree {
	tree := NewAVLTree(comparator, safe...)
	for k, v := range data {
		tree.put(k, v, nil, &tree.root)
	}
	return tree
}

// Clone 返回一个新的树，其中包含当前树的副本。. md5:256477216ae712b7
func (tree *AVLTree) Clone() *AVLTree {
	newTree := NewAVLTree(tree.comparator, tree.mu.IsSafe())
	newTree.Sets(tree.Map())
	return newTree
}

// Set 将节点插入到树中。. md5:8af837873ed60d6a
func (tree *AVLTree) Set(key interface{}, value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.put(key, value, nil, &tree.root)
}

// 设置批处理将键值对添加到树中。. md5:70c6ec85c8b7476c
func (tree *AVLTree) Sets(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for key, value := range data {
		tree.put(key, value, nil, &tree.root)
	}
}

// Search 函数使用给定的 `key` 在树中进行查找。
// 第二个返回参数 `found` 为 true 表示找到了键，否则为 false。
// md5:d151c3783cadda2c
func (tree *AVLTree) Search(key interface{}) (value interface{}, found bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	if node, found := tree.doSearch(key); found {
		return node.Value, true
	}
	return nil, false
}

// doSearch 使用给定的 `key` 在树中进行搜索。
// 第二个返回参数 `found` 如果找到了键，则为 true，否则为 false。
// md5:46c1cf26ceea8b4b
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

// Get 通过`key`在树中搜索节点，并返回其值，如果`key`在树中未找到，则返回nil。. md5:2e2483db20a69167
func (tree *AVLTree) Get(key interface{}) (value interface{}) {
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

// GetOrSet 通过键返回值，
// 如果该键不存在，则使用给定的`value`设置值，然后返回这个值。
// md5:d8f89b6dec47292b
func (tree *AVLTree) GetOrSet(key interface{}, value interface{}) interface{} {
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
func (tree *AVLTree) GetOrSetFunc(key interface{}, f func() interface{}) interface{} {
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
func (tree *AVLTree) GetOrSetFuncLock(key interface{}, f func() interface{}) interface{} {
	if v, ok := tree.Search(key); !ok {
		return tree.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// GetVar 函数通过给定的 `key` 返回一个 gvar.Var，其值为对应的变量。
// 返回的 gvar.Var 不是并发安全的。
// md5:a04747902e4bf242
func (tree *AVLTree) GetVar(key interface{}) *gvar.Var {
	return gvar.New(tree.Get(key))
}

// GetVarOrSet 返回一个从 GetVarOrSet 获取的结果的 gvar.Var。返回的 gvar.Var 不是线程安全的。
// md5:089beb08264e18cf
func (tree *AVLTree) GetVarOrSet(key interface{}, value interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSet(key, value))
}

// GetVarOrSetFunc 返回一个 gvar.Var，其结果来自 GetOrSetFunc。
// 返回的 gvar.Var 不是线程安全的。
// md5:8c97b145faade5ae
func (tree *AVLTree) GetVarOrSetFunc(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFunc(key, f))
}

// GetVarOrSetFuncLock 返回一个gvar.Var，其结果来自GetOrSetFuncLock。
// 返回的gvar.Var是非并发安全的。
// md5:90c22300c2187ce4
func (tree *AVLTree) GetVarOrSetFuncLock(key interface{}, f func() interface{}) *gvar.Var {
	return gvar.New(tree.GetOrSetFuncLock(key, f))
}

// SetIfNotExist 如果键`key`不存在，则将`value`设置到映射中，并返回true。如果键`key`已存在，且`value`将被忽略，函数返回false。
// md5:f80895920828f03e
func (tree *AVLTree) SetIfNotExist(key interface{}, value interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, value)
		return true
	}
	return false
}

// SetIfNotExistFunc 使用回调函数`f`的返回值设置值，并返回true。
// 如果`key`已存在，则返回false，且`value`会被忽略。
// md5:326c0b7c63d813e7
func (tree *AVLTree) SetIfNotExistFunc(key interface{}, f func() interface{}) bool {
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
func (tree *AVLTree) SetIfNotExistFuncLock(key interface{}, f func() interface{}) bool {
	if !tree.Contains(key) {
		tree.doSetWithLockCheck(key, f)
		return true
	}
	return false
}

// Contains 检查键 `key` 是否存在于树中。. md5:77fd85af8e586867
func (tree *AVLTree) Contains(key interface{}) bool {
	_, ok := tree.Search(key)
	return ok
}

// Remove 通过键从树中移除节点。
// 键应符合比较器的类型断言，否则方法将 panic。
// md5:23794cd4708d8756
func (tree *AVLTree) Remove(key interface{}) (value interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	value, _ = tree.remove(key, &tree.root)
	return
}

// 通过`keys`移除树中的批量删除值。. md5:4620c81ac88b2936
func (tree *AVLTree) Removes(keys []interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	for _, key := range keys {
		tree.remove(key, &tree.root)
	}
}

// IsEmpty 如果树中不包含任何节点，则返回true。. md5:8f7ae813360d880b
func (tree *AVLTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size 返回树中的节点数量。. md5:d437d5852f80de5c
func (tree *AVLTree) Size() int {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	return tree.size
}

// Keys 返回所有键，按升序排列。. md5:c2a692ea3491e160
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

// Values返回根据键值升序排列的所有值。. md5:6268d60d7aa20c91
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

// Left 返回 AVL 树中的最小元素
// 如果树为空，则返回 nil。
// md5:d6b4c070feb60521
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

// Right 返回AVL树中的最大元素，如果树为空则返回nil。
// md5:7f0d34ae61ed561f
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

// Floor 找到输入键的地板节点，如果没有找到地板节点，则返回nil。第二个返回参数表示是否找到了地板，如果找到为true，否则为false。
// 
// 地板节点定义为大于或等于给定节点的最大节点。可能找不到地板节点，原因可能是树为空，或者树中的所有节点都大于给定节点。
// 
// 键应遵循比较器的类型断言，否则方法会 panic。
// md5:720f6000179912eb
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

// Ceiling 找到输入键的天花板节点，如果没有找到天花板节点，则返回nil。第二个返回参数表示是否找到了天花板（true）或未找到（false）。
// 
// 定义天花板节点为大于或等于给定节点的最小节点。可能找不到天花板节点，因为树为空，或者树中的所有节点都小于给定节点。
// 
// 键应遵循比较器的类型断言，否则方法会 panic。
// md5:6b92342a03f9f586
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

// Clear 从树中移除所有节点。. md5:a7db742922264980
func (tree *AVLTree) Clear() {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
}

// 使用给定的`data`替换树中的数据。. md5:ff636c579597f294
func (tree *AVLTree) Replace(data map[interface{}]interface{}) {
	tree.mu.Lock()
	defer tree.mu.Unlock()
	tree.root = nil
	tree.size = 0
	for key, value := range data {
		tree.put(key, value, nil, &tree.root)
	}
}

// String 返回容器的字符串表示形式. md5:7daf925d59987319
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

// Print 将树打印到标准输出。. md5:24fd6288549a501b
func (tree *AVLTree) Print() {
	fmt.Println(tree.String())
}

// Map 返回所有键值对项作为一个映射。. md5:c12ca822a6c71dc1
func (tree *AVLTree) Map() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[key] = value
		return true
	})
	return m
}

// MapStrAny 将所有键值对作为 map[string]interface{} 返回。. md5:412456aafc43f7a8
func (tree *AVLTree) MapStrAny() map[string]interface{} {
	m := make(map[string]interface{}, tree.Size())
	tree.IteratorAsc(func(key, value interface{}) bool {
		m[gconv.String(key)] = value
		return true
	})
	return m
}

// Flip 将树中的键值对交换为值键。
// 请注意，你应该确保值的类型与键相同，否则比较器会panic。
//
// 如果值的类型与键不同，你需要传递新的 `comparator`。
// md5:e71ceac22aee55f1
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

// Iterator 是 IteratorAsc 的别名。. md5:1bfdea306db62845
func (tree *AVLTree) Iterator(f func(key, value interface{}) bool) {
	tree.IteratorAsc(f)
}

// IteratorFrom是IteratorAscFrom的别名。. md5:6d3d506bcb5fe942
func (tree *AVLTree) IteratorFrom(key interface{}, match bool, f func(key, value interface{}) bool) {
	tree.IteratorAscFrom(key, match, f)
}

// IteratorAsc 使用给定的回调函数 `f` 以升序遍历树（只读）。如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c13b99ae40add3b0
func (tree *AVLTree) IteratorAsc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorAsc(tree.bottom(0), f)
}

// IteratorAscFrom 从给定的回调函数 `f` 以升序遍历树。
// 参数 `key` 指定了遍历的起始条目。`match` 参数指定如果 `key` 完全匹配时是否开始遍历，否则使用索引搜索进行遍历。
// 如果 `f` 返回 true，则继续遍历；如果返回 false，则停止遍历。
// md5:c04855bbd3989808
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

// IteratorDesc 以降序方式遍历树，使用给定的回调函数 `f`。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:f6740ea55dafe4bb
func (tree *AVLTree) IteratorDesc(f func(key, value interface{}) bool) {
	tree.mu.RLock()
	defer tree.mu.RUnlock()
	tree.doIteratorDesc(tree.bottom(1), f)
}

// IteratorDescFrom 以降序方式遍历树，使用给定的回调函数 `f`。参数 `key` 指定开始遍历的条目。`match` 表示是否在 `key` 完全匹配时开始遍历，否则使用索引搜索遍历。如果 `f` 返回 true，则继续遍历；否则停止。
// md5:e6bb2f7d12ab34f6
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

// Prev 返回AVL树中序遍历的上一个元素。
// md5:d859f5a91f8afa30
func (node *AVLTreeNode) Prev() *AVLTreeNode {
	return node.walk1(0)
}

// Next 返回AVL树中序遍历的下一个元素。
// md5:bf33a084df1455d4
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

// MarshalJSON 实现了接口 MarshalJSON 以供 json.Marshal 使用。. md5:43c3b36e60a18f9a
func (tree AVLTree) MarshalJSON() (jsonBytes []byte, err error) {
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
func (tree *AVLTree) getComparator() func(a, b interface{}) int {
	if tree.comparator == nil {
		panic("comparator is missing for tree")
	}
	return tree.comparator
}
