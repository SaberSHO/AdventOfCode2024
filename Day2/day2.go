package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
//import "sort"
import "math"

func main() {
	var levels [][]int


	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a reader
	reader := bufio.NewReader(file)

	// Read the file line by line v
	var i = 0
	for {
	
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Reached end of file
		}

		lineSlice := strings.Fields(line)
		//fmt.Println(lineSlice)
		
		ints := make([]int, len(lineSlice))
		for i, s := range lineSlice {
            // convert string to int
            val, err := strconv.Atoi(s)
            if err != nil {
                panic(err)
            }

            // update the corresponding position in the
            // ints slice
            ints[i] = val
        }

        levels = append(levels,ints)
		i++
	}

	var ReportsSafe = 0
	var ReportsSafePart2 = 0

	for j := range levels {
		//check each level for safe/unsafe
		//First check all increasing or all decreasing
		fmt.Println(levels[j],"is increasing = ",isSafelyIncreasing(levels[j]))
		fmt.Println(levels[j],"is decreasing = ",isSafelyDecreasing(levels[j]))

		if (isSafelyDecreasing(levels[j]) || isSafelyIncreasing(levels[j])) {
			ReportsSafe += 1
		}
	}

	fmt.Println("Safe Reports (part 1): ", ReportsSafe)

	
	for j := range levels {
		//check each level for safe/unsafe
		//First check all increasing or all decreasing
		//fmt.Println(levels[j],"is increasing = ",isSafelyIncreasing(levels[j]))
		//fmt.Println(levels[j],"is decreasing = ",isSafelyDecreasing(levels[j]))

		if (isSafelyDecreasing(levels[j]) || isSafelyIncreasing(levels[j])) {
			ReportsSafePart2 += 1
		} else {
			//Now we try with Toleration
			var subTest []int
			for k := range levels[j] {
				subTest = RemoveIndex(levels[j],k)
				//fmt.Println(levels[j],k, subTest)
				if (isSafelyDecreasing(subTest) || isSafelyIncreasing(subTest)){
					ReportsSafePart2 +=1
					break
				}
			}
		}
	}

	fmt.Println("Safe Reports (part 2): ", ReportsSafePart2)
}

func isSafelyIncreasing(s []int) bool {
    for i := 1; i < len(s); i++ {
        if s[i] <= s[i-1]  || math.Abs(float64(s[i]-s[i-1])) >3 {
            return false
        }
    }
    return true
}

func isSafelyDecreasing(s []int) bool {
    for i := 1; i < len(s); i++ {
        if s[i] >= s[i-1] || math.Abs(float64(s[i]-s[i-1])) >3 {
            return false
        }

    }
    return true
}

func RemoveIndex(s []int, index int) []int {
    ret := make([]int, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}