package hashtableWithBuckets

import (
	"hashtables/hashtable"
	"math"
)

type mapNode[K comparable, V any] struct {
	key     K
	value   V
	deleted bool
}

type Hashtable[K comparable, V any] struct {
	mapNodes   []*mapNode[K, V]
	size       int
	emptyValue V
}

const LOAD_COEFF = .5
const INITIAL_NODES_SIZE = 11

func NewHashtable[K comparable, V any]() *Hashtable[K, V] {
	table := &Hashtable[K, V]{}
	table.mapNodes = make([]*mapNode[K, V], INITIAL_NODES_SIZE)
	return table
}

func (table *Hashtable[K, V]) Put(key K, value V) {
	hashCode := hashtable.GetHashCode(key)

	if table.isReadyToRehash() {
		table.rehash()
	}

	idx := hashCode % len(table.mapNodes)
	deletedIdx := -1
	for ; table.mapNodes[idx] != nil; idx = (idx + 1) % len(table.mapNodes) {
		if table.mapNodes[idx].key == key {
			table.mapNodes[idx].value = value
			return
		} else if table.mapNodes[idx].deleted && deletedIdx != -1 {
			deletedIdx = idx
		}
	}

	idxToWrite := idx
	if deletedIdx != -1 {
		idxToWrite = deletedIdx
	}
	table.mapNodes[idxToWrite] = &mapNode[K, V]{key, value, false}
	table.size++
}

func (table *Hashtable[K, V]) Get(key K) (V, bool) {
	hashCode := hashtable.GetHashCode(key)

	idx := hashCode % len(table.mapNodes)
	for ; table.mapNodes[idx] != nil; idx = (idx + 1) % len(table.mapNodes) {
		if !table.mapNodes[idx].deleted && table.mapNodes[idx].key == key {
			return table.mapNodes[idx].value, true
		}
	}

	return table.emptyValue, false
}

func (table *Hashtable[K, V]) Remove(key K) {
	hashCode := hashtable.GetHashCode(key)

	idx := hashCode % len(table.mapNodes)
	for ; table.mapNodes[idx] != nil; idx = (idx + 1) % len(table.mapNodes) {
		if !table.mapNodes[idx].deleted && table.mapNodes[idx].key == key {
			table.mapNodes[idx].deleted = true
			table.size--
			return
		}
	}
}

func (table *Hashtable[K, V]) Size() int {
	return table.size
}

func (table *Hashtable[K, V]) rehash() {
	newMapNodes := make([]*mapNode[K, V], len(table.mapNodes)*2+1)
	for _, oldNode := range table.mapNodes {
		if oldNode != nil {
			hashCode := hashtable.GetHashCode(oldNode.key)
			newIdx := hashCode % len(newMapNodes)
			for ; newMapNodes[newIdx] != nil; newIdx = (newIdx + 1) % len(newMapNodes) {
			}
			newMapNodes[newIdx] = &mapNode[K, V]{oldNode.key, oldNode.value, false}
		}
	}
	table.mapNodes = newMapNodes
}

func (table *Hashtable[K, V]) isReadyToRehash() bool {
	threshhold := math.Floor(float64(len(table.mapNodes)) * LOAD_COEFF)
	return table.size > int(threshhold)
}
