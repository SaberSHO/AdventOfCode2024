package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule map[int][]int

func (r Rule) compare(n, p int) int {
	if slices.Contains(r[n], p) {
		return -1
	}
	if slices.Contains(r[p], n) {
		return 1
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//parse rules first into a map
	rules := make(Rule)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		p, pErr := strconv.Atoi(line[:2])
		n, nErr := strconv.Atoi(line[3:])
		if pErr == nil && nErr == nil {
			//add to map, taking into account p could have multiple n
			curVal := rules[p]
			if curVal == nil {
				rules[p] = []int{n}
			} else {
				rules[p] = append(curVal, n)
			}
		}
	}

	//parse updates next into a slice of slices[int]
	var updates [][]int
	for scanner.Scan() {
		line := scanner.Text()
		pages := strings.Split(line, ",")
		var update []int
		for _, page := range pages {
			pageInt, _ := strconv.Atoi(page)
			update = append(update, pageInt)
		}

		updates = append(updates, update)
	}

	fmt.Println(rules)
	fmt.Println(updates)

	//do sorting magic
	var part1 int = 0
	var part2 int = 0
	for _, update := range updates {
		if slices.IsSortedFunc(update, rules.compare) {
			part1 += update[len(update)/2]
		} else {
			slices.SortFunc(update, rules.compare)
			//fmt.Println(update)
			part2 += update[len(update)/2]
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
