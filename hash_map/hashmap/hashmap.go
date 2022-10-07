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
	// hash function
	Hash func(string) int
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

// hash takes a string and returns a hash object from it
func hash(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return int(hash.Sum32())
}

// NewLPTable constructs a new lpTable with the given capacity
func NewLPTable[T comparable](capacity int) lpTable[T] {
	table := lpTable[T]{}
	table.records = make([]record[T], capacity)
	table.Hash = hash

	// set all records to empty
	for i := 0; i < capacity; i++ {
		table.records[i].isEmpty = true
	}
	table.Capacity = capacity
	return table
}

// Update updates a record in the table with the given key and value. Returns true if update is successful, false other wise
func (lp *lpTable[T]) Update(key string, value T) bool {
	hash := lp.Hash(key) % lp.Capacity
	r := record[T]{key: key, data: value}

	startingHash := hash

	// if the current index is not empty, and the key does not match, move to the next hash
	for !lp.records[hash].isEmpty &&
		lp.records[hash].key != key {
		hash = (hash + 1) % lp.Capacity // treat the table as a circular array, when we reach the end, go back to the beginning
		if hash == startingHash {
			break
		}
	}

	if lp.NumRecords == lp.Capacity {
		return false
	}

	// if the record is empty, then we are inserting
	if lp.records[hash].isEmpty {
		lp.records[hash] = r
		lp.NumRecords++
	} else {
		// otherwise, we are updating, NumRecords does not increase
		lp.records[hash] = r
	}

	return true
}

// Find finds a record in the table with the given key, and assign the value to value passed in.
// Returns true and the record it self if the record is found, false and default value of T otherwise
func (lp *lpTable[T]) Find(key string) (bool, T) {
	startingHash := lp.Hash(key) % lp.Capacity

	// if current index contains the key we are looking for, return
	if lp.records[startingHash].key == key {
		return true, lp.records[startingHash].data
	}

	// start at the next hash over, while keeping track of the old hash
	hash := (startingHash + 1) % lp.Capacity
	// keep looking until we find the key we are looking for. If we've made a full circle, then we havent found anything
	for hash != lp.Capacity && key != lp.records[hash].key && hash != startingHash {
		hash = (hash + 1) % lp.Capacity // treat the table as a circular array, when we reach the end, go back to the beginning
	}

	// if the record we found is empty, or key does not match, then return false
	// the table could make full circle without finding the correct data member, hence the second condition is needed
	if lp.records[hash].isEmpty || lp.records[hash].key != key {
		var empty T
		return false, empty
	}

	return true, lp.records[hash].data
}

// Remove removes a record from the table with the given key. Returns true if the record is found and removed, false otherwise
func (lp *lpTable[T]) Remove(key string) bool {
	hash := lp.Hash(key) % lp.Capacity
	// keep track of the starting hash before modification
	oldHash := hash
	// emptyRecord empties out a record
	emptyRecord := func(r *record[T]) {
		r.isEmpty = true
		lp.NumRecords--
		r.key = ""
		var empty T
		r.data = empty
	}

	// if the current index is the one containing the key, then remove it
	if lp.records[hash].key == key {
		emptyRecord(&lp.records[hash])
	} else {
		// we look for the key
		startingHash := hash
		// set the current hash to the hash after the starting. Since we know the starting one is not the one we are looking for.
		hash := (startingHash + 1) % lp.Capacity
		// keep looking until we find the key we are looking for. If we've made a full circle, then we havent found anything
		for key != lp.records[hash].key && hash != startingHash {
			hash = (hash + 1) % lp.Capacity // treat the table as a circular array, when we reach the end, go back to the beginning
		}
		// if the record we found is empty, or is not the key we are looking for, then return false
		if lp.records[hash].isEmpty || lp.records[hash].key != key {
			return false
		}
		// otherwise, remove the record
		emptyRecord(&lp.records[hash])
	}
	// shift records that needs shifting over
	// skip the empty record we just removed
	for correctHash := oldHash + 1; correctHash < lp.Capacity; correctHash++ {
		currentHash := lp.Hash(lp.records[correctHash].key)
		// as long as the current index hash doesnt match the index its suppose to be at, then we shift it forward one index
		if currentHash != correctHash {
			lp.records[correctHash-1] = lp.records[correctHash]
			emptyRecord(&lp.records[correctHash]) // and empty out the current index, that has been shifted over
		}
	}
	return true
}
