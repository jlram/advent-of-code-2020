package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func readBagsOne(newBags []string, validBags map[string]bool, input []string) int {
	reg := regexp.MustCompile("(bags|bag)|[^a-zA-Z,]+")
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
		readBagsOne(newBags, validBags, input)
	}
	// If not, return the number of bags that could eventually contain one shiny gold bag
	return len(validBags)
}

func readBagsTwo(newBags []string, validBags map[string]int, mapBags map[string]string) int {
	var aux []string // Var used to fill validBags each time
	for _, v := range newBags {
		totalValue := 1
		for _, value := range strings.Split(mapBags[v], ",") {
			re := regexp.MustCompile("[0-9]+")
			charvalue := re.ReplaceAllString(value, "")
			reFind := re.FindAllString(value, -1)
			if len(reFind) > 0 {
				numvalue, _ := strconv.Atoi(reFind[0])
				validBags[charvalue] += numvalue
				totalValue = totalValue * numvalue
				fmt.Println(totalValue)
				aux = append(aux, charvalue)
			}
		}
	}

	if len(aux) > 0 {
		newBags = aux
		readBagsTwo(newBags, validBags, mapBags)
	}

	total := 0
	for _, v := range validBags {
		total += v
	}

	// fmt.Println(validBags)
	return total
}

func partOne(input []string) (int, time.Duration) {
	start := time.Now()
	newBags := []string{"shinygold"}
	validBags := make(map[string]bool)

	readBagsOne(newBags, validBags, input)

	return len(validBags), time.Since(start)
}

func partTwo(input []string) int {
	newBags := []string{"shinygold"}
	validBags := make(map[string]int)

	mapBags := make(map[string]string)
	reg := regexp.MustCompile("(bags|bag)|[^a-zA-Z0-9,]+")

	for _, rule := range input {
		elements := strings.Split(rule, "bags contain")
		bag := strings.Replace(elements[0], " ", "", -1)
		colors := reg.ReplaceAllString(elements[1], "")
		mapBags[bag] = colors
	}

	return readBagsTwo(newBags, validBags, mapBags)
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n")

	result, runtime := partOne(givenInput)
	fmt.Printf("[#] Part 1: RESULT: %d RUNTIME: %s\n", result, runtime)

	fmt.Println(partTwo(givenInput))
}
