package hashtableWithBuckets

import (
	"hash/fnv"
	"math"
)

type bucketNode struct {
	key, value string
	next       *bucketNode
}

type Hashtable struct {
	buckets []*bucketNode
	size    int
}

const LOAD_COEFF = .75
const INITIAL_BUCKETS = 11
const SUPPOSED_BUCKET_LENGTH = 8

func NewHashtable() *Hashtable {
	table := &Hashtable{}
	table.buckets = make([]*bucketNode, INITIAL_BUCKETS)
	return table
}

func (table *Hashtable) Put(key, value string) {
	hashCode := getHashCode(key)

	if table.isReadyToRehash() {
		table.rehash()
	}

	idx := hashCode % len(table.buckets)
	node := table.buckets[idx].getNode(key)
	if node != nil {
		node.value = value
		return
	}

	newNode := &bucketNode{key, value, table.buckets[idx]}
	table.buckets[idx] = newNode
	table.size++
}

func (table *Hashtable) Get(key string) (string, bool) {
	hashCode := getHashCode(key)

	idx := hashCode % len(table.buckets)
	node := table.buckets[idx].getNode(key)
	if node != nil {
		return node.value, true
	}

	return "", false
}

func (table *Hashtable) Remove(key string) {
	hashCode := getHashCode(key)

	idx := hashCode % len(table.buckets)
	if table.buckets[idx] == nil {
		return
	}

	if table.buckets[idx].key == key {
		table.buckets[idx] = table.buckets[idx].next
		table.size--
		return
	}

	for prevNode, node := table.buckets[idx], table.buckets[idx].next; node != nil; prevNode, node = node, node.next {
		if node.key == key {
			prevNode.next = node.next
			table.size--
			return
		}
	}
}

func (head *bucketNode) getNode(key string) *bucketNode {
	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.next {
		if node.key == key {
			return node
		}
	}

	return nil
}

func (table *Hashtable) Size() int {
	return table.size
}

func (table *Hashtable) rehash() {
	newBuckets := make([]*bucketNode, len(table.buckets)*2+1)
	for _, oldBucket := range table.buckets {
		for oldNode := oldBucket; oldNode != nil; oldNode = oldNode.next {
			hashCode := getHashCode(oldNode.key)
			newIdx := hashCode % len(newBuckets)
			newBuckets[newIdx] = &bucketNode{oldNode.key, oldNode.value, newBuckets[newIdx]}
		}
	}
	table.buckets = newBuckets
}

func (table *Hashtable) isReadyToRehash() bool {
	threshhold := math.Floor(float64(len(table.buckets)*SUPPOSED_BUCKET_LENGTH) * LOAD_COEFF)
	return table.size > int(threshhold)
}

func getHashCode(value string) int {
	h := fnv.New32a()
	h.Write([]byte(value))
	return int(h.Sum32())
}
