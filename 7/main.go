package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var filename string = "example.txt"

// var filename string = "input.txt"
var myBag = "shiny gold"

type bag struct {
	children      map[*bag]int
	selfDesc      string
	containsMyBag bool
}

type allBags []bag

func main() {
	fmt.Println("beginning")
	_ = parseAllBags()
}

func parseAllBags() allBags {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	all := make(allBags, 0)
	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines); i++ {
		selfBag := strings.Split(string(lines[i]), " bags ")[0]
		log.Println("selfBag:", selfBag)
	}
	return all

}
