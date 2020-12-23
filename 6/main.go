package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type plane []group
type group []person
type person []rune

// var filename string = "input.txt"

var filename string = "example.txt"

func main() {
	var parsedPlane plane
	parsedPlane = parsePlane(filename)
	totalYeses := sumPlaneYeses(&parsedPlane)
	fmt.Println("totalYeses:", totalYeses)
	fmt.Println("part1 complete")
	sharedYeses := sumSharedYeses(&parsedPlane)
	fmt.Println("sharedYeses:", sharedYeses)

}

func parsePlane(file string) plane {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	formedPlane := make(plane, 0)
	formedGroup := make(group, 0)
	formedPerson := make(person, 0)

	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			formedPlane = append(formedPlane, formedGroup)
			formedGroup = group{}
			continue
		}
		for _, char := range lines[i] {
			formedPerson = append(formedPerson, char)
		}
		formedGroup = append(formedGroup, formedPerson)
		formedPerson = person{}
	}

	formedPlane = append(formedPlane, formedGroup)
	return formedPlane
}

func sumPlaneYeses(pl *plane) int {
	totalYeses := 0
	for _, grp := range *pl {
		thisGroupYeses := getGroupYeses(&grp)
		totalYeses = totalYeses + thisGroupYeses
	}
	return totalYeses
}

func getGroupYeses(grp *group) int {
	uniqueYeses := make(map[rune]int)
	res := make([]rune, 0)
	var rawJoined []rune
	for _, prs := range *grp {
		rawJoined = append(rawJoined, prs...)
	}
	for _, char := range rawJoined {
		uniqueYeses[char] = 1
	}
	for letter := range uniqueYeses {
		res = append(res, letter)
	}
	return len(res)
}

func sumSharedYeses(pl *plane) int {
	totalYeses := 0
	for _, grp := range *pl {
		thisGroupYeses := groupSharedYeses(&grp)
		totalYeses = totalYeses + thisGroupYeses
	}
	return totalYeses
}

func groupSharedYeses(grp *group) int {
	uniqueYeses := make(map[rune]int)
	var rawJoined []rune
	res := make([]rune, 0)
	for _, prs := range *grp {
		rawJoined = append(rawJoined, prs...)
	}
	for _, char := range rawJoined {
		uniqueYeses[char] = uniqueYeses[char] + 1
	}
	for letter, i := range uniqueYeses {
		if i == len(*grp) {
			res = append(res, letter)
		}
	}
	return len(res)

	return 0
}

func printPlane(pl *plane) {
	for _, grp := range *pl {
		printGrp(&grp)
		fmt.Println("")
	}
}

func printGrp(grp *group) {
	for _, prsn := range *grp {
		printPerson(&prsn)
		fmt.Println("")
	}
}

func printPerson(prs *person) {
	for _, char := range *prs {
		fmt.Printf(string(char))
	}
}
