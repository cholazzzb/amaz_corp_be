package algo_test

import (
	"fmt"
	"testing"

	"github.com/cholazzzb/amaz_corp_be/pkg/algo"
)

func TestKhanAlgorithmSuccess(t *testing.T) {
	IDs := []string{"0", "1", "2", "3", "4", "5"}
	adj := map[string][]string{}
	for _, ID := range IDs {
		adj[ID] = []string{}
	}

	adj["5"] = append(adj["5"], "2")
	adj["5"] = append(adj["5"], "0")

	adj["4"] = append(adj["4"], "0")
	adj["4"] = append(adj["4"], "1")

	adj["2"] = append(adj["2"], "3")

	adj["3"] = append(adj["3"], "1")

	indegree := map[string]int{
		"2": 1, "3": 1, "4": 2, "5": 2,
	}
	g := algo.NewGraph(IDs, adj, indegree)

	order, err := g.TopologicalSort()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Topological Sort Order:", order)
	}

}

func TestKhanAlgorithmCycle(t *testing.T) {
	IDs := []string{"0", "1", "2"}
	adj := map[string][]string{}

	for _, ID := range IDs {
		adj[ID] = []string{}
	}

	adj["0"] = append(adj["0"], "1")
	adj["1"] = append(adj["1"], "2")
	adj["2"] = append(adj["2"], "0")

	indegree := map[string]int{
		"0": 1, "1": 1, "2": 1,
	}

	g2 := algo.NewGraph(IDs, adj, indegree)

	order, err := g2.TopologicalSort()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Topological Sort Order:", order)
	}
}
