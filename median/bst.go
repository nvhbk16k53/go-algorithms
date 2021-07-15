package main

import (
	"errors"
)

// BSNode ...
type BSNode struct {
	Parent *BSNode
	Left   *BSNode
	Right  *BSNode
	Key    int
	Size   int
}

// BSTree ...
type BSTree struct {
	root *BSNode
}

// NewBSTree ...
func NewBSTree() *BSTree {
	return &BSTree{
		root: nil,
	}
}

// Search ...
func (t *BSTree) Search(k int) (int, error) {
	n := t.root
	for n != nil {
		if n.Key == k {
			return n.Key, nil
		}

		if k < n.Key {
			n = n.Left
		} else {
			n = n.Right
		}
	}

	return 0, errors.New("key not found")
}

// Insert ...
func (t *BSTree) Insert(k int) {
	var p *BSNode
	var dir int
	for n := t.root; n != nil; {
		if k == n.Key {
			break // Ignore existing key.
		}

		p = n
		if k < n.Key {
			dir = 0
			n = n.Left
		} else {
			dir = 1
			n = n.Right
		}
	}

	n := &BSNode{
		Parent: p,
		Left:   nil,
		Right:  nil,
		Key:    k,
		Size:   1,
	}

	if p == nil {
		t.root = n
	} else {
		if dir == 0 {
			p.Left = n
		} else {
			p.Right = n
		}
	}

	for p = n.Parent; p != nil; p = p.Parent {
		p.Size++
	}
}

// GetMedian ...
func (t *BSTree) GetMedian() int {
	if t.root == nil {
		panic("tree is empty")
	}

	mpos := (t.root.Size + 1) / 2
	for n := t.root; n != nil; {
		lsize := 0
		if n.Left != nil {
			lsize = n.Left.Size
		}

		if lsize+1 == mpos {
			return n.Key
		}

		if lsize+1 > mpos {
			n = n.Left
		} else {
			mpos -= lsize + 1
			n = n.Right
		}
	}

	return 0
}

// RBNode color constants
const (
	RBNodeColorBlack int = iota
	RBNodeColorRed
)

// RBNode child direction constants
const (
	RBNodeChildLeft  = 0
	RBNodeChildRight = 1
)

// RBNode ...
type RBNode struct {
	parent *RBNode
	child  [2]*RBNode
	key    int
	color  int
	size   int
}

// RBTree ...
type RBTree struct {
	root *RBNode
}

// NewRBTree ...
func NewRBTree() *RBTree {
	return &RBTree{}
}

func childDir(n *RBNode) int {
	if n == n.parent.child[RBNodeChildRight] {
		return RBNodeChildRight
	}
	return RBNodeChildLeft
}

func getParent(n *RBNode) *RBNode {
	if n == nil {
		return nil
	}
	return n.parent
}

func getGrandParent(n *RBNode) *RBNode {
	return getParent(getParent(n))
}

func getSibling(n *RBNode) *RBNode {
	p := getParent(n)
	if p == nil {
		return nil
	}
	return p.child[1-childDir(n)]
}

func getUncle(n *RBNode) *RBNode {
	p := getParent(n)
	if p == nil {
		return nil
	}

	return getSibling(p)
}

func getCloseNephew(n *RBNode) *RBNode {
	p := getParent(n)
	if p == nil {
		return nil
	}

	dir := childDir(n)
	s := p.child[1-dir]
	if s == nil {
		return nil
	}

	return s.child[dir]
}

func getDistanceNephew(n *RBNode) *RBNode {
	p := getParent(n)
	if p == nil {
		return nil
	}

	dir := childDir(n)
	s := p.child[1-dir]
	if s == nil {
		return nil
	}

	return s.child[1-dir]
}

func getSize(n *RBNode) int {
	if n == nil {
		return 0
	}
	return n.size
}

func (t *RBTree) rotateDirRoot(p *RBNode, dir int) *RBNode {
	g := getParent(p)
	s := p.child[1-dir]
	if s == nil {
		panic("rotate root to <nil> node")
	}

	c := s.child[dir]
	p.child[1-dir] = c
	if c != nil {
		c.parent = p
	}
	s.child[dir] = p
	p.parent = s
	s.parent = g
	if g != nil {
		if g.child[RBNodeChildRight] == p {
			g.child[RBNodeChildRight] = s
		} else {
			g.child[RBNodeChildLeft] = s
		}
	} else {
		t.root = s
	}

	p.size = getSize(p.child[0]) + getSize(p.child[1]) + 1
	s.size = getSize(s.child[0]) + getSize(s.child[1]) + 1

	return s
}

func (t *RBTree) insert(p *RBNode, n *RBNode, dir int) {
	if p == nil {
		t.root = n
		return
	}
	p.child[dir] = n

	// Rebalancing loop.
	for {
		if p.color == RBNodeColorBlack {
			return
		}

		g := getParent(p)
		if g == nil {
			p.color = RBNodeColorBlack
			return
		}

		dir = childDir(p)
		u := g.child[1-dir]
		if u == nil || u.color == RBNodeColorBlack {
			if n == p.child[1-dir] {
				t.rotateDirRoot(p, dir)
				n = p
				p = g.child[dir]
			}

			t.rotateDirRoot(g, 1-dir)
			p.color = RBNodeColorBlack
			g.color = RBNodeColorRed
			return
		}

		p.color = RBNodeColorBlack
		u.color = RBNodeColorBlack
		g.color = RBNodeColorRed
		n = g
		p = n.parent
		if p == nil {
			return
		}
	}
}

// Insert ...
func (t *RBTree) Insert(k int) {
	var p *RBNode
	var dir int
	for n := t.root; n != nil; {
		if k == n.key {
			break // Ignore existing key.
		}

		p = n
		if k < n.key {
			dir = RBNodeChildLeft
		} else {
			dir = RBNodeChildRight
		}
		n = n.child[dir]
	}

	n := &RBNode{
		parent: p,
		child:  [2]*RBNode{nil, nil},
		key:    k,
		size:   1,
		color:  RBNodeColorRed,
	}
	for pp := n.parent; pp != nil; pp = pp.parent {
		pp.size++
	}

	t.insert(p, n, dir)
}

// GetMedian ...
func (t *RBTree) GetMedian() int {
	if t.root == nil {
		panic("tree is empty")
	}

	mpos := (t.root.size + 1) / 2
	for n := t.root; n != nil; {
		lsize := 0
		if n.child[RBNodeChildLeft] != nil {
			lsize = n.child[RBNodeChildLeft].size
		}

		if lsize+1 == mpos {
			return n.key
		}

		if lsize+1 > mpos {
			n = n.child[RBNodeChildLeft]
		} else {
			mpos -= lsize + 1
			n = n.child[RBNodeChildRight]
		}
	}

	return 0
}
