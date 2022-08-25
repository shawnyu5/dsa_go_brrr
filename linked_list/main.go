package main

import (
	"github.com/shawnyu5/linked_list/linkedlist"
)

func main() {
	list := linkedlist.NewList[string]()
	it := list.It.Begin()
	list.Insert(&it, "hello")
	it.Increment()
	list.Insert(&it, "world")
	it.Increment()
	list.Insert(&it, "Apple")
}
