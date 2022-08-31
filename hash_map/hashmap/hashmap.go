package hashmap

import (
	"hash/fnv"
)

type record[T comparable] struct {
	// the unhashed key of the record
	key  string
	data T
	// keeps track of whether this record stores data
	isEmpty bool
}

// base hash table to drive other hash tables from
type hashTable[T comparable] struct {
	// // updates a record in the table with the given key and data
	// Update func(string, T) bool
	// // removes a record from the table with the given key
	// Remove func(string) bool
	// // find a record in the table with the given key
	// Find func(string) bool
	// number of records in the table
	NumRecords int
	// maximum number of records allowed in the table
	Capacity int
}

// hash table using linear probing as the collision resolution strategy
type lpTable[T comparable] struct {
	// the array of records
	records []record[T]
	hashTable[T]
	// number of records in the table
	front *record[T]
	back  *record[T]
}

// IsEmpty returns true if the table is empty, false otherwise
func (lp *hashTable[T]) IsEmpty() bool {
	return lp.NumRecords == 0
}

// Hash takes a string and returns a hash object from it
func Hash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32()
}

// NewLPTable constructs a new lpTable with the given capacity
func NewLPTable[T comparable](capacity int) lpTable[T] {
	table := lpTable[T]{}
	table.records = make([]record[T], capacity)
	// set all records to empty
	for i := 0; i < capacity; i++ {
		table.records[i].isEmpty = true
	}
	table.Capacity = capacity
	return table
}

// Update updates a record in the table with the given key and value. Returns true if update is successful, false other wise
func (lp *lpTable[T]) Update(key string, value T) bool {
	hsh := Hash(key) % uint32(lp.Capacity)
	r := record[T]{key: key, data: value}

	// if the current index is full, find the next empty index
	if !lp.records[hsh].isEmpty {
		// if the current index is not empty, and the key does not match, move to the next hash
		for hsh < uint32(lp.Capacity) &&
			!lp.records[hsh].isEmpty &&
			lp.records[hsh].key != key {
			hsh++
			hsh = (hsh + 1) % uint32(lp.Capacity) // treat the table as a circular array, when we reach the end, go back to the beginning
			// i = (1 + i) % lp.Capacity
		}
	}

	if lp.NumRecords == lp.Capacity {
		return false
	}

	lp.records[hsh] = r
	lp.NumRecords++
	return true
}

// Find finds a record in the table with the given key, and assign the value to value passed in. Returns true if the record is found, false otherwise
func (lp *lpTable[T]) Find(key string) (bool, T) {
	startingHash := Hash(key) % uint32(lp.Capacity)

	// if current index contains the key we are looking for, return
	if lp.records[startingHash].key == key {
		return true, lp.records[startingHash].data
	}

	// start at the next hash over, while keeping track of the old hash
	hash := startingHash + 1
	// keep looking until we find the key we are looking for. If we've made a full circle, then we havent found anything
	for hash != uint32(lp.Capacity) && key != lp.records[hash].key && hash != startingHash {
		hash++
		hash = (hash + 1) % uint32(lp.Capacity) // treat the table as a circular array, when we reach the end, go back to the beginning
	}

	// if the record we found is empty, then return false
	if lp.records[hash].isEmpty {
		var empty T
		return false, empty
	}

	return true, lp.records[hash].data
}
