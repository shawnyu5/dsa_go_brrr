package hashmap_test

import (
	"github.com/shawnyu5/hash_map/hashmap"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LPTable", func() {
	Context("Hash()", func() {
		It("should return a hash object from a string", func() {
			hash := hashmap.Hash("hello")
			Expect(hash).ToNot(BeNil())
		})
	})

	Context("new table", func() {
		It("can create a new table", func() {
			table := hashmap.NewLPTable[int](10)
			Expect(table).ToNot(BeNil())
			Expect(table.Capacity).To(Equal(10))
			Expect(table.NumRecords).To(Equal(0))
		})
	})
})
