package aoc

func SolveDay4(grid []string) (int64, error) {
	// accessibleRolls := solveDay4Part1(grid)
	accessibleRolls := solveDay4Part2(grid)
	return accessibleRolls, nil
}

func solveDay4Part1(grid []string) int {
	accessibleRolls := getAccessilbeRolls(grid)
	return len(accessibleRolls)
}

func getAccessilbeRolls(grid []string) [][]int {
	rows, cols := len(grid), len(grid[0])
	accessibleRolls := make([][]int, 0)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			if grid[row][col] != '@' {
				continue
			}

			// Check current row and column and if it is accessible
			if getNeighborCount(grid, row, col) < 4 {
				accessibleRolls = append(accessibleRolls, []int{row, col})
			}

		}
	}
	return accessibleRolls
}

func solveDay4Part2(grid []string) int64 {
	totalAccessibleRolls := int64(0)
	for {
		accessibleRolls := getAccessilbeRolls(grid)
		if len(accessibleRolls) == 0 {
			break
		}
		totalAccessibleRolls += int64(len(accessibleRolls))

		// Mark each row and col with '.'
		for _, accessibleRole := range accessibleRolls {
			row, col := accessibleRole[0], accessibleRole[1]
			grid[row] = grid[row][:col] + "." + grid[row][col+1:]
		}
	}
	return totalAccessibleRolls
}

func getNeighborCount(grid []string, row, col int) int {
	rows, cols := len(grid), len(grid[0])
	neighborCount := 0

	for dRow := -1; dRow <= 1; dRow++ {
		for dCol := -1; dCol <= 1; dCol++ {
			if dRow == dCol && dRow == 0 {
				continue
			}

			// Check if within bounds
			if row+dRow < 0 || row+dRow >= rows || col+dCol < 0 || col+dCol >= cols {
				continue
			}

			// Check if it is a roll
			if grid[row+dRow][col+dCol] == '@' {
				neighborCount++
			}
		}
	}

	return neighborCount
}
