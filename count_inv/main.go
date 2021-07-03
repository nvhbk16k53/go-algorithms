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

func readFile(path string) ([]int, error) {
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

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	a, err := readFile(args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("Input array len:", len(a))

	start := time.Now()
	invCount, _ := CountInv(a)
	fmt.Println("Number of inversions:", invCount)
	fmt.Println("Wall time:", time.Since(start))
}
