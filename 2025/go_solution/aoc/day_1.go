package aoc

import (
	"strconv"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

func SolveDay1(commands []string) (int64, error) {
	var zeroesCounter int64 = 0
	var currentDialPosition int64 = 50

	for _, command := range commands {
		// Parse command
		isLeftRotation := strings.HasPrefix(command, "L")
		turns, err := strconv.ParseInt(command[1:], 10, 64)
		if err != nil {
			utils.GetLogger().Printf("Error while parsing turns, %w", err)
			return 0, err
		}

		// Perform rotation
		newDialPosition := currentDialPosition
		if isLeftRotation {
			newDialPosition -= turns % 100
			newDialPosition += 100
			newDialPosition %= 100
		} else {
			newDialPosition += turns % 100
			newDialPosition %= 100
		}

		// zeroesCounter += day1Part1Solution(newDialPosition)

		zeroesCounter += day1Part2Solution(currentDialPosition, turns, isLeftRotation)

		currentDialPosition = newDialPosition

	}

	return zeroesCounter, nil
}

func day1Part1Solution(currentDialPosition int64) int64 {
	// Check if dial's last position is zero.
	if currentDialPosition == 0 {
		return 1
	}
	return 0
}

// day1Part2Solution returns count the number of times current dial was on 0
func day1Part2Solution(currentDialPosition int64, turns int64, isLeftRotation bool) int64 {
	if isLeftRotation {
		// Simulate as if the turns are made in positive direction with a different starting point which will hit 0 same number of times.
		// starting at 20 with left rotation is the same as starting at 80 with right rotation. (you just need 20 to reach 100 and ...)
		currentDialPosition = (100 - currentDialPosition) % 100
		return (currentDialPosition + turns) / 100
	} else {
		// If turns were positive (right), we can count the number of times we triggered 0, not counting if we are already on zero.
		return (currentDialPosition + turns) / 100
	}

}
