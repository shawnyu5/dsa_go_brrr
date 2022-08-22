package linkedlist

import "testing"

// TestIteratorGet tests get function of the iterator returns the correct data inside the iterator
func TestIteratorGet(t *testing.T) {
	n := node[int]{data: 10}
	it := &Iterator[int]{current: &n}
	data := it.get()

	if data != 10 {
		t.Error("Iterator get() does not return the correct data")
	}
}

// TestIteratorIncreament tests increment function of the iterator increments the iterator to the next node
func TestIteratorIncreament(t *testing.T) {
	n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
	it := &Iterator[int]{current: &n}
	if it.get() != 10 {
		t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 10, it.get())
	}
	it.increment()
	if it.get() != 20 {
		t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 20, it.get())
	}

	it.increment()
	if it.get() != 30 {
		t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 30, it.get())
	}
}

// TestIteratorIncreament tests increment function of the const_iterator increments the iterator to the next node
func TestConstIteratorIncreament(t *testing.T) {
	n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
	list := LinkedList[int]{front: &node[int]{data: 10}, back: &node[int]{data: 20}}

	it := NewConstIterator(&n, &list)

	if it.get() != 10 {
		t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 10, it.get())
	}
	it.increment()
	if it.get() != 20 {
		t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 20, it.get())
	}

	it.increment()
	if it.get() != 30 {
		t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 30, it.get())
	}
}

// TestBegin tests begin function of the iterator and const_iterator returns the beginning of the list
func TestBegin(t *testing.T) {
	n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
	// construct a new list with sentinel nodes
	list := NewList[int]()

	it := NewIterator(&n, &list)
	// begin and end should point to the same sentinel node
	if it.begin() != it.end() {
		t.Error("Iterator begin() does not return the correct data")
	}
}

// TestCbegin tests begin function of the iterator and const_iterator returns the beginning of the list
func TestCbegin(t *testing.T) {
	n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
	// construct a new list with sentinel nodes
	list := NewList[int]()

	it := NewConstIterator(&n, &list)
	if it.cbegin() != it.cend() {
		t.Error("Iterator begin() does not return the correct data")
	}
}
