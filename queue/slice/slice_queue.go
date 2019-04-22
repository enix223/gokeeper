package slice

import "github.com/enix223/gokeeper/queue"

// QueueImpl queue implementation with slice
type QueueImpl struct {
	elements []interface{}
	head     uint
	tail     uint
	size     int
}

// NewQueue create a queue with given size
// If size is zero, then ErrQueueInvalidSize is return
func NewQueue(size int) (queue.Queue, error) {
	if size <= 0 {
		return nil, queue.ErrQueueSizeInvalid
	}

	q := new(QueueImpl)
	q.head = 0
	q.tail = 0
	q.size = 0
	q.elements = make([]interface{}, size, size)
	return q, nil
}

// IsEmpty check queue is empty or not
func (q *QueueImpl) IsEmpty() bool {
	return q.size == 0
}

// IsFull check queue is full or not
func (q *QueueImpl) IsFull() bool {
	return q.size == cap(q.elements)
}

// MakeEmpty clear the queue
func (q *QueueImpl) MakeEmpty() {
	q.tail = 0
	q.head = 0
	q.size = 0
}

// Enqueue enqueue an element into the back of the queue
func (q *QueueImpl) Enqueue(elem interface{}) error {
	if q.IsFull() {
		return queue.ErrQueueFull
	}

	q.elements[q.tail] = elem
	q.size++
	q.tail = (q.tail + 1) % uint(cap(q.elements))

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

	return q.elements[q.head], nil
}

// FrontAndDequeue dequeue the element at the front of the queue, and return it if exist
// If queue is empty, then ErrQueueEmpty is returned
func (q *QueueImpl) FrontAndDequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, queue.ErrQueueEmpty
	}

	val := q.elements[q.head]
	q.head = (q.head + 1) % uint(cap(q.elements))
	q.size--

	return val, nil
}
