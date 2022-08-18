package linkedlist

import "testing"

func TestNewList(t *testing.T) {
	list := LinkedList[int]{}
	list = list.NewList()
	if list.front == nil {
		t.Error("Expected list.front to be a node")
	}
	if list.back == nil {
		t.Error("Expected list.back to be a node")
	}
}
