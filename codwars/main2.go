package main

import "math"

func TwoOldestAges(ages []int) [2]int {
	var result [2]int

	oldestAge := -math.MaxInt32
	secondOldestAge := oldestAge

	for _, age := range ages {
		if age > oldestAge {
			secondOldestAge = oldestAge
			oldestAge = age
		} else if age > secondOldestAge {
			secondOldestAge = age
		}
	}
	result[0] = secondOldestAge
	result[1] = oldestAge

	return result
}
