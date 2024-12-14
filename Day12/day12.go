package main

import (
	"bufio"
	"fmt"
	"os"
)

type Region struct {
	Rune      rune
	Cells     [][]int
	Area      int
	Perimeter int
	Sides     int
}

type Adjancies struct {
	up        bool
	down      bool
	left      bool
	right     bool
	upleft    bool
	downleft  bool
	downright bool
	upright   bool
}

func findRegions(matrix [][]rune) []Region {
	if len(matrix) == 0 {
		return nil
	}

	rows, cols := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var regions []Region
	// Depth-first search to find connected regions
	var dfs func(int, int, rune, *Region)
	dfs = func(r, c int, targetRune rune, region *Region) {
		// Check bounds and visited status
		if r < 0 || r >= rows || c < 0 || c >= cols ||
			visited[r][c] || matrix[r][c] != targetRune {
			return
		}

		// Mark as visited and add to region
		visited[r][c] = true
		region.Cells = append(region.Cells, []int{r, c})

		//Look for corners
		adj := buildAdjancies(matrix, r, c, targetRune)

		//check outer corners
		if !adj.up && !adj.left {
			region.Sides++
		}
		if !adj.up && !adj.right {
			region.Sides++
		}
		if !adj.down && !adj.left {
			region.Sides++
		}
		if !adj.down && !adj.right {
			region.Sides++
		}
		//check inner corners
		if adj.up && adj.left && !adj.upleft {
			region.Sides++
		}
		if adj.up && adj.right && !adj.upright {
			region.Sides++
		}
		if adj.down && adj.left && !adj.downleft {
			region.Sides++
		}
		if adj.down && adj.right && !adj.downright {
			region.Sides++
		}

		// Explore 4-directionally adjacent cells
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]

			if newR < 0 || newR >= rows || newC < 0 || newC >= cols || matrix[newR][newC] != targetRune {
				region.Perimeter++
			}

			dfs(newR, newC, targetRune, region)
		}
	}

	// Iterate through each cell to find regions
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] {
				region := Region{Rune: matrix[r][c]}
				dfs(r, c, matrix[r][c], &region)

				// Only add regions with at least one cell
				if len(region.Cells) > 0 {
					region.Area = len(region.Cells)
					regions = append(regions, region)
				}
			}
		}
	}

	return regions
}

func buildAdjancies(matrix [][]rune, r int, c int, targetRune rune) Adjancies {
	var adj Adjancies
	if !inBounds(matrix, r-1, c) {
		adj.up = false
	} else if matrix[r-1][c] == targetRune {
		adj.up = true
	}

	if !inBounds(matrix, r+1, c) {
		adj.down = false
	} else if matrix[r+1][c] == targetRune {
		adj.down = true
	}

	if !inBounds(matrix, r, c-1) {
		adj.left = false
	} else if matrix[r][c-1] == targetRune {
		adj.left = true
	}

	if !inBounds(matrix, r, c+1) {
		adj.right = false
	} else if matrix[r][c+1] == targetRune {
		adj.right = true
	}
	if !inBounds(matrix, r-1, c-1) {
		adj.upleft = false
	} else if matrix[r-1][c-1] == targetRune {
		adj.upleft = true
	}
	if !inBounds(matrix, r-1, c+1) {
		adj.upright = false
	} else if matrix[r-1][c+1] == targetRune {
		adj.upright = true
	}
	if !inBounds(matrix, r+1, c-1) {
		adj.downleft = false
	} else if matrix[r+1][c-1] == targetRune {
		adj.downleft = true
	}
	if !inBounds(matrix, r+1, c+1) {
		adj.downright = false
	} else if matrix[r+1][c+1] == targetRune {
		adj.downright = true
	}

	return adj
}
func inBounds(matrix [][]rune, r int, c int) bool {
	rows, cols := len(matrix), len(matrix[0])
	inbound := r >= 0 && r < rows && c >= 0 && c < cols
	return inbound
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]rune
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	regions := findRegions(matrix)

	var totalCostPart1 = 0
	var totalCostPart2 = 0
	// Go through each region and calc total cost
	for i, region := range regions {
		fmt.Printf("Region %d:\n", i+1)
		fmt.Printf("  Rune: %c\n", region.Rune)
		fmt.Printf("  Cells: %v\n", region.Cells)
		fmt.Printf("  Area: %v\n", region.Area)
		fmt.Printf("  Perimeter: %v\n", region.Perimeter)
		fmt.Printf("  Sides: %v\n", region.Sides)
		totalCostPart1 += region.Area * region.Perimeter
		totalCostPart2 += region.Area * region.Sides
	}

	fmt.Println("Total Cost (Part1): ", totalCostPart1)
	fmt.Println("Total Cost (Part2): ", totalCostPart2)

}
