package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

type Problem struct {
	ax      int
	ay      int
	bx      int
	by      int
	targetX int
	targetY int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var problems []Problem
	reg_val := regexp.MustCompile(`\+(\d+)`)
	reg_eq := regexp.MustCompile(`=(\d+)`)

	i := 0
	var problemTemp Problem
	for {
		line, err := reader.ReadString('\n')

		i++
		if i%4 == 1 {
			values := reg_val.FindAllStringSubmatch(line, -1)
			problemTemp.ax, _ = strconv.Atoi(values[0][1])
			problemTemp.ay, _ = strconv.Atoi(values[1][1])
		}
		if i%4 == 2 {
			values := reg_val.FindAllStringSubmatch(line, -1)
			problemTemp.bx, _ = strconv.Atoi(values[0][1])
			problemTemp.by, _ = strconv.Atoi(values[1][1])
		}
		if i%4 == 3 {
			values := reg_eq.FindAllStringSubmatch(line, -1)
			problemTemp.targetX, _ = strconv.Atoi(values[0][1])
			problemTemp.targetY, _ = strconv.Atoi(values[1][1])
		}
		if i%4 == 0 {
			problems = append(problems, problemTemp)
			problemTemp = Problem{ax: 0, ay: 0, bx: 0, by: 0, targetX: 0, targetY: 0}
		}
		//hacky bullshit
		if err != nil {
			problems = append(problems, problemTemp)
			problemTemp = Problem{ax: 0, ay: 0, bx: 0, by: 0, targetX: 0, targetY: 0}
			break // Reached end of file
		}
	}

	//fmt.Println(problems)
	totalCost := 0
	for i, problem := range problems {
		solutionX, solutionY := solveLinearEquation(problem, false)
		if solutionX == -1 || solutionY == -1 {
			fmt.Println("Problem ", i+1, " is not solvable")
		} else {
			fmt.Println("Problem ", i+1, " has solution x:", solutionX, " y:", solutionY, " with cost: ", solutionX*3+solutionY)
			totalCost += solutionX*3 + solutionY
		}
	}
	fmt.Println("Part 1: The minimum cost is", totalCost)

	totalCost2 := 0
	for i, problem := range problems {
		solutionX, solutionY := solveLinearEquation(problem, true)
		if solutionX == -1 || solutionY == -1 {
			fmt.Println("Problem ", i+1, " is not solvable")
		} else {
			fmt.Println("Problem ", i+1, " has solution x:", solutionX, " y:", solutionY, " with cost: ", solutionX*3+solutionY)
			totalCost2 += solutionX*3 + solutionY
		}
	}
	fmt.Println("Part 2: The minimum cost is", totalCost2)
}

func solveLinearEquation(problem Problem, part2 bool) (int, int) {
	xy := mat.NewDense(2, 2, []float64{
		float64(problem.ax), float64(problem.bx),
		float64(problem.ay), float64(problem.by),
	})
	targetX := problem.targetX
	targetY := problem.targetY
	if part2 {
		targetX += 10000000000000
		targetY += 10000000000000
	}
	eq := mat.NewVecDense(2, []float64{
		float64(targetX),
		float64(targetY),
	})
	//fmt.Println(xy)
	//fmt.Println(eq)
	var sol mat.VecDense
	err := sol.SolveVec(xy, eq)
	if err != nil {
		return -1, -1
	} else {
		solIntX, errX := FloatToIntIfClose(sol.RawVector().Data[0], .001)
		if !errX {
			return -1, 0
		}
		solIntY, errY := FloatToIntIfClose(sol.RawVector().Data[1], .001)
		if !errY {
			return 0, -1
		}
		return solIntX, solIntY
	}

}

func FloatToIntIfClose(f float64, tolerance float64) (int, bool) {
	// Check if the float is within the tolerance of its nearest integer
	if math.Abs(f-math.Round(f)) < tolerance {
		return int(math.Round(f)), true
	}
	return 0, false
}
