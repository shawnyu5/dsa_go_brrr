package linkedlist

// const iterator
type Const_iterator[T comparable] struct {
	myList  *LinkedList[T]
	current *node[T]
}

// non const iterator
type Iterator[T comparable] struct {
	MyList  *LinkedList[T]
	Current *node[T]
}

// NewConstIterator constructs a new Const_iterator
func NewConstIterator[T comparable](curr *node[T], theList *LinkedList[T]) Const_iterator[T] {
	return Const_iterator[T]{myList: theList, current: curr}
}

// get returns the data from an iterator
func (it *Iterator[T]) get() T {
	return it.Current.data
}

// get returns the data from an Const_iterator
func (it *Const_iterator[T]) get() T {
	return it.current.data
}

// increment increments the iterator to the next node
func (it *Iterator[T]) increment() {
	if it.Current.next != nil {
		it.Current = it.Current.next
	}
}

// begin returns the first node in the list, excluding the sentinel node
func (it *Const_iterator[T]) begin() Const_iterator[T] {
	return Const_iterator[T]{myList: it.myList, current: it.myList.front.next}
}

// begin returns the first node in the list, excluding the sentinel node
func (it *Iterator[T]) cbegin() Iterator[T] {
	return Iterator[T]{MyList: it.MyList, Current: it.MyList.front.next}
}

// end returns the last node in the list, the sentinel node
func (it *Const_iterator[T]) cend() Iterator[T] {
	return Iterator[T]{MyList: it.myList, Current: it.myList.back}
}

// end returns the last node in the list, the sentinel node
func (it *Iterator[T]) end() Iterator[T] {
	return Iterator[T]{MyList: it.MyList, Current: it.MyList.back}
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
	if it.Current.prev != nil {
		it.Current = it.Current.prev
	}
}
