package cache

/**
=============
LRU Algorithm
=============

LRU is short for 'Least recently used'. LRU algorithm is based on the
priciple of:
The probability of least used items in the past which being access in
the future is also not high.

So we can sort the items from left to right by last accessing time.
New node is being inserted in the end of the list. If the list size
exceed the size limit, then we can discard the node on the first position.

================
Hash Linked List
================

Hash Linked List is a structure combining Hash table with double linked list.

Access Time recently --------->  Access time early

  HEAD                                   TAIL
   |                                      |
   v                                      v
+------+     +------+     +------+     +------+
| Key1 |<--->| Key2 |<--->| Key3 |<--->| Key4 |
+------+     +------+     +------+     +------+
| Val1 |     | Val2 |     | Val3 |     | Val4 |
+------+     +------+     +------+     +------+

The LRU algorithm describe as below:

1. Search Hash linked list with given key, eg., Key2, when the key is found
in the hash table, then pull the node out from the hash linked list, and
insert it back to the head of the list, so the Hash linked list would look
like this now:

  HEAD                                 TAIL
   |                                    |
   v                                    v
+------+    +------+    +------+    +------+
| Key2 |<-->| Key1 |<-->| Key3 |<-->| Key4 |
+------+    +------+    +------+    +------+
| Val2 |    | Val1 |    | Val3 |    | Val4 |
+------+    +------+    +------+    +------+

2. Suppose the key is not in the hash linked list, we need to insert the new
key-value node to the head of the list, and check whether the size of list
after inserted excceed the limit size or not. eg., we need to insert a new
key-value pair (Key5, Val5), and suppose our list limit is equal to 4:

Hash linked list after (Key5, Val5) being inserted:

  HEAD                                            TAIL
   |                                               |
   v                                               v
+------+    +------+    +------+    +------+    +------+
| Key5 |<-->| Key2 |<-->| Key1 |<-->| Key3 |<-->| Key4 |
+------+    +------+    +------+    +------+    +------+
| Val5 |    | Val2 |    | Val1 |    | Val3 |    | Val4 |
+------+    +------+    +------+    +------+    +------+

And we need to remove the right most node from the list to keep the list size
not excceeding the list limit:

  HEAD                                 TAIL
   |                                    |
   v                                    v
+------+    +------+    +------+    +------+
| Key5 |<-->| Key2 |<-->| Key1 |<-->| Key3 |
+------+    +------+    +------+    +------+
| Val5 |    | Val2 |    | Val1 |    | Val3 |
+------+    +------+    +------+    +------+
*/

const defaultSize = 10

// node is a node struct for double link list
type node struct {
	prev  *node
	next  *node
	key   interface{}
	value interface{}
}

// LRUCache a cache implemented by LRU algorithm
type LRUCache struct {
	Size    int                   // Cache size
	hashMap map[interface{}]*node // Map to keep the key and the value node
	head    *node                 // the head of the link list
	tail    *node                 // the tail of the link list
}

// NewLRUCache create a LRUCache instance
func NewLRUCache(size int) *LRUCache {
	if size == 0 {
		size = defaultSize
	}

	return &LRUCache{
		Size:    size,
		hashMap: make(map[interface{}]*node, size),
		head:    nil,
		tail:    nil,
	}
}

// Get get value from LRUCache by key
// NOTE: key must be hashable, or it will panic
func (l *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	n := l.hashMap[key]
	if n != nil {
		// Reload the position base on LRU
		l.moveNodeToHead(n)

		return n.value, true
	}

	return nil, false
}

// Set add value into the LRUCache with given key
// NOTE: key must be hashable, or it will panic
func (l *LRUCache) Set(key interface{}, value interface{}) {
	n := l.hashMap[key]
	if n != nil {
		// key already exist
		l.moveNodeToHead(n)
		n.value = value
		return
	}

	n = &node{
		key:   key,
		value: value,
		prev:  nil,
		next:  l.head,
	}
	l.hashMap[key] = n

	if l.head != nil {
		// not a empty list
		l.head.prev = n
	}

	l.head = n

	last := l.tail
	if l.tail == nil {
		// an empty list
		l.tail = n
		return
	}

	if l.Size < len(l.hashMap) {
		// Need to drop the earliest node
		delete(l.hashMap, last.key)
		l.tail = last.prev
		last.next = nil
		last.key = nil
		last.value = nil
		last.prev = nil
	}
}

func (l *LRUCache) moveNodeToHead(n *node) {
	if n.prev == nil {
		// n is the head no need to refresh
		return
	} else if n.next == nil {
		// n is the tail
		n.prev.next = n.next
		n.next = l.head
		l.head.prev = n
		l.head = n
		l.tail = n.prev
		n.prev = nil
	} else {
		// n is in the middle of the list
		n.prev.next = n.next
		n.next.prev = n.prev
		n.prev = nil
		n.next = l.head
		l.head = n
	}
}
