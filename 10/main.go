package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// var filename string = "example.txt"

// var filename string = "example2.txt"

var filename string = "input.txt"

type allAdaptors []int

func main() {
	fmt.Println("beginning")
	aA := parseInput()
	// printAllAdaptors(&aA)
	differences := findStepDifferences(&aA)
	fmt.Println("differences:", differences)

}

func findStepDifferences(aA *allAdaptors) int {
	diffByOne := 0
	diffByThree := 1
	currentJoltage := 0
	for {
		oneStep := isOneStepDifference(aA, currentJoltage)
		if oneStep {
			fmt.Printf("diff by one, currentJoltage: %d, diffByOne: %d, diffByThree: %d\n", currentJoltage, diffByOne, diffByThree)
			currentJoltage++
			diffByOne++
			continue
		}
		twoStep := isTwoStepDifference(aA, currentJoltage)
		if twoStep == true {
			currentJoltage = currentJoltage + 2
			continue
		}
		threeStep := isThreeStepDifference(aA, currentJoltage)
		if threeStep == true {
			fmt.Printf("diff by three, currentJoltage: %d, diffByOne: %d, diffByThree: %d\n", currentJoltage, diffByOne, diffByThree)
			currentJoltage = currentJoltage + 3
			diffByThree++
			continue
		}
		break
	}
	fmt.Printf("currentJoltage: %d, diffByOne: %d, diffByThree: %d\n", currentJoltage, diffByOne, diffByThree)
	return diffByOne * diffByThree
}

func isOneStepDifference(aA *allAdaptors, current int) bool {
	for _, adaptor := range *aA {
		if adaptor == current+1 {
			return true
		}
	}
	return false
}
func isTwoStepDifference(aA *allAdaptors, current int) bool {
	for _, adaptor := range *aA {
		if adaptor == current+2 {
			return true
		}
	}
	return false
}
func isThreeStepDifference(aA *allAdaptors, current int) bool {
	for _, adaptor := range *aA {
		if adaptor == current+3 {
			return true
		}
	}
	return false
}

func parseInput() allAdaptors {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	all := make([]int, 0)
	lines := strings.Split(string(dat), "\n")

	//parse entire input text
	for _, joltage := range lines {
		toI, err := strconv.Atoi(joltage)
		if err != nil {
			panic(err)
		}
		all = append(all, toI)
	}

	return all
}

func printAllAdaptors(aA *allAdaptors) {
	for _, jolt := range *aA {
		fmt.Println("joltage:", jolt)
	}
}
