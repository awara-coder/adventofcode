package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

// Input file path
const day2InputFilePath = "day_2/input.input"

// const day2InputFilePath = "day_2/sample_input"

func day2Solution() {
	// Read file contents
	lines, err := utils.ReadFileContents(day2InputFilePath)
	if err != nil {
		utils.GetLogger().Fatalf("Error while reading contents of input file for day 2 problem, %v", err)
	}

	utils.GetLogger().Println("Starting solver for day 2")
	output, err := solveDay2(lines[0])
	utils.GetLogger().Println("Complted solver for day 2")

	if err != nil {
		utils.GetLogger().Fatalf("Error when solving day 2 problem %w", err)
	}

	fmt.Println("Output for day 2 problem: ", output)
}

func solveDay2(input string) (int64, error) {
	invalidIDSum := int64(0)

	ranges := strings.Split(input, ",")

	for _, currentRange := range ranges {
		// Parse the range
		parsedRange := strings.Split(currentRange, "-")
		if len(parsedRange) != 2 {
			utils.GetLogger().Fatalf("Length of parsed string is not 2, it's %v, String: %s", len(parsedRange), currentRange)
		}

		lowerRange, parseErr := strconv.ParseInt(parsedRange[0], 10, 64)
		if parseErr != nil {
			utils.GetLogger().Fatalf("Failed to parse string range, %v", parsedRange[0])

		}
		upperRange, parseErr := strconv.ParseInt(parsedRange[1], 10, 64)
		if parseErr != nil {
			utils.GetLogger().Fatalf("Failed to parse string range, %v", parsedRange[1])

		}

		for id := lowerRange; id <= upperRange; id += 1 {
			// Check if ID is repeated twice
			// if isRepeatedTwice(id) {
			// 	invalidIDSum += id
			// }

			// Part two solution
			if isRepeated(id) {
				utils.GetLogger().Println(id)
				invalidIDSum += id
			}
		}
	}

	return invalidIDSum, nil
}

func isRepeatedTwice(ID int64) bool {
	IDStr := strconv.FormatInt(ID, 10)
	if len(IDStr)%2 != 0 || len(IDStr) == 0 {
		return false
	}

	// check if first half is same as second half
	IDLen := len(IDStr)
	return IDStr[:IDLen/2] == IDStr[IDLen/2:]

}

func isRepeated(ID int64) bool {
	IDStr := strconv.FormatInt(ID, 10)
	IDLen := int64(len(IDStr))
	for divisor := int64(1); divisor*divisor <= IDLen; divisor += 1 {
		if IDLen%divisor == 0 {
			// check if string is a repeated sequence of length divisor
			if isRepeatedX(IDStr, int(divisor)) {
				return true
			}
			// check if string is a repeated sequence of id % length divisor
			if isRepeatedX(IDStr, int(IDLen/divisor)) {
				return true
			}
		}
	}

	return false
}

func isRepeatedX(ID string, sequenceLen int) bool {
	if len(ID) <= sequenceLen || len(ID)%sequenceLen != 0 {
		return false
	}
	baseSequence := ID[:sequenceLen]

	for startIdx := sequenceLen; startIdx < len(ID); startIdx += sequenceLen {
		if baseSequence != ID[startIdx:startIdx+sequenceLen] {
			return false
		}
	}

	return true
}
