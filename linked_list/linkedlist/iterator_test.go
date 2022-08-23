package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Iterator", func() {
	It("get() should return the data within the iterator", func() {
		n := node[int]{data: 10}
		it := &Iterator[int]{current: &n}
		data := it.get()

		Expect(data).To(Equal(10))
	})

	It("should increment the iterator to the next node", func() {
		n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
		it := &Iterator[int]{current: &n}
		Expect(it.get()).To(Equal(10))
		it.increment()
		Expect(it.get()).To(Equal(20))
		it.increment()
		Expect(it.get()).To(Equal(30))
	})

	Context("(c)begin()", func() {
		It("should return the begining of the list - iterator", func() {
			n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
			// construct a new list with sentinel nodes
			list := NewList[int]()

			it := newIterator(&n, &list)
			// begin and end should point to the same sentinel node
			Expect(it.begin()).To(Equal(it.end()))
		})

		It("should return the begining of the list - const iterator", func() {
			n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
			// construct a new list with sentinel nodes
			list := NewList[int]()

			it := newConstIterator(&n, &list)
			// begin and end should point to the same sentinel node
			Expect(it.cbegin()).To(Equal(it.cend()))
		})
	})
})
