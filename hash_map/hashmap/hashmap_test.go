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
})
