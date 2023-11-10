package schedule

type Graph struct {
	Node         map[string]TaskWithDetailQuery
	Dependencies TasksDependency
}

func TopologicalSort(graph Graph) []TaskWithDetailQuery {
	sorted := []TaskWithDetailQuery{}
	visited := map[string]bool{}

	var visit func(node TaskWithDetailQuery)
	visit = func(node TaskWithDetailQuery) {
		if _, ok := visited[node.TaskID]; !ok {
			visited[node.TaskID] = true
			if deps, ok := graph.Dependencies[node.TaskID]; ok {
				for _, dep := range deps {
					visit(graph.Node[dep])
				}
			}
			sorted = append(sorted, node)
		}
	}

	for _, node := range graph.Node {
		visit(node)
	}

	// TODO: Handle Cyclic graph

	return sorted
}

func CreateGraph(tasks []TaskWithDetailQuery) Graph {
	node := map[string]TaskWithDetailQuery{}
	dependencies := TasksDependency{}

	for _, task := range tasks {
		node[task.TaskID] = task
		dependencies[task.TaskID] = task.Dependencies
	}

	return Graph{
		Node:         node,
		Dependencies: dependencies,
	}
}
