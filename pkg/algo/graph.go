package algo

import (
	"errors"
)

// Graph represents a directed graph using adjacency list representation
type Graph struct {
	IDs      []string
	adj      map[string][]string
	indegree map[string]int
}

// NewGraph initializes a new directed graph with vertices
func NewGraph(
	IDs []string,
	adj map[string][]string,
	indegree map[string]int,
) *Graph {
	return &Graph{
		IDs:      IDs,
		adj:      adj,
		indegree: indegree,
	}
}

// TopologicalSort sort vertices using Kahn's algorithm
func (g *Graph) TopologicalSort() ([]string, error) {
	result := []string{}
	queue := []string{}

	for _, ID := range g.IDs {
		if g.indegree[ID] == 0 {
			queue = append(queue, ID)
		}
	}

	count := 0

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]
		result = append(result, el)
		count++

		for _, nei := range g.adj[el] {
			g.indegree[nei]--

			if g.indegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	if count != len(g.IDs) {
		return nil, errors.New("graph has a cycle")
	}

	return result, nil
}
