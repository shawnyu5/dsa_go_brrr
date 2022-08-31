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
		})
	})
})
