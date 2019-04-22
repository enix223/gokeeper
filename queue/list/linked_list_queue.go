package list

import "github.com/enix223/gokeeper/queue"

type node struct {
	next  *node
	value interface{}
}

// QueueImpl queue implementation with linked list
type QueueImpl struct {
	head *node
	tail *node
	cap  int
	size int
}

// NewQueue create a queue with given size
// If size is zero, then ErrQueueInvalidSize is return
func NewQueue(size int) (queue.Queue, error) {
	if size <= 0 {
		return nil, queue.ErrQueueSizeInvalid
	}

	q := new(QueueImpl)
	q.head = nil
	q.tail = nil
	q.size = 0
	q.cap = size
	return q, nil
}

// IsEmpty check queue is empty or not
func (q *QueueImpl) IsEmpty() bool {
	return q.size == 0
}

// IsFull check queue is full or not
func (q *QueueImpl) IsFull() bool {
	return q.size == q.cap
}

// MakeEmpty clear the queue
func (q *QueueImpl) MakeEmpty() {
	q.tail = nil
	q.head = nil
	q.size = 0
}

// Enqueue enqueue an element into the back of the queue
func (q *QueueImpl) Enqueue(elem interface{}) error {
	if q.IsFull() {
		return queue.ErrQueueFull
	}

	n := new(node)
	n.next = nil
	n.value = elem

	if q.tail == nil {
		q.head = n
	} else {
		q.tail.next = n
	}

	q.tail = n
	q.size++

	return nil
}

// Dequeue dequeue an element from the queue
func (q *QueueImpl) Dequeue() error {
	_, err := q.FrontAndDequeue()
	return err
}

// Front return the element at the front of the queue
// If queue is empty, then ErrQeueueEmpty is returned
func (q *QueueImpl) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, queue.ErrQueueEmpty
	}

	return q.head.value, nil
}

// FrontAndDequeue dequeue the element at the front of the queue, and return it if exist
// If queue is empty, then ErrQueueEmpty is returned
func (q *QueueImpl) FrontAndDequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, queue.ErrQueueEmpty
	}

	val := q.head.value
	q.head = q.head.next
	q.size--

	return val, nil
}
