package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//parse input
	var line string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line = scanner.Text()
	}

	// PART 1
	fmt.Println("Part1")
	//build diskmap (Slice of Int). Represent free space with -1 to keep it all ints.
	diskMap := buildDiskMap(line)
	//fmt.Println(diskMap)

	//compact the Disk Map
	compactDiskMapPart1(diskMap)
	//fmt.Println(diskMap)

	//calculate checksum
	checksum := calculateCheckSum(diskMap)
	fmt.Println(checksum)

	// PART 2
	fmt.Println("Part2")
	//build diskmap (Slice of Int). Represent free space with -1 to keep it all ints.
	diskMap2 := buildDiskMap(line)
	//fmt.Println(diskMap2)

	//compact the Disk Map
	compactDiskMapPart2(diskMap2)
	//fmt.Println(diskMap2)

	//calculate checksum
	checksum2 := calculateCheckSum(diskMap2)
	fmt.Println(checksum2)
}

func buildDiskMap(line string) []int {
	var diskMap []int
	var fileValue int
	for i, value := range line {
		numOfValue, _ := strconv.Atoi(string(value))
		for range numOfValue {
			if i%2 == 0 {
				diskMap = append(diskMap, fileValue)
			} else {
				diskMap = append(diskMap, -1)
			}
		}
		if i%2 == 0 {
			fileValue += 1
		}
	}
	return diskMap
}

func compactDiskMapPart1(diskMap []int) []int {
	firstFree := 0
	lastNonFree := len(diskMap) - 1
	for firstFree < lastNonFree {
		if diskMap[firstFree] == -1 {
			//if we are at a free space, do swap
			for diskMap[lastNonFree] == -1 {
				//keep going until we find last non-free
				lastNonFree -= 1
			}
			if lastNonFree < firstFree {
				break
			}

			// do swap
			diskMap[firstFree] = diskMap[lastNonFree]
			diskMap[lastNonFree] = -1
		}
		firstFree += 1
	}
	return diskMap
}

func compactDiskMapPart2(diskMap []int) []int {
	//Find last file "block" (start at highest num, decrease each time). We need the start index, end index and size
	//move it if we have a free space block that is big enough, if not, leave it
	for i := len(diskMap) - 1; i >= 0; {
		start, _, size, value := findFileBlock(diskMap, i)
		emptyBlockIndex := findEmptyBlock(diskMap, size, start)
		//fmt.Println(start, end, size, value)
		if emptyBlockIndex != -1 {
			//move the block
			for s := range size {
				diskMap[emptyBlockIndex+s] = value
				diskMap[start+s] = -1
			}
		}
		//fmt.Println(diskMap)
		i = start - 1
	}

	return diskMap
}

func findFileBlock(diskMap []int, index int) (int, int, int, int) {
	//first go back from index to next non "empty" index
	for index > 0 && diskMap[index] == -1 {
		index -= 1
	}
	rightMost := diskMap[index]
	i := index
	for i >= 0 && diskMap[i] == rightMost {
		i -= 1
	}

	return i + 1, index, index - i, rightMost

}

func findEmptyBlock(diskMap []int, size int, limit int) int {
	i := 0
	j := 0

	for j-i < size && i < limit {
		for diskMap[i] != -1 && i < limit {
			i += 1
		}
		if i >= limit || j >= limit {
			break
		}
		j = i
		for diskMap[j] == -1 && j < limit {
			j += 1
		}
		if j-i >= size {
			return i
		}
		i = j
	}
	return -1
}

func calculateCheckSum(diskMap []int) int {
	var checksum int = 0
	for i, fileValue := range diskMap {
		if fileValue != -1 {
			checksum += i * fileValue
		}
	}
	return checksum
}
