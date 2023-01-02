package priorityQueue

import (
	"errors"
	factorArray "kosaraju/priority-queue/factor-array"
)

type PriorityQueueItem[T any] struct {
	value T
	next  *PriorityQueueItem[T]
}

type PriorityQueueNode[T any] struct {
	length, priority int
	head             *PriorityQueueItem[T]
}

type PriorityQueue[T any] struct {
	length    int
	nodes     factorArray.FactorArray[*PriorityQueueNode[T]]
	zeroValue T
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	newPriorityQueue := new(PriorityQueue[T])
	newPriorityQueue.nodes.Initialize()
	return newPriorityQueue
}

func NewPriorityQueueNode[T any](priority int) *PriorityQueueNode[T] {
	newNode := new(PriorityQueueNode[T])
	newNode.priority = priority
	return newNode
}

func NewPriorityQueueItem[T any](value T) *PriorityQueueItem[T] {
	newItem := new(PriorityQueueItem[T])
	newItem.value = value
	return newItem
}

func (pq *PriorityQueue[T]) Enqueue(item T, priority int) error {
	added := false
	for i := 0; i < pq.nodes.Length(); i++ {
		node, error := pq.nodes.Get(i)
		if error != nil {
			return error
		}

		if priority == node.priority {
			node.append(item)
			added = true
			break
		} else if priority < node.priority {
			newPqn := NewPriorityQueueNode[T](priority)
			newPqn.append(item)
			pq.nodes.Add(newPqn, i)
			added = true
			break
		}
	}

	if !added {
		newPqn := NewPriorityQueueNode[T](priority)
		newPqn.append(item)
		pq.nodes.Push(newPqn)
	}

	pq.length++

	return nil
}

func (pq *PriorityQueue[T]) Dequeue() (T, int, error) {
	if pq.length == 0 {
		return pq.zeroValue, 0, errors.New("empty priority queue")
	}

	for i := 0; i < pq.nodes.Length(); i++ {
		node, error := pq.nodes.Get(i)
		if error != nil {
			return pq.zeroValue, 0, error
		}

		if node.length > 0 {
			res := node.pull()
			pq.length--
			return res, node.priority, nil
		}
	}

	return pq.zeroValue, 0, errors.New("failed to dequeue")
}

func (pq *PriorityQueue[T]) Length() int {
	return pq.length
}

func (pqn *PriorityQueueNode[T]) append(item T) {
	newItem := NewPriorityQueueItem(item)
	newItem.next = pqn.head
	pqn.head = newItem
	pqn.length++
}

func (pqn *PriorityQueueNode[T]) pull() T {
	res := pqn.head
	pqn.head = pqn.head.next
	pqn.length--
	return res.value
}
