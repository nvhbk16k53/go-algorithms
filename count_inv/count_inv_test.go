package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountInv(t *testing.T) {
	data := []struct {
		a      []int
		expect int
	}{
		{
			a:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expect: 0,
		},
		{
			a:      []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expect: 45,
		},
		{
			a:      []int{54044, 14108, 79294, 29649, 25260, 60660, 2995, 53777, 49689, 9083},
			expect: 28,
		},
	}

	for _, test := range data {
		count, _ := CountInv(test.a)
		assert.Equal(t, test.expect, count)
	}
}
