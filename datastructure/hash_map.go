package datastructure

import "fmt"

// ArraySize size of in array in has table
const ArraySize = 7

// HashTable structure
type HashTable struct {
	Array [ArraySize]*bucket
}

// bucket strucuture
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.Array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.Array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.Array[index].delete(key)
}

// insert
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(key, "already exists")
	}
}

// search
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next

	}
}

// Init
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.Array {
		result.Array[i] = &bucket{}
	}

	return result
}

// hash sum of each charactor
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}
