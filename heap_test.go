package collection_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/collection"
)

func ExampleHeap() {
	var h collection.Heap[int, string]
	h.Push(42, "foo")
	h.Push(17, "bar")
	key, value, _ := h.Pop()
	fmt.Println(key, value)
	// Output: 17 bar
}

func TestHeap(t *testing.T) {
	var h collection.Heap[int, string]
	h.Push(3, "banana")
	h.Push(2, "apple")
	h.Push(4, "pear")
	{
		key, value, ok := h.Pop()
		require.True(t, ok)
		require.Equal(t, 2, key)
		require.Equal(t, "apple", value)
	}
	{
		key, value, ok := h.Pop()
		require.True(t, ok)
		require.Equal(t, 3, key)
		require.Equal(t, "banana", value)
	}

	{
		key, value, ok := h.Pop()
		require.True(t, ok)
		require.Equal(t, 4, key)
		require.Equal(t, "pear", value)
	}
	{
		_, _, ok := h.Pop()
		require.False(t, ok)
	}
}

func BenchmarkHeapMaxPush(b *testing.B) {
	b.ReportAllocs()
	var h collection.Heap[int, string]
	for i := 0; i < b.N; i++ {
		h.Push(i, "x")
	}
	for {
		_, _, ok := h.Pop()
		if !ok {
			break
		}
	}
}

func BenchmarkHeapMinPush(b *testing.B) {
	b.ReportAllocs()
	var h collection.Heap[int, string]
	for i := 0; i < b.N; i++ {
		h.Push(b.N-i, "x")
	}
	for {
		_, _, ok := h.Pop()
		if !ok {
			break
		}
	}
}
