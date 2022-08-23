package linkedlist

import "fmt"

type LinkedList[T comparable] struct {
	front *node[T]
	back  *node[T]
	// number of records in the list
	numRecords int
}

type node[T comparable] struct {
	data T
	next *node[T]
	prev *node[T]
}

// NewList constructs a new LinkedList with sentinel nodes
func NewList[T comparable]() LinkedList[T] {
	list := LinkedList[T]{front: &node[T]{}, back: &node[T]{}}
	list.front.next = list.back
	list.back.prev = list.front
	return list
}

// begin returns the first node in the list, excluding the sentinel node
func (l *LinkedList[T]) begin() node[T] {
	return *l.front.next
}

// Insert inserts the data before the node that it is pointing to
func (l *LinkedList[T]) Insert(it *Iterator[T], data T) {
	// the node the iterator is pointing to
	ogNode := it.current
	// find where the iterator is pointing to
	// for ogNode.next != nil && ogNode.data != it.get() {
	// ogNode = *ogNode.next
	// }
	fmt.Println(fmt.Sprintf("insert ogNode: %+v", ogNode)) // __AUTO_GENERATED_PRINT_VAR__

	// construct a new node from the data passed in
	nn := node[T]{data: data, next: ogNode, prev: ogNode.prev}

	// if we are at the tail, add the new node between front and back sentinel
	if ogNode.next == nil {
		l.back.prev.next = &nn
		l.back.prev = &nn
	} else { // if we are not at end, insert normally
		ogNode.prev.next = &nn
		ogNode.prev = &nn
	}
	l.numRecords++
	// move the iterator to the new node
	it.decrement()
}

// size returns the number of records in the list
func (l *LinkedList[T]) size() int {
	return l.numRecords
}
