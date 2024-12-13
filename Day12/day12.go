package main

import (
	"bufio"
	"fmt"
	"os"
)

// Region represents a connected group of runes of the same type
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

// findRegions identifies all connected regions of the same rune in a matrix
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

		// Explore 4-directionally adjacent cells
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]

			if newR < 0 || newR >= rows || newC < 0 || newC >= cols || matrix[newR][newC] != targetRune {
				region.Perimeter++
			}

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
			if !adj.up && !adj.left {
				region.Sides++
			}

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

			//upR, upC := r-1, c
			//downR, downC := r+1, c
			//leftR, leftC := r, c-1
			//rightR, rightC := r, c+1
			//upleftR, upleftC := r-1, c-1
			//uprightR, uprightC := r-1, c+1
			//downleftR, downleftC := r+1, c-1
			//downrightR, downrightC := r+1, c+1

			//boundUp := !inBounds(matrix, upR, upC) || matrix[upR][upC] != targetRune
			//boundDown := !inBounds(matrix, downR, downC) || matrix[downR][downC] != targetRune
			//boundLeft := !inBounds(matrix, leftR, leftC) || matrix[leftR][leftC] != targetRune
			//boundRight := !inBounds(matrix, rightR, rightC) || matrix[rightR][rightC] != targetRune
			//boundUpLeft := !inBounds(matrix, upleftR, upleftC) || matrix[upleftR][upleftC] != targetRune
			//boundUpRight := !inBounds(matrix, uprightR, uprightC) || matrix[uprightR][uprightC] != targetRune
			//boundDownLeft := !inBounds(matrix, downleftR, downleftC) || matrix[downleftR][downleftC] != targetRune
			//boundDownRight := !inBounds(matrix, downrightR, downrightC) || matrix[downrightR][downrightC] != targetRune

			// if visited[upR][upC] {
			// 	if boundUp && boundLeft {
			// 		region.Sides++
			// 	} else if boundUpLeft {
			// 		region.Sides++
			// 	}
			// }
			// 	region.Sides++
			// } else if matrix[upR][upC] != targetRune && matrix[leftR][leftC] != targetRune {
			// 	region.Sides++
			// } else if inBounds(matrix, upleftR, upleftC) {
			// 	if matrix[upleftR][upleftC] != targetRune {
			// 		region.Sides++
			// 	}
			// }

			//check for corners
			//first check if its inbounds. if it isnt, its a corner.
			//check up and left, if they are different than current cell, its convex corner
			//if they are the same, and the up-left diag is different, its a concave corner
			// if matrix[r+upleft[1][0]][c+upleft[1][1]] != targetRune && matrix[r+upleft[2][0]][c+upleft[2][1]] != targetRune {
			// 	region.Sides++
			// } else if matrix[r+upleft[0][0]][c+upleft[0][1]] != targetRune {
			// 	region.Sides++
			// }

			// if matrix[r+upright[1][0]][c+upright[1][1]] != targetRune && matrix[r+upright[2][0]][c+upright[2][1]] != targetRune {
			// 	region.Sides++
			// } else if matrix[r+upright[0][0]][c+upright[0][1]] != targetRune {
			// 	region.Sides++
			// }

			// if matrix[r+downleft[1][0]][c+downleft[1][1]] != targetRune && matrix[r+downleft[2][0]][c+downleft[2][1]] != targetRune {
			// 	region.Sides++
			// } else if matrix[r+downleft[0][0]][c+downleft[0][1]] != targetRune {
			// 	region.Sides++
			// }

			// if matrix[r+downright[1][0]][c+downright[1][1]] != targetRune && matrix[r+downright[2][0]][c+downright[2][1]] != targetRune {
			// 	region.Sides++
			// } else if matrix[r+downright[0][0]][c+upleft[0][1]] != targetRune {
			// 	region.Sides++
			// }

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
		adj.up = true
	} else if matrix[r-1][c] == targetRune {
		adj.up = true
	}

	if !inBounds(matrix, r+1, c) {
		adj.down = true
	} else if matrix[r+1][c] == targetRune {
		adj.down = true
	}

	if !inBounds(matrix, r, c-1) {
		adj.left = true
	} else if matrix[r][c-1] == targetRune {
		adj.left = true
	}

	if !inBounds(matrix, r, c+1) {
		adj.right = true
	} else if matrix[r][c+1] == targetRune {
		adj.right = true
	}
	if !inBounds(matrix, r-1, c-1) {
		adj.upleft = true
	} else if matrix[r-1][c-1] == targetRune {
		adj.upleft = true
	}
	if !inBounds(matrix, r-1, c+1) {
		adj.upright = true
	} else if matrix[r-1][c+1] == targetRune {
		adj.upright = true
	}
	if !inBounds(matrix, r+1, c-1) {
		adj.downleft = true
	} else if matrix[r+1][c-1] == targetRune {
		adj.downleft = true
	}
	if !inBounds(matrix, r+1, c+1) {
		adj.downright = true
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
