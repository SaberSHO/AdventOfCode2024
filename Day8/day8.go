package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//parse rules first into a grid
	var grid [][]rune
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	//fmt.Println(grid)

	//now lets build a map of antennas of each type with points
	antennaMap := map[rune][]Point{}
	for iX := 0; iX < len(grid); iX++ {
		for iY := 0; iY < len(grid[0]); iY++ {
			current := grid[iX][iY]
			if current == '.' {
				continue
			}
			if value, exists := antennaMap[current]; exists {
				value = append(value, Point{iX, iY})
				antennaMap[current] = value
			} else {
				antennaMap[current] = []Point{{iX, iY}}
			}
		}
	}
	//fmt.Println(antennaMap)
	//fmt.Println(calcAntiNodes(Point{3, 4}, Point{5, 5}))

	//Find all combinations of 2 antennas of each type
	var antiNodes []Point
	for key := range antennaMap {
		//fmt.Println(generatePairs(antennaMap[key]))
		pairs := generatePairs(antennaMap[key])
		//calculate antiNodes of each combination and add them to a result slice
		for i := range pairs {
			antiNodes = append(antiNodes, calcAntiNodes(pairs[i][0], pairs[i][1], len(grid), len(grid[0]))...)
		}
		//fmt.Println(antiNodes)
	}

	//finally find unique antiNodes
	uniqueAntiNodes := removeDuplicateAntiNodes(antiNodes)
	fmt.Println(len(uniqueAntiNodes))

	var antiNodes2 []Point
	for key := range antennaMap {
		//fmt.Println(generatePairs(antennaMap[key]))
		pairs := generatePairs(antennaMap[key])
		//calculate antiNodes of each combination and add them to a result slice
		for i := range pairs {
			antiNodes2 = append(antiNodes2, calcAllAntiNodes(pairs[i][0], pairs[i][1], len(grid), len(grid[0]))...)
		}
		//fmt.Println(antiNodes)
	}

	//finally find unique antiNodes
	uniqueAntiNodes2 := removeDuplicateAntiNodes(antiNodes2)
	fmt.Println(len(uniqueAntiNodes2))
}

func calcAntiNodes(p1 Point, p2 Point, maxX int, maxY int) []Point {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	antiNodeP1 := Point{p1.X - dx, p1.Y - dy}
	antiNodeP2 := Point{p2.X + dx, p2.Y + dy}

	var result []Point

	if antiNodeP1.X >= 0 && antiNodeP1.X < maxX && antiNodeP1.Y >= 0 && antiNodeP1.Y < maxY {
		result = append(result, antiNodeP1)
	}

	if antiNodeP2.X >= 0 && antiNodeP2.X < maxX && antiNodeP2.Y >= 0 && antiNodeP2.Y < maxY {
		result = append(result, antiNodeP2)
	}

	return result
}

func calcAllAntiNodes(p1 Point, p2 Point, maxX int, maxY int) []Point {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	var result []Point
	var mul int = 0
	for {
		antiNodeP1 := Point{p1.X - mul*dx, p1.Y - mul*dy}

		if antiNodeP1.X >= 0 && antiNodeP1.X < maxX && antiNodeP1.Y >= 0 && antiNodeP1.Y < maxY {
			result = append(result, antiNodeP1)
			mul += 1
			continue
		}

		break
	}
	mul = 0
	for {
		antiNodeP2 := Point{p2.X + mul*dx, p2.Y + mul*dy}

		if antiNodeP2.X >= 0 && antiNodeP2.X < maxX && antiNodeP2.Y >= 0 && antiNodeP2.Y < maxY {
			result = append(result, antiNodeP2)
			mul += 1
			continue
		}

		break
	}

	return result
}

func generatePairs(points []Point) [][]Point {
	var pairs [][]Point
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pair := []Point{points[i], points[j]}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func removeDuplicateAntiNodes(slice []Point) []Point {
	seen := make(map[Point]bool)
	result := []Point{}

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
