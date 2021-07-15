package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	dtype = flag.String("type", "heap", "Data structure type used to maintain median")
)

// MedianMaintainer ...
type MedianMaintainer interface {
	Insert(int)
	GetMedian() int
}

type heapMedian struct {
	left  *Heap
	right *Heap
}

func newHeapMedian() MedianMaintainer {
	return &heapMedian{
		left:  NewHeap(),
		right: NewHeap(),
	}
}

func (h *heapMedian) Insert(num int) {
	if h.left.Empty() {
		h.left.Insert(-num)
		return
	}
	lk := -h.left.FindMin()

	if h.right.Empty() {
		if lk > num {
			h.right.Insert(-h.left.ExtractMin())
			h.left.Insert(-num)
		} else {
			h.right.Insert(num)
		}
		return
	}
	rk := h.right.FindMin()

	if num < lk {
		if h.left.Len() > h.right.Len() {
			h.right.Insert(-h.left.ExtractMin())
		}
		h.left.Insert(-num)
	} else if num > rk {
		if h.right.Len() > h.left.Len() {
			h.left.Insert(-h.right.ExtractMin())
		}
		h.right.Insert(num)
	} else {
		if h.left.Len() <= h.right.Len() {
			h.left.Insert(-num)
		} else {
			h.right.Insert(num)
		}
	}
}

func (h *heapMedian) GetMedian() int {
	if h.right.Len() > h.left.Len() {
		return h.right.FindMin()
	}
	return -h.left.FindMin()
}

func loadData(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		s := strings.TrimSpace(line)
		if s == "" {
			continue
		}

		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Printf("Usage: %s [options] <file>\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var mm MedianMaintainer
	switch *dtype {
	case "heap":
		mm = newHeapMedian()
	case "bst":
		mm = NewBSTree()
	case "rbt":
		mm = NewRBTree()
	default:
		fmt.Println("Invalid data structure type:", *dtype)
		os.Exit(1)
	}

	a, err := loadData(args[0])
	if err != nil {
		fmt.Printf("Could not load data from file %s: %v", args[0], err)
		os.Exit(1)
	}

	start := time.Now()
	sum := 0
	for _, v := range a {
		mm.Insert(v)
		m := mm.GetMedian()
		sum += m
	}

	fmt.Println("Sum of kth medians:", sum)
	fmt.Println("Running time:", time.Since(start))
}
