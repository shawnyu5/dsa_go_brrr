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

// Search searches the list for the data passed in. If the data is found, this function returns an iterator pointing to the node containing the data. Else returns an iterator pointing to the end of the list
func (l *LinkedList[T]) Search(data T) Iterator[T] {
	it := l.It.Begin()
	found := false

	// as long as the next node is not the sentinel node, we keep looking
	for it.current.next != l.It.End().current {
		// if we have found the data, return the iterator
		if it.Get() == data {
			found = true
			break
		}
		it.Increment()
	}

	if !found {
		return l.It.End()
	}
	return it
}

// Split splits the list into two halves, returns 2 iterators pointing to the beginning of each half
func (l *LinkedList[T]) Split(begin Iterator[T], end Iterator[T]) (Iterator[T], Iterator[T]) {
	midPoint := begin
	nav := begin
	nav.Increment() // nav is one ahead of midPoint

	for nav != l.It.End() {
		nav.Increment()
		if nav != l.It.End() {
			midPoint.Increment()
			nav.Increment()
		}
	}

	midPoint.Increment()
	return l.It.Begin(), midPoint
}

// Erase erases the node that the iterator is pointing to. And returns an iterator pointing to the next node
func (l *LinkedList[T]) Erase(it Iterator[T]) Iterator[T] {
	found := false
	rmPosition := l.It.Begin()

	// look till the end of the list
	for !found && rmPosition != l.It.End() {
		if it == rmPosition {
			found = true
			break
		} else {
			rmPosition.Increment()
		}
	}

	if found {
		current := rmPosition.current
		prev := current.prev
		next := current.next

		prev.next = next
		next.prev = prev

		it.Increment()
		current = nil
		l.numRecords--
	}

	return it
}

// TODO: implement merge sort later
func (l *LinkedList[T]) Sort(first *Iterator[T], last *Iterator[T]) {
	panic("Not implemented")

}

func (l *LinkedList[T]) mergeSort(n node[T]) {

}
