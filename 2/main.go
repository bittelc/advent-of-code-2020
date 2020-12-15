package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type input struct {
	min      int
	max      int
	char     string
	password string
}

// var filename = "example.txt"
var filename = "input.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalAcceptable := 0
	for scanner.Scan() {
		candidate := parseLine(scanner.Text())
		count := strings.Count(candidate.password, candidate.char)
		if count >= candidate.min && count <= candidate.max {
			totalAcceptable++
		}
		fmt.Println("totalAcceptable =", totalAcceptable)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func parseLine(line string) input {
	dashSplit := strings.Split(line, "-")
	min, err := strconv.Atoi(dashSplit[0])
	if err != nil {
		log.Fatal(err)
	}
	spaceSplit := strings.Split(dashSplit[1], " ")
	max, err := strconv.Atoi(spaceSplit[0])
	if err != nil {
		log.Fatal(err)
	}
	charIndex := strings.Index(line, ":")
	char := string(line[charIndex-1])
	password := spaceSplit[len(spaceSplit)-1]
	return input{min, max, char, password}
}
