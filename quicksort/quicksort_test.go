package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	data := []struct {
		path      string
		pivotType int
		expect    int
	}{
		{
			path:      "problem5.6test1.txt",
			pivotType: pivotTypeFirst,
			expect:    25,
		},
		{
			path:      "problem5.6test1.txt",
			pivotType: pivotTypeLast,
			expect:    31,
		},
		{
			path:      "problem5.6test1.txt",
			pivotType: pivotTypeMedianOf3,
			expect:    21,
		},
		{
			path:      "problem5.6test1.txt",
			pivotType: pivotTypeRandom,
		},
		{
			path:      "problem5.6test2.txt",
			pivotType: pivotTypeFirst,
			expect:    620,
		},
		{
			path:      "problem5.6test2.txt",
			pivotType: pivotTypeLast,
			expect:    573,
		},
		{
			path:      "problem5.6test2.txt",
			pivotType: pivotTypeMedianOf3,
			expect:    502,
		},
		{
			path:      "problem5.6test2.txt",
			pivotType: pivotTypeRandom,
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, test := range data {
		a, err := loadData(test.path)
		assert.NoError(t, err)

		totalComps := QuickSort(a, test.pivotType)
		assert.True(t, sort.SliceIsSorted(a, func(i, j int) bool {
			return a[i] <= a[j]
		}))
		if test.pivotType != pivotTypeRandom {
			assert.Equal(t, test.expect, totalComps)
		}
		fmt.Printf("Input length: %d, pivot type: %d\n", len(a), test.pivotType)
		fmt.Println("Total number of comparisions:", totalComps)
	}
}
