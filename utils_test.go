package sortedmap_test

import "strconv"

func safeAtoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}
