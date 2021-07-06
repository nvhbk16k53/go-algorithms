package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	pivotTypeFirst int = iota
	pivotTypeLast
	pivotTypeRandom
	pivotTypeMedianOf3
)

func loadData(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	nums := strings.Fields(string(data))
	a := make([]int, 0, len(nums))
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		a = append(a, n)
	}

	return a, nil
}

func medianOf3(a []int) int {
	n := len(a)
	mid := (n - 1) / 2

	x, y, z := a[0], a[mid], a[n-1]
	if (x > y && x < z) || (x > z && x < y) {
		return 0
	}
	if (y > x && y < z) || (y < x && y > z) {
		return mid
	}
	return n - 1
}

func choosePivot(a []int, pivotType int) int {
	n := len(a)
	switch pivotType {
	case pivotTypeLast:
		return n - 1
	case pivotTypeRandom:
		return rand.Intn(n)
	case pivotTypeMedianOf3:
		return medianOf3(a)
	default:
		return 0
	}
}

func partition(a []int, pivot int) int {
	// Swap pivot element with first element of array.
	a[0], a[pivot] = a[pivot], a[0]

	// Partition the array around pivot element.
	i := 1
	for j := 1; j < len(a); j++ {
		if a[j] < a[0] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	// Swap pivot element with last element of left part.
	a[0], a[i-1] = a[i-1], a[0]

	return i - 1
}

// QuickSort ...
func QuickSort(a []int, pivotType int) int {
	if len(a) <= 1 {
		return 0
	}

	pivot := choosePivot(a, pivotType)
	if pivot >= len(a) {
		panic(fmt.Sprintf("Invalid pivot index %d for array length %d", pivot, len(a)))
	}
	pivot = partition(a, pivot)
	if pivot >= len(a) {
		panic(fmt.Sprintf("Invalid pivot index %d for array length %d", pivot, len(a)))
	}

	totalComps := len(a) - 1
	if pivot > 0 {
		totalComps += QuickSort(a[:pivot], pivotType)
	}
	if pivot < len(a)-1 {
		totalComps += QuickSort(a[pivot+1:], pivotType)
	}

	return totalComps
}
