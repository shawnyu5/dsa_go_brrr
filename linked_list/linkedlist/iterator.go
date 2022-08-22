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

// NewConstIterator constructs a new Const_iterator
func NewConstIterator[T comparable](curr *node[T], theList *LinkedList[T]) Const_iterator[T] {
	return Const_iterator[T]{myList: theList, current: curr}
}

// NewIterator constructs a new Iterator
func NewIterator[T comparable](curr *node[T], theList *LinkedList[T]) Iterator[T] {
	return Iterator[T]{myList: theList, current: curr}
}

// get returns the data from an iterator
func (it *Iterator[T]) get() T {
	return it.current.data
}

// get returns the data from an Const_iterator
func (it *Const_iterator[T]) get() T {
	return it.current.data
}

// increment increments the iterator to the next node
func (it *Iterator[T]) increment() {
	if it.current.next != nil {
		it.current = it.current.next
	}
}

// begin returns the first node in the list, excluding the sentinel node
func (it *Iterator[T]) begin() Iterator[T] {
	return Iterator[T]{myList: it.myList, current: it.myList.front.next}
}

// end returns the last node in the list, the sentinel node
func (it *Iterator[T]) end() Iterator[T] {
	return Iterator[T]{myList: it.myList, current: it.myList.back}
}

// cbegin returns the first node in the list, excluding the sentinel node
func (it *Const_iterator[T]) cbegin() Const_iterator[T] {
	return Const_iterator[T]{myList: it.myList, current: it.myList.front.next}
}

// cend returns the last node in the list, the sentinel node
func (it *Const_iterator[T]) cend() Const_iterator[T] {
	return Const_iterator[T]{myList: it.myList, current: it.myList.back}
}

// decrement decrements the Const_iterator to the previous node
func (it *Const_iterator[T]) decrement() {
	if it.current.prev != nil {
		it.current = it.current.prev
	}
}

// increment increments the Const_iterator to the next node
func (it *Const_iterator[T]) increment() {
	if it.current.next != nil {
		it.current = it.current.next
	}
}

// decrement decrements the iterator to the previous node
func (it *Iterator[T]) decrement() {
	if it.current.prev != nil {
		it.current = it.current.prev
	}
}
