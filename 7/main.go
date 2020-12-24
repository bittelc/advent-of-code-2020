package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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
	aB := parseAllBags()
	// printAllBags(&aB)
	createTree(&aB)
}

func createTree(aB *allBags) {
	for i, baag := range *aB {

	}
}

func parseAllBags() allBags {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	all := make(allBags, 0)
	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines); i++ {

		// establish first level of tree
		selfDesc := strings.Split(string(lines[i]), " bags ")[0]
		parsedBag := bag{selfDesc: selfDesc}
		parsedBag.children = make(map[*bag]int)

		// parse children of first level
		rawChildBags := strings.Split(string(lines[i]), " contain ")[1]
		splitter := regexp.MustCompile(` bag(s*)((\.)|(, ))`)
		allChildBags := splitter.Split(rawChildBags, -1)
		log.Printf("selfBag: \"%s\", childBags: \"%#v\", len(childBags): %d", selfDesc, allChildBags, len(allChildBags))
		for i := 0; i < len(allChildBags)-1; i++ {
			bagStr := allChildBags[i]
			numBag := strings.Split(bagStr, " ")[0]
			if numBag == "no" {
				continue
			}
			typeBag := strings.SplitN(bagStr, " ", 2)[1]
			baag := bag{selfDesc: typeBag}
			log.Printf("numBag: %s, typeBag: %s", numBag, typeBag)
			parsedBag.children[&baag], err = strconv.Atoi(numBag)
			if err != nil {
				panic(err)
			}
			log.Printf("parsedBag.children: %#v", parsedBag.children)
		}

		all = append(all, parsedBag)

	}
	log.Println("len(allBags):", len(all))
	return all
}

func printAllBags(aB *allBags) {
	fmt.Println("len(allBags):", len(*aB))
	for _, baag := range *aB {
		fmt.Println("----")
		printBag(&baag, "", 1)
	}
}

func printBag(b *bag, prefix string, depth int) {
	fmt.Printf("%s  depth: %d\n", prefix, depth)
	fmt.Printf("%s  bag.selfDesc: %s\n", prefix, (*b).selfDesc)
	fmt.Printf("%s  bag.containsMyBag: %t\n", prefix, (*b).containsMyBag)
	for childBag := range (*b).children {
		fmt.Printf("%s  bag.child:\n", prefix)
		printBag(childBag, fmt.Sprintf("%s  ", prefix), depth+1)
	}
}
