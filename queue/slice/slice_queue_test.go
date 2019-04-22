package slice

import (
	"testing"

	"github.com/enix223/gokeeper/queue"
)

func TestQueueIsEmpty(t *testing.T) {
	q, err := NewQueue(1)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("exp true, but got false")
	}
}

func TestQueueCreate(t *testing.T) {
	q, err := NewQueue(0)
	if err != queue.ErrQueueSizeInvalid {
		t.Fatalf("exp nil, got %v", err)
	}

	q, err = NewQueue(-1)
	if err != queue.ErrQueueSizeInvalid {
		t.Fatalf("exp nil, got %v", err)
	}

	if q != nil {
		t.Fatalf("exp nil, got %v", q)
	}
}

func TestQueueIsFull(t *testing.T) {
	q, err := NewQueue(1)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	err = q.Enqueue(1)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if !q.IsFull() {
		t.Fatalf("exp true, got false")
	}
}

func TestQueueMakeEmpty(t *testing.T) {
	q, err := NewQueue(2)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if err := q.Enqueue(1); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}
	if err := q.Enqueue(1); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	q.MakeEmpty()

	if !q.IsEmpty() {
		t.Fatalf("exp true, got false")
	}
}

func TestQueueEnqueue(t *testing.T) {
	q, err := NewQueue(2)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if err := q.Enqueue(1); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}
	if err := q.Enqueue(2); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if err := q.Enqueue(3); err != queue.ErrQueueFull {
		t.Fatalf("exp ErrQueueFull, got %v", err)
	}
}

func TestQueueFrontAndDequeue(t *testing.T) {
	q, err := NewQueue(3)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, err := q.FrontAndDequeue()
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}
	if val, ok := v.(int); !ok || val != 1 {
		t.Fatalf("exp %v, got %v", 1, val)
	}

	v, _ = q.FrontAndDequeue()
	if val, ok := v.(int); !ok || val != 2 {
		t.Fatalf("exp %v, got %v", 2, val)
	}

	v, _ = q.FrontAndDequeue()
	if val, ok := v.(int); !ok || val != 3 {
		t.Fatalf("exp %v, got %v", 3, val)
	}

	if _, err = q.FrontAndDequeue(); err != queue.ErrQueueEmpty {
		t.Fatalf("exp %v, got %v", queue.ErrQueueEmpty, err)
	}
}

func TestQueueFront(t *testing.T) {
	q, err := NewQueue(3)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	v, err := q.Front()
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if val, ok := v.(int); !ok || val != 1 {
		t.Fatalf("exp %v, got %v", 1, val)
	}

	q, _ = NewQueue(3)
	if _, err := q.Front(); err != queue.ErrQueueEmpty {
		t.Fatalf("exp %v, got %v", queue.ErrQueueEmpty, err)
	}

}

func TestQueueDequeue(t *testing.T) {
	q, err := NewQueue(3)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if err := q.Dequeue(); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if err := q.Dequeue(); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if err := q.Dequeue(); err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	if err := q.Dequeue(); err != queue.ErrQueueEmpty {
		t.Fatalf("exp %v, got %v", queue.ErrQueueEmpty, err)
	}
}

func TestQueueEnqueueDequeue(t *testing.T) {
	iterations := 1234
	q, err := NewQueue(iterations)
	if err != nil {
		t.Fatalf("exp nil, got %v", err)
	}

	for i := 0; i < iterations; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < iterations; i++ {
		v, err := q.FrontAndDequeue()
		if err != nil {
			t.Fatalf("exp %v, got %v", nil, err)
		}

		if val, ok := v.(int); !ok || val != i {
			t.Fatalf("exp %v, got %v", i, val)
		}
	}
}

func BenchmarkEnqueueDequeue(b *testing.B) {
	q, err := NewQueue(b.N)
	if err != nil {
		b.Fatalf("exp nil, got %v", err)
	}

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		q.FrontAndDequeue()
	}
}
