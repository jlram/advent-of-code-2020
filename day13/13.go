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

// Bus --- struct we use later for our channel
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

func partTwo(input []string) int {
	timeStamp := 0
	initValue, _ := strconv.Atoi(input[0]) // 41
	found := false

	for !found {
		timeStamp += initValue
		for i := 0; i < len(input); {
			gap := calcGap(input[i+1:])
			nextNumber, err := strconv.Atoi(input[i+gap])

			if err == nil {
				if (timeStamp+i+gap)%nextNumber == 0 {
					i += gap
					if i == len(input)-1 {
						found = true
						break
					}
				} else {
					break
				}
			} else {
				println(err)
				break
			}
		}
	}

	return timeStamp
}

func calcGap(slice []string) int {
	total := 1
	for _, v := range slice {
		if v == "x" {
			total++
		} else {
			return total
		}
	}
	return 0
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
	fmt.Println(partTwo(strings.Split(givenInput[1], ",")))
}
