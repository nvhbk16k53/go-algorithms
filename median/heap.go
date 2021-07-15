package main

// Heap ...
type Heap struct {
	keys []int
}

// NewHeap ...
func NewHeap() *Heap {
	return &Heap{
		keys: make([]int, 0),
	}
}

// Empty ...
func (h *Heap) Empty() bool {
	return h.Len() == 0
}

// Len ...
func (h *Heap) Len() int {
	return len(h.keys)
}

func (h *Heap) bubbleUp(pos int) {
	for pos > 0 {
		parent := (pos - 1) / 2
		if h.keys[parent] <= h.keys[pos] {
			break
		}

		h.keys[pos], h.keys[parent] = h.keys[parent], h.keys[pos]

		pos = parent
	}
}

// Insert ...
func (h *Heap) Insert(k int) {
	h.keys = append(h.keys, k)
	h.bubbleUp(h.Len() - 1)
}

func (h *Heap) bubbleDown(pos int) {
	n := h.Len()
	for {
		left := 2*pos + 1
		right := 2*pos + 2
		if left >= n {
			break // pos is a leaf node.
		}

		child := left
		if right < n && h.keys[right] < h.keys[left] {
			child = right
		}

		if h.keys[pos] <= h.keys[child] {
			break // heap invariants satisfied.
		}

		h.keys[pos], h.keys[child] = h.keys[child], h.keys[pos]

		pos = child
	}
}

// ExtractMin ...
func (h *Heap) ExtractMin() int {
	if h.Empty() {
		panic("heap is empty")
	}

	n := h.Len()
	h.keys[0], h.keys[n-1] = h.keys[n-1], h.keys[0]
	k := h.keys[n-1]
	h.keys = h.keys[:n-1]
	h.bubbleDown(0)

	return k
}

// FindMin ...
func (h *Heap) FindMin() int {
	if h.Empty() {
		panic("heap is empty")
	}

	return h.keys[0]
}
