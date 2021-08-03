package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func twoSum(a []int, m map[int]bool, t int) bool {
	for _, x := range a {
		y := t - x
		if y != x && m[y] {
			return true
		}
	}
	return false
}

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
	for _, s := range nums {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		a = append(a, num)
	}

	return a, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <file> <min> <max>\n", os.Args[0])
		os.Exit(1)
	}

	tmin, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid value for t min:", err)
		os.Exit(1)
	}

	tmax, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid value for t max:", err)
		os.Exit(1)
	}

	if tmin >= tmax {
		fmt.Println("t min must be less than t max", tmin, tmax)
		os.Exit(1)
	}

	a, err := loadData(os.Args[1])
	if err != nil {
		fmt.Println("Could not load data from file:", err)
		os.Exit(1)
	}

	m := make(map[int]bool)
	for _, n := range a {
		m[n] = true
	}

	total := 0
	start := time.Now()
	for t := tmin; t <= tmax; t++ {
		if twoSum(a, m, t) {
			total++
		}
	}
	fmt.Println("Total number of t:", total)
	fmt.Println("Running time:", time.Since(start))
}
