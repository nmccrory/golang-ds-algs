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