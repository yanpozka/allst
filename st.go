package allmst

import (
	"container/list"
	"errors"
	"fmt"
)

type Graph struct {
	nodes []int
	adj   []*list.List
}

var ErrInvalidIndex = errors.New("invalid index node")

func NewGraph(size int) *Graph {
	g := Graph{
		nodes: make([]int, size),
		adj:   make([]*list.List, size),
	}

	for ix := 0; ix < size; ix++ {
		g.adj[ix] = list.New()
	}

	return &g
}

type path []int

func (g *Graph) AllSpanningTrees(root int) []path {
	set := map[int]bool{root: true}

	return g.allST(root, set)
}

func (g *Graph) allST(root int, set map[int]bool) []path {

	var p []path

	for e := g.adj[root].Front(); e != nil; e = e.Next() {
		a := e.Value.(int)

		if _, contains := set[a]; !contains {

			new_set := cloneSet(set)
			new_set[a] = true

			childPaths := g.allST(a, set)

			for _, chp := range childPaths {
				p = append(p, append(chp, root))
			}
		}
	}

	if len(p) == 0 { // leaf?
		p = append(p, []int{root})
	}

	return p
}

func cloneSet(set map[int]bool) map[int]bool {
	new_set := make(map[int]bool)

	for k, v := range set {
		new_set[k] = v
	}

	return new_set
}

func (g *Graph) PrintBFS(root int) {
	visited := make([]bool, len(g.nodes))
	visited[root] = true

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		r := queue.Remove(queue.Front()).(int)
		fmt.Print(r, " ")

		for e := g.adj[r].Front(); e != nil; e = e.Next() {
			if a := e.Value.(int); !visited[a] {
				queue.PushBack(a)
				visited[a] = true
			}
		}
	}
}

func (g *Graph) AddEdge(src, dst int) error {
	if src >= len(g.nodes) || dst >= len(g.nodes) {
		return ErrInvalidIndex
	}

	g.adj[src].PushBack(dst)

	return nil
}
