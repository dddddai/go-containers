package linkedhashset

import m "github.com/dddddai/go-utils/linkedhashmap"

type LinkedHashSet struct {
	lhmap *m.LinkedHashMap
}

func New() *LinkedHashSet {
	return &LinkedHashSet{lhmap: m.New()}
}

func NewWithCapacity(capacity int) *LinkedHashSet {
	return &LinkedHashSet{lhmap: m.NewWithCapacity(capacity)}
}

func (lhset *LinkedHashSet) Has(key interface{}) bool {
	return lhset.lhmap.Has(key)
}

func (lhset *LinkedHashSet) First() (interface{}, error) {
	n,err:=lhset.lhmap.First()
	if err!=nil {
		return nil,err
	}
	return n.Key,nil
}

func (lhset *LinkedHashSet) Last() (interface{}, error) {
	n,err:=lhset.lhmap.Last()
	if err!=nil {
		return nil,err
	}
	return n.Key,nil
}

func (lhset *LinkedHashSet) AddFirst(val interface{}) {
	lhset.lhmap.PutFirst(val, struct{}{})
}

func (lhset *LinkedHashSet) AddLast(val interface{}) {
	lhset.lhmap.PutLast(val, struct{}{})
}

func (lhset *LinkedHashSet) Remove(val interface{}) bool {
	return lhset.lhmap.Remove(val)
}

func (lhset *LinkedHashSet) PollFirst() (interface{}, error) {
	n,err:=lhset.lhmap.PollFirst()
	if err!=nil {
		return nil,err
	}
	return n.Key,nil
}

func (lhset *LinkedHashSet) PollLast() (interface{}, error) {
	n,err:=lhset.lhmap.PollLast()
	if err!=nil {
		return nil,err
	}
	return n.Key,nil
}

func (lhset *LinkedHashSet) Size() int {
	return lhset.lhmap.Size()
}

func (lhset *LinkedHashSet) Empty() bool {
	return lhset.lhmap.Size() == 0
}

func (lhset *LinkedHashSet) Head() *Node {
	return &Node{lhset.lhmap.Head()}
}

func (lhset *LinkedHashSet) Tail() *Node {
	return &Node{lhset.lhmap.Tail()}
}