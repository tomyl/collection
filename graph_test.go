package collection_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tomyl/collection"
)

func ExampleGraph() {
	var g collection.Graph[string, string, uint]
	g.SetEdge("a", "b", "ab", 1)
	g.SetEdge("a", "c", "ac", 2)
	g.SetEdge("b", "c", "bc", 3)
	g.SetEdge("b", "d", "bd", 4)
	nodes, edges := g.ShortestPath("a", "d")
	fmt.Println(nodes, edges)
	// Output: [a b d] [ab bd]
}

func TestGraph(t *testing.T) {
	var g collection.Graph[string, string, uint]
	g.SetEdge("a", "b", "ab", 40)
	g.SetEdge("a", "c", "ac", 30)
	g.SetEdge("b", "c", "bc", 1)
	g.SetEdge("b", "d", "bd", 20)
	g.SetEdge("c", "d", "cd", 10)
	g.SetEdge("d", "e", "de", 100)
	g.SetEdge("x", "y", "xy", 100)
	{
		nodes, edges := g.ShortestPath("a", "y")
		require.Nil(t, nodes)
		require.Nil(t, edges)
	}
	{
		nodes, edges := g.ShortestPath("a", "a")
		require.Equal(t, []string{"a"}, nodes)
		require.Nil(t, edges)
	}
	{
		nodes, edges := g.ShortestPath("a", "e")
		require.Equal(t, []string{"a", "c", "d", "e"}, nodes)
		require.Equal(t, []string{"ac", "cd", "de"}, edges)
	}
}

func BenchmarkGraphBuild(b *testing.B) {
	b.ReportAllocs()
	var g collection.Graph[int, struct{}, uint]
	for i := 0; i < b.N; i++ {
		g.SetEdge(i, i+1, struct{}{}, 1)
	}
}

func BenchmarkGraphShortedPath(b *testing.B) {
	var g collection.Graph[int, struct{}, uint]
	for i := 0; i < b.N; i++ {
		g.SetEdge(i, i+1, struct{}{}, 1)
	}
	b.ReportAllocs()
	b.ResetTimer()
	nodes, edges := g.ShortestPath(0, b.N)
	require.Equal(b, b.N+1, len(nodes))
	require.Equal(b, b.N, len(edges))
}
