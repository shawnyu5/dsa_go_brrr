package linkedlist_test

import (
	"github.com/shawnyu5/linked_list/linkedlist"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// createLinkedList creates a new linked list for testing. Containing the data 1 to 6
func createLinkedList() linkedlist.LinkedList[int] {
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
	it.Increment()
	list.Insert(&it, 6)
	return list
}

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
		list := createLinkedList()
		Expect(list.Size()).To(Equal(6))
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
		list := createLinkedList()
		// set the iterator to end of list
		it := list.It.End()

		it.Decrement()

		for i := 5; i < 5; i-- {
			Expect(it.Get()).To(Equal(i))
			it.Decrement()
		}
	})

	Context("Search", func() {
		It("Should return the node data if found", func() {
			list := createLinkedList()
			foundIt := list.Search(3)
			Expect(foundIt.Get()).To(Equal(3))
		})

		It("Did not find the node, should return iterator pointing to end of list", func() {
			list := createLinkedList()
			foundIt := list.Search(6)
			Expect(foundIt).To(Equal(list.It.End()))
		})
	})

	Context("Sort", func() {
		It("Split should split the list in half with even number of elements", func() {
			list := createLinkedList()
			begin := list.It.Begin()
			end := list.It.End()
			end.Decrement() // exculude the sentinel node, for split

			firstHalf, mid, endPos := list.Split(begin, end)

			Expect(firstHalf.Get()).To(Equal(1)) // first half should start at 1
			Expect(mid.Get()).To(Equal(4))       // middle should be 3
			Expect(endPos.Get()).To(Equal(6))    // list should end at 6

		})
		It("Split should split the list in half with odd number of elements", func() {
			list := createLinkedList()

			// second half starts at 3
			// first half should be bigger than second half
			begin := list.It.Begin()
			end := list.It.End()
			end.Decrement()
			end.Decrement() // list ends at 5

			first, mid, endPos := list.Split(begin, end)

			Expect(first.Get()).To(Equal(1))
			Expect(mid.Get()).To(Equal(3))
			Expect(endPos.Get()).To(Equal(5))
		})

		It("Should sort the list", func() {
			list := createLinkedList()
			begin := list.It.Begin()
			end := list.It.End()
			end.Decrement()
			list.Sort(&begin, &end)
		})
	})

	Context("erase", func() {
		It("Should erase an existing node", func() {
			list := createLinkedList()
			list.Erase(list.Search(3))

			// check the node no longer exists within the list
			Expect(list.Search(3)).To(Equal(list.It.End()))
			// the size of the list should also shrink
			Expect(list.Size()).To(Equal(5))
		})
	})
})
