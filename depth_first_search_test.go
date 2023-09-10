package datastructures

import (
	"testing"
)

func TestDFS(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(3, 5)

	edges := g.DFS(0)
	result := []int{0, 1, 3, 5, 4, 2}
	if !areSlicesEqual(edges, result) {
		t.Errorf("Error with DFS: result %v expected %v", result, edges)
	}
}

func areSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
