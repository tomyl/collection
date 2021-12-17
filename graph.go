package collection

import (
	"constraints"
)

// Graph implements a directed graph.
type Graph[K comparable, V any, W constraints.Unsigned] struct {
	edges map[K]map[K]edge[V, W]
}

// SetEdge inserts an edge into the graph. If an edge already exists between
// the nodes, that edge is overwritten.
func (g *Graph[K, V, W]) SetEdge(from, to K, value V, weight W) {
	if g.edges == nil {
		g.edges = make(map[K]map[K]edge[V, W], 1)
	}
	outgoing, ok := g.edges[from]
	if !ok {
		outgoing = make(map[K]edge[V, W], 1)
		g.edges[from] = outgoing
	}
	outgoing[to] = edge[V, W]{value, weight}
}

// ShortestPath returns the shorted path between two nodes. Returns both the node path
// and the edge path. Return nil if the goal node can't be reached from the
// start node. If the start node and the goal node is the same node, that node is returned and the edge path is nil.
func (g *Graph[K, V, W]) ShortestPath(start, goal K) (nodes []K, edges []V) {
	visited := map[K]struct{}{start: struct{}{}}
	dist := map[K]W{}   // lowest distance per node seen so far
	origin := map[K]K{} // previous node for lowest distance seen so far

	var queue Heap[W, K]
	queue.Push(0, start)

	for {
		total, node, ok := queue.Pop()
		if !ok {
			return nil, nil
		}
		if node == goal {
			break
		}
		outgoing, ok := g.edges[node]
		if ok {
			for next, edge := range outgoing {
				if _, ok := visited[next]; !ok {
					visited[next] = struct{}{}
					newTotal := total + edge.weight
					if oldTotal, ok := dist[next]; !ok || newTotal < oldTotal {
						dist[next] = newTotal
						origin[next] = node
					}
					queue.Push(newTotal, next)
				}
			}
		}
	}

	// Found the goal. Reconstruct the nod path.
	nodes = append(nodes, goal)
	for {
		prev, ok := origin[nodes[len(nodes)-1]]
		if !ok {
			break
		}
		nodes = append(nodes, prev)
	}

	// Reverse the node path.
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}

	// Build the edge path
	for i := 1; i < len(nodes); i++ {
		outgoing := g.edges[nodes[i-1]]
		edges = append(edges, outgoing[nodes[i]].value)
	}

	return nodes, edges
}

type edge[V any, W constraints.Ordered] struct {
	value  V
	weight W
}
