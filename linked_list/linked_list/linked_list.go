package linkedlist

type linkedList[T any] struct {
	front *node[T]
	back  *node[T]
}

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

func (linkedList[T]) NewList() linkedList[T] {
	list := linkedList[T]{front: &node[T]{}, back: &node[T]{}}
	return list
}
