package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calibration struct {
	testVal int
	nums    []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//parse rules first into a map
	var calibrations []Calibration
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		testVal, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		nums := strings.Split(parts[1], " ")
		calibration := Calibration{testVal, make([]int, 0, len(nums))}
		for _, part := range nums {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			calibration.nums = append(calibration.nums, num)
		}

		calibrations = append(calibrations, calibration)

	}
	fmt.Println(calibrations)

	ops := []rune{'+', '*'}
	validEquations := 0
	totalCalibrationResults := 0

	for i := range calibrations {
		if testSolve(calibrations[i], ops) == true {
			validEquations += 1
			totalCalibrationResults += calibrations[i].testVal
		}
	}
	fmt.Println("Part 1 Answers")
	fmt.Println("Number of valid equations:", validEquations)
	fmt.Println("Total Calibration Results:", totalCalibrationResults)

	ops2 := []rune{'+', '*', '|'}
	validEquations2 := 0
	totalCalibrationResults2 := 0

	for i := range calibrations {
		if testSolve(calibrations[i], ops2) == true {
			validEquations2 += 1
			totalCalibrationResults2 += calibrations[i].testVal
		}
	}
	fmt.Println("Part 2 Answers")
	fmt.Println("Number of valid equations:", validEquations2)
	fmt.Println("Total Calibration Results:", totalCalibrationResults2)
}

func testSolve(calibration Calibration, ops []rune) bool {

	//generate all combos of operations
	allCombinations := generateAllCombinations(ops, len(calibration.nums)-1)

	//test each combination of operations on the numbers and return true if it works.
	for _, combination := range allCombinations {
		result := calibration.nums[0]
		for i := 0; i < len(combination); i++ {
			switch combination[i] {
			case '+':
				result += calibration.nums[i+1]
			case '*':
				result *= calibration.nums[i+1]
			case '|':
				result = concatInts(result, calibration.nums[i+1])
			}
		}

		if result == calibration.testVal {
			//fmt.Println(calibration)
			return true
		}
	}

	return false
}

func generateAllCombinations(runes []rune, length int) [][]rune {
	// Store all combinations
	var combinations [][]rune

	// Helper function to recursively generate combinations
	var backtrack func(current []rune)

	backtrack = func(current []rune) {
		// If we've reached the desired length, add the current combination
		if len(current) == length {
			// Create a copy of the current combination to avoid pointer issues
			combination := make([]rune, len(current))
			copy(combination, current)
			combinations = append(combinations, combination)
			return
		}

		// Try adding each rune to the current combination
		for _, r := range runes {
			// Append the current rune
			current = append(current, r)

			// Recurse
			backtrack(current)

			// Backtrack by removing the last added rune
			current = current[:len(current)-1]
		}
	}

	// Start the backtracking process with an empty combination
	backtrack([]rune{})

	return combinations
}

func concatInts(a, b int) int {
	// Convert integers to strings
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	// Concatenate the strings
	concatenatedStr := aStr + bStr

	// Convert back to integer
	result, err := strconv.Atoi(concatenatedStr)
	if err != nil {
		// Handle potential overflow or conversion error
		panic(err)
	}

	return result
}
