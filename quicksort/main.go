package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func deepCopy(a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	input, err := loadData(os.Args[1])
	if err != nil {
		fmt.Printf("Could not load input array from file: %s\n", os.Args[1])
	}
	fmt.Println("Length of the input array:", len(input))

	a := deepCopy(input)
	totalComps := QuickSort(a, pivotTypeFirst)
	fmt.Println("Total number of comparisions for pivot first:", totalComps)
	if !isSorted(a) {
		fmt.Println("The output array is not in sorted order")
		os.Exit(1)
	}

	a = deepCopy(input)
	totalComps = QuickSort(a, pivotTypeLast)
	fmt.Println("Total number of comparisions for pivot last:", totalComps)
	if !isSorted(a) {
		fmt.Println("The output array is not in sorted order")
		os.Exit(1)
	}

	a = deepCopy(input)
	totalComps = QuickSort(a, pivotTypeMedianOf3)
	fmt.Println("Total number of comparisions for pivot median of 3:", totalComps)
	if !isSorted(a) {
		fmt.Println("The output array is not in sorted order")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	totalComps = 0
	for i := 0; i < 10; i++ {
		a = deepCopy(input)
		totalComps += QuickSort(a, pivotTypeRandom)
	}
	fmt.Println("Total number of comparisions for pivot random:", totalComps/10)
	if !isSorted(a) {
		fmt.Println("The output array is not in sorted order")
		os.Exit(1)
	}

	a = deepCopy(input)
	totalComps = FastQuickSort(a)
	fmt.Println("Total number of comparisions for FastQuickSort:", totalComps)
	if !isSorted(a) {
		fmt.Println("The output array is not in sorted order")
		os.Exit(1)
	}
}
