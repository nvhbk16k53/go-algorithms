package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKasaraju(t *testing.T) {
	data := []struct {
		path   string
		expect []int
	}{
		{
			path:   "problem8.10test1.txt",
			expect: []int{3, 3, 3, 0, 0},
		},
		{
			path:   "problem8.10test2.txt",
			expect: []int{3, 3, 2, 0, 0},
		},
		{
			path:   "problem8.10test3.txt",
			expect: []int{3, 3, 1, 1, 0},
		},
		{
			path:   "problem8.10test4.txt",
			expect: []int{7, 1, 0, 0, 0},
		},
		{
			path:   "problem8.10test5.txt",
			expect: []int{6, 3, 2, 1, 0},
		},
	}

	for _, test := range data {
		g, err := loadGraph(test.path)
		if assert.NoError(t, err) {
			Kasaraju(g)
			top5 := topSCC(g, 5)
			assert.Equal(t, test.expect, top5)
		}
	}
}
