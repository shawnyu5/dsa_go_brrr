package linkedlist_test

import (
	"github.com/shawnyu5/linked_list/linkedlist"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linked list", func() {
	data := [5]int{1, 2, 3, 4, 5}

	It("should create a new list with sentinel nodes", func() {
		list := linkedlist.NewList[int]()
		// front next should point to the back sentinel node
		Expect(list.Cit.CBegin()).To(Equal(list.Cit.CEnd()))
		// back prev should point to the front sentinel node
		Expect(list.Cit.CEnd()).To(Equal(list.Cit.CBegin()))

		// also check the non const iterator
		Expect(list.It.Begin()).To(Equal(list.It.End()))
		// back prev should point to the front sentinel node
		Expect(list.It.End()).To(Equal(list.It.Begin()))
	})

	It("Should insert the correct number of records", func() {
		list := linkedlist.NewList[int]()
		// begin := list.begin()
		it := list.It.Begin()
		for i := 0; i < 5; i++ {
			list.Insert(&it, data[i])
			it.Increment()
		}
		Expect(list.Size()).To(Equal(5))
	})

	It("Should insert the correct nodes", func() {
		list := linkedlist.NewList[int]()
		it := list.It.Begin()
		for i := 0; i < 5; i++ {
			GinkgoWriter.Printf("Inserting %d\n", data[i])
			list.Insert(&it, data[i])
			Expect(it.Get()).To(Equal(data[i]))
			it.Increment()
		}
	})

	It("Decrement should return the correct data", func() {
		list := linkedlist.NewList[int]()
		it := list.It.Begin()
		list.Insert(&it, 1)
		it.Increment()
		list.Insert(&it, 2)
		it.Increment()
		list.Insert(&it, 3)
		it.Increment()
		list.Insert(&it, 4)
		it.Increment()
		list.Insert(&it, 5)

		// set the iterator to end of list
		it = list.It.End()
		it.Decrement()

		for i := 5; i < 5; i-- {
			Expect(it.Get()).To(Equal(i))
			it.Decrement()
		}
	})

	Context("Search", func() {
		It("Should return the node data if found", func() {
			list := linkedlist.NewList[int]()
			it := list.It.Begin()

			list.Insert(&it, 1)
			it.Increment()
			list.Insert(&it, 2)
			it.Increment()
			list.Insert(&it, 3)
			it.Increment()
			list.Insert(&it, 4)
			it.Increment()
			list.Insert(&it, 5)

			foundIt := list.Search(3)
			Expect(foundIt.Get()).To(Equal(3))
		})

		It("Did not find the node, should return iterator pointing to end of list", func() {
			list := linkedlist.NewList[int]()
			it := list.It.Begin()

			list.Insert(&it, 1)
			it.Increment()
			list.Insert(&it, 2)
			it.Increment()
			list.Insert(&it, 3)
			it.Increment()
			list.Insert(&it, 4)
			it.Increment()
			list.Insert(&it, 5)

			foundIt := list.Search(6)
			Expect(foundIt).To(Equal(list.It.End()))

		})
	})
})
