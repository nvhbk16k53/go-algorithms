package main

import (
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
	pivotTypeMedianOf9
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

//func medianOf3(a []int) int {
//	n := len(a)
//	lo, mid, hi := 0, (n-1)/2, n-1
//
//	if a[lo] > a[hi] {
//		a[lo], a[hi] = a[hi], a[lo]
//	}
//	if a[lo] > a[mid] {
//		a[lo], a[mid] = a[mid], a[lo]
//	}
//	if a[mid] > a[hi] {
//		a[mid], a[hi] = a[hi], a[mid]
//	}
//
//	return mid
//}

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

func medianOf9(a []int) int {
	n := len(a)
	k := n / 3
	lo := medianOf3(a[:k])
	mid := medianOf3(a[k : 2*k])
	hi := medianOf3(a[2*k:])

	a[0], a[lo] = a[lo], a[0]
	a[(n+1)/2], a[k+mid] = a[k+mid], a[(n+1)/2]
	a[2*k+hi], a[n-1] = a[n-1], a[2*k+hi]

	return medianOf3(a)
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
	case pivotTypeMedianOf9:
		return medianOf9(a)
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
	pivot = partition(a, pivot)

	totalComps := len(a) - 1
	if pivot > 0 {
		totalComps += QuickSort(a[:pivot], pivotType)
	}
	if pivot < len(a)-1 {
		totalComps += QuickSort(a[pivot+1:], pivotType)
	}

	return totalComps
}

func insertionSort(a []int) int {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
	return n * (n - 1) / 2
}

// FastQuickSort ...
func FastQuickSort(a []int) int {
	if len(a) <= 10 {
		return insertionSort(a)
	}

	pivot := choosePivot(a, pivotTypeMedianOf9)
	pivot = partition(a, pivot)

	totalComps := len(a) - 1
	if pivot > 0 {
		totalComps += FastQuickSort(a[:pivot])
	}
	if pivot < len(a)-1 {
		totalComps += FastQuickSort(a[pivot+1:])
	}

	return totalComps
}
