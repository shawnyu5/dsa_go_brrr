package hashmap_test

import (
	"bufio"
	"bytes"
	"os"

	"github.com/shawnyu5/hash_map/hashmap"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// createData reads data.txt and returns a slice of strings to act as keys for testing
func createData() []string {
	data, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}

	bytesReader := bytes.NewReader(data)
	bufReader := bufio.NewReader(bytesReader)
	arr := []string{}

	for {
		value1, _, _ := bufReader.ReadLine()
		if value1 == nil {
			break
		}
		arr = append(arr, string(value1))
	}
	return arr
}

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

			// for _, v := range table.records {
			// }
		})
	})

	Context("Update()", func() {
		It("should insert a single record", func() {
			table := hashmap.NewLPTable[int](10)
			Expect(table.Capacity).To(Equal(10))
			table.Update("hello", 1)
			Expect(table.NumRecords).To(Equal(1))
		})

		It("should insert a bunch of records", func() {
			keys := createData()
			table := hashmap.NewLPTable[int](10)
			Expect(table.IsEmpty()).To(BeTrue()) // table should be empty to begin with
			Expect(table.Capacity).To(Equal(10))

			for i := 0; i < 10; i++ {
				Expect(table.Update(keys[i], i)).To(BeTrue())
				Expect(table.NumRecords).To(Equal(i + 1))
			}
			Expect(table.IsEmpty()).To(BeFalse())
		})
	})
})
