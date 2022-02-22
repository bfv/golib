package collections

import "fmt"

type Node[K comparable, V any] struct {
	prev *Node[K, V]
	next *Node[K, V]
	key  K
	val  V
}

type LinkedList[K comparable, V any] struct {
	head *Node[K, V]
	tail *Node[K, V]
}

func (l *LinkedList[K, V]) Insert(k K, v V) {
	list := &Node[K, V]{
		next: l.head,
		key:  k,
		val:  v,
	}

	if l.head != nil {
		l.head.prev = list
	}
	l.head = list

	list = l.head
	for list.next != nil {
		list = list.next
	}
	l.tail = list
}

func (l *LinkedList[K, V]) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}
