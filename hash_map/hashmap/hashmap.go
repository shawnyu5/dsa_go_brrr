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
	// updates a record in the table with the given key and data
	Update func(string, T) bool
	// removes a record from the table with the given key
	Remove func(string) bool
	// find a record in the table with the given key
	Find func(string) bool
	// number of records in the table
	NumRecords int
	// weather the table is empty
	IsEmpty func() bool
	// maximum number of records allowed in the table
	Capacity int
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
