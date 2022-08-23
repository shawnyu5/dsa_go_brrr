package linkedlist

// const iterator
type Const_iterator[T comparable] struct {
	myList  *LinkedList[T]
	current *node[T]
}

// non const iterator
type Iterator[T comparable] struct {
	myList  *LinkedList[T]
	current *node[T]
}

// newConstIterator constructs a new Const_iterator
func newConstIterator[T comparable](curr *node[T], theList *LinkedList[T]) Const_iterator[T] {
	return Const_iterator[T]{myList: theList, current: curr}
}

// newIterator constructs a new Iterator
func newIterator[T comparable](curr *node[T], theList *LinkedList[T]) Iterator[T] {
	return Iterator[T]{myList: theList, current: curr}
}

// Get returns the data from an iterator
func (it *Iterator[T]) Get() T {
	return it.current.data
}

// Get returns the data from an Const_iterator
func (it *Const_iterator[T]) Get() T {
	return it.current.data
}

// Begin returns the first node in the list, excluding the sentinel node
func (it *Iterator[T]) Begin() Iterator[T] {
	return Iterator[T]{myList: it.myList, current: it.myList.front.next}
}

// End returns the last node in the list, the sentinel node
func (it *Iterator[T]) End() Iterator[T] {
	return Iterator[T]{myList: it.myList, current: it.myList.back}
}

// CBegin returns the first node in the list, excluding the sentinel node
func (it *Const_iterator[T]) CBegin() Const_iterator[T] {
	return Const_iterator[T]{myList: it.myList, current: it.myList.front.next}
}

// CEnd returns the last node in the list, the sentinel node
func (it *Const_iterator[T]) CEnd() Const_iterator[T] {
	return Const_iterator[T]{myList: it.myList, current: it.myList.back}
}

// Decrement decrements the Const_iterator to the previous node
func (it *Const_iterator[T]) Decrement() {
	if it.current.prev != nil {
		it.current = it.current.prev
	}
}

// Increment increments the Const_iterator to the next node
func (it *Const_iterator[T]) Increment() {
	if it.current.next != nil {
		it.current = it.current.next
	}
}

// Increment increments the iterator to the next node
func (it *Iterator[T]) Increment() {
	if it.current.next != nil {
		it.current = it.current.next
	}
}

// Decrement decrements the iterator to the previous node
func (it *Iterator[T]) Decrement() {
	if it.current.prev != nil {
		it.current = it.current.prev
	}
}
