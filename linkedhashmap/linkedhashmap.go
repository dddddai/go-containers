package linkedhashmap

import "errors"

type LinkedHashMap struct {
	dummyHead *Node
	dummyTail *Node
	m         map[interface{}]*Node
}

func New() *LinkedHashMap {
	dh,dt:=new(Node),new(Node)
	dh.next=dt
	dt.pre=dh
	return &LinkedHashMap{
		dummyHead: dh,
		dummyTail: dt,
		m:         make(map[interface{}]*Node),
	}
}

func NewWithCapacity(capacity int) *LinkedHashMap {
	dh,dt:=new(Node),new(Node)
	dh.next=dt
	dt.pre=dh
	return &LinkedHashMap{
		dummyHead: dh,
		dummyTail: dt,
		m:         make(map[interface{}]*Node, capacity),
	}
}

func (lhmap *LinkedHashMap) Has(key interface{}) bool {
	_, ok := lhmap.m[key]
	return ok
}

func (lhmap *LinkedHashMap) Get(key interface{}) (interface{}, bool) {
	n, ok := lhmap.m[key]
	return n.Value, ok
}

func (lhmap *LinkedHashMap) GetOrDefault(key, dft interface{}) interface{} {
	if val, ok := lhmap.Get(key); ok {
		return val
	}
	return dft
}

func (lhmap *LinkedHashMap) First() (*Node, error) {
	if len(lhmap.m) == 0 {
		return nil, errors.New("First() was called when empty!")
	}
	return lhmap.dummyHead.next, nil
}

func (lhmap *LinkedHashMap) Last() (*Node, error) {
	if len(lhmap.m) == 0 {
		return nil, errors.New("Last() was called when empty!")
	}
	return lhmap.dummyTail.pre, nil
}

func (lhmap *LinkedHashMap) Put(key, val interface{}) {
	n, ok := lhmap.m[key]
	if ok {
		n.Value=val
	}else{
		n = &Node{Key: key, Value: val}
		lhmap.dummyTail.pre.add(n)
		lhmap.m[key] = n
	}
}


func (lhmap *LinkedHashMap) PutFirst(key, val interface{}) {
	n, ok := lhmap.m[key]
	if ok {
		n.Value =val
		lhmap.removeNode(n)
		lhmap.dummyHead.add(n)
	} else {
		n = &Node{Key: key, Value: val}
		lhmap.dummyHead.add(n)
		lhmap.m[key] = n
	}
}

func (lhmap *LinkedHashMap) PutLast(key, val interface{}) {
	n, ok := lhmap.m[key]
	if ok {
		n.Value =val
		lhmap.removeNode(n)
		lhmap.dummyTail.pre.add(n)
	} else {
		n = &Node{Key: key, Value: val}
		lhmap.dummyTail.pre.add(n)
		lhmap.m[key] = n
	}
}

func (lhmap *LinkedHashMap) Remove(key interface{}) bool {
	n, ok := lhmap.m[key]
	if !ok {
		return false
	}
	delete(lhmap.m, key)
	lhmap.removeNode(n)
	return true
}

func (lhmap *LinkedHashMap) PollFirst() (*Node, error) {
	result, err := lhmap.First()
	if err != nil {
		return nil, err
	}
	delete(lhmap.m, result.Key)
	lhmap.removeNode(lhmap.dummyHead.next)
	return result, nil
}

func (lhmap *LinkedHashMap) PollLast() (*Node, error) {
	result, err := lhmap.Last()
	if err != nil {
		return nil, err
	}
	delete(lhmap.m, result.Key)
	lhmap.removeNode(lhmap.dummyTail.pre)
	return result, nil
}

func (lhmap *LinkedHashMap) Size() int {
	return len(lhmap.m)
}

func (lhmap *LinkedHashMap) Empty() bool {
	return len(lhmap.m) == 0
}

func (lhmap *LinkedHashMap) Head() *Node {
	return lhmap.dummyHead.next
}

func (lhmap *LinkedHashMap) Tail() *Node {
	return lhmap.dummyTail.pre
}

func (lhmap *LinkedHashMap) removeNode(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}
