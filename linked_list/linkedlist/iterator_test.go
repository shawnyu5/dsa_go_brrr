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

			it := NewIterator(&n, &list)
			// begin and end should point to the same sentinel node
			Expect(it.begin()).To(Equal(it.end()))
		})

		It("should return the begining of the list - const iterator", func() {
			n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
			// construct a new list with sentinel nodes
			list := NewList[int]()

			it := NewConstIterator(&n, &list)
			// begin and end should point to the same sentinel node
			Expect(it.cbegin()).To(Equal(it.cend()))
		})
	})
})

// // TestIteratorIncreament tests increment function of the const_iterator increments the iterator to the next node
// func TestConstIteratorIncreament(t *testing.T) {
// n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
// list := LinkedList[int]{front: &node[int]{data: 10}, back: &node[int]{data: 20}}

// it := NewConstIterator(&n, &list)

// if it.get() != 10 {
// t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 10, it.get())
// }
// it.increment()
// if it.get() != 20 {
// t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 20, it.get())
// }

// it.increment()
// if it.get() != 30 {
// t.Errorf("Iterator get() does not return the correct data. Expected %d, got %d", 30, it.get())
// }
// }

// // TestBegin tests begin and end function of the iterator. Should both point to the same sentinel node
// func TestBegin(t *testing.T) {
// n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
// // construct a new list with sentinel nodes
// list := NewList[int]()

// it := NewIterator(&n, &list)
// // begin and end should point to the same sentinel node
// if it.begin() != it.end() {
// t.Error("Iterator begin() does not return the correct data")
// }
// }

// // TestCbegin tests cbegin and cend function of the iterator. Should both point to the same sentinel node
// func TestCbegin(t *testing.T) {
// n := node[int]{data: 10, next: &node[int]{data: 20, next: &node[int]{data: 30}}, prev: &node[int]{data: 0}}
// // construct a new list with sentinel nodes
// list := NewList[int]()

// it := NewConstIterator(&n, &list)
// if it.cbegin() != it.cend() {
// t.Error("Iterator begin() does not return the correct data")
// }
// }
