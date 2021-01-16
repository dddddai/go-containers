package linkedhashmap

type Node struct {
	pre   *Node
	next  *Node
	Key   interface{}
	Value interface{}
}

func (n *Node) add(x *Node) {
	n.next.pre = x
	x.next = n.next
	n.next = x
	x.pre = n
}

func (n *Node) Pre() *Node {
	if n.pre.pre == nil {
		return nil
	}
	return n.pre
}
func (n *Node) Next() *Node {
	if n.next.next == nil {
		return nil
	}
	return n.next
}
