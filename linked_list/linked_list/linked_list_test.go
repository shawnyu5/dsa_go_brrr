package linkedlist

import "testing"

// TestNewList tests the NewList function returns a linked list with sentinel nodes
func TestNewList(t *testing.T) {
	list := NewList[int]()
	if list.front == nil {
		t.Error("Expected list.front to be a node")
	}
	if list.back == nil {
		t.Error("Expected list.back to be a node")
	}
}

// TestNewConstIterator tests the NewConstIterator function returns an empty Const_iterator
func TestNewConstIterator(t *testing.T) {
	iterator := NewConstIterator[int]()
	if iterator.MyList != nil {
		t.Error("Expected iterator.MyList to be nil")
	}

	if iterator.Current != nil {
		t.Error("Expected iterator.Current to be nil")
	}
}
