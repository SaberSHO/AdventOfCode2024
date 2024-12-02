package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "sort"
import "math"

func main() {
	var list1 []int
	var list2 []int
	var distances []int
	var difference int = 0
	var sumDistances int = 0


	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a reader
	reader := bufio.NewReader(file)

	// Read the file line by line and split into 2 lists
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Reached end of file
		}
		//fmt.Println(line)

		lineSlice := strings.Fields(line)
		//fmt.Println(lineSlice)

		num1, err := strconv.Atoi(lineSlice[0])
		num2, err := strconv.Atoi(lineSlice[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	//fmt.Println(list1)
	//fmt.Println(list2)

	sort.Ints(list1)
	sort.Ints(list2)

	fmt.Println(list1)
	fmt.Println(list2)

	for i := range list1 {
		difference = int(math.Abs(float64(list1[i] - list2[i])))
		distances = append(distances,difference)
		sumDistances = sumDistances + difference
	//	fmt.Println(sumDistances)
	}
	//fmt.Println(distances)
	fmt.Println("\n",sumDistances)
}
