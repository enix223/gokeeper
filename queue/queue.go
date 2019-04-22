package queue

import "errors"

var (
	// ErrQueueEmpty indicates empty queue
	ErrQueueEmpty = errors.New("Queue is empty")
	// ErrQueueFull indicates queue is full
	ErrQueueFull = errors.New("Queue is full")
	// ErrQueueSizeInvalid indicates invalid queue size
	ErrQueueSizeInvalid = errors.New("Invalid queue size")
)

// Queue queue interface
type Queue interface {
	IsEmpty() bool
	IsFull() bool
	MakeEmpty()
	Enqueue(elem interface{}) error
	Dequeue() error
	Front() (interface{}, error)
	FrontAndDequeue() (interface{}, error)
}
