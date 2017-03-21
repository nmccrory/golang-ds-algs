package golang

import (
	"fmt"
	"hash/fnv"
)

type Entry struct {
	key string
	val interface{}
}

type Bucket struct {
	entries []Entry
}

type HashTable struct {
	size int
	entries int
	buckets []Bucket
}

func Hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (h *HashTable) PrintHash(key string) {
	e := Hash(key)
	fmt.Printf("Hash value: %d \n", e)
}