package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	type testData struct {
		data   []int
		answer int
	}

	data := []testData{
		{[]int{2, 3}, 5},
		{[]int{1, 2}, 3},
		testData{[]int{1, 0, -1}, 0},
	}

	for _, d := range data {

		x := sum(d.data...)
		if x != d.answer {
			t.Error("Expected", d.answer, "Got", x)
		}
	}
}
