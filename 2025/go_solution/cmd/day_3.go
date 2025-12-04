package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

// Input file path
// const day3InputFilePath = "day_3/sample_input"

const day3InputFilePath = "day_3/input.input"

func day3Solution() {
	// Read file contents
	lines, err := utils.ReadFileContents(day3InputFilePath)
	if err != nil {
		utils.GetLogger().Fatalf("Error while reading contents of input file for day 3 problem, %v", err)
	}

	utils.GetLogger().Println("Starting solver for day 3")
	output, err := solveDay3(lines)
	utils.GetLogger().Println("Complted solver for day 3")

	if err != nil {
		utils.GetLogger().Fatalf("Error when solving day 3 problem %w", err)
	}

	fmt.Println("Output for day 2 problem", output)
}

func solveDay3(banks []string) (int64, error) {
	// Iterate through each bank and find the max battery joltage
	totalJoltage := int64(0)

	for _, bank := range banks {
		// Part 1
		// totalJoltage += getMaxJoltage(bank, 2)

		// Part 2
		totalJoltage += getMaxJoltage(bank, 12)
	}

	return totalJoltage, nil
}

func getMaxJoltage(bank string, totalDigits int) int64 {
	// Base case: if Len is 0 or total digits is 0
	if len(bank) == 0 || totalDigits == 0 {
		return 0
	}

	// Get the largest digit first occurence
	largestDigit, largestDigitPos := getLargestDigitAndPos(bank[:len(bank)-totalDigits+1])

	// Recursively fetch the remaining digits.
	remainingMaxJoltage := getMaxJoltage(bank[largestDigitPos+1:], totalDigits-1)

	return largestDigit*int64(math.Pow10(totalDigits-1)) + remainingMaxJoltage

}

func getLargestDigitAndPos(bank string) (int64, int) {
	largestDigit := int64(-1)
	pos := -1
	for idx := 0; idx < len(bank); idx += 1 {
		digit, err := strconv.ParseInt(bank[idx:idx+1], 10, 64)
		if err != nil {
			utils.GetLogger().Fatalf("error while parsing digit: %s, index: %v", bank, idx)
		}

		if digit > largestDigit {
			largestDigit = digit
			pos = idx
		}
	}

	return largestDigit, pos
}
