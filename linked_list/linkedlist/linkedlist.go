package linkedlist

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

// insert inserts the data before the node that is refered to by it
func (l *LinkedList[T]) insert(it *Iterator[T], data T) {
	// the starting node
	ogNode := l.front.next
	// find where the iterator is pointing to
	for ogNode.next != nil && ogNode.data != it.get() {
		ogNode = ogNode.next
	}

	// construct a new node from the data passed in
	nn := node[T]{data: data, next: ogNode.next, prev: ogNode.prev}

	if ogNode.next == nil {
		l.back.prev.next = &nn
		l.back.prev = &nn
	} else {
		ogNode.prev.next = &nn
		ogNode.prev = &nn
	}
	l.numRecords++
}

// NumRecords returns the number of records in the list
func (l *LinkedList[T]) NumRecords() int {
	return l.numRecords
}
