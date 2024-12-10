package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		lineSliceString := strings.Split(line, "")
		lineSliceInt := make([]int, len(lineSliceString))
		for i, strVal := range lineSliceString {
			intVal, _ := strconv.Atoi(strVal)
			lineSliceInt[i] = intVal
		}

		matrix = append(matrix, lineSliceInt)
	}
	//fmt.Println(matrix)
	trailHeads := findPointsInMatrix(matrix, 0)
	summits := findPointsInMatrix(matrix, 9)
	paths := findPaths(matrix, trailHeads)

	fmt.Println("Total Score: ", calculateScore(paths, trailHeads, summits))
	fmt.Println("Total Rating: ", calculateRating(paths, trailHeads, summits))
}
func findPointsInMatrix(matrix [][]int, value int) []Point {
	rows, cols := len(matrix), len(matrix[0])
	var points []Point

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == value {
				points = append(points, Point{row, col})
			}
		}
	}
	//fmt.Println(points)
	return points
}

func findPaths(matrix [][]int, trailheads []Point) [][]Point {
	var allPaths [][]Point

	//find starting points

	directions := []Point{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}

	// DFS function
	var dfs func(current Point, path []Point)
	dfs = func(current Point, path []Point) {
		//if the current point is 9 we have reached the end
		if matrix[current.x][current.y] == 9 {
			completePath := make([]Point, len(path))
			copy(completePath, path)
			allPaths = append(allPaths, completePath)
			return
		}

		//try all 4 directions
		for _, dir := range directions {
			next := Point{x: current.x + dir.x, y: current.y + dir.y}

			//check if the next point is valid
			if isValid(matrix, path, current, next) {
				path = append(path, next)
				dfs(next, path)
				path = path[:len(path)-1] //backtrack
			}
		}

	}

	for _, start := range trailheads {
		dfs(start, []Point{start})
	}

	return allPaths
}

func isValid(matrix [][]int, path []Point, current Point, next Point) bool {
	// a valid point is in bounds and must be increment of 1 and must be new or else we get stuck in a loop)
	rows, cols := len(matrix), len(matrix[0])

	// check in bounds
	if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= cols {
		return false
	}

	// check if already in path
	for _, p := range path {
		if p.x == next.x && p.y == next.y {
			return false
		}
	}

	//check value increment is 1
	return matrix[next.x][next.y] == matrix[current.x][current.y]+1

}

func calculateScore(paths [][]Point, trailHeads []Point, summits []Point) int {
	//create a map of key trailhead value summits[]
	scores := make(map[Point][]Point)
	//for each trailhead, find the unique summits in the map
	for _, trailhead := range trailHeads {
		var allSummitsInPath []Point
		for _, path := range paths {
			if path[0] == trailhead {
				allSummitsInPath = append(allSummitsInPath, path[len(path)-1])
			}
		}
		scores[trailhead] = removeDuplicatePoints(allSummitsInPath)
	}

	totalScore := 0
	for i := range scores {
		totalScore += len(scores[i])
	}
	return totalScore
}

func calculateRating(paths [][]Point, trailHeads []Point, summits []Point) int {
	//create a map of key trailhead value summits[]
	scores := make(map[Point][]Point)
	//for each trailhead, find the unique summits in the map
	for _, trailhead := range trailHeads {
		var allSummitsInPath []Point
		for _, path := range paths {
			if path[0] == trailhead {
				allSummitsInPath = append(allSummitsInPath, path[len(path)-1])
			}
		}
		scores[trailhead] = allSummitsInPath
	}

	totalScore := 0
	for i := range scores {
		totalScore += len(scores[i])
	}
	return totalScore
}

func removeDuplicatePoints(slice []Point) []Point {
	// Create a map to track unique elements
	seen := make(map[Point]bool)
	result := []Point{}

	// Iterate over the slice
	for _, value := range slice {
		// Check if the element is already seen
		if !seen[value] {
			// Add the element to the result slice
			result = append(result, value)
			// Mark the element as seen
			seen[value] = true
		}
	}
	return result
}
