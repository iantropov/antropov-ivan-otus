package priorityQueue

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[string]()

	testEnqueue(pq, t)
	testDequeue(pq, t)
}

func testEnqueue(pq *PriorityQueue[string], t *testing.T) {
	error := pq.Enqueue("very", 7)
	if error != nil {
		t.Errorf("Failed to enqueue - %v", error)
	}
	error = pq.Enqueue("is", 8)
	if error != nil {
		t.Errorf("Failed to enqueue - %v", error)
	}
	error = pq.Enqueue("interesting!", 3)
	if error != nil {
		t.Errorf("Failed to enqueue - %v", error)
	}
	error = pq.Enqueue("This", 10)
	if error != nil {
		t.Errorf("Failed to enqueue - %v", error)
	}
	error = pq.Enqueue("test", 8)
	if error != nil {
		t.Errorf("Failed to enqueue - %v", error)
	}
}

func testDequeue(pq *PriorityQueue[string], t *testing.T) {
	value, priority, error := pq.Dequeue()
	if value != "interesting!" {
		t.Errorf("Invalid value - %v", value)
	}
	if priority != 3 {
		t.Errorf("Invalid priority - %v", priority)
	}
	if error != nil {
		t.Errorf("Failed to dequeue - %v", error)
	}

	value, priority, error = pq.Dequeue()
	if value != "very" {
		t.Errorf("Invalid value - %v", value)
	}
	if priority != 7 {
		t.Errorf("Invalid priority - %v", priority)
	}
	if error != nil {
		t.Errorf("Failed to dequeue - %v", error)
	}

	value, priority, error = pq.Dequeue()
	if value != "test" {
		t.Errorf("Invalid value - %v", value)
	}
	if priority != 8 {
		t.Errorf("Invalid priority - %v", priority)
	}
	if error != nil {
		t.Errorf("Failed to dequeue - %v", error)
	}

	value, priority, error = pq.Dequeue()
	if value != "is" {
		t.Errorf("Invalid value - %v", value)
	}
	if priority != 8 {
		t.Errorf("Invalid priority - %v", priority)
	}
	if error != nil {
		t.Errorf("Failed to dequeue - %v", error)
	}

	value, priority, error = pq.Dequeue()
	if value != "This" {
		t.Errorf("Invalid value - %v", value)
	}
	if priority != 10 {
		t.Errorf("Invalid priority - %v", priority)
	}
	if error != nil {
		t.Errorf("Failed to dequeue - %v", error)
	}
}
