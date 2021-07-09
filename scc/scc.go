package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Vertex ...
type Vertex struct {
	Name     int
	Edges    []*Edge
	RevEdges []*Edge
	visited  bool
	Label    int
	SCC      int
}

// NewVertex ...
func NewVertex(name int) *Vertex {
	return &Vertex{
		Name:     name,
		Edges:    make([]*Edge, 0),
		RevEdges: make([]*Edge, 0),
	}
}

// String ...
func (v *Vertex) String() string {
	if v == nil {
		return "<nil>"
	}
	return strconv.Itoa(v.Name)
}

// Edge ...
type Edge struct {
	Tail *Vertex
	Head *Vertex
}

// String ...
func (e *Edge) String() string {
	if e == nil {
		return "<nil>"
	}

	return e.Tail.String() + " " + e.Head.String()
}

// Graph ...
type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

// NewGraph ...
func NewGraph() *Graph {
	return &Graph{
		Vertices: make([]*Vertex, 0),
		Edges:    make([]*Edge, 0),
	}
}

// String ...
func (g *Graph) String() string {
	if g == nil {
		return "<nil>"
	}

	buf := strings.Builder{}
	for _, e := range g.Edges {
		_, _ = buf.WriteString(e.String())
		_, _ = buf.WriteString("\n")
	}

	return buf.String()
}

func loadGraph(path string) (*Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	vids := make(map[int]int)
	g := NewGraph()
	for {
		s, err2 := r.ReadString('\n')
		if err2 != nil && err2 != io.EOF {
			return nil, err2
		}

		s = strings.TrimSpace(s)
		if s != "" {
		}

		if s != "" {
			parts := strings.Fields(strings.TrimSpace(s))
			if len(parts) != 2 {
				return nil, errors.New("invalid file format")
			}

			tail, err3 := strconv.Atoi(parts[0])
			if err3 != nil {
				return nil, err3
			}

			head, err3 := strconv.Atoi(parts[1])
			if err3 != nil {
				return nil, err3
			}

			i, ok := vids[tail]
			if !ok {
				i = len(g.Vertices)
				vids[tail] = i
				g.Vertices = append(g.Vertices, NewVertex(tail))
			}
			j, ok := vids[head]
			if !ok {
				j = len(g.Vertices)
				vids[head] = j
				g.Vertices = append(g.Vertices, NewVertex(head))
			}

			e := &Edge{
				Tail: g.Vertices[i],
				Head: g.Vertices[j],
			}
			g.Edges = append(g.Edges, e)

			g.Vertices[i].Edges = append(g.Vertices[i].Edges, e)
			g.Vertices[j].RevEdges = append(g.Vertices[j].RevEdges, e)
		}

		if err2 == io.EOF {
			break
		}
	}

	return g, nil
}

func pushVertex(s *Stack, v *Vertex) {
	s.Push(v)
}

func popVertex(s *Stack) *Vertex {
	v, err := s.Pop()
	if err != nil {
		panic(err)
	}
	return v.(*Vertex)
}

func dfsTopoRev(g *Graph, s *Vertex, curLabel *int) {
	vs := NewStack() // Stack used for visiting vertices.
	ls := NewStack() // Stack used for assigning label.
	pushVertex(vs, s)

	// Search for all reachable vertices from `s`.
	for !vs.Empty() {
		v := popVertex(vs)
		if !v.visited {
			v.visited = true

			for _, e := range v.RevEdges {
				pushVertex(vs, e.Tail)
			}

			pushVertex(ls, v) // Add v to stack for assigning label.
		}
	}

	// Assign topological ordering.
	for !ls.Empty() {
		v := popVertex(ls)
		v.Label = *curLabel
		*curLabel--
	}
}

func topoSortRev(g *Graph) {
	// Mark all vertices unvisited.
	for _, v := range g.Vertices {
		v.visited = false
	}

	curLabel := len(g.Vertices)
	for _, v := range g.Vertices {
		if !v.visited {
			dfsTopoRev(g, v, &curLabel)
		}
	}
}

func dfsSCC(g *Graph, s *Vertex, numSCC int) {
	ss := NewStack()
	pushVertex(ss, s)
	for !ss.Empty() {
		v := popVertex(ss)
		if !v.visited {
			v.visited = true
			v.SCC = numSCC
			for _, e := range v.Edges {
				pushVertex(ss, e.Head)
			}
		}
	}
}

// Kasaraju ...
func Kasaraju(g *Graph) {
	// Calculate topological ordering for reverse of graph `g`.
	topoSortRev(g)

	// Sort vertices by topological ordering.
	sortedIds := make([]int, len(g.Vertices))
	for i, v := range g.Vertices {
		sortedIds[v.Label-1] = i
	}

	// Mark all vertices unvisited.
	for _, v := range g.Vertices {
		v.visited = false
	}

	// Assign SCCs for graph `g`.
	numSCC := 0
	for _, i := range sortedIds {
		v := g.Vertices[i]
		if !v.visited {
			numSCC++
			dfsSCC(g, v, numSCC)
		}
	}
}

func topSCC(g *Graph, n int) []int {
	sccCountMap := make(map[int]int)
	for _, v := range g.Vertices {
		count := sccCountMap[v.SCC]
		sccCountMap[v.SCC] = count + 1
	}

	counts := make([]int, 0, len(sccCountMap))
	for _, count := range sccCountMap {
		counts = append(counts, count)
	}

	if len(counts) < n {
		cs := len(counts)
		for i := 0; i < n-cs; i++ {
			counts = append(counts, 0)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return counts[:n]
}
