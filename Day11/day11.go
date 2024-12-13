package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//parse input
	var line string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line = scanner.Text()
	}
	var initialStones []int
	splitLine := strings.Fields(line)
	for _, str := range splitLine {
		stone, _ := strconv.Atoi(str)
		initialStones = append(initialStones, stone)
	}
	fmt.Println(initialStones)
	//Part 1
	start := time.Now()
	part1Counts := blink(initialStones, 25)
	fmt.Println("Part 1 Stone Count (25 iterations): ", sumCounts(part1Counts))
	fmt.Println("Exec Time:", time.Since(start))

	//Part 2
	start2 := time.Now()
	part2Counts := blink(initialStones, 75)
	fmt.Println("Part 2 Stone Count(75 iterations): ", sumCounts(part2Counts))
	fmt.Println("Exec Time:", time.Since(start2))

	//Pushing the limit for funsies
	start3 := time.Now()
	part3Counts := blink(initialStones, 5000)
	fmt.Println("Pushing the Limit Stone Count (5000 iterations): ", sumCounts(part3Counts))
	fmt.Println("Exec Time:", time.Since(start3))
}

func blink(stones []int, iterations int) map[int]int {
	stoneCountMap := make(map[int]int)
	for _, stone := range stones {
		stoneCountMap[stone] += 1
	}

	for i := 0; i < iterations; i += 1 {
		nextStoneCountMap := make(map[int]int)
		for stone, count := range stoneCountMap {
			if stone == 0 {
				nextStoneCountMap[1] += count
			} else {
				stringStone := strconv.Itoa(stone)
				if len(stringStone)%2 == 0 {
					firstHalf, _ := strconv.Atoi(stringStone[:len(stringStone)/2])
					secondHalf, _ := strconv.Atoi(stringStone[len(stringStone)/2:])
					nextStoneCountMap[firstHalf] += count
					nextStoneCountMap[secondHalf] += count
				} else {
					nextStoneCountMap[stone*2024] += count
				}
			}

		}
		stoneCountMap = nextStoneCountMap
	}

	return stoneCountMap
}

func sumCounts(mapToCount map[int]int) int {
	totalCount := 0
	for _, count := range mapToCount {
		totalCount += count
	}
	return totalCount
}
