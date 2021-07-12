package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	query = flag.String("query", "", "List of vertices for querying distances")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Printf("Usage: %s <file> <source>\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	g, err := loadGraph(args[0])
	if err != nil {
		fmt.Println("Could not load graph:", err)
		os.Exit(1)
	}

	s, err := strconv.Atoi(args[1])

	vertices, err := parseQuery(*query)
	if err != nil {
		fmt.Println("Could not parse query: %s", *query)
		os.Exit(1)
	}

	start := time.Now()
	dists := Dijkstra(g, NewVertex(s))
	elapsed := time.Since(start)
	if len(vertices) == 0 {
		fmt.Printf("Shortest path from %d: %v\n", s, dists)
	} else {
		ds := make([]int, 0, len(vertices))
		for _, v := range vertices {
			ds = append(ds, dists[v])
		}
		fmt.Printf("Shortest path from %d to (%s): %v\n", s, *query, ds)
	}
	fmt.Println("Running time:", elapsed)
}

func parseQuery(q string) ([]int, error) {
	if q == "" {
		return nil, nil
	}

	parts := strings.Split(q, ",")
	if len(parts) == 0 {
		return nil, nil
	}

	vs := make([]int, 0, len(parts))
	for _, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		vs = append(vs, v)
	}

	return vs, nil
}
