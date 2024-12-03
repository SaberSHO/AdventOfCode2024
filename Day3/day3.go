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

	sum := sumMuls(allMuls)

	fmt.Println("Sum is (part 1): ", sum)

	// Read the file line by line v
	file2, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	var re2 = regexp.MustCompile("mul\\([{0-9}]*,[{0-9}]*\\)|do\\(\\)|don't\\(\\)")
	reader2 := bufio.NewReader(file2)
	var allInst []string
	for {

		line2, err := reader2.ReadString('\n')

		rawInst := re2.FindAllString(line2, -1)
		//fmt.Println(rawMuls)
		allInst = append(allInst, rawInst...)
		//fmt.Println(allMuls)

		if err != nil {
			break // Reached end of file
		}
	}

	//loop through all Inst and create new list to multiply
	doOrDoNotThereIsNoTry := true
	var finalMuls []string
	for _, s := range allInst {
		//fmt.Println(s)

		switch s {
		case "do()":
			doOrDoNotThereIsNoTry = true
		case "don't()":
			doOrDoNotThereIsNoTry = false
		default:
			if doOrDoNotThereIsNoTry {
				finalMuls = append(finalMuls, s)
			}
		}

		//fmt.Println(finalMuls)
	}

	sum2 := sumMuls(finalMuls)
	fmt.Println("The sum is (part2):", sum2)
}

func sumMuls(allMuls []string) int {
	var sum int
	var redigits = regexp.MustCompile("[0-9]+")

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
	return sum
}
