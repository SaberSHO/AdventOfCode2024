package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a reader
	reader := bufio.NewReader(file)

	// Read the file line by line v
	var re = regexp.MustCompile("mul\\([{0-9}]*,[{0-9}]*\\)")

	var allMuls []string
	for {

		line, err := reader.ReadString('\n')

		rawMuls := re.FindAllString(line, -1)
		//fmt.Println(rawMuls)
		allMuls = append(allMuls, rawMuls...)
		//fmt.Println(allMuls)

		if err != nil {
			break // Reached end of file
		}
	}

	var redigits = regexp.MustCompile("[0-9]+")

	var sum int
	for _, s := range allMuls {
		nums := redigits.FindAllString(s, -1)
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		product := num1 * num2
		//fmt.Println(nums, product)
		sum = sum + product
	}

	fmt.Println("Sum is: ", sum)
}
