package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

type Robot struct {
	PositionX int
	PositionY int
	VelocityX int
	VelocityY int
}

type Point struct {
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

	//parse robots
	var robots []Robot
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		line := scanner.Text()
		// Regex pattern explanation:
		// p=           : literally matches "p="
		// (-?\d+)     : captures one or more digits, optionally preceded by a minus sign
		// ,           : literally matches ","
		// (-?\d+)     : captures one or more digits, optionally preceded by a minus sign
		// \s+         : matches one or more whitespace characters
		// v=          : literally matches "v="
		// (-?\d+)     : captures one or more digits, optionally preceded by a minus sign
		// ,           : literally matches ","
		// (-?\d+)     : captures one or more digits, optionally preceded by a minus sign
		pattern := `p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`

		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(line)

		if len(matches) == 5 { // Full match plus 4 capture groups
			p1, _ := strconv.Atoi(matches[1])
			p2, _ := strconv.Atoi(matches[2])
			v1, _ := strconv.Atoi(matches[3])
			v2, _ := strconv.Atoi(matches[4])
			robot := Robot{PositionX: p1, PositionY: p2, VelocityX: v1, VelocityY: v2}
			robots = append(robots, robot)
			// fmt.Printf("Position: %s, %s\n", p1, p2)
			// fmt.Printf("Velocity: %s, %s\n", v1, v2)
		}

	}

	cols := 101
	rows := 103

	// fmt.Println(robots[10])
	// fmt.Println(robotPositionAtTime(robots[10], 1, cols, rows))
	// fmt.Println(robotPositionAtTime(robots[10], 2, cols, rows))
	// fmt.Println(robotPositionAtTime(robots[10], 3, cols, rows))
	// fmt.Println(robotPositionAtTime(robots[10], 4, cols, rows))
	// fmt.Println(robotPositionAtTime(robots[10], 5, cols, rows))

	//loop through robots at time, record what quadrant they are in in a map
	quadrantMap := make(map[string][]Point)
	time := 100
	for i, _ := range robots {
		point := robotPositionAtTime(robots[i], time, cols, rows)
		//fmt.Println(point)
		//NE Qudrant x < cols/2, y < rows/2
		if point.X < cols/2 && point.Y < rows/2 {
			quadrantMap["NE"] = append(quadrantMap["NE"], point)
		}
		//NW Quadrant
		if point.X > cols/2 && point.Y < rows/2 {
			quadrantMap["NW"] = append(quadrantMap["NW"], point)
		}
		//SW Quadrant
		if point.X > cols/2 && point.Y > rows/2 {
			quadrantMap["SW"] = append(quadrantMap["SW"], point)
		}
		//SE Quadrant
		if point.X < cols/2 && point.Y > rows/2 {
			quadrantMap["SE"] = append(quadrantMap["SE"], point)
		}
	}

	fmt.Println("NorthEast has", len(quadrantMap["NE"]))
	fmt.Println("NorthWest has", len(quadrantMap["NW"]))
	fmt.Println("SouthEast has", len(quadrantMap["SE"]))
	fmt.Println("SouthWest has", len(quadrantMap["SW"]))
	safetyFactor := len(quadrantMap["NE"]) * len(quadrantMap["NW"]) * len(quadrantMap["SW"]) * len(quadrantMap["SE"])

	fmt.Println("Safety Factor of ", safetyFactor)

	encoding.Register()
	scn, err := tcell.NewScreen()
	scn.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	scn.Clear()

	// first i did this starting at 0 and incrementing by 1. this had some patterns visible after looking for a while, but not tree
	// this is a magic number specific to my input (pretty sure). this is the first appearance of "clustering" of robots.
	// not yet quite a tree. but i also found clustering at 114 and 215 (114+101).
	// there is a different clustering at 182 285 (182+103).
	// the answer is the first time these clusterings converge
	// how we could figure out this seed number i dont know
	time = 114 // or 182
	for {
		scn.Show()
		ev := scn.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'n' {
				scn.Clear()
				for i, _ := range robots {
					point := robotPositionAtTime(robots[i], time, cols, rows)
					scn.SetContent(point.X, point.Y, rune('#'), []rune(""), tcell.StyleDefault)
				}
				drawText(scn, 0, 0, 6, 0, tcell.StyleDefault, strconv.Itoa(time))
				time += 101 // or 103
			}
			if ev.Rune() == 'p' {
				scn.Clear()
				time -= 101 // or 101
				for i, _ := range robots {
					point := robotPositionAtTime(robots[i], time, cols, rows)
					scn.SetContent(point.X, point.Y, rune('#'), []rune(""), tcell.StyleDefault)
				}
				drawText(scn, 0, 0, 6, 0, tcell.StyleDefault, strconv.Itoa(time))
			}
			if ev.Rune() == 'x' {
				scn.Fini()
				fmt.Println(time)
				os.Exit(0)
			}
		}
	}
}

func robotPositionAtTime(robot Robot, time int, cols int, rows int) Point {
	newX := (((robot.PositionX + (time * robot.VelocityX)) % cols) + cols) % cols
	newY := (((robot.PositionY + (time * robot.VelocityY)) % rows) + rows) % rows

	return Point{X: newX, Y: newY}

}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
