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
		}
		allVisited = append(allVisited, curPos)
		curPos = moveInDirection(curPos, curDir)

	}

	//we only want unique points
	visited := map[Point]bool{}
	for _, v := range allVisited {
		visited[v] = true
	}

	fmt.Println("Visited: ", len(visited))

	//Part 2
	//Take the visited path from above. Add an obstable somewhere along the path. run through the visit algo again, but this time check if tryNewPos is already visited. if it is, we are looping. break and increment

	for visit := range visited {
		//ignore starting position
		if visit == startingPos {
			continue
		}
		tempGrid := make([][]string, len(grid))
		copy(tempGrid, grid)
		tempGrid[visit.Y][visit.X] = "O"
		fmt.Println(tempGrid)

	}
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
