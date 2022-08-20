package linkedlist

import "testing"

// makeData creates and returns an array of data for testing
func makeData() [5]int {
	data := make([]int, 5)
	expected := [5]int{0, 1, 2, 3, 4}

	for i := 0; i < 5; i++ {
		data[i] = expected[i]
	}

	return expected
}

// TestInsert test the linked list interface
func TestInsert(t *testing.T) {
	data := makeData()
	list := NewList[int]()
	n := &node[int]{data: 20}
	it := NewIterator(n, &list)
	for i := 0; i < 5; i++ {
		list.insert(&it, data[i])
	}
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
