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
	probe      func(int, int, int) int
}

const LOAD_COEFF = .5
const INITIAL_NODES_SIZE = 11

func NewHashtable[K comparable, V any]() *Hashtable[K, V] {
	table := &Hashtable[K, V]{}
	table.mapNodes = make([]*mapNode[K, V], INITIAL_NODES_SIZE)
	table.probe = linearProbe
	return table
}

func NewHashtableWithQuadraticProbe[K comparable, V any]() *Hashtable[K, V] {
	table := &Hashtable[K, V]{}
	table.mapNodes = make([]*mapNode[K, V], INITIAL_NODES_SIZE)
	table.probe = quadraticProbe
	return table
}

func (table *Hashtable[K, V]) Put(key K, value V) {
	hashCode := hashtable.GetHashCode(key)

	if table.isReadyToRehash() {
		table.rehash()
	}

	var idx int
	deletedIdx := -1
	for i := 0; ; i++ {
		idx = table.probe(hashCode, i, len(table.mapNodes))
		if table.mapNodes[idx] == nil {
			break
		}

		if table.mapNodes[idx].key == key && !table.mapNodes[idx].deleted {
			table.mapNodes[idx].value = value
			return
		} else if table.mapNodes[idx].key == key {
			table.mapNodes[idx].value = value
			table.mapNodes[idx].deleted = false
			table.size++
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
	deletedIdx := -1
	var idx int
	for i := 0; ; i++ {
		idx = table.probe(hashCode, i, len(table.mapNodes))
		if table.mapNodes[idx] == nil {
			break
		}

		if table.mapNodes[idx].key == key && !table.mapNodes[idx].deleted {
			break
		} else if table.mapNodes[idx].deleted && deletedIdx != -1 {
			deletedIdx = idx
		}
	}

	if table.mapNodes[idx] == nil {
		return table.emptyValue, false
	}

	if deletedIdx != -1 {
		table.mapNodes[idx], table.mapNodes[deletedIdx] = table.mapNodes[deletedIdx], table.mapNodes[idx]
		idx, deletedIdx = deletedIdx, idx
	}

	return table.mapNodes[idx].value, true
}

func (table *Hashtable[K, V]) Remove(key K) {
	hashCode := hashtable.GetHashCode(key)
	idx := table.findIdxForKey(key, hashCode)
	if idx != -1 {
		table.mapNodes[idx].deleted = true
		table.size--
	}
}

func (table *Hashtable[K, V]) Size() int {
	return table.size
}

func (table *Hashtable[K, V]) rehash() {
	newMapNodes := make([]*mapNode[K, V], len(table.mapNodes)*2+1)
	for _, oldNode := range table.mapNodes {
		if oldNode != nil {
			newHashCode := hashtable.GetHashCode(oldNode.key)
			newIdx := table.findNewIdx(newHashCode, newMapNodes)
			newMapNodes[newIdx] = &mapNode[K, V]{oldNode.key, oldNode.value, false}
		}
	}
	table.mapNodes = newMapNodes
}

func (table *Hashtable[K, V]) findIdxForKey(key K, hashCode int) int {
	for i := 0; ; i++ {
		idx := table.probe(hashCode, i, len(table.mapNodes))
		if table.mapNodes[idx] == nil {
			return -1
		}
		if !table.mapNodes[idx].deleted && table.mapNodes[idx].key == key {
			return idx
		}
	}
}

func (table *Hashtable[K, V]) findNewIdx(hashCode int, mapNodes []*mapNode[K, V]) int {
	for i := 0; ; i++ {
		idx := table.probe(hashCode, i, len(mapNodes))
		if mapNodes[idx] == nil {
			return idx
		}
	}
}

func (table *Hashtable[K, V]) isReadyToRehash() bool {
	threshhold := math.Floor(float64(len(table.mapNodes)) * LOAD_COEFF)
	return table.size > int(threshhold)
}

func linearProbe(hashCode, idx, mod int) int {
	return (hashCode + idx) % mod
}

func quadraticProbe(hashCode, idx, mod int) int {
	return (hashCode + idx + idx*idx) % mod
}
