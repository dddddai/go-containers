package main

import (
	"fmt"
	"go-utils/linkedhashmap"
	"go-utils/singlylinkedlist"
)

const n=10

func main() {
	testLinkedList()
	fmt.Println()
	testLinkedHashMap()
}

func testLinkedList() {
	list:=singlylinkedlist.New()
	for i:=0;i<n;i++ {
		list.AddFirst(i)
	}
	for i:=0;i<n;i++ {
		list.AddLast(i)
	}
	for i:=0;i<n*2;i++ {
		fmt.Print(list.PollFirst()," ")
	}
	fmt.Println("\n",list.Size(),list.Empty())
	for i:=0;i<n;i++ {
		list.AddLast(i)
	}
	for p:=list.Begin();p!=nil;p=p.Next(){
		if p.Next()!=nil && p.Next().Value==5 {
			list.RemoveNext(p)
		}
		fmt.Print(p.Value," ")
	}
}

func testLinkedHashMap(){
	lhmap:=linkedhashmap.New()
	for i:=0;i<n;i++ {
		lhmap.PutLast(i,i<<1)
	}
	for i:=100;i<100+n;i++ {
		lhmap.PutFirst(i,i<<1)
	}
	for i:=0;i<n*2;i++ {
		node,_:= lhmap.PollFirst()
		fmt.Println(node.Key,node.Val)
	}
	for i:=100;i<100+n;i++ {
		lhmap.PutFirst(i,i<<1)
	}
	for p:=lhmap.Tail();p!=nil;p=p.Pre() {
		fmt.Println(p.Key,p.Val)
	}
}