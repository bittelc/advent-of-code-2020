package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var filename = "input.txt"

// var filename = "example.txt"

func main() {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("looks good")
	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines)-2; i++ {
		for j := i + 1; j < len(lines); j++ {
			for k := j + 1; k < len(lines); k++ {
				a, err := strconv.Atoi(lines[i])
				if err != nil {
					panic(err)
				}
				b, err := strconv.Atoi(lines[j])
				if err != nil {
					panic(err)
				}
				c, err := strconv.Atoi(lines[k])
				if err != nil {
					panic(err)
				}
				if a+b+c == 2020 {
					fmt.Println("found it!")
					fmt.Println("a = ", a)
					fmt.Println("b = ", b)
					fmt.Println("c = ", c)
					fmt.Println("final answer:", a*b*c)
				}
			}
		}
	}

	fmt.Println("all done")
}
