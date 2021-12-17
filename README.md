# collection

[![GoDoc](https://godoc.org/github.com/tomyl/collection?status.png)](http://godoc.org/github.com/tomyl/collection)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomyl/collection)](https://goreportcard.com/report/github.com/tomyl/collection)

A few generic data structures for Go. Tested with `go1.18beta1`.

Work in progress. Expect bugs, poor performance, API instability.

```sh
go get github.com/tomyl/collection
```

# Examples

## Deque 

A double-ended queue based on a double-linked list.

```go
    var q collection.Deque[int]
    q.PushFront(42)
    q.PushFront(17)
    value, ok := q.PopBack() // 42 true
```

## Graph

A directed graph.

```go
    var g collection.Graph[string, string, uint]
    g.SetEdge("a", "b", "ab", 1)
    g.SetEdge("a", "c", "ac", 2)
    g.SetEdge("b", "c", "bc", 3)
    g.SetEdge("b", "d", "bd", 4)
    nodes, edges := g.ShortestPath("a", "d") // [a b d] [ab bd]
```

## Heap

A binary min heap. Trivial to use as a priority queue, just set the key to the negative priority.

```go
    var h collection.Heap[int, string]
    h.Push(42, "foo")
    h.Push(17, "bar")
    key, value, ok := h.Pop()  // 17 "bar" true
```

## Slice

```go
    var s collection.Slice[int]
    s.PushBack(42)
    s.PushBack(17)
    value, ok := s.PopBack() // 17 true
```
