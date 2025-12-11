package main

import (
	"fmt"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/aoc"
	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

const SAMPLE_INPUT_FILE_NAME = "sample_input"
const CHALLENGE_INPUT_FILE_NAME = "input.input"

type DailyChallengeSolverFunction func([]string) (int64, error)

var daySolverFunctionMapping map[int]DailyChallengeSolverFunction = map[int]DailyChallengeSolverFunction{
	1: aoc.SolveDay1,
	2: aoc.SolveDay2,
	3: aoc.SolveDay3,
	4: aoc.SolveDay4,
	5: aoc.SolveDay5,
	6: aoc.SolveDay6,
	7: aoc.SolveDay7,
	8: aoc.SolveDay8,
}

func main() {
	fmt.Println("Advent of code 2025 Go solutions!")

	// Call solver functions
	dayNSolution(8, CHALLENGE_INPUT_FILE_NAME)
}

func dayNSolution(day int, fileName string) {
	if _, ok := daySolverFunctionMapping[day]; !ok {
		utils.GetLogger().Fatalf("incorrect value of day for dayNSolution: %v", day)
	}
	// Read file contents
	inputFilePath := fmt.Sprintf("day_%v/%s", day, fileName)
	lines, err := utils.ReadFileContents(inputFilePath)
	if err != nil {
		utils.GetLogger().Fatalf("Error while reading contents of input file for day %v problem, %v", day, err)
	}

	utils.GetLogger().Printf("Starting solver for day %v\n", day)
	output, err := daySolverFunctionMapping[day](lines)
	utils.GetLogger().Printf("Complted solver for day %v\n", day)

	if err != nil {
		utils.GetLogger().Fatalf("Error when solving day %v problem %v", day, err)
	}

	fmt.Printf("Output for day %v problem: %v\n", day, output)
}
