package singlylinkedlist

import "errors"

type Node struct {
	next  *Node
	Value interface{}
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func New() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

func (s *SinglyLinkedList) AddLast(val interface{}) {
	n := &Node{Value: val}
	if s.size == 0 {
		s.head = n
		s.tail = n
	} else {
		s.tail.next = n
		s.tail = n
	}
	s.size++
}

func (s *SinglyLinkedList) AddFirst(val interface{}) {
	n := &Node{Value: val}
	if s.size == 0 {
		s.head = n
		s.tail = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.size++
}

func (s *SinglyLinkedList) Last() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.tail.Value
}

func (s *SinglyLinkedList) First() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.head.Value
}

func (s *SinglyLinkedList) PollFirst() interface{} {
	if s.size == 0 {
		return nil
	}
	result := s.First()
	s.size--
	if s.size == 0 {
		s.head = nil
		s.tail = nil
	} else {
		s.head = s.head.next
	}
	return result
}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func (s *SinglyLinkedList) Empty() bool {
	return s.size == 0
}

func (s *SinglyLinkedList) Begin() *Node {
	return s.head
}

func (s *SinglyLinkedList) End() *Node {
	return s.tail
}

func (s *SinglyLinkedList) AddNext(n *Node, v interface{}) {
	n.next = &Node{Value: v, next: n.next}
	s.size++
}

func (s *SinglyLinkedList) RemoveNext(n *Node) error {
	if s.size == 0 {
		return errors.New("The SinglyLinkedList is empty!")
	}
	if n.next != nil {
		n.next = n.next.next
	}
	s.size--
	return nil
}

func (s *SinglyLinkedList) Clear() {
	s.head = nil
	s.tail = nil
	s.size = 0
}

func (n *Node) Next() *Node {
	return n.next
}
