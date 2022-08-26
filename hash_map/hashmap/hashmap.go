package hashmap

import (
	"crypto/sha256"
	"hash"
)

// Hash takes a string and returns a hash object from it
func Hash(key string) hash.Hash {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hash
}
