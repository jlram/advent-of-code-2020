package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(input []int) int {
	min, max := 0, 25
	toIterate := input[25:]

	for _, value := range toIterate {
		arr := make([]int, 25)
		copy(arr, input[min:max])
		valid := false
		for _, number := range arr {
			for _, number2 := range arr {
				if number+number2 == value && number != number2 {
					valid = true
				}
			}
		}

		min++
		max++

		if !valid {
			return value
		}
	}
	return 0
}

func partTwo(input []int, value int) int {
	min, max := 0, 24
	valid := false
	for !valid {
		arr := make([]int, 25)
		copy(arr, input[min:max])
		for _, number := range arr {
			for _, number2 := range arr {
				if number+number2 == value && number != number2 {
					valid = true
				}
			}
		}
		if valid {
			return input[min] + input[max-1]
		}
		min++
		max++
	}

	return 0
}

func main() {
	var intInput []int
	file, _ := os.Open("input.txt")
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		intInput = append(intInput, num)
	}
	resultOne := partOne(intInput)
	fmt.Println(resultOne)
	fmt.Println(partTwo(intInput, resultOne))
}
