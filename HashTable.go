package golang

import (
	"fmt"
	"hash/fnv"
)

type Entry struct {
	key string
	val interface{}
}

func NewEntry(key string, val interface{}) *Entry {
	e := Entry{key, val}
	return &e
}

type Bucket struct {
	entries []*Entry
}

func NewBucket() *Bucket {
	b := Bucket{}
	return &b
}

type HashTable struct {
	size uint32
	entries uint32
	buckets []*Bucket
}

func NewHashTable(size uint32) *HashTable {
	h := HashTable{}
	h.size = size
	h.entries = 0
	h.buckets = make([]*Bucket, size)
	for i, _ := range h.buckets {
		h.buckets[i] = NewBucket()
	}
	return &h
}

func Hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (b *Bucket) RetrieveEntry(key string) *Entry {
	// If there are no entries in the bucket, return nil
	if b.entries == nil {
		return nil
	}
	// Find entry by iterating through the bucket
	for _, entry := range b.entries {
		if(entry.key == key) {
			return entry
		}
	}
	// Entry is not in the bucket
	return nil
}

func (h *HashTable) Set(key string, val interface{}) bool {
	hash := Hash(key)
	index := hash % h.size
	if index >= 0 && index < uint32(len(h.buckets)) {
		bucket := h.buckets[index]
		entry := bucket.RetrieveEntry(key)
		// If there is no entry for that key, add it
		if entry == nil {
			bucket.entries = append(bucket.entries, NewEntry(key, val))
			h.entries++
			return true
		}else{
			// Modify existing entry
			entry.val = val
			return true
		}
	}else{
		// Present error if key is out of range
		fmt.Printf("Could not Set(key: %s): out of bucket range. ", key)
		return false
	}
}

func (h *HashTable) Get(key string) interface{} {
	hash := Hash(key)
	index := hash % h.size
	if index >= 0 && index < uint32(len(h.buckets)) {
		bucket := h.buckets[index]
		entry := bucket.RetrieveEntry(key)
		if entry != nil {
			return entry.val
		}
		return nil
	}else{
		// Index is out of range so we return nil
		return nil
	}
}
