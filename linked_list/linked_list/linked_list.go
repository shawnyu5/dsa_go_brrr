package linkedlist

type linkedList[T comparable] struct {
	front *node[T]
	back  *node[T]
}

type node[T comparable] struct {
	next *node[T]
	prev *node[T]
	data T
}

type Const_iterator[T comparable] struct {
	MyList  *linkedList[T]
	Current *node[T]
}

type Iterator[T comparable] struct {
	Const_iterator[T]
}

// NewList constructs a new list with sentinel nodes
// returns a new linkedlist
func (linkedList[T]) NewList() linkedList[T] {
	list := linkedList[T]{front: &node[T]{}, back: &node[T]{}}
	return list
}
