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

})
