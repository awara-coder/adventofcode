package main

import (
	"strconv"
	"strings"

	"github.com/awara-coder/adventofcode/tree/main/2025/go_solution/utils"
)

func parseDay6Part1Input(lines []string) ([][]int64, []string) {
	operands := make([][]int64, 0)
	for i := 0; i < len(lines)-1; i++ {
		operand, err := convertStringSliceToIntSlice(splitBySpace(lines[i]))
		if err != nil {
			utils.GetLogger().Fatalf("Error while parsing day 6 input for operand %v", i+1)
		}
		operands = append(operands, operand)
	}

	// Transpose operands so that operands[i] refers to operands for operators[i]
	operands = transpose2DSlice(operands)
	operators := splitBySpace(lines[len(lines)-1])
	return operands, operators
}

func transpose2DSlice[T int64 | float64 | string](input [][]T) [][]T {
	if len(input) == 0 {
		return [][]T{}
	}

	rows := len(input)
	cols := len(input[0])
	transpose := make([][]T, 0, cols)
	for col := range cols {
		transposedCol := make([]T, 0, rows)
		for row := range rows {
			transposedCol = append(transposedCol, input[row][col])
		}
		transpose = append(transpose, transposedCol)
	}

	return transpose
}

func parseDay6Part2Input(lines []string) ([][]int64, []string) {
	operands := make([][]string, 0)
	operators := make([]string, 0)
	// read each operand using by using operators line as references
	for pos, char := range lines[len(lines)-1] {
		if char != ' ' {
			// We have found an operator
			operators = append(operators, string(char))

			// Start reading from this position to the end, till all the previous lines at read ptr are empty spaces
			operandList := make([]string, len(lines)-1)
			col := pos
			for {
				// Go through each row from 0 ... lines - 2, and add data to operand list
				areAllCellsEmpty := true
				for row := 0; row < len(lines)-1; row++ {
					if col < len(lines[row]) && lines[row][col] != ' ' {
						areAllCellsEmpty = false
					}
					if col < len(lines[row]) {
						operandList[row] = operandList[row] + string(lines[row][col])
					} else {
						// Force empty space to equalize
						operandList[row] = operandList[row] + " "
					}
				}
				if areAllCellsEmpty {
					break
				}
				col++
			}

			if col != len(lines[0]) {
				// remove extra space from each operand
				for i := range len(operandList) {
					operandList[i] = operandList[i][:len(operandList[i])-1]
				}
			}

			// Now this contains, everything including spaces.
			operands = append(operands, operandList)
		}
	}

	parsedOperands := make([][]int64, 0, len(operands))
	for _, operandList := range operands {
		parsedOperands = append(parsedOperands, parseCephalopodMathNumbers(operandList))
	}
	return parsedOperands, operators
}

func parseCephalopodMathNumbers(numbers []string) []int64 {
	maxNumberOfDigits := 0
	for _, number := range numbers {
		maxNumberOfDigits = max(maxNumberOfDigits, len(number))
	}

	// Parse numbers from right to left
	operands := make([]int64, 0, maxNumberOfDigits)

	for j := range maxNumberOfDigits {
		currentNumber := int64(0)
		for i := range len(numbers) {
			if numbers[i][j] != ' ' {
				currentNumber *= 10
				currentNumber += int64(numbers[i][j] - '0')
			}
		}
		operands = append(operands, currentNumber)
	}

	return operands
}

// convertStringSliceToIntSlice converts string slice to integer slice
func convertStringSliceToIntSlice(strSlice []string) ([]int64, error) {
	intSlice := make([]int64, 0, len(strSlice))

	for _, strValue := range strSlice {
		intValue, parseErr := strconv.ParseInt(strValue, 10, 64)
		if parseErr != nil {
			utils.GetLogger().Println("Error while converting string to integer, ", strValue)
			return nil, parseErr
		}

		intSlice = append(intSlice, intValue)
	}

	return intSlice, nil
}

// splitBySpace splits by space and removes all the empty strings
func splitBySpace(line string) []string {
	splitString := strings.Split(line, " ")

	// Think of a better name
	finalResult := make([]string, 0)
	for _, str := range splitString {
		if str != "" {
			finalResult = append(finalResult, str)
		}
	}

	return finalResult

}

func solveDay6(lines []string) (int64, error) {
	// Parse the input.
	// operands, operators := parseDay6Part1Input(lines)
	operands, operators := parseDay6Part2Input(lines)

	// Assert input
	if len(operands) != len(operators) {
		utils.GetLogger().Fatal("operator and operand lenghts are not matching")
	}

	// Call solver function
	totalSum := solveDay6Expression(operands, operators)
	return totalSum, nil
}

func solveDay6Expression(operands [][]int64, operators []string) int64 {
	totalSum := int64(0)

	for i, operator := range operators {
		if operator == "+" {
			currentExpressionValue := int64(0)
			for _, operand := range operands[i] {
				currentExpressionValue += operand
			}
			totalSum += currentExpressionValue
		} else if operators[i] == "*" {
			currentExpressionValue := int64(1)
			for _, operand := range operands[i] {
				currentExpressionValue *= operand
			}
			totalSum += currentExpressionValue
		} else {
			utils.GetLogger().Fatalf("Unknown operator encountered: %v", operator)
		}
	}

	return totalSum
}
