package collection

// Deque is a generic double-ended queue based on a double-linked list. The
// zero value is can be used without any further initialization. The
// implementation is heavily inspired by the standard container/list package.
type Deque[T any] struct {
	root element[T]
	len  int
}

// Len returns the number of elements.
func (d *Deque[T]) Len() int {
	return d.len
}

// Clear removes all elements.
func (d *Deque[T]) Clear() {
	d.root.next = &d.root
	d.root.prev = &d.root
	d.len = 0
}

// Front returns the first element or false if the queue is empty.
func (d *Deque[T]) Front() (value T, ok bool) {
	if d.len == 0 {
		return value, false
	}
	return d.root.next.value, true
}

// Back returns the last element or false if the queue is empty.
func (d *Deque[T]) Back() (value T, ok bool) {
	if d.len == 0 {
		return value, false
	}
	return d.root.prev.value, true
}

// PushFront insert an element in the beginning of the queue.
func (d *Deque[T]) PushFront(v T) {
	d.lazyInit()
	d.insertValue(v, &d.root)
}

// PushBack appends an element to end of the queue.
func (d *Deque[T]) PushBack(v T) {
	d.lazyInit()
	d.insertValue(v, d.root.prev)
}

// PopFront removes and returns the first element in the queue or false if the
// queue is empty.
func (d *Deque[T]) PopFront() (value T, ok bool) {
	if d.len == 0 {
		return value, false
	}
	value = d.root.next.value
	d.remove(d.root.next)
	return value, true
}

// PopBack removes and returns the last element in the queue or false if the
// queue is empty.
func (d *Deque[T]) PopBack() (value T, ok bool) {
	if d.len == 0 {
		return value, false
	}
	value = d.root.prev.value
	d.remove(d.root.prev)
	return value, true
}

func (d *Deque[T]) lazyInit() {
	if d.root.next == nil {
		d.Clear()
	}
}

func (d *Deque[T]) insert(e, at *element[T]) {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	d.len++
}

func (d *Deque[T]) insertValue(v T, at *element[T]) {
	d.insert(&element[T]{value: v}, at)
}

func (d *Deque[T]) remove(e *element[T]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	d.len--
}

type element[T any] struct {
	prev  *element[T]
	next  *element[T]
	value T
}
