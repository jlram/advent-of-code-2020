package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// struct we use later for our channel
type Bus struct {
	key   int
	value int
}

func partOne(input []string) int {
	queue := make(chan Bus, 1)

	arrival, _ := strconv.Atoi(input[0])
	// extract numbers from chain
	re := regexp.MustCompile("[0-9]+")
	buses := re.FindAllString(input[1], -1)

	// init waitgroup
	wg.Add(len(buses))

	// calculate min interval for each bus
	for _, bus := range buses {
		go calcTime(arrival, bus, queue)
	}

	// get min interval overall
	minInterval := 9223372036854775807
	minID := 0
	go func() {
		for t := range queue {
			if t.value < minInterval {
				minInterval = t.value
				minID = t.key
			}
			wg.Done()
		}
	}()

	wg.Wait()

	return (minInterval - arrival) * minID
}

func calcTime(arrival int, item string, queue chan Bus) {
	numItem, _ := strconv.Atoi(item)
	initValue := numItem

	for numItem < arrival {
		numItem += initValue
	}

	bus := new(Bus)
	bus.key = initValue
	bus.value = numItem

	// add bus to channel
	queue <- *bus
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	givenInput := strings.Split(string(file), "\n")

	fmt.Println(partOne(givenInput))
}
