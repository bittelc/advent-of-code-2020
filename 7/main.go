package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// var filename string = "example.txt"

var filename string = "input.txt"
var myBag = "shiny gold"

type allBags map[string][]string

func main() {
	fmt.Println("beginning")
	aB := parseAllBags()
	// printAllBags(&aB)
	count := findCount(&aB)
	log.Println("count:", count)
}

func findCount(aB *allBags) int {
	count := 0

	for bag := range *aB {
		found := traverseChildren(aB, bag)
		if found {
			count++
		}
	}
	return count
}

func traverseChildren(aB *allBags, bag string) bool {
	for _, j := range (*aB)[bag] {
		if j == myBag {
			return true
		}
		found := traverseChildren(aB, j)
		if found {
			return true
		}
	}
	return false
}

func parseAllBags() allBags {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	all := make(allBags, 0)
	lines := strings.Split(string(dat), "\n")

	//parse entire input text
	for i := 0; i < len(lines); i++ {

		// establish first level of tree
		selfDesc := strings.Split(string(lines[i]), " bags ")[0]
		all[selfDesc] = make([]string, 0)

		// parse second level
		rawChildBags := strings.Split(string(lines[i]), " contain ")[1]
		splitter := regexp.MustCompile(` bag(s*)((\.)|(, ))`)
		allChildBags := splitter.Split(rawChildBags, -1)
		for i := 0; i < len(allChildBags)-1; i++ {
			typeBag := strings.SplitN(allChildBags[i], " ", 2)[1]
			all[selfDesc] = append(all[selfDesc], typeBag)
		}
	}
	return all
}
func printAllBags(aB *allBags) {
	for a, b := range *aB {
		log.Println(a)
		log.Println("children:")
		for _, str := range b {
			log.Println("  ", str)
		}
		log.Println("------")
	}
}

// func printAllBags(aB *allBags) {
// 	fmt.Println("len(allBags):", len(*aB))
// 	for _, baag := range *aB {
// 		fmt.Println("----")
// 		printBag(&baag, "", 1)
// 	}
// }

// func printBag(b *bag, prefix string, depth int) {
// 	fmt.Printf("%s  depth: %d\n", prefix, depth)
// 	fmt.Printf("%s  bag.selfDesc: %s\n", prefix, (*b).selfDesc)
// 	fmt.Printf("%s  bag.containsMyBag: %t\n", prefix, (*b).containsMyBag)
// 	for childBag := range (*b).children {
// 		fmt.Printf("%s  bag.child:\n", prefix)
// 		printBag(childBag, fmt.Sprintf("%s  ", prefix), depth+1)
// 	}
// }
