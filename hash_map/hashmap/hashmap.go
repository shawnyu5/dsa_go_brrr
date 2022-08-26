package hashmap

import (
	"crypto/sha256"
	"hash"
)

type record[T comparable] struct {
	// the unhashed key of the record
	key  string
	data T
}

// base hash table to drive other hash tables from
type hashTable[T comparable] struct {
	Update     func(string, T) bool
	Remove     func(string) bool
	Find       func(string) bool
	NumRecords int
	IsEmpty    func() bool
	Capacity   int
}

// hash table using linear probing as the collision resolution strategy
type lpTable[T comparable] struct {
	hashTable[T]
	// capability of the table
	capability int
	// number of records in the table
	numRecords int
	front      *record[T]
	back       *record[T]
}

// Hash takes a string and returns a hash object from it
func Hash(key string) hash.Hash {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hash
}
