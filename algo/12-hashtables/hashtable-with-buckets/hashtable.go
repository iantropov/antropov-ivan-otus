package hashtableWithBuckets

import (
	"hashtables/hashtable"
	"math"
)

type bucketNode[K comparable, V any] struct {
	key   K
	value V
	next  *bucketNode[K, V]
}

type Hashtable[K comparable, V any] struct {
	buckets    []*bucketNode[K, V]
	size       int
	emptyValue V
}

const LOAD_COEFF = .75
const INITIAL_BUCKETS = 11
const SUPPOSED_BUCKET_LENGTH = 8

func NewHashtable[K comparable, V any]() *Hashtable[K, V] {
	table := &Hashtable[K, V]{}
	table.buckets = make([]*bucketNode[K, V], INITIAL_BUCKETS)
	return table
}

func (table *Hashtable[K, V]) Put(key K, value V) {
	hashCode := hashtable.GetHashCode(key)

	if table.isReadyToRehash() {
		table.rehash()
	}

	idx := hashCode % len(table.buckets)
	node := table.buckets[idx].getNode(key)
	if node != nil {
		node.value = value
		return
	}

	newNode := &bucketNode[K, V]{key, value, table.buckets[idx]}
	table.buckets[idx] = newNode
	table.size++
}

func (table *Hashtable[K, V]) Get(key K) (V, bool) {
	hashCode := hashtable.GetHashCode(key)

	idx := hashCode % len(table.buckets)
	node := table.buckets[idx].getNode(key)
	if node != nil {
		return node.value, true
	}

	return table.emptyValue, false
}

func (table *Hashtable[K, V]) Remove(key K) {
	hashCode := hashtable.GetHashCode(key)

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

func (head *bucketNode[K, V]) getNode(key K) *bucketNode[K, V] {
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

func (table *Hashtable[K, V]) Size() int {
	return table.size
}

func (table *Hashtable[K, V]) rehash() {
	newBuckets := make([]*bucketNode[K, V], len(table.buckets)*2+1)
	for _, oldBucket := range table.buckets {
		for oldNode := oldBucket; oldNode != nil; oldNode = oldNode.next {
			hashCode := hashtable.GetHashCode(oldNode.key)
			newIdx := hashCode % len(newBuckets)
			newBuckets[newIdx] = &bucketNode[K, V]{oldNode.key, oldNode.value, newBuckets[newIdx]}
		}
	}
	table.buckets = newBuckets
}

func (table *Hashtable[K, V]) isReadyToRehash() bool {
	threshhold := math.Floor(float64(len(table.buckets)*SUPPOSED_BUCKET_LENGTH) * LOAD_COEFF)
	return table.size > int(threshhold)
}
