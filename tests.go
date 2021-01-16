package main

import (
	"fmt"
	"github.com/dddddai/go-utils/linkedhashmap"
	"github.com/dddddai/go-utils/linkedhashset"
	"github.com/dddddai/go-utils/singlylinkedlist"
)

const n = 3

func main() {
	testLinkedList()
	fmt.Println()
	testLinkedHashMap()
	fmt.Println()
	testLinkedHashSet()
	fmt.Println()
}

func testLinkedList() {
	list := singlylinkedlist.New()
	for i := 0; i < n; i++ {
		list.AddFirst(i)
	}
	for i := 0; i < n; i++ {
		list.AddLast(i)
	}
	for i := 0; i < n*2; i++ {
		fmt.Println(list.PollFirst())
	}
	fmt.Println(list.Size(), list.Empty())
	for i := 0; i < n; i++ {
		list.AddLast(i)
	}
	for p := list.Begin(); p != nil; p = p.Next() {
		if p.Next() != nil && p.Next().Value == 2 {
			list.RemoveNext(p)
		}
		fmt.Println(p.Value, " ")
	}
}

func testLinkedHashMap() {
	lhmap := linkedhashmap.New()
	for i := 0; i < n; i++ {
		lhmap.PutLast(i, i<<1)
	}
	for i := 100; i < 100+n; i++ {
		lhmap.PutFirst(i, i<<1)
	}
	for i := 0; i < n*2; i++ {
		node, _ := lhmap.PollFirst()
		fmt.Println(node.Key, node.Value)
	}
	for i := 100; i < 100+n; i++ {
		lhmap.PutFirst(i, i<<1)
	}
	lhmap.Remove(100)
	lhmap.Remove(102)
	for p := lhmap.Tail(); p != nil; p = p.Pre() {
		fmt.Println(p.Key, p.Value)
	}
}

func testLinkedHashSet() {
	lhset := linkedhashset.New()
	for i := 0; i < n; i++ {
		lhset.AddLast(i)
	}
	for i := 100; i < 100+n; i++ {
		lhset.AddFirst(i)
	}
	for i := 0; i < n*2; i++ {
		x, _ := lhset.PollFirst()
		fmt.Println(x)
	}
	for i := 100; i < 100+n; i++ {
		lhset.AddFirst(i)
	}
	lhset.Remove(101)
	lhset.Remove(100)
	for p := lhset.Tail(); p != nil; p = p.Pre() {
		fmt.Println(p.Value())
	}
}
