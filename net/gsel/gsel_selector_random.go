// 版权归GoFrame作者(https://goframe.org)所有。保留所有权利。
//
// 本源代码形式受MIT许可证条款约束。
// 如果未随本文件一同分发MIT许可证副本，
// 您可以在https://github.com/gogf/gf处获取。
// md5:a9832f33b234e3f3

package gsel

import (
	"context"
	"sync"

	"github.com/888go/goframe/internal/intlog"
	grand "github.com/888go/goframe/util/grand"
)

type selectorRandom struct {
	mu    sync.RWMutex
	nodes Nodes
}

func NewSelectorRandom() Selector {
	return &selectorRandom{
		nodes: make([]Node, 0),
	}
}

func (s *selectorRandom) Update(ctx context.Context, nodes Nodes) error {
	intlog.Printf(ctx, `Update nodes: %s`, nodes.String())
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nodes = nodes
	return nil
}

func (s *selectorRandom) Pick(ctx context.Context) (node Node, done DoneFunc, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.nodes) == 0 {
		return nil, nil, nil
	}
	node = s.nodes[grand.X整数(len(s.nodes))]
	intlog.Printf(ctx, `Picked node: %s`, node.Address())
	return node, nil, nil
}
