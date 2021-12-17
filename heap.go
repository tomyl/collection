package collection

import (
	"constraints"
)

// Heap is a generic binary min heap. The implementation is inspired by the
// standard container/heap package.
type Heap[K constraints.Ordered, V any] struct {
	items Slice[item[K, V]]
}

// Len returns the number of elements.
func (h *Heap[K, V]) Len() int {
	return h.items.Len()
}

// Push inserts an element into the heap.
func (h *Heap[K, V]) Push(key K, value V) {
	h.items.PushBack(item[K, V]{key, value})
	h.up(h.items.Len() - 1)
}

// Pop returns the element with smallest key or false if the heap is empty.
func (h *Heap[K, V]) Pop() (key K, value V, ok bool) {
	n := h.items.Len() - 1
	if n < 0 {
		return key, value, false
	}
	h.items.Swap(0, n)
	h.down(0, n)
	item, ok := h.items.PopBack()
	return item.key, item.value, ok
}

func (h *Heap[K, V]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.items.Swap(i, j)
		j = i
	}
}

func (h *Heap[K, V]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.items.Swap(i, j)
		i = j
	}
	return i > i0
}

func (h Heap[K, V]) less(i, j int) bool {
	return h.items[i].key < h.items[j].key
}

type item[K any, V any] struct {
	key   K
	value V
}
