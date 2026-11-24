package main

import "fmt"

// MaxConsecutiveOnes returns the length of the longest run of 1s.
func MaxConsecutiveOnes(nums []int) int {
	best, run := 0, 0
	for _, v := range nums {
		if v == 1 {
			run++
			if run > best {
				best = run
			}
		} else {
			run = 0
		}
	}
	return best
}

func main() {
	fmt.Println(MaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
