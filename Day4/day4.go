package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a reader
	reader := bufio.NewReader(file)

	var wordSearchGrid []string
	for {

		line, err := reader.ReadString('\n')
		wordSearchGrid = append(wordSearchGrid, line)
		if err != nil {
			break // Reached end of file
		}
	}

	fmt.Println(wordSearchGrid)
	//fmt.Println(wordSearchGrid[1][2])

	maxRow := len(wordSearchGrid) - 1
	maxCol := len(wordSearchGrid[0]) - 1

	//fmt.Println(maxRow)
	//fmt.Println(maxCol)

	count := 0
	for i, searchLine := range wordSearchGrid {
		for j, searchChar := range searchLine {
			//first start at X
			if searchChar == 'X' {
				// check E
				if j+3 < maxCol && searchLine[j+1:j+4] == "MAS" {
					//fmt.Println("found E")
					count += 1
				}
				// check W
				if j-3 >= 0 && searchLine[j-3:j] == "SAM" {
					count += 1
					//fmt.Println("found W")
				}
				// Northward checks. first check in bound, then check N, NE, NW
				if i-3 >= 0 {
					// check N
					if wordSearchGrid[i-3][j] == 'S' && wordSearchGrid[i-2][j] == 'A' && wordSearchGrid[i-1][j] == 'M' {
						count += 1
						//fmt.Println("found N")
					}
					//check NE
					if j-3 >= 0 && wordSearchGrid[i-3][j-3] == 'S' && wordSearchGrid[i-2][j-2] == 'A' && wordSearchGrid[i-1][j-1] == 'M' {
						count += 1
						//fmt.Println("found NE")
					}
					//check NW
					if j+3 < maxCol && wordSearchGrid[i-3][j+3] == 'S' && wordSearchGrid[i-2][j+2] == 'A' && wordSearchGrid[i-1][j+1] == 'M' {
						count += 1
						//fmt.Println("found NW")
					}
				}
				// Southward checks. first check in bound, then check S, SE, SW
				if i+3 < maxRow {
					// check S
					if wordSearchGrid[i+3][j] == 'S' && wordSearchGrid[i+2][j] == 'A' && wordSearchGrid[i+1][j] == 'M' {
						count += 1
						//fmt.Println("found S")
					}
					//check SE
					if j-3 >= 0 && wordSearchGrid[i+3][j-3] == 'S' && wordSearchGrid[i+2][j-2] == 'A' && wordSearchGrid[i+1][j-1] == 'M' {
						count += 1
						//fmt.Println("found SE")
					}
					//check SW
					if j+3 < maxRow && wordSearchGrid[i+3][j+3] == 'S' && wordSearchGrid[i+2][j+2] == 'A' && wordSearchGrid[i+1][j+1] == 'M' {
						count += 1
						//fmt.Println("found SW")
					}
				}

			}

		}
	}
	fmt.Println(count)

	//part 2
	// valid is		M.M		M.S		S.S		S.M
	//				.A.		.A.		.A.		.A.
	// 				S.S		M.S		M.M		S.M

	countP2 := 0
	for i, searchLine := range wordSearchGrid {
		for j, searchChar := range searchLine {
			//first start at A
			if searchChar == 'A' {
				// check OOB
				if i-1 >= 0 && i+1 < maxCol && j-1 >= 0 && j+1 < maxRow {
					if wordSearchGrid[i-1][j-1] == 'M' && wordSearchGrid[i-1][j+1] == 'M' && wordSearchGrid[i+1][j-1] == 'S' && wordSearchGrid[i+1][j+1] == 'S' {
						countP2 += 1
					}
					if wordSearchGrid[i-1][j-1] == 'M' && wordSearchGrid[i-1][j+1] == 'S' && wordSearchGrid[i+1][j-1] == 'M' && wordSearchGrid[i+1][j+1] == 'S' {
						countP2 += 1
					}
					if wordSearchGrid[i-1][j-1] == 'S' && wordSearchGrid[i-1][j+1] == 'S' && wordSearchGrid[i+1][j-1] == 'M' && wordSearchGrid[i+1][j+1] == 'M' {
						countP2 += 1
					}
					if wordSearchGrid[i-1][j-1] == 'S' && wordSearchGrid[i-1][j+1] == 'M' && wordSearchGrid[i+1][j-1] == 'S' && wordSearchGrid[i+1][j+1] == 'M' {
						countP2 += 1
					}
				}
			}
		}
	}

	fmt.Println(countP2)

}
