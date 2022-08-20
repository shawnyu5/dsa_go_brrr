package linkedlist

import "testing"

// TestList test the linked list interface
func TestList(t *testing.T) {
	list := NewList[int]()
	n := &node[int]{data: 20}
	it := NewIterator(n, &list)
	list.insert(&it, 20)
}

func TestIteratorGet(t *testing.T) {
	n := node[int]{data: 10}
	it := &Iterator[int]{Current: &n}
	data := it.get()

	if data != 10 {
		t.Fatal("Iterator get() does not return the correct data")
	}
}

// TestNewList tests the NewList function returns a LinkedList with sentinel nodes, and the front and back nodes are pointing to the right nodes
func TestNewList(t *testing.T) {
	list := NewList[int]()
	if list.front.next != list.back {
		t.Fatal("Front is not set to back")
	}

	if list.back.prev != list.front {
		t.Fatal("back prev is not set to front")
	}
}

// // TestNewConstIterator tests the NewConstIterator function returns an Const iterator
// func TestNewConstIterator(t *testing.T) {
// iterator := NewConstIterator[int](nil, nil)
// if iterator.MyList != nil {
// t.Error("Expected iterator.MyList to be nil")
// }

// if iterator.Current != nil {
// t.Error("Expected iterator.Current to be nil")
// }

// curr := &node[int]{data: 10, next: nil, prev: nil}
// iterator = NewConstIterator(curr, nil)
// }
