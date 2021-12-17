package collection

// Slice is generic slice implementation.
type Slice[T any] []T

// Len returns the number of elements.
func (s *Slice[T]) Len() int {
	return len(*s)
}

// Clear removes all elements.
func (s *Slice[T]) Clear() {
	*s = (*s)[:0]
}

// Front returns the first element or false if the queue is empty.
func (s Slice[T]) Front() (value T, ok bool) {
	if len(s) == 0 {
		return value, false
	}
	return s[0], true
}

// Back returns the last element or false if the queue is empty.
func (s Slice[T]) Back() (value T, ok bool) {
	if len(s) == 0 {
		return value, false
	}
	return s[len(s)-1], true
}

// PushFront insert an element in the beginning of the slice.
func (s *Slice[T]) PushFront(v T) {
	*s = append([]T{v}, *s...)
}

// PushBack appends an element to end of the slice.
func (s *Slice[T]) PushBack(v T) {
	*s = append(*s, v)
}

// PopFront removes and returns the first element of the slice or false if the
// slice is empty.
func (s *Slice[T]) PopFront() (value T, ok bool) {
	if len(*s) == 0 {
		return value, false
	}
	value, *s = (*s)[0], (*s)[1:]
	return value, true
}

// PopBack removes and returns the last element of the slice or false if the
// slice is empty.
func (s *Slice[T]) PopBack() (value T, ok bool) {
	if len(*s) == 0 {
		return value, false
	}
	value, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return value, true
}

// Swap exchanges position of two elements.
func (s *Slice[T]) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
