package aoc

import (
	"fmt"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

func SolveDay7(tachyonManifoldDiagram []string) (int64, error) {

	// solution := solveDay7Part1(tachyonManifoldDiagram)
	solution := solveDay7Part2(tachyonManifoldDiagram)

	return solution, nil
}

func solveDay7Part1(tachyonManifoldDiagram []string) int64 {
	// currentBeamRow denotes the current row at which beam is current present.
	currentBeamRow := 0
	totalSplitCount := int64(0)

	// run simulation until current beam has reached the last row.
	for ; currentBeamRow < len(tachyonManifoldDiagram); currentBeamRow++ {
		nextTachyonManifoldDiagram, splitCount := getNextStageOfTachyonManifold(tachyonManifoldDiagram, currentBeamRow)
		printTachyonManifoldDiagram(nextTachyonManifoldDiagram)
		totalSplitCount += splitCount

		tachyonManifoldDiagram = nextTachyonManifoldDiagram
	}
	return totalSplitCount
}

// getNextStageOfTachyonManifold takes the diagram and current row in which beam is current present,
// and returns the next tachyonManifoldDiagram after beam has travelled one more row as well as splits that happend.
// This is part 1 solution.
func getNextStageOfTachyonManifold(tachyonManifoldDiagram []string, currentBeamRow int) ([]string, int64) {
	splitCount := int64(0)
	updatedTachyonManifoldDiagram := make([]string, len(tachyonManifoldDiagram))
	copy(updatedTachyonManifoldDiagram, tachyonManifoldDiagram)

	// Edge case: if current beam row is last row.
	if currentBeamRow == len(tachyonManifoldDiagram)-1 {
		return tachyonManifoldDiagram, 0
	}

	for col := 0; col < len(tachyonManifoldDiagram[currentBeamRow]); col++ {
		// Check if current character is a beam
		if updatedTachyonManifoldDiagram[currentBeamRow][col] == '|' || updatedTachyonManifoldDiagram[currentBeamRow][col] == 'S' {
			// If next row is splitter
			if updatedTachyonManifoldDiagram[currentBeamRow+1][col] == '^' {
				splitCount++

				// Split the beam to left and right
				if col-1 >= 0 {
					updatedTachyonManifoldDiagram[currentBeamRow+1] = updatedTachyonManifoldDiagram[currentBeamRow+1][:col-1] + "|" + updatedTachyonManifoldDiagram[currentBeamRow+1][col:]
				}
				if col+1 < len(tachyonManifoldDiagram[currentBeamRow]) {
					updatedTachyonManifoldDiagram[currentBeamRow+1] = updatedTachyonManifoldDiagram[currentBeamRow+1][:col+1] + "|" + updatedTachyonManifoldDiagram[currentBeamRow+1][col+2:]
				}
			} else {
				// Pass the beam as it is to next row.
				updatedTachyonManifoldDiagram[currentBeamRow+1] = updatedTachyonManifoldDiagram[currentBeamRow+1][:col] + "|" + updatedTachyonManifoldDiagram[currentBeamRow+1][col+1:]
			}
		}
	}

	return updatedTachyonManifoldDiagram, splitCount
}

func solveDay7Part2(diagram []string) int64 {
	// DFS through the manifold and return all the possible paths (timelines)
	startingRow := 0
	startingCol := strings.IndexAny(diagram[startingRow], "S")
	if startingCol == -1 {
		utils.GetLogger().Fatalf("starting column not found")
	}

	// We need to memoize the pre computed solutions in order to prevent exponential findTachyonTimeline calls and slow.
	rows := len(diagram)
	cols := len(diagram[0])
	memo := make([][]int64, rows)
	for row, _ := range memo {
		memo[row] = make([]int64, cols)
		for col, _ := range memo[row] {
			memo[row][col] = -1
		}
	}

	return findTachyonTimeline(diagram, startingRow, startingCol, memo)
}

func findTachyonTimeline(diagram []string, currentRow, currentCol int, memo [][]int64) int64 {
	// If we are out of bounds
	if currentCol < 0 || currentCol >= len(diagram[currentRow]) {
		return int64(0)
	}

	// If we have reached last row
	if currentRow == len(diagram)-1 {
		return int64(1)
	}

	// Check if we have already precomputed the solution in memo
	if memo[currentRow][currentCol] != -1 {
		return memo[currentRow][currentCol]
	}

	// check if next row is a splitter
	if diagram[currentRow+1][currentCol] == '^' {
		// Go left and right
		memo[currentRow][currentCol] = findTachyonTimeline(diagram, currentRow+1, currentCol-1, memo) + findTachyonTimeline(diagram, currentRow+1, currentCol+1, memo)
	} else {
		// Go as it is to next row
		memo[currentRow][currentCol] = findTachyonTimeline(diagram, currentRow+1, currentCol, memo)
	}

	return memo[currentRow][currentCol]
}

func printTachyonManifoldDiagram(diagram []string) {
	for _, row := range diagram {
		fmt.Println(row)
	}
	fmt.Println()
}
