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

type Const_iterator[T comparable] struct {
	MyList  *LinkedList[T]
	Current *node[T]
}

type Iterator[T comparable] struct {
	MyList  *LinkedList[T]
	Current *node[T]
}

// NewConstIterator constructs a new Const_iterator
func NewConstIterator[T comparable](curr *node[T], theList *LinkedList[T]) Const_iterator[T] {
	return Const_iterator[T]{MyList: theList, Current: curr}
}

// NewList constructs a new LinkedList with sentinel nodes
func NewList[T comparable]() LinkedList[T] {
	list := LinkedList[T]{front: &node[T]{}, back: &node[T]{}}
	list.front.next = list.back
	list.back.prev = list.front
	return list
}

// NewIterator constructs a new Iterator
func NewIterator[T comparable](curr *node[T], theList *LinkedList[T]) Iterator[T] {
	return Iterator[T]{MyList: theList, Current: curr}
}

// get returns the data from an iterator
func (it Iterator[T]) get() T {
	return it.Current.data
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
