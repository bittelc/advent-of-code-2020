package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// var filename = "example.txt"

var filename = "input.txt"

type input struct {
	inputString string
	row         int
	column      int
	seatID      int
}

func main() {
	formedInput := parseToMem()
	assignRow(&formedInput)
	assignColumn(&formedInput)
	assignIndex(&formedInput)
	largestInput := findLargestSeatID(&formedInput)
	fmt.Println("pt 1. max:", largestInput.seatID)
	fmt.Println("pt 2. sublime text sort")
	fmt.Println(formedInput)
}

// func parseLine(line string) input {
func parseToMem() []input {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	formedInput := make([]input, len(lines))
	for i := 0; i < len(lines); i++ {
		formedInput[i].inputString = lines[i]
	}

	return formedInput
}

func assignRow(f *[]input) {
	for i := 0; i < len(*f); i++ {
		rowStr := (*f)[i].inputString[:7]
		rowNum, err := binarySplit(rowStr, "F", "B")
		if err != nil {
			panic(err)
		}
		(*f)[i].row = rowNum
	}
}

func assignColumn(f *[]input) {
	for i := 0; i < len(*f); i++ {
		columnStr := (*f)[i].inputString[7:10]
		columnNum, err := binarySplit(columnStr, "L", "R")
		if err != nil {
			panic(err)
		}
		(*f)[i].column = columnNum
	}
}

func assignIndex(f *[]input) {
	for i := 0; i < len(*f); i++ {
		(*f)[i].seatID = (*f)[i].row*8 + (*f)[i].column
	}
}

func findLargestSeatID(f *[]input) input {
	var max input
	for i := 0; i < len(*f); i++ {
		if (*f)[i].seatID > max.seatID {
			max = (*f)[i]
		}
	}
	return max
}

func binarySplit(raw, low, high string) (int, error) {
	l := math.Pow(2, float64(len(raw)))
	within := []int{0, int(l) - 1}
	for i := 0; i < len(raw); i++ {
		c := string(raw[i])
		diff := within[1] - within[0]
		if c == low {
			within[1] = within[1] - (diff / 2) - 1
		}
		if c == high {
			within[0] = within[0] + (diff / 2) + 1
		}
	}
	if within[1] != within[0] {
		return 0, fmt.Errorf("couldn't settle on single val, low = %d, high = %d", within[0], within[1])
	}
	return within[0], nil
}
