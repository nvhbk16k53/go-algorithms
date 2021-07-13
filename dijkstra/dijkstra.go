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

// Heap ...
type Heap struct {
	keys     []int
	vertices []*Vertex
	posMap   map[int]int
}

// NewHeap ...
func NewHeap() *Heap {
	return &Heap{
		keys:     make([]int, 0),
		vertices: make([]*Vertex, 0),
		posMap:   make(map[int]int),
	}
}

func (h *Heap) bubbleUp(pos int) {
	for pos > 0 {
		parent := (pos - 1) / 2
		if h.keys[pos] >= h.keys[parent] {
			break
		}

		h.keys[pos], h.keys[parent] = h.keys[parent], h.keys[pos]
		h.vertices[pos], h.vertices[parent] = h.vertices[parent], h.vertices[pos]
		h.posMap[h.vertices[pos].Name] = pos
		h.posMap[h.vertices[parent].Name] = parent
		pos = parent
	}
}

func (h *Heap) bubbleDown(pos int) {
	n := len(h.keys)
	for {
		left := pos*2 + 1
		right := pos*2 + 2
		if left >= n {
			break
		}

		child := left
		if right < n && h.keys[left] > h.keys[right] {
			child = right
		}

		if h.keys[pos] <= h.keys[child] {
			break
		}

		h.keys[pos], h.keys[child] = h.keys[child], h.keys[pos]
		h.vertices[pos], h.vertices[child] = h.vertices[child], h.vertices[pos]
		h.posMap[h.vertices[pos].Name] = pos
		h.posMap[h.vertices[child].Name] = child
		pos = child
	}
}

// Empty ...
func (h *Heap) Empty() bool {
	return len(h.keys) == 0
}

// Insert ...
func (h *Heap) Insert(k int, v *Vertex) {
	h.keys = append(h.keys, k)
	h.vertices = append(h.vertices, v)
	h.posMap[v.Name] = len(h.keys) - 1
	h.bubbleUp(len(h.keys) - 1)
}

// ExtractMin ...
func (h *Heap) ExtractMin() (int, *Vertex) {
	n := len(h.keys)
	h.keys[0], h.keys[n-1] = h.keys[n-1], h.keys[0]
	h.vertices[0], h.vertices[n-1] = h.vertices[n-1], h.vertices[0]
	h.posMap[h.vertices[0].Name] = 0
	h.posMap[h.vertices[n-1].Name] = n - 1

	k := h.keys[n-1]
	v := h.vertices[n-1]

	h.keys = h.keys[:n-1]
	h.vertices = h.vertices[:n-1]
	h.bubbleDown(0)

	return k, v
}

// Delete ...
func (h *Heap) Delete(pos int) (int, *Vertex) {
	n := len(h.keys)
	if pos >= n {
		panic("heap out of range")
	}

	if pos < n-1 {
		h.keys[pos], h.keys[n-1] = h.keys[n-1], h.keys[pos]
		h.vertices[pos], h.vertices[n-1] = h.vertices[n-1], h.vertices[pos]
		h.posMap[h.vertices[pos].Name] = pos
	}

	k := h.keys[n-1]
	v := h.vertices[n-1]

	h.keys = h.keys[:n-1]
	h.vertices = h.vertices[:n-1]
	delete(h.posMap, v.Name)

	parent := (pos - 1) / 2
	if pos > 0 && pos < n-1 && h.keys[pos] > h.keys[parent] {
		h.bubbleUp(pos)
	} else {
		h.bubbleDown(pos)
	}

	return k, v
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

// FastDijkstra ...
func FastDijkstra(g *Graph, s *Vertex) map[int]int {
	x := make(map[int]bool)
	dists := map[int]int{s.Name: 0}

	h := NewHeap()
	h.Insert(0, s)

	for !h.Empty() {
		k, v := h.ExtractMin()
		x[v.Name] = true
		dists[v.Name] = k

		for _, e := range v.Edges {
			if x[e.Head.Name] {
				continue
			}

			pos, ok := h.posMap[e.Head.Name]
			if !ok {
				h.Insert(k+e.Weight, e.Head)
			} else {
				his, _ := h.Delete(pos)
				if k+e.Weight < his {
					his = k + e.Weight
				}
				h.Insert(his, e.Head)
			}
		}
	}

	return dists
}
