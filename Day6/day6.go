package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Direction struct {
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

	//parse rules first into a map
	var grid [][]string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	//set up our pathing Points/Directions
	var curPos Point

	curDir := Direction{0, -1}

	//find starting point
	curPos = findCharIndex(grid, "^")
	startingPos := curPos

	var allVisited []Point
	for {
		tryNewPos := moveInDirection(curPos, curDir)
		if tryNewPos.X < 0 || tryNewPos.Y < 0 || tryNewPos.X >= len(grid[0]) || tryNewPos.Y >= len(grid) {
			allVisited = append(allVisited, curPos)
			break
		}
		if grid[tryNewPos.Y][tryNewPos.X] == "#" {
			curDir = nextDir(curDir)
			continue
		}
		allVisited = append(allVisited, curPos)
		curPos = moveInDirection(curPos, curDir)

	}

	//we only want unique points
	allVisitedUnique := []Point{}
	visited := map[Point]bool{}
	for _, v := range allVisited {
		if !visited[v] {
			visited[v] = true
			allVisitedUnique = append(allVisitedUnique, v)
		}
	}

	fmt.Println("Visited: ", len(allVisitedUnique))

	//Part 2
	//Take the visited path from above. Add an obstable somewhere along the path. run through the visit algo again, but this time check if tryNewPos is already visited in the same dir. if it is, we are looping. break and increment
	looping := 0
	for visit := range allVisitedUnique {
		//ignore starting position
		if allVisitedUnique[visit] == startingPos {
			continue
		}
		tempGrid := copyStringMatrix(grid)
		//fmt.Println(allVisitedUnique[visit])
		tempGrid[allVisitedUnique[visit].Y][allVisitedUnique[visit].X] = "O"
		//fmt.Println(tempGrid)
		curPos = startingPos
		curDir := Direction{0, -1}
		pointAndDirectionVisited := map[Point][]Direction{}
		for {
			tryNewPos := moveInDirection(curPos, curDir)
			if tryNewPos.X < 0 || tryNewPos.Y < 0 || tryNewPos.X >= len(grid[0]) || tryNewPos.Y >= len(grid) {
				break
			}
			if tempGrid[tryNewPos.Y][tryNewPos.X] == "O" {
				curDir = nextDir(curDir)
				continue
			}
			if tempGrid[tryNewPos.Y][tryNewPos.X] == "#" {
				curDir = nextDir(curDir)
				continue
			}

			if containsValue(pointAndDirectionVisited, curPos, curDir) {
				looping += 1
				break
			}
			pointAndDirectionVisited[curPos] = append(pointAndDirectionVisited[curPos], curDir)
			curPos = moveInDirection(curPos, curDir)
		}

	}
	fmt.Println(looping)

}

func findCharIndex(matrix [][]string, char string) Point {
	for row, slice := range matrix {
		for col, val := range slice {
			if val == char {
				p := Point{col, row}
				return p
			}
		}
	}
	return Point{0, 0} // Return -1 if character not found
}

func moveInDirection(current Point, direction Direction) Point {
	return Point{current.X + direction.X, current.Y + direction.Y}
}

func nextDir(currentDirection Direction) Direction {
	up := Direction{0, -1}
	down := Direction{0, 1}
	left := Direction{-1, 0}
	right := Direction{1, 0}

	if currentDirection == up {
		return right
	}
	if currentDirection == right {
		return down
	}
	if currentDirection == down {
		return left
	}
	if currentDirection == left {
		return up
	}
	return Direction{0, 0}
}

func copyStringMatrix(matrix [][]string) [][]string {
	newMatrix := make([][]string, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]string, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}

func containsValue(m map[Point][]Direction, key Point, value Direction) bool {
	list, ok := m[key]
	if !ok {
		return false // Key not found in the map
	}

	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}
