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

func (lhset *LinkedHashSet) GetFirst() interface{} {
	return lhset.lhmap.GetFirst().Key
}

func (lhset *LinkedHashSet) GetLast() interface{} {
	return lhset.lhmap.GetLast().Key
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

func (lhset *LinkedHashSet) PollFirst() interface{} {
	return lhset.lhmap.PollFirst()
}

func (lhset *LinkedHashSet) PollLast() interface{} {
	return lhset.lhmap.PollLast()
}

func (lhset *LinkedHashSet) Size() int {
	return lhset.lhmap.Size()
}

func (lhset *LinkedHashSet) Empty() bool {
	return lhset.lhmap.Size() == 0
}

func (lhset *LinkedHashSet) Clear() {
	for !lhset.Empty() {
		lhset.PollFirst()
	}
}