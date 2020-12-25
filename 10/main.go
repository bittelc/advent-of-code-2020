package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// var filename string = "example.txt"

// var filename string = "example2.txt"

var filename string = "input.txt"

type allAdaptors []int

func main() {
	aA := parseInput()
	// printAllAdaptors(&aA)
	differences := findStepDifferences(&aA)
	fmt.Println("differences:", differences)
	fmt.Println("----- pt II ------:")
	allVariations := findVariations(&aA)
	fmt.Println("allVariations:", allVariations)
}

func findVariations(aA *allAdaptors) int64 {
	possible := make(map[int][]int)
	possible[0] = []int{1, 2, 3}
	for _, val := range *aA {
		possible[val] = []int{val + 1, val + 2, val + 3}
	}
	return connections(possible, make(map[int]int), (*aA)[len(*aA)-1]+3, 0)
}

func connections(possible map[int][]int, memo map[int]int, target int, currPos int) int64 {
	if value, seen := memo[currPos]; seen {
		return int64(value)
	}

	value := int64(0)
	for _, current := range possible[currPos] {
		if current != target {
			value += connections(possible, memo, target, current)
			continue
		}

		value++
	}

	memo[currPos] = int(value)
	return value
}

func findStepDifferences(aA *allAdaptors) int {
	diffByOne := 0
	diffByThree := 1
	currentJoltage := 0
	for {
		oneStep := isOneStepDifference(aA, currentJoltage)
		if oneStep {
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
			currentJoltage = currentJoltage + 3
			diffByThree++
			continue
		}
		break
	}
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
	sort.Ints(all)
	return all
}

func printAllAdaptors(aA *allAdaptors) {
	for _, jolt := range *aA {
		fmt.Println("joltage:", jolt)
	}
}
