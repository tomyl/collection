package collection_test

import (
	"fmt"
	"testing"

	"github.com/tomyl/collection"
)

func ExampleSlice() {
	var s collection.Slice[int]
	s.PushBack(42)
	s.PushBack(17)
	value, _ := s.PopBack()
	fmt.Println(value)
	// Output: 17
}

func TestSlice(t *testing.T) {
	var s collection.Slice[int]
	testList(t, &s)
}

func BenchmarkSlicePushFrontPopFront(b *testing.B) {
	var s collection.Slice[int]
	benchmarkPushFrontPopFront(b, &s)
}

func BenchmarkSlicePushFrontPopBack(b *testing.B) {
	var s collection.Slice[int]
	benchmarkPushFrontPopBack(b, &s)
}

func BenchmarkSlicePushBackPopFront(b *testing.B) {
	var s collection.Slice[int]
	benchmarkPushBackPopFront(b, &s)
}

func BenchmarkSlicePushBackPopBack(b *testing.B) {
	var s collection.Slice[int]
	benchmarkPushBackPopBack(b, &s)
}
