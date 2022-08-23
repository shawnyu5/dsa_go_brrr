package linkedlist

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linked list", func() {
	data := [5]int{1, 2, 3, 4, 5}

	It("should create a new list with sentinel nodes", func() {
		list := NewList[int]()
		// front next should point to the back sentinel node
		Expect(list.front.next).To(Equal(list.back))
		// back prev should point to the front sentinel node
		Expect(list.back.prev).To(Equal(list.front))
	})

	It("Should insert the correct number of records", func() {
		list := NewList[int]()
		begin := list.begin()
		it := NewIterator(&begin, &list)
		for i := 0; i < 5; i++ {
			GinkgoWriter.Printf("Inserting %d\n", data[i])
			list.insert(&it, data[i])
			// Expect(it.get()).To(Equal(data[i]))
			it.increment()
		}
		Expect(list.size()).To(Equal(5))
	})

	It("Should insert the correct nodes", func() {
		list := NewList[int]()
		begin := list.begin()
		it := NewIterator(&begin, &list)
		for i := 0; i < 5; i++ {
			GinkgoWriter.Printf("Inserting %d\n", data[i])
			list.insert(&it, data[i])
			Expect(it.get()).To(Equal(data[i]))
			it.increment()
		}
	})

})
