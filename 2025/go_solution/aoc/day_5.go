package aoc

import (
	"slices"
	"strconv"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

func SolveDay5(lines []string) (int64, error) {
	// Parse the input.
	ranges, _ := parseDay5Input(lines)

	// Call solver function
	// freshIngridients := solveDay5Part1(ranges, queries)
	freshIngridients := solveDay5Part2(ranges)
	return freshIngridients, nil
}

func parseDay5Input(lines []string) ([][]int64, []int64) {
	ranges := make([][]int64, 0)

	lineIdx := 0
	for ; lineIdx < len(lines); lineIdx += 1 {

		if lines[lineIdx] == "" {
			break
		}

		rangeStr := strings.Split(lines[lineIdx], "-")
		if len(rangeStr) != 2 {
			utils.GetLogger().Fatalf("length or rangeStr != 2, len: %v", len(rangeStr))
		}
		rangeStart, err := strconv.ParseInt(rangeStr[0], 10, 64)
		if err != nil {
			utils.GetLogger().Fatalf("Failed to parse begin range value: %v", rangeStr[0])
		}
		rangeEnd, err := strconv.ParseInt(rangeStr[1], 10, 64)
		if err != nil {
			utils.GetLogger().Fatalf("Failed to parse end range value: %v", rangeStr[1])
		}

		ranges = append(ranges, []int64{rangeStart, rangeEnd})
	}

	lineIdx += 1

	queries := make([]int64, 0)
	for ; lineIdx < len(lines); lineIdx++ {
		query, err := strconv.ParseInt(lines[lineIdx], 10, 64)
		if err != nil {
			utils.GetLogger().Fatalf("Failed to parse end range value: %v", query)
		}
		queries = append(queries, query)
	}

	return ranges, queries
}

func sortAndMergeIngridientRanges(ranges [][]int64) [][]int64 {
	// Sort the ranges and merge them.
	slices.SortFunc(ranges, func(a, b []int64) int {
		if a[0] == b[0] {
			// Sort by end time in ascending order
			return int(a[1] - b[1])
		}

		// Sort by start time.
		return int(a[0] - b[0])
	})

	// Merge the sorted ranges
	mergedRanges := make([][]int64, 0, len(ranges))

	for _, currentRange := range ranges {
		// Check if we should merge this to last range
		if len(mergedRanges) > 0 && mergedRanges[len(mergedRanges)-1][1] >= currentRange[0] {
			// merge because last range's end time is >= current range's start time
			mergedRanges[len(mergedRanges)-1][1] = max(mergedRanges[len(mergedRanges)-1][1], currentRange[1])
		} else {
			// Add new range
			mergedRanges = append(mergedRanges, currentRange)
		}
	}
	return mergedRanges
}

func solveDay5Part1(ranges [][]int64, queries []int64) int64 {

	mergedRanges := sortAndMergeIngridientRanges(ranges)

	freshIngridents := int64(0)

	// Answer the queries
	for _, query := range queries {
		if isIngridentInFreshRange(mergedRanges, query) {
			freshIngridents++
		}
	}

	return freshIngridents
}

func solveDay5Part2(ranges [][]int64) int64 {
	mergedRanges := sortAndMergeIngridientRanges(ranges)

	// return sum merged ranges
	freshIngridients := int64(0)
	for _, currentRange := range mergedRanges {
		freshIngridients += currentRange[1] - currentRange[0] + 1
	}

	return freshIngridients

}

func isIngridentInFreshRange(ranges [][]int64, query int64) bool {
	// Binary search to find the last range whose start <= query
	lo := 0
	hi := len(ranges) - 1
	rangeIdx := -1

	for lo <= hi {
		mid := (hi-lo)/2 + lo

		if ranges[mid][0] <= query {
			rangeIdx = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if rangeIdx == -1 {
		return false
	}

	return ranges[rangeIdx][0] <= query && query <= ranges[rangeIdx][1]
}
