package linkedhashset

import "github.com/dddddai/go-utils/linkedhashmap"

type Node struct {
	inner *linkedhashmap.Node
}

func (n *Node) Pre() *Node {
	tn := n.inner.Pre()
	if tn == nil {
		return nil
	}
	return &Node{tn}
}
func (n *Node) Next() *Node {
	tn := n.inner.Next()
	if tn == nil {
		return nil
	}
	return &Node{tn}
}

func (n *Node) Value() interface{} {
	return n.inner.Key
}
