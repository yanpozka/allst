package allmst

import "testing"

func TestNewGraph(t *testing.T) {
	g := NewGraph(4)

	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	g.PrintBFS(0)

	ps := g.AllSpanningTrees(0)

	for _, p := range ps {
		t.Log(p)
	}
}
