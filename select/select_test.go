package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRSelect(t *testing.T) {
	data := []struct {
		path   string
		expect int
	}{
		{
			path:   "problem6.5test1.txt",
			expect: 5469,
		},
		{
			path:   "problem6.5test2.txt",
			expect: 4715,
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, test := range data {
		a, err := loadData(test.path)
		assert.NoError(t, err)

		res := RSelect(a, (len(a)-1)/2)
		assert.Equal(t, test.expect, res)
	}
}

func TestDSelect(t *testing.T) {
	data := []struct {
		path   string
		expect int
	}{
		{
			path:   "problem6.5test1.txt",
			expect: 5469,
		},
		{
			path:   "problem6.5test2.txt",
			expect: 4715,
		},
	}

	rand.Seed(time.Now().UnixNano())

	for _, test := range data {
		a, err := loadData(test.path)
		assert.NoError(t, err)

		res := DSelect(a, (len(a)-1)/2)
		assert.Equal(t, test.expect, res)
	}
}
