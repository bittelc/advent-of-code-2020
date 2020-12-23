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
	printPlane(&parsedPlane)
	int := sumGroupYeses(&parsedPlane)
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

func sumGroupYeses(pl *plane) int {
	totalYeses := 0
	for i, grp := range *pl {
		thisGroupYeses := getGroupYeses(&grp)
		fmt.Printf("grp %d yeses: %d", i, thisGroupYeses)
		totalYeses = +thisGroupYeses
	}
	return totalYeses
}

func getGroupYeses(grp *group) int {
	uniqueYeses := make([]string, 0)
	for _, prsn := range *grp {
	}
	return len(uniqueYeses)
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
