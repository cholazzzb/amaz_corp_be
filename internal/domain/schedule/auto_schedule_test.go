package schedule_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cholazzzb/amaz_corp_be/internal/domain/schedule"
)

func TestTopologicalSort(t *testing.T) {
	g := schedule.CreateGraph(
		[]schedule.TaskWithDetailQuery{
			{TaskID: "3", Dependencies: []string{"2"}},
			{TaskID: "5", Dependencies: []string{"4"}},
			{TaskID: "1", Dependencies: []string{}},
			{TaskID: "4", Dependencies: []string{"3"}},
			{TaskID: "2", Dependencies: []string{"1"}},
		},
	)
	res := schedule.TopologicalSort(g)
	expected := []schedule.TaskWithDetailQuery{
		{TaskID: "1", Dependencies: []string{}},
		{TaskID: "2", Dependencies: []string{"1"}},
		{TaskID: "3", Dependencies: []string{"2"}},
		{TaskID: "4", Dependencies: []string{"3"}},
		{TaskID: "5", Dependencies: []string{"4"}},
	}

	assert.Equal(t, res, expected, "Basic Test Case")
}
