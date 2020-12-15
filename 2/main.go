package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type input struct {
	min      int
	max      int
	char     string
	password string
}

var filename = "example.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func parseLine(line string) input {
	fmt.Println(line)
	var a input
	dashSplit := strings.Split(line, "-")
	min := dashSplit[0]
	max := strings.Split(dashSplit[1], " ")[0]
	charIndex := strings.Index(line, ":")
	fmt.Println("min =", min)
	fmt.Println("max =", max)
	fmt.Println("charIndex =", charIndex)
	return a
}
