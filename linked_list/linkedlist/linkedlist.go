package linkedlist

type LinkedList[T comparable] struct {
	front *node[T]
	back  *node[T]
	// number of records in the list
	numRecords int
	// const iterator
	Cit Const_iterator[T]
	// iterator
	It Iterator[T]
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
	begin := list.begin()
	list.Cit = newConstIterator(&begin, &list)
	list.It = newIterator(&begin, &list)
	return list
}

// begin returns the first node in the list, excluding the sentinel node
func (l *LinkedList[T]) begin() node[T] {
	return *l.front.next
}

// Insert inserts the data before the node that it is pointing to
func (l *LinkedList[T]) Insert(it *Iterator[T], data T) {
	// the node the iterator is pointing to
	insertLocation := it.current

	// construct a new node from the data passed in
	nn := node[T]{data: data, next: insertLocation, prev: insertLocation.prev}

	insertLocation.prev.next = &nn
	insertLocation.prev = &nn
	l.numRecords++
	// move the iterator to the new node
	it.Decrement()
}

// Size returns the number of records in the list
func (l *LinkedList[T]) Size() int {
	return l.numRecords
}
