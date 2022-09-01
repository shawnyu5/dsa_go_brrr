package hashmap_test

import (
	"bufio"
	"bytes"
	"fmt"
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
			table := hashmap.NewLPTable[int](10)
			hash := table.Hash("hello")
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

	Context("Update()", func() {
		It("should insert a single record", func() {
			table := hashmap.NewLPTable[int](10)
			Expect(table.Capacity).To(Equal(10))
			table.Update("hello", 1)
			Expect(table.NumRecords).To(Equal(1))
		})

		// TODO: idk how to cause a collision
		It("Should handle linear collisions correctly", func() {
			// collisions that do not wrap around the table
			table := hashmap.NewLPTable[int](2)
			table.Hash = func(key string) int {
				return 0
			}

			inserted := table.Update("apple", 1)
			Expect(inserted).To(BeTrue())
			inserted = table.Update("orange", 2)
			Expect(inserted).To(BeTrue())

			_, value := table.Find("apple")
			Expect(value).To(Equal(1))

			_, value = table.Find("orange")
			Expect(value).To(Equal(2))
			fmt.Println(fmt.Sprintf(" table: %+v", table)) // __AUTO_GENERATED_PRINT_VAR__
		})

		It("should update a value without increasing the number of records", func() {
			table := hashmap.NewLPTable[int](10)
			table.Update("hello", 1)
			table.Update("hello", 1)
			Expect(table.NumRecords).To(Equal(1))
		})

		// TODO: test this later
		// It("should handle circular collisions correctly", func() {
		// // collisions that wrap around the table
		// table := hashmap.NewLPTable[int](3)
		// table.Update("apple", 1)
		// table.Update("orange", 2)
		// table.Update("banana", 3)

		// })

		It("should insert a bunch of records", func() {
			// keys := createData()
			keys := []string{"apple", "orange", "banana", "grape", "kiwi", "pear", "pineapple", "strawberry", "watermelon", "blueberry"}
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

	Context("Find()", func() {
		It("should find a record", func() {
			keys := createData()
			table := hashmap.NewLPTable[int](10)
			for i := 0; i < 10; i++ {
				Expect(table.Update(keys[i], i)).To(BeTrue())
			}
			found, _ := table.Find(keys[0])
			Expect(found).To(BeTrue())
		})

		It("should find a record with a collision", func() {
			table := hashmap.NewLPTable[int](10)
			table.Update("apple", 1)
			table.Update("orange", 2)
			// table.Update("bee", 3)
			found, value := table.Find("orange")
			Expect(found).To(BeTrue())
			Expect(value).To(Equal(2), "orange should have value 2")
		})

		It("should find a record with a lot of data", func() {
			keys := createData()
			table := hashmap.NewLPTable[int](len(keys))
			for i := 0; i < len(keys); i++ {
				Expect(table.Update(keys[i], i)).To(BeTrue())
			}

			for i := 0; i < len(keys); i++ {
				found, value := table.Find(keys[i])
				Expect(found).To(BeTrue())
				Expect(value).To(Equal(i), "Iteration: %d/%d", i, len(keys))
			}
		})
	})
})
