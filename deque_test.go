package collection_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/collection"
)

func ExampleDeque() {
	var q collection.Deque[int]
	q.PushFront(42)
	value, _ := q.PopBack()
	fmt.Println(value)
	// Output: 42
}

type list[T any] interface {
	Clear()
	Len() int
	Front() (value T, ok bool)
	Back() (value T, ok bool)
	PushFront(value T)
	PushBack(value T)
	PopFront() (value T, ok bool)
	PopBack() (value T, ok bool)
}

func TestDeque(t *testing.T) {
	var q collection.Deque[int]
	testList(t, &q)
}

func testList(t *testing.T, l list[int]) {
	// First test run
	testListOnce(t, l)
	// Repeat
	testListOnce(t, l)
	// Add element, clear, repeat
	l.PushBack(13)
	l.Clear()
	testListOnce(t, l)
}

func testListOnce(t *testing.T, l list[int]) {
	// Empty list
	{
		_, ok := l.Front()
		require.False(t, ok)
	}
	{
		_, ok := l.Back()
		require.False(t, ok)
	}
	{
		_, ok := l.PopFront()
		require.False(t, ok)
	}
	{
		_, ok := l.PopBack()
		require.False(t, ok)
	}

	// Push back
	l.PushBack(42)
	require.Equal(t, 1, l.Len())
	{
		value, ok := l.Front()
		require.Equal(t, 42, value)
		require.True(t, ok)
	}
	{
		value, ok := l.Back()
		require.Equal(t, 42, value)
		require.True(t, ok)
	}

	// Push front twice
	l.PushFront(17)
	l.PushFront(4711)
	require.Equal(t, 3, l.Len())
	{
		value, ok := l.Front()
		require.Equal(t, 4711, value)
		require.True(t, ok)
	}
	{
		value, ok := l.Back()
		require.Equal(t, 42, value)
		require.True(t, ok)
	}

	// Pop back
	{
		value, ok := l.PopBack()
		require.Equal(t, 42, value)
		require.True(t, ok)
		require.Equal(t, 2, l.Len())
	}
	{
		value, ok := l.Front()
		require.Equal(t, 4711, value)
		require.True(t, ok)
	}
	{
		value, ok := l.Back()
		require.Equal(t, 17, value)
		require.True(t, ok)
	}

	// Pop front
	{
		value, ok := l.PopFront()
		require.Equal(t, 4711, value)
		require.True(t, ok)
		require.Equal(t, 1, l.Len())
	}
	{
		value, ok := l.Front()
		require.Equal(t, 17, value)
		require.True(t, ok)
	}
	{
		value, ok := l.Back()
		require.Equal(t, 17, value)
		require.True(t, ok)
	}

	// Pop back again. The list is empty again.
	{
		value, ok := l.PopBack()
		require.Equal(t, 17, value)
		require.True(t, ok)
		require.Equal(t, 0, l.Len())
	}
	{
		_, ok := l.Front()
		require.False(t, ok)
	}
	{
		_, ok := l.Back()
		require.False(t, ok)
	}
	{
		_, ok := l.PopFront()
		require.False(t, ok)
	}
	{
		_, ok := l.PopBack()
		require.False(t, ok)
	}
}

func BenchmarkDequePushFrontPopFront(b *testing.B) {
	var l collection.Deque[int]
	benchmarkPushFrontPopFront(b, &l)
}

func BenchmarkDequePushFrontPopBack(b *testing.B) {
	var l collection.Deque[int]
	benchmarkPushBackPopBack(b, &l)
}

func BenchmarkDequePushBackPopFront(b *testing.B) {
	var l collection.Deque[int]
	benchmarkPushBackPopFront(b, &l)
}

func BenchmarkDequePushBackPopBack(b *testing.B) {
	var l collection.Deque[int]
	benchmarkPushBackPopBack(b, &l)
}

func benchmarkPushFrontPopFront(b *testing.B, l list[int]) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.PushFront(42)
	}
	for {
		_, ok := l.PopFront()
		if !ok {
			break
		}
	}
}

func benchmarkPushFrontPopBack(b *testing.B, l list[int]) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.PushFront(42)
	}
	for {
		_, ok := l.PopBack()
		if !ok {
			break
		}
	}
}

func benchmarkPushBackPopFront(b *testing.B, l list[int]) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.PushBack(42)
	}
	for {
		_, ok := l.PopFront()
		if !ok {
			break
		}
	}
}

func benchmarkPushBackPopBack(b *testing.B, l list[int]) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l.PushBack(42)
	}
	for {
		_, ok := l.PopBack()
		if !ok {
			break
		}
	}
}
