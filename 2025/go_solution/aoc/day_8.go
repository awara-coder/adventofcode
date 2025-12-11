package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/pkg/datastructures"
	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

type JunctionBox struct {
	X int64
	Y int64
	Z int64
}

func (j *JunctionBox) getSquaredStraightLineDistance(other *JunctionBox) int64 {
	return (j.X-other.X)*(j.X-other.X) + (j.Y-other.Y)*(j.Y-other.Y) + (j.Z-other.Z)*(j.Z-other.Z)
}

func SolveDay8(junctionBoxPositions []string) (int64, error) {
	junctionBoxes, err := parseAndReturnJunctionBoxes(junctionBoxPositions)
	if err != nil {
		utils.GetLogger().Fatalf("error while parsing junctin boxes input: %v", err)
	}

	// Find the pair of shortest distance junction boxes
	// Each entry stores a tuple of ID1, ID2, distance_between_them_squared
	shortestDistanceJunctionBoxPairs := make([][]int64, 0)
	for i := range len(junctionBoxes) {
		for j := i + 1; j < len(junctionBoxes); j++ {
			shortestDistanceJunctionBoxPairs = append(shortestDistanceJunctionBoxPairs, []int64{int64(i), int64(j), junctionBoxes[i].getSquaredStraightLineDistance(&junctionBoxes[j])})
		}
	}

	// sort them in increasing order of distances
	slices.SortFunc(shortestDistanceJunctionBoxPairs, func(a, b []int64) int {
		return int(a[2] - b[2])
	})

	// Iterate over them and keep joining the junction boxes into circuits
	// return solveDay8Part1(junctionBoxes, shortestDistanceJunctionBoxPairs)

	return solveDay8Part2(junctionBoxes, shortestDistanceJunctionBoxPairs)
}

func solveDay8Part1(junctionBoxes []JunctionBox, shortestDistanceJunctionBoxPairs [][]int64) (int64, error) {
	circuitsToJoin := 10
	// challenge input case.
	if len(junctionBoxes) > 20 {
		circuitsToJoin = 1000
	}
	// Create circuit, and represent it using DSU data structure.
	circuit := datastructures.NewDSU(len(junctionBoxes))
	for i := range circuitsToJoin {
		// Join both pairs
		circuit.Add(int(shortestDistanceJunctionBoxPairs[i][0]), int(shortestDistanceJunctionBoxPairs[i][1]))
	}

	// Find unique circuit sizes and multiply 3 largest ones
	circuitSize := make(map[int]int64)
	for i := range len(junctionBoxes) {
		parentI := circuit.Find(i)
		circuitSize[parentI] = int64(circuit.GetSize(parentI))
	}

	// We now have unique circuit parent's and their sizes,
	topThree := getTopThreeCircuits(circuitSize)
	if topThree == nil {
		utils.GetLogger().Printf("got nil while fetching top three circuits")
		return 0, fmt.Errorf("got nil while fetching top three circuits")
	}

	return topThree[0] * topThree[1] * topThree[2], nil
}

func solveDay8Part2(junctionBoxes []JunctionBox, shortestDistanceJunctionBoxPairs [][]int64) (int64, error) {
	// Create circuit, and represent it using DSU data structure.
	circuit := datastructures.NewDSU(len(junctionBoxes))
	sol := int64(-1)

	for i := range len(shortestDistanceJunctionBoxPairs) {
		// Join both pairs
		if circuit.Add(int(shortestDistanceJunctionBoxPairs[i][0]), int(shortestDistanceJunctionBoxPairs[i][1])) {
			// If this was a successful join, save product of thier x coordinates
			sol = junctionBoxes[int(shortestDistanceJunctionBoxPairs[i][0])].X * junctionBoxes[int(shortestDistanceJunctionBoxPairs[i][1])].X
		}
	}

	return sol, nil
}

func getTopThreeCircuits(circuitSize map[int]int64) []int64 {
	if len(circuitSize) < 3 {
		return nil
	}

	allCircuitSizes := make([]int64, 0, len(circuitSize))
	for _, size := range circuitSize {
		allCircuitSizes = append(allCircuitSizes, size)
	}

	slices.Sort(allCircuitSizes)
	slices.Reverse(allCircuitSizes)

	return allCircuitSizes[:3]
}

func parseAndReturnJunctionBoxes(junctionBoxPositions []string) ([]JunctionBox, error) {
	junctionBoxes := make([]JunctionBox, 0, len(junctionBoxPositions))

	for _, junctionBoxPosition := range junctionBoxPositions {
		tokens := strings.Split(junctionBoxPosition, ",")
		if len(tokens) != 3 {
			return nil, fmt.Errorf("tokens not equal 3 when parsing junction box position, len: %v", len(tokens))
		}

		junctionBox := JunctionBox{}

		// Parse 3 tokens to get x,y,z coordinate
		var parseError error
		junctionBox.X, parseError = strconv.ParseInt(tokens[0], 10, 64)
		if parseError != nil {
			utils.GetLogger().Printf("Error while parsing string integer: %v", tokens[0])
			return nil, fmt.Errorf("Error while parsing string integer: %v", tokens[0])
		}

		junctionBox.Y, parseError = strconv.ParseInt(tokens[1], 10, 64)
		if parseError != nil {
			utils.GetLogger().Printf("Error while parsing string integer: %v", tokens[1])
			return nil, fmt.Errorf("Error while parsing string integer: %v", tokens[1])
		}

		junctionBox.Z, parseError = strconv.ParseInt(tokens[2], 10, 64)
		if parseError != nil {
			utils.GetLogger().Printf("Error while parsing string integer: %v", tokens[2])
			return nil, fmt.Errorf("Error while parsing string integer: %v", tokens[2])
		}

		junctionBoxes = append(junctionBoxes, junctionBox)
	}

	return junctionBoxes, nil
}
