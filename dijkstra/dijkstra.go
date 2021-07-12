package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Vertex ...
type Vertex struct {
	Name    int
	Edges   []*Edge
	visited bool
}

// NewVertex ...
func NewVertex(name int) *Vertex {
	return &Vertex{
		Name:  name,
		Edges: make([]*Edge, 0),
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
	Tail   *Vertex
	Head   *Vertex
	Weight int
}

// String ...
func (e *Edge) String() string {
	if e == nil {
		return "<nil>"
	}

	return e.Tail.String() + " " + e.Head.String() + " " + strconv.Itoa(e.Weight)
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
			parts := strings.Fields(s)
			if len(parts) < 2 {
				return nil, errors.New("invalid file format")
			}

			tail, err3 := strconv.Atoi(parts[0])
			if err3 != nil {
				return nil, err3
			}

			i, ok := vids[tail]
			if !ok {
				i = len(g.Vertices)
				vids[tail] = i
				g.Vertices = append(g.Vertices, NewVertex(tail))
			}

			for _, p := range parts[1:] {
				pp := strings.Split(strings.TrimSpace(p), ",")
				if len(pp) != 2 {
					return nil, errors.New("invalid file format")
				}

				head, err4 := strconv.Atoi(pp[0])
				if err4 != nil {
					return nil, err4
				}

				w, err4 := strconv.Atoi(pp[1])
				if err4 != nil {
					return nil, err4
				}

				j, ok := vids[head]
				if !ok {
					j = len(g.Vertices)
					vids[head] = j
					g.Vertices = append(g.Vertices, NewVertex(head))
				}

				e := &Edge{
					Tail:   g.Vertices[i],
					Head:   g.Vertices[j],
					Weight: w,
				}
				g.Edges = append(g.Edges, e)

				g.Vertices[i].Edges = append(g.Vertices[i].Edges, e)
			}
		}

		if err2 == io.EOF {
			break
		}
	}

	return g, nil
}

// Dijkstra ...
func Dijkstra(g *Graph, s *Vertex) map[int]int {
	x := map[int]bool{s.Name: true}
	dists := map[int]int{s.Name: 0}

	for {
		min := -1
		v := -1
		for _, e := range g.Edges {
			if x[e.Tail.Name] && !x[e.Head.Name] {
				d := dists[e.Tail.Name] + e.Weight
				if min == -1 || d < min {
					min = d
					v = e.Head.Name
				}
			}
		}

		if v == -1 {
			break
		}
		x[v] = true
		dists[v] = min
	}

	return dists
}
