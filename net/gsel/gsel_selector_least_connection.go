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

	gtype "github.com/888go/goframe/container/gtype"
	"github.com/888go/goframe/internal/intlog"
)

type selectorLeastConnection struct {
	mu    sync.RWMutex
	nodes []*leastConnectionNode
}

type leastConnectionNode struct {
	Node
	inflight *gtype.Int
}

func NewSelectorLeastConnection() Selector {
	return &selectorLeastConnection{
		nodes: make([]*leastConnectionNode, 0),
	}
}

func (s *selectorLeastConnection) Update(ctx context.Context, nodes Nodes) error {
	intlog.Printf(ctx, `Update nodes: %s`, nodes.String())
	var newNodes []*leastConnectionNode
	for _, v := range nodes {
		node := v
		newNodes = append(newNodes, &leastConnectionNode{
			Node:     node,
			inflight: gtype.NewInt(),
		})
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nodes = newNodes
	return nil
}

func (s *selectorLeastConnection) Pick(ctx context.Context) (node Node, done DoneFunc, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var pickedNode *leastConnectionNode
	if len(s.nodes) == 1 {
		pickedNode = s.nodes[0]
	} else {
		for _, v := range s.nodes {
			if pickedNode == nil {
				pickedNode = v
			} else if v.inflight.X取值() < pickedNode.inflight.X取值() {
				pickedNode = v
			}
		}
	}
	pickedNode.inflight.Add(1)
	done = func(ctx context.Context, di DoneInfo) {
		pickedNode.inflight.Add(-1)
	}
	node = pickedNode.Node
	intlog.Printf(ctx, `Picked node: %s`, node.Address())
	return node, done, nil
}
