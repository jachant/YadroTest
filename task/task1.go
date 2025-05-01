package task

import (
	"sort"
)

func Task1(containers [][]uint64) string {
	n := len(containers)
	colorsSum := make([]uint64, n)
	containerSums := make([]uint64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			containerSums[i] += containers[i][j]
			colorsSum[j] += containers[i][j]
		}
	}
	sort.Slice(containerSums, func(i, j int) bool {
		return containerSums[i] < containerSums[j]
	})
	sort.Slice(colorsSum, func(i, j int) bool {
		return colorsSum[i] < colorsSum[j]
	})
	for i := 0; i < n; i++ {
		if colorsSum[i] != containerSums[i] {
			return "no"
		}
	}
	return "yes"
}
