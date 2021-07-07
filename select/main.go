package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	size  = flag.Int("size", 1000, "Size of input array")
	stype = flag.String("type", "r", "Type of select algorithm to use")
	ftype = flag.String("file-type", "t", "Type of input file")
)

func readPi(path, ftype string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	if ftype == "t" {
		return string(buf)[2:], nil
	}

	s := make([]byte, 0, len(buf)*2)
	for _, b := range buf {
		s = append(s, '0'+b/16, '0'+b%16)
	}

	return string(s), nil
}

func buildInput(path string, ftype string, n int) ([]int, error) {
	s, err := readPi(path, ftype)
	if err != nil {
		return nil, err
	}

	a := make([]int, n)
	for i := range a {
		if len(s) == 0 {
			break
		}

		k := 10
		if len(s) < 10 {
			k = len(s)
		}

		num := s[:k]
		v, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		a[i] = v

		s = s[k:]
	}

	return a, nil
}

func printUsage(pname string) {
	fmt.Printf("Usage: %s [options] <file>\n", pname)
	fmt.Printf("\nOptions:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		printUsage(os.Args[0])
	}

	rand.Seed(time.Now().UnixNano())

	a, err := buildInput(args[0], *ftype, *size)
	if err != nil {
		fmt.Println("Could not build input array:", err)
		os.Exit(1)
	}

	start := time.Now()
	median := 0
	if *stype == "r" {
		median = RSelect(a, (len(a)-1)/2)
	} else if *stype == "d" {
		median = DSelect(a, (len(a)-1)/2)
	}
	fmt.Printf("Input length: %d, median: %d, time: %v\n", len(a), median, time.Since(start))
}
