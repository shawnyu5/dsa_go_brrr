package linkedlist_test

import (
	"github.com/shawnyu5/linked_list/linkedlist"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Iterator", func() {
	It("get() should return the data within the iterator", func() {
		// n := node[int]{data: 10}
		list := linkedlist.NewList[int]()
		it := list.It.Begin()
		list.Insert(&it, 10)
		data := it.Get()

		Expect(data).To(Equal(10))
	})

	It("should increment the iterator to the next node", func() {
		list := linkedlist.NewList[int]()
		it := list.It.Begin()
		list.Insert(&it, 10)
		it.Increment()
		list.Insert(&it, 20)
		it.Increment()
		list.Insert(&it, 30)

		it = list.It.Begin()
		Expect(it.Get()).To(Equal(10))
		it.Increment()
		Expect(it.Get()).To(Equal(20))
		it.Increment()
		Expect(it.Get()).To(Equal(30))
	})

	Context("(c)begin()", func() {
		It("Itertor begin should return the begining of list", func() {
			// construct a new list with sentinel nodes
			list := linkedlist.NewList[int]()

			it := list.It.Begin()
			// begin and end should point to the same sentinel node
			Expect(it.Begin()).To(Equal(it.End()))
		})

		It("cbegin() should return the begining of list", func() {
			// construct a new list with sentinel nodes
			list := linkedlist.NewList[int]()

			cIt := list.CIt.CBegin()
			// begin and end should point to the same sentinel node
			Expect(cIt.CBegin()).To(Equal(cIt.CEnd()))
		})
	})
})
