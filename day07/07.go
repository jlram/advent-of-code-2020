package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

func readBags(newBags []string, validBags map[string]bool, input []string) int {
	reg, _ := regexp.Compile("(bags|bag)|[^a-zA-Z,]+")
	var aux []string // Var used to fill validBags each time

	for _, v := range newBags { // Iterate new level of bags
		for _, rule := range input { // Iterate input
			elements := strings.Split(rule, "bags contain")
			bag := strings.Replace(elements[0], " ", "", -1)
			colors := reg.ReplaceAllString(elements[1], "")

			// If the current bag is contained in any of these bags, we add the bag to validBags
			if strings.Contains(colors, v) {
				validBags[bag] = true
				aux = append(aux, bag)
			}
		}
	}
	// Check if there are new bags to iterate
	if len(aux) > 0 {
		newBags = aux
		readBags(newBags, validBags, input)
	}
	// If not, return the number of bags that could eventually contain one shiny gold bag
	return len(validBags)
}

func partOne(input []string) (int, time.Duration) {
	start := time.Now()
	newBags := []string{"shinygold"}
	validBags := make(map[string]bool)

	readBags(newBags, validBags, input)

	return len(validBags), time.Since(start)
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n")

	result, runtime := partOne(givenInput)
	fmt.Printf("[#] Part 1: RESULT: %d RUNTIME: %s", result, runtime)
}
