package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	start := time.Now()
	g, err := loadGraph(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Finish loading graph data from file:", time.Since(start))

	Kasaraju(g)
	top5 := topSCC(g, 5)
	fmt.Println("Finish finding top 5 SCCs:", top5, time.Since(start))
}
