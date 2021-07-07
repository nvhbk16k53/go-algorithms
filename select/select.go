package main

import (
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func choosePivot(a []int) int {
	return a[rand.Intn(len(a))]
}

func partition(a []int, pivot int) int {
	// Partition the array around pivot element.
	i := 0
	pidx := 0
	for j := 0; j < len(a); j++ {
		if a[j] == pivot {
			a[pidx], a[j] = a[j], a[pidx]
			pidx++
		} else if a[j] < pivot {
			a[pidx], a[j] = a[j], a[pidx]
			a[i], a[pidx] = a[pidx], a[i]
			pidx++
			i++
		}
	}

	return pidx - 1
}

// RSelect ...
func RSelect(a []int, i int) int {
	if len(a) == 0 {
		panic("a is an empty array")
	}

	if len(a) == 1 {
		return a[0]
	}

	p := choosePivot(a)
	j := partition(a, p)

	if j == i {
		return p
	}
	if j > i {
		return RSelect(a[:j], i)
	}
	return RSelect(a[j+1:], i-j-1)
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

func medianOf5(a []int) int {
	if len(a) == 1 {
		return a[0]
	}

	_ = insertionSort(a)
	return a[(len(a)-1)/2]
}

// DSelect ...
func DSelect(a []int, i int) int {
	n := len(a)

	if n == 1 {
		return a[0]
	}

	k := (n + 4) / 5
	medians := make([]int, 0, k)
	for j := 0; j < k; j++ {
		si := j * 5
		ei := si + 5
		if ei > n {
			ei = n
		}

		medians = append(medians, medianOf5(a[si:ei]))
	}

	p := DSelect(medians, (k-1)/2)
	j := partition(a, p)

	if j == i {
		return p
	}
	if j > i {
		return DSelect(a[:j], i)
	}
	return DSelect(a[j+1:], i-j-1)
}
