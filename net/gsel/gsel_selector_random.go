	// 版权归GoFrame作者(https:	//goframe.org)所有。保留所有权利。
	//
	// 本源代码形式受MIT许可证条款约束。
	// 如果未随本文件一同分发MIT许可证副本，
	// 您可以在https:	//github.com/gogf/gf处获取。
	// md5:a9832f33b234e3f3

package gsel

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/internal/intlog"
	"github.com/gogf/gf/v2/util/grand"
)

type selectorRandom struct {
	mu    sync.RWMutex
	nodes Nodes
}

// ff:
func NewSelectorRandom() Selector {
	return &selectorRandom{
		nodes: make([]Node, 0),
	}
}

// ff:
// s:
// ctx:
// nodes:
func (s *selectorRandom) Update(ctx context.Context, nodes Nodes) error {
	intlog.Printf(ctx, `Update nodes: %s`, nodes.String())
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nodes = nodes
	return nil
}

// ff:
// s:
// ctx:
// node:
// done:
// err:
func (s *selectorRandom) Pick(ctx context.Context) (node Node, done DoneFunc, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.nodes) == 0 {
		return nil, nil, nil
	}
	node = s.nodes[grand.Intn(len(s.nodes))]
	intlog.Printf(ctx, `Picked node: %s`, node.Address())
	return node, nil, nil
}
