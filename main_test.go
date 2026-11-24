package main

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 1, 1}, 3},
		{[]int{}, 0},
	}
	for _, c := range cases {
		if got := MaxConsecutiveOnes(c.in); got != c.want {
			t.Errorf("MaxConsecutiveOnes(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
