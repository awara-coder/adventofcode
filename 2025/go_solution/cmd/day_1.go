package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

// Input file path
const day1InputFilePath = "day_1/sample_input"

// const day1InputFilePath = "day_1/input.input"

func day1Solution() {
	// Read file contents
	lines, err := utils.ReadFileContents(day1InputFilePath)
	if err != nil {
		utils.GetLogger().Fatalf("Error while reading contents of input file for day 1 problem, %w", err)
	}

	utils.GetLogger().Println("Starting solver for day 1")
	output, err := solveDay1(lines)
	utils.GetLogger().Println("Complted solver for day 1")

	if err != nil {
		utils.GetLogger().Fatalf("Error when solving day 1 problem %w", err)
	}

	fmt.Println("Output for day 1 problem", output)
}

func solveDay1(commands []string) (int64, error) {
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
		if isLeftRotation {
			currentDialPosition -= turns % 100
			currentDialPosition += 100
			currentDialPosition %= 100
		} else {
			currentDialPosition += turns % 100
			currentDialPosition %= 100
		}

		// Check if dial is at zero
		if currentDialPosition == 0 {
			zeroesCounter++
		}
	}

	return zeroesCounter, nil
}
