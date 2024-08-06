package schedule

import (
	"github.com/cholazzzb/amaz_corp_be/pkg/algo"
)

type Scheduler struct {
	taskMap map[string]TaskWithDetailQuery
	graph   *algo.Graph
}

func NewScheduler(tasks []TaskWithDetailQuery) *Scheduler {
	taskMap := map[string]TaskWithDetailQuery{}

	taskIDs := []string{}
	adj := map[string][]string{}
	indegree := map[string]int{}
	for _, task := range tasks {
		taskMap[task.TaskID] = task
		taskIDs = append(taskIDs, task.TaskID)
		adj[task.TaskID] = task.Dependencies
		indegree[task.TaskID] = 0
	}

	for _, task := range tasks {
		for _, dep := range task.Dependencies {
			if len(dep) > 0 {
				indegree[dep]++
			}
		}
	}

	graph := algo.NewGraph(
		taskIDs,
		adj,
		indegree,
	)

	return &Scheduler{
		taskMap: taskMap,
		graph:   graph,
	}
}

func (sch *Scheduler) GetScheduledTask() ([]TaskWithDetailQuery, error) {
	sorted, err := sch.graph.TopologicalSort()
	if err != nil {
		return []TaskWithDetailQuery{}, err
	}

	out := []TaskWithDetailQuery{}
	for idx := len(sorted) - 1; idx >= 0; idx-- {
		taskID := sorted[idx]
		task := sch.taskMap[taskID]
		out = append(out, task)
	}
	return out, nil
}
